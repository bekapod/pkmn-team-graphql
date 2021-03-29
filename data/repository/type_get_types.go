package repository

import (
	"bekapod/pkmn-team-graphql/data/model"
	"bekapod/pkmn-team-graphql/log"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

var (
	ErrNoTypes = errors.New("no types found")
	ErrNoType  = errors.New("no type found")
)

func (r Type) GetTypes(ctx context.Context) (*model.TypeList, error) {
	types := model.NewEmptyTypeList()

	rows, err := r.db.QueryContext(
		ctx,
		"SELECT id, name, slug FROM types ORDER BY slug ASC",
	)
	if err != nil {
		return &types, fmt.Errorf("error fetching all types: %w", err)
	}

	defer rows.Close()
	for rows.Next() {
		var t model.Type
		err := rows.Scan(&t.ID, &t.Name, &t.Slug)
		if err != nil {
			if err == sql.ErrNoRows {
				return &types, ErrNoTypes
			}
			return &types, fmt.Errorf("error scanning result in GetAllTypes: %w", err)
		}
		types.AddType(&t)
	}

	return &types, nil
}

func (r Type) GetTypeById(ctx context.Context, id string) (*model.Type, error) {
	t := model.Type{}

	err := r.db.QueryRowContext(
		ctx,
		"SELECT id, name, slug FROM types WHERE id = $1",
		id,
	).Scan(&t.ID, &t.Name, &t.Slug)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNoType
		}
		return nil, fmt.Errorf("error scanning result in GetTypeById %s: %w", id, err)
	}

	return &t, nil
}

func (r Type) TypesByPokemonIdDataLoader(ctx context.Context) func(pokemonIds []string) ([]*model.TypeList, []error) {
	return func(pokemonIds []string) ([]*model.TypeList, []error) {
		typesByPokemonId := map[string]*model.TypeList{}
		placeholders := make([]string, len(pokemonIds))
		args := make([]interface{}, len(pokemonIds))
		for i := 0; i < len(pokemonIds); i++ {
			placeholders[i] = fmt.Sprintf("$%d", i+1)
			args[i] = pokemonIds[i]
		}

		query := "SELECT id, name, slug, pokemon_type.pokemon_id FROM types LEFT JOIN pokemon_type ON types.id = pokemon_type.type_id WHERE pokemon_type.pokemon_id IN (" + strings.Join(placeholders, ",") + ")"

		log.Logger.WithField("args", args).Debug(query)
		rows, err := r.db.QueryContext(ctx,
			query,
			args...,
		)
		if err != nil {
			panic(fmt.Errorf("error fetching types for pokemon: %w", err))
		}

		defer rows.Close()
		for rows.Next() {
			var t model.Type
			var pokemonId string
			err := rows.Scan(&t.ID, &t.Name, &t.Slug, &pokemonId)
			if err != nil {
				panic(fmt.Errorf("error scanning result in TypesByPokemonId: %w", err))
			}

			_, ok := typesByPokemonId[pokemonId]
			if !ok {
				tl := model.NewEmptyTypeList()
				typesByPokemonId[pokemonId] = &tl
			}

			typesByPokemonId[pokemonId].AddType(&t)
		}

		typeList := make([]*model.TypeList, len(pokemonIds))
		for i, id := range pokemonIds {
			typeList[i] = typesByPokemonId[id]
			i++
		}

		return typeList, nil
	}
}

func (r Type) TypesByTypeIdDataLoader(ctx context.Context) func(typeIds []string) ([]*model.Type, []error) {
	return func(typeIds []string) ([]*model.Type, []error) {
		typesByTypeId := map[string]*model.Type{}
		placeholders := make([]string, len(typeIds))
		args := make([]interface{}, len(typeIds))
		for i := 0; i < len(typeIds); i++ {
			placeholders[i] = fmt.Sprintf("$%d", i+1)
			args[i] = typeIds[i]
		}

		query := "SELECT id, name, slug FROM types WHERE id IN (" + strings.Join(placeholders, ",") + ")"

		log.Logger.WithField("args", args).Debug(query)
		rows, err := r.db.QueryContext(ctx,
			query,
			args...,
		)
		if err != nil {
			panic(fmt.Errorf("error fetching types: %w", err))
		}

		defer rows.Close()
		for rows.Next() {
			var t model.Type
			err := rows.Scan(&t.ID, &t.Name, &t.Slug)
			if err != nil {
				panic(fmt.Errorf("error scanning result in TypeByTypeId: %w", err))
			}

			typesByTypeId[t.ID] = &t
		}

		types := make([]*model.Type, len(typeIds))
		for i, id := range typeIds {
			types[i] = typesByTypeId[id]
			i++
		}

		return types, nil
	}
}
