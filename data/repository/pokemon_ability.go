package repository

import (
	"bekapod/pkmn-team-graphql/data/db"
	"bekapod/pkmn-team-graphql/data/model"
	"context"
	"fmt"
)

type PokemonAbility struct {
	client *db.PrismaClient
}

func NewPokemonAbility(client *db.PrismaClient) PokemonAbility {
	return PokemonAbility{
		client: client,
	}
}

func (r PokemonAbility) PokemonAbilityByPokemonIdDataLoader(ctx context.Context) func(ids []string) ([]*model.PokemonAbilityConnection, []error) {
	return func(ids []string) ([]*model.PokemonAbilityConnection, []error) {
		pokemonAbilityConnectionsById := map[string]*model.PokemonAbilityConnection{}
		results, err := r.client.PokemonAbility.
			FindMany(db.PokemonAbility.PokemonID.In(ids)).
			Exec(ctx)
		pokemonAbilityConnections := make([]*model.PokemonAbilityConnection, len(ids))

		if err != nil {
			errors := make([]error, len(ids))
			for i, id := range ids {
				al := model.NewEmptyPokemonAbilityConnection()
				pokemonAbilityConnections[i] = &al
				errors[i] = fmt.Errorf("error loading pokemon ability by pokemon id %s in dataloader %w", id, err)
			}

			return pokemonAbilityConnections, errors
		}

		if len(results) == 0 {
			for i := range ids {
				c := model.NewEmptyPokemonAbilityConnection()
				pokemonAbilityConnections[i] = &c
			}

			return pokemonAbilityConnections, nil
		}

		for _, result := range results {
			al := pokemonAbilityConnectionsById[result.PokemonID]
			if al == nil {
				empty := model.NewEmptyPokemonAbilityConnection()
				pokemonAbilityConnectionsById[result.PokemonID] = &empty
			}
			e := model.NewPokemonAbilityEdgeFromDb(result)
			pokemonAbilityConnectionsById[result.PokemonID].AddEdge(&e)
		}

		for i, id := range ids {
			pokemonAbilityConnection := pokemonAbilityConnectionsById[id]

			if pokemonAbilityConnection == nil {
				empty := model.NewEmptyPokemonAbilityConnection()
				pokemonAbilityConnection = &empty
			}

			pokemonAbilityConnections[i] = pokemonAbilityConnection
		}

		return pokemonAbilityConnections, nil
	}
}

func (r PokemonAbility) PokemonAbilityByAbilityIdDataLoader(ctx context.Context) func(ids []string) ([]*model.PokemonWithAbilityConnection, []error) {
	return func(ids []string) ([]*model.PokemonWithAbilityConnection, []error) {
		pokemonAbilityConnectionsById := map[string]*model.PokemonWithAbilityConnection{}
		results, err := r.client.PokemonAbility.
			FindMany(db.PokemonAbility.AbilityID.In(ids)).
			Exec(ctx)
		pokemonAbilityConnections := make([]*model.PokemonWithAbilityConnection, len(ids))

		if err != nil {
			errors := make([]error, len(ids))
			for i, id := range ids {
				al := model.NewEmptyPokemonWithAbilityConnection()
				pokemonAbilityConnections[i] = &al
				errors[i] = fmt.Errorf("error loading pokemon ability by ability id %s in dataloader %w", id, err)
			}

			return pokemonAbilityConnections, errors
		}

		if len(results) == 0 {
			for i := range ids {
				al := model.NewEmptyPokemonWithAbilityConnection()
				pokemonAbilityConnections[i] = &al
			}

			return pokemonAbilityConnections, nil
		}

		for _, result := range results {
			al := pokemonAbilityConnectionsById[result.AbilityID]
			if al == nil {
				empty := model.NewEmptyPokemonWithAbilityConnection()
				pokemonAbilityConnectionsById[result.AbilityID] = &empty
			}
			a := model.NewPokemonWithAbilityEdgeFromDb(result)
			pokemonAbilityConnectionsById[result.AbilityID].AddEdge(&a)
		}

		for i, id := range ids {
			pokemonAbilityList := pokemonAbilityConnectionsById[id]

			if pokemonAbilityList == nil {
				empty := model.NewEmptyPokemonWithAbilityConnection()
				pokemonAbilityList = &empty
			}

			pokemonAbilityConnections[i] = pokemonAbilityList
		}

		return pokemonAbilityConnections, nil
	}
}
