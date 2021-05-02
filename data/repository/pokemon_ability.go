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

func (r PokemonAbility) PokemonAbilityByPokemonIdDataLoader(ctx context.Context) func(ids []string) ([]*model.PokemonAbilityList, []error) {
	return func(ids []string) ([]*model.PokemonAbilityList, []error) {
		pokemonAbilityListsById := map[string]*model.PokemonAbilityList{}
		results, err := r.client.PokemonAbility.
			FindMany(db.PokemonAbility.PokemonID.In(ids)).
			Exec(ctx)
		pokemonAbilityLists := make([]*model.PokemonAbilityList, len(ids))

		if err != nil {
			errors := make([]error, len(ids))
			for i, id := range ids {
				al := model.NewEmptyPokemonAbilityList()
				pokemonAbilityLists[i] = &al
				errors[i] = fmt.Errorf("error loading pokemon ability by pokemon id %s in dataloader %w", id, err)
			}

			return pokemonAbilityLists, errors
		}

		if len(results) == 0 {
			for i := range ids {
				al := model.NewEmptyPokemonAbilityList()
				pokemonAbilityLists[i] = &al
			}

			return pokemonAbilityLists, nil
		}

		for _, result := range results {
			al := pokemonAbilityListsById[result.PokemonID]
			if al == nil {
				empty := model.NewEmptyPokemonAbilityList()
				pokemonAbilityListsById[result.PokemonID] = &empty
			}
			a := model.NewPokemonAbilityFromDb(result)
			pokemonAbilityListsById[result.PokemonID].AddPokemonAbility(&a)
		}

		for i, id := range ids {
			pokemonAbilityList := pokemonAbilityListsById[id]

			if pokemonAbilityList == nil {
				empty := model.NewEmptyPokemonAbilityList()
				pokemonAbilityList = &empty
			}

			pokemonAbilityLists[i] = pokemonAbilityList
		}

		return pokemonAbilityLists, nil
	}
}

func (r PokemonAbility) PokemonAbilityByAbilityIdDataLoader(ctx context.Context) func(ids []string) ([]*model.PokemonAbilityList, []error) {
	return func(ids []string) ([]*model.PokemonAbilityList, []error) {
		pokemonAbilityListsById := map[string]*model.PokemonAbilityList{}
		results, err := r.client.PokemonAbility.
			FindMany(db.PokemonAbility.AbilityID.In(ids)).
			Exec(ctx)
		pokemonAbilityLists := make([]*model.PokemonAbilityList, len(ids))

		if err != nil {
			errors := make([]error, len(ids))
			for i, id := range ids {
				al := model.NewEmptyPokemonAbilityList()
				pokemonAbilityLists[i] = &al
				errors[i] = fmt.Errorf("error loading pokemon ability by ability id %s in dataloader %w", id, err)
			}

			return pokemonAbilityLists, errors
		}

		if len(results) == 0 {
			for i := range ids {
				al := model.NewEmptyPokemonAbilityList()
				pokemonAbilityLists[i] = &al
			}

			return pokemonAbilityLists, nil
		}

		for _, result := range results {
			al := pokemonAbilityListsById[result.AbilityID]
			if al == nil {
				empty := model.NewEmptyPokemonAbilityList()
				pokemonAbilityListsById[result.AbilityID] = &empty
			}
			a := model.NewPokemonAbilityFromDb(result)
			pokemonAbilityListsById[result.AbilityID].AddPokemonAbility(&a)
		}

		for i, id := range ids {
			pokemonAbilityList := pokemonAbilityListsById[id]

			if pokemonAbilityList == nil {
				empty := model.NewEmptyPokemonAbilityList()
				pokemonAbilityList = &empty
			}

			pokemonAbilityLists[i] = pokemonAbilityList
		}

		return pokemonAbilityLists, nil
	}
}
