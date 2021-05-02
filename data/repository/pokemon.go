package repository

import (
	"bekapod/pkmn-team-graphql/data/db"
	"bekapod/pkmn-team-graphql/data/model"
	"context"
	"errors"
	"fmt"
)

type Pokemon struct {
	client *db.PrismaClient
}

func NewPokemon(client *db.PrismaClient) Pokemon {
	return Pokemon{
		client: client,
	}
}

func (r Pokemon) GetPokemon(ctx context.Context) (*model.PokemonList, error) {
	pokemon := model.NewEmptyPokemonList()

	results, err := r.client.Pokemon.FindMany().
		With(db.Pokemon.EggGroups.Fetch()).
		Exec(ctx)

	if errors.Is(err, db.ErrNotFound) {
		return &pokemon, nil
	}

	if err != nil {
		return &pokemon, fmt.Errorf("error getting pokemon: %s", err)
	}

	for _, result := range results {
		p := model.NewPokemonFromDb(result)
		pokemon.AddPokemon(&p)
	}

	return &pokemon, nil
}

func (r Pokemon) GetPokemonById(ctx context.Context, id string) (*model.Pokemon, error) {
	result, err := r.client.Pokemon.
		FindUnique(db.Pokemon.ID.Equals(id)).
		With(db.Pokemon.EggGroups.Fetch()).
		Exec(ctx)

	if errors.Is(err, db.ErrNotFound) {
		return nil, fmt.Errorf("couldn't find pokemon by id: %s", id)
	}

	if err != nil {
		return nil, fmt.Errorf("error getting pokemon by id: %s, error: %s", id, err)
	}

	p := model.NewPokemonFromDb(*result)
	return &p, nil
}

func (r Pokemon) PokemonByIdDataLoader(ctx context.Context) func(ids []string) ([]*model.Pokemon, []error) {
	return func(ids []string) ([]*model.Pokemon, []error) {
		pokemonById := map[string]*model.Pokemon{}
		results, err := r.client.Pokemon.FindMany(db.Pokemon.ID.In(ids)).With(db.Pokemon.EggGroups.Fetch()).Exec(ctx)
		pokemon := make([]*model.Pokemon, len(ids))

		if err != nil {
			errors := make([]error, len(ids))
			for i, id := range ids {
				errors[i] = fmt.Errorf("error loading type by id %s in dataloader %w", id, err)
			}

			return pokemon, errors
		}

		for _, result := range results {
			p := model.NewPokemonFromDb(result)
			pokemonById[p.ID] = &p
		}

		for i, id := range ids {
			pokemon[i] = pokemonById[id]
		}

		return pokemon, nil
	}
}
