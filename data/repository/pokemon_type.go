package repository

import (
	"bekapod/pkmn-team-graphql/data/db"
	"bekapod/pkmn-team-graphql/data/model"
	"context"
	"fmt"
)

type PokemonType struct {
	client *db.PrismaClient
}

func NewPokemonType(client *db.PrismaClient) PokemonType {
	return PokemonType{
		client: client,
	}
}

func (r PokemonType) PokemonTypeByPokemonIdDataLoader(ctx context.Context) func(ids []string) ([]*model.PokemonTypeList, []error) {
	return func(ids []string) ([]*model.PokemonTypeList, []error) {
		pokemonTypeListsById := map[string]*model.PokemonTypeList{}
		results, err := r.client.PokemonType.
			FindMany(db.PokemonType.PokemonID.In(ids)).
			Exec(ctx)
		pokemonTypeLists := make([]*model.PokemonTypeList, len(ids))

		if err != nil {
			errors := make([]error, len(ids))
			for i, id := range ids {
				tl := model.NewEmptyPokemonTypeList()
				pokemonTypeLists[i] = &tl
				errors[i] = fmt.Errorf("error loading pokemon type by pokemon id %s in dataloader %w", id, err)
			}

			return pokemonTypeLists, errors
		}

		if len(results) == 0 {
			for i := range ids {
				tl := model.NewEmptyPokemonTypeList()
				pokemonTypeLists[i] = &tl
			}

			return pokemonTypeLists, nil
		}

		for _, result := range results {
			tl := pokemonTypeListsById[result.PokemonID]
			if tl == nil {
				empty := model.NewEmptyPokemonTypeList()
				pokemonTypeListsById[result.PokemonID] = &empty
			}
			t := model.NewPokemonTypeFromDb(result)
			pokemonTypeListsById[result.PokemonID].AddPokemonType(&t)
		}

		for i, id := range ids {
			pokemonTypeList := pokemonTypeListsById[id]

			if pokemonTypeList == nil {
				empty := model.NewEmptyPokemonTypeList()
				pokemonTypeList = &empty
			}

			pokemonTypeLists[i] = pokemonTypeList
		}

		return pokemonTypeLists, nil
	}
}

func (r PokemonType) PokemonTypeByTypeIdDataLoader(ctx context.Context) func(ids []string) ([]*model.PokemonTypeList, []error) {
	return func(ids []string) ([]*model.PokemonTypeList, []error) {
		pokemonTypeListsById := map[string]*model.PokemonTypeList{}
		results, err := r.client.PokemonType.
			FindMany(db.PokemonType.TypeID.In(ids)).
			Exec(ctx)
		pokemonTypeLists := make([]*model.PokemonTypeList, len(ids))

		if err != nil {
			errors := make([]error, len(ids))
			for i, id := range ids {
				tl := model.NewEmptyPokemonTypeList()
				pokemonTypeLists[i] = &tl
				errors[i] = fmt.Errorf("error loading pokemon type by type id %s in dataloader %w", id, err)
			}

			return pokemonTypeLists, errors
		}

		if len(results) == 0 {
			for i := range ids {
				tl := model.NewEmptyPokemonTypeList()
				pokemonTypeLists[i] = &tl
			}

			return pokemonTypeLists, nil
		}

		for _, result := range results {
			tl := pokemonTypeListsById[result.TypeID]
			if tl == nil {
				empty := model.NewEmptyPokemonTypeList()
				pokemonTypeListsById[result.TypeID] = &empty
			}
			t := model.NewPokemonTypeFromDb(result)
			pokemonTypeListsById[result.TypeID].AddPokemonType(&t)
		}

		for i, id := range ids {
			pokemonTypeList := pokemonTypeListsById[id]

			if pokemonTypeList == nil {
				empty := model.NewEmptyPokemonTypeList()
				pokemonTypeList = &empty
			}

			pokemonTypeLists[i] = pokemonTypeList
		}

		return pokemonTypeLists, nil
	}
}
