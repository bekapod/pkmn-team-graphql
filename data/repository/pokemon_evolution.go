package repository

import (
	"bekapod/pkmn-team-graphql/data/db"
	"bekapod/pkmn-team-graphql/data/model"
	"context"
	"fmt"
)

type PokemonEvolution struct {
	client *db.PrismaClient
}

func NewPokemonEvolution(client *db.PrismaClient) PokemonEvolution {
	return PokemonEvolution{
		client: client,
	}
}

func (r PokemonEvolution) PokemonEvolutionByFromPokemonIdDataLoader(ctx context.Context) func(ids []string) ([]*model.PokemonEvolutionList, []error) {
	return func(ids []string) ([]*model.PokemonEvolutionList, []error) {
		pokemonEvolutionListsById := map[string]*model.PokemonEvolutionList{}
		results, err := r.client.PokemonEvolution.
			FindMany(db.PokemonEvolution.FromPokemonID.In(ids)).
			Exec(ctx)
		pokemonEvolutionLists := make([]*model.PokemonEvolutionList, len(ids))

		if err != nil {
			errors := make([]error, len(ids))
			for i, id := range ids {
				el := model.NewEmptyPokemonEvolutionList()
				pokemonEvolutionLists[i] = &el
				errors[i] = fmt.Errorf("error loading pokemon evolution by pokemon id %s in dataloader %w", id, err)
			}

			return pokemonEvolutionLists, errors
		}

		if len(results) == 0 {
			for i := range ids {
				el := model.NewEmptyPokemonEvolutionList()
				pokemonEvolutionLists[i] = &el
			}

			return pokemonEvolutionLists, nil
		}

		for _, result := range results {
			el := pokemonEvolutionListsById[result.FromPokemonID]
			if el == nil {
				empty := model.NewEmptyPokemonEvolutionList()
				pokemonEvolutionListsById[result.FromPokemonID] = &empty
			}
			e := model.NewPokemonEvolutionFromDb(result)
			e.PokemonID = e.ToPokemonID
			pokemonEvolutionListsById[result.FromPokemonID].AddPokemonEvolution(&e)
		}

		for i, id := range ids {
			pokemonEvolutionList := pokemonEvolutionListsById[id]

			if pokemonEvolutionList == nil {
				empty := model.NewEmptyPokemonEvolutionList()
				pokemonEvolutionList = &empty
			}

			pokemonEvolutionLists[i] = pokemonEvolutionList
		}

		return pokemonEvolutionLists, nil
	}
}

func (r PokemonEvolution) PokemonEvolutionByToPokemonIdDataLoader(ctx context.Context) func(ids []string) ([]*model.PokemonEvolutionList, []error) {
	return func(ids []string) ([]*model.PokemonEvolutionList, []error) {
		pokemonEvolutionListsById := map[string]*model.PokemonEvolutionList{}
		results, err := r.client.PokemonEvolution.
			FindMany(db.PokemonEvolution.ToPokemonID.In(ids)).
			Exec(ctx)
		pokemonEvolutionLists := make([]*model.PokemonEvolutionList, len(ids))

		if err != nil {
			errors := make([]error, len(ids))
			for i, id := range ids {
				el := model.NewEmptyPokemonEvolutionList()
				pokemonEvolutionLists[i] = &el
				errors[i] = fmt.Errorf("error loading pokemon evolution by pokemon id %s in dataloader %w", id, err)
			}

			return pokemonEvolutionLists, errors
		}

		if len(results) == 0 {
			for i := range ids {
				el := model.NewEmptyPokemonEvolutionList()
				pokemonEvolutionLists[i] = &el
			}

			return pokemonEvolutionLists, nil
		}

		for _, result := range results {
			el := pokemonEvolutionListsById[result.ToPokemonID]
			if el == nil {
				empty := model.NewEmptyPokemonEvolutionList()
				pokemonEvolutionListsById[result.ToPokemonID] = &empty
			}
			e := model.NewPokemonEvolutionFromDb(result)
			e.PokemonID = e.FromPokemonID
			pokemonEvolutionListsById[result.ToPokemonID].AddPokemonEvolution(&e)
		}

		for i, id := range ids {
			pokemonEvolutionList := pokemonEvolutionListsById[id]

			if pokemonEvolutionList == nil {
				empty := model.NewEmptyPokemonEvolutionList()
				pokemonEvolutionList = &empty
			}

			pokemonEvolutionLists[i] = pokemonEvolutionList
		}

		return pokemonEvolutionLists, nil
	}
}
