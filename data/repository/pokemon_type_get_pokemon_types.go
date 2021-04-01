package repository

import (
	"bekapod/pkmn-team-graphql/data/model"
	"bekapod/pkmn-team-graphql/log"
	"context"
	"fmt"
	"strings"
)

func (r PokemonType) PokemonTypesByPokemonIdDataLoader(ctx context.Context) func(pokemonIds []string) ([]*model.PokemonTypeList, []error) {
	return func(pokemonIds []string) ([]*model.PokemonTypeList, []error) {
		pokemonTypesByPokemonId := map[string]*model.PokemonTypeList{}
		placeholders := make([]string, len(pokemonIds))
		args := make([]interface{}, len(pokemonIds))
		for i := 0; i < len(pokemonIds); i++ {
			placeholders[i] = fmt.Sprintf("$%d", i+1)
			args[i] = pokemonIds[i]
		}

		query := "SELECT pokemon_id, type_id, slot FROM pokemon_type WHERE pokemon_type.pokemon_id IN (" + strings.Join(placeholders, ",") + ") ORDER BY slot"

		log.Logger.WithField("args", args).Debug(query)
		rows, err := r.db.QueryContext(ctx,
			query,
			args...,
		)
		if err != nil {
			typeList := make([]*model.PokemonTypeList, len(pokemonIds))
			emptyPokemonTypeList := model.NewEmptyPokemonTypeList()
			errors := make([]error, len(pokemonIds))
			for i := range pokemonIds {
				typeList[i] = &emptyPokemonTypeList
				errors[i] = fmt.Errorf("error fetching types for pokemon in PokemonTypesByPokemonIdDataLoader: %w", err)
			}
			return typeList, errors
		}

		defer rows.Close()
		for rows.Next() {
			var t model.PokemonType
			err := rows.Scan(&t.PokemonID, &t.TypeID, &t.Slot)
			if err != nil {
				typeList := make([]*model.PokemonTypeList, len(pokemonIds))
				emptyPokemonTypeList := model.NewEmptyPokemonTypeList()
				errors := make([]error, len(pokemonIds))
				for i := range pokemonIds {
					typeList[i] = &emptyPokemonTypeList
					errors[i] = fmt.Errorf("error scanning result in PokemonTypesByPokemonIdDataLoader: %w", err)
				}
				return typeList, errors
			}

			_, ok := pokemonTypesByPokemonId[t.PokemonID]
			if !ok {
				tl := model.NewEmptyPokemonTypeList()
				pokemonTypesByPokemonId[t.PokemonID] = &tl
			}

			pokemonTypesByPokemonId[t.PokemonID].AddPokemonType(&t)
		}

		pokemonTypeList := make([]*model.PokemonTypeList, len(pokemonIds))
		for i, id := range pokemonIds {
			pokemonTypeList[i] = pokemonTypesByPokemonId[id]
			i++
		}

		err = rows.Err()
		if err != nil {
			errors := make([]error, len(pokemonIds))
			for i := range pokemonIds {
				errors[i] = fmt.Errorf("error after fetching types for pokemon in PokemonTypesByPokemonIdDataLoader: %w", err)
			}
			return pokemonTypeList, errors
		}

		return pokemonTypeList, nil
	}
}
