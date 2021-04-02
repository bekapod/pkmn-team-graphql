package repository

import (
	"bekapod/pkmn-team-graphql/data/model"
	"bekapod/pkmn-team-graphql/log"
	"context"
	"fmt"
	"strings"
)

func (r EggGroup) EggGroupsByPokemonIdDataLoader(ctx context.Context) func(pokemonIds []string) ([]*model.EggGroupList, []error) {
	return func(pokemonIds []string) ([]*model.EggGroupList, []error) {
		eggGroupsByPokemonId := map[string]*model.EggGroupList{}
		placeholders := make([]string, len(pokemonIds))
		args := make([]interface{}, len(pokemonIds))
		for i := 0; i < len(pokemonIds); i++ {
			placeholders[i] = fmt.Sprintf("$%d", i+1)
			args[i] = pokemonIds[i]
		}

		query := "SELECT id, name, slug, pokemon_egg_group.pokemon_id FROM egg_groups LEFT JOIN pokemon_egg_group ON egg_groups.id = pokemon_egg_group.egg_group_id WHERE pokemon_egg_group.pokemon_id IN (" + strings.Join(placeholders, ",") + ")"

		log.Logger.WithField("args", args).Debug(query)
		rows, err := r.db.QueryContext(ctx,
			query,
			args...,
		)
		if err != nil {
			eggGroupList := make([]*model.EggGroupList, len(pokemonIds))
			emptyEggGroupList := model.NewEmptyEggGroupList()
			errors := make([]error, len(pokemonIds))
			for i := range pokemonIds {
				eggGroupList[i] = &emptyEggGroupList
				errors[i] = fmt.Errorf("error fetching egg groups for pokemon in EggGroupsByPokemonIdDataLoader: %w", err)
			}
			return eggGroupList, errors
		}

		defer rows.Close()
		for rows.Next() {
			var e model.EggGroup
			var pokemonId string
			err := rows.Scan(&e.ID, &e.Name, &e.Slug, &pokemonId)
			if err != nil {
				eggGroupList := make([]*model.EggGroupList, len(pokemonIds))
				emptyEggGroupList := model.NewEmptyEggGroupList()
				errors := make([]error, len(pokemonIds))
				for i := range pokemonIds {
					eggGroupList[i] = &emptyEggGroupList
					errors[i] = fmt.Errorf("error scanning result egg groups for pokemon in EggGroupsByPokemonIdDataLoader: %w", err)
				}
				return eggGroupList, errors
			}

			_, ok := eggGroupsByPokemonId[pokemonId]
			if !ok {
				el := model.NewEmptyEggGroupList()
				eggGroupsByPokemonId[pokemonId] = &el
			}

			eggGroupsByPokemonId[pokemonId].AddEggGroup(&e)
		}

		eggGroupList := make([]*model.EggGroupList, len(pokemonIds))
		for i, id := range pokemonIds {
			eggGroupList[i] = eggGroupsByPokemonId[id]
			i++
		}

		err = rows.Err()
		if err != nil {
			errors := make([]error, len(pokemonIds))
			for i := range pokemonIds {
				errors[i] = fmt.Errorf("error after fetching egg groups for pokemon in EggGroupsByPokemonIdDataLoader: %w", err)
			}
			return eggGroupList, errors
		}

		return eggGroupList, nil
	}
}
