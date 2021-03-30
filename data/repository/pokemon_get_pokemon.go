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
	ErrNoPokemon = errors.New("no pokemon found")
)

func (r Pokemon) GetPokemon(ctx context.Context) (*model.PokemonList, error) {
	pokemon := model.NewEmptyPokemonList()

	rows, err := r.db.QueryContext(
		ctx,
		"SELECT id, name, slug, pokedex_id, sprite, hp, attack, defense, special_attack, special_defense, speed, is_baby, is_legendary, is_mythical, description FROM pokemon ORDER BY pokedex_id, slug ASC",
	)
	if err != nil {
		return &pokemon, fmt.Errorf("error fetching all pokemon in GetAllPokemon: %w", err)
	}

	defer rows.Close()
	for rows.Next() {
		var pkmn model.Pokemon
		err := rows.Scan(&pkmn.ID, &pkmn.Name, &pkmn.Slug, &pkmn.PokedexId, &pkmn.Sprite, &pkmn.HP, &pkmn.Attack, &pkmn.Defense, &pkmn.SpecialAttack, &pkmn.SpecialDefense, &pkmn.Speed, &pkmn.IsBaby, &pkmn.IsLegendary, &pkmn.IsMythical, &pkmn.Description)
		if err != nil {
			return &pokemon, fmt.Errorf("error scanning result in GetAllPokemon: %w", err)
		}
		pokemon.AddPokemon(&pkmn)
	}

	err = rows.Err()
	if err != nil {
		return &pokemon, fmt.Errorf("error after fetching all pokemon in GetAllPokemon: %w", err)
	}

	return &pokemon, nil
}

func (r Pokemon) GetPokemonById(ctx context.Context, id string) (*model.Pokemon, error) {
	pkmn := model.Pokemon{}

	err := r.db.QueryRowContext(
		ctx,
		"SELECT id, name, slug, pokedex_id, sprite, hp, attack, defense, special_attack, special_defense, speed, is_baby, is_legendary, is_mythical, description FROM pokemon WHERE id = $1",
		id,
	).Scan(&pkmn.ID, &pkmn.Name, &pkmn.Slug, &pkmn.PokedexId, &pkmn.Sprite, &pkmn.HP, &pkmn.Attack, &pkmn.Defense, &pkmn.SpecialAttack, &pkmn.SpecialDefense, &pkmn.Speed, &pkmn.IsBaby, &pkmn.IsLegendary, &pkmn.IsMythical, &pkmn.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNoPokemon
		}
		return nil, fmt.Errorf("error scanning result in GetPokemonById %s: %w", id, err)
	}

	return &pkmn, nil
}

func (r Pokemon) PokemonByMoveIdDataLoader(ctx context.Context) func(moveIds []string) ([]*model.PokemonList, []error) {
	return func(moveIds []string) ([]*model.PokemonList, []error) {
		pokemonByMoveId := map[string]*model.PokemonList{}
		placeholders := make([]string, len(moveIds))
		args := make([]interface{}, len(moveIds))
		for i := 0; i < len(moveIds); i++ {
			placeholders[i] = fmt.Sprintf("$%d", i+1)
			args[i] = moveIds[i]
		}

		query := "SELECT id, name, slug, pokedex_id, sprite, hp, attack, defense, special_attack, special_defense, speed, is_baby, is_legendary, is_mythical, description, pokemon_move.move_id FROM pokemon LEFT JOIN pokemon_move ON pokemon.id = pokemon_move.pokemon_id WHERE pokemon_move.move_id IN (" + strings.Join(placeholders, ",") + ")"

		log.Logger.WithField("args", args).Debug(query)
		rows, err := r.db.QueryContext(ctx,
			query,
			args...,
		)
		if err != nil {
			pokemonList := make([]*model.PokemonList, len(moveIds))
			emptyPokemonList := model.NewEmptyPokemonList()
			errors := make([]error, len(moveIds))
			for i := range moveIds {
				pokemonList[i] = &emptyPokemonList
				errors[i] = fmt.Errorf("error fetching pokemon for move in PokemonByMoveIdDataLoader: %w", err)
			}
			return pokemonList, errors
		}

		defer rows.Close()
		for rows.Next() {
			var pkmn model.Pokemon
			var moveId string
			err := rows.Scan(&pkmn.ID, &pkmn.Name, &pkmn.Slug, &pkmn.PokedexId, &pkmn.Sprite, &pkmn.HP, &pkmn.Attack, &pkmn.Defense, &pkmn.SpecialAttack, &pkmn.SpecialDefense, &pkmn.Speed, &pkmn.IsBaby, &pkmn.IsLegendary, &pkmn.IsMythical, &pkmn.Description, &moveId)
			if err != nil {
				pokemonList := make([]*model.PokemonList, len(moveIds))
				emptyPokemonList := model.NewEmptyPokemonList()
				errors := make([]error, len(moveIds))
				for i := range moveIds {
					pokemonList[i] = &emptyPokemonList
					errors[i] = fmt.Errorf("error scanning result moves for type in PokemonByMoveIdDataLoader: %w", err)
				}
				return pokemonList, errors
			}

			_, ok := pokemonByMoveId[moveId]
			if !ok {
				pl := model.NewEmptyPokemonList()
				pokemonByMoveId[moveId] = &pl
			}

			pokemonByMoveId[moveId].AddPokemon(&pkmn)
		}

		pokemonList := make([]*model.PokemonList, len(moveIds))
		for i, id := range moveIds {
			pokemonList[i] = pokemonByMoveId[id]
			i++
		}

		err = rows.Err()
		if err != nil {
			errors := make([]error, len(moveIds))
			for i := range moveIds {
				errors[i] = fmt.Errorf("error after fetching pokemon for move in PokemonByMoveIdDataLoader: %w", err)
			}
			return pokemonList, errors
		}

		return pokemonList, nil
	}
}

func (r Pokemon) PokemonByTypeIdDataLoader(ctx context.Context) func(typeIds []string) ([]*model.PokemonList, []error) {
	return func(typeIds []string) ([]*model.PokemonList, []error) {
		pokemonByTypeId := map[string]*model.PokemonList{}
		placeholders := make([]string, len(typeIds))
		args := make([]interface{}, len(typeIds))
		for i := 0; i < len(typeIds); i++ {
			placeholders[i] = fmt.Sprintf("$%d", i+1)
			args[i] = typeIds[i]
		}

		query := "SELECT id, name, slug, pokedex_id, sprite, hp, attack, defense, special_attack, special_defense, speed, is_baby, is_legendary, is_mythical, description, pokemon_type.type_id FROM pokemon LEFT JOIN pokemon_type ON pokemon.id = pokemon_type.pokemon_id WHERE pokemon_type.type_id IN (" + strings.Join(placeholders, ",") + ")"

		log.Logger.WithField("args", args).Debug(query)
		rows, err := r.db.QueryContext(ctx,
			query,
			args...,
		)
		if err != nil {
			pokemonList := make([]*model.PokemonList, len(typeIds))
			emptyPokemonList := model.NewEmptyPokemonList()
			errors := make([]error, len(typeIds))
			for i := range typeIds {
				pokemonList[i] = &emptyPokemonList
				errors[i] = fmt.Errorf("error fetching pokemon for type in PokemonByTypeIdDataLoader: %w", err)
			}
			return pokemonList, errors
		}

		defer rows.Close()
		for rows.Next() {
			var pkmn model.Pokemon
			var typeId string
			err := rows.Scan(&pkmn.ID, &pkmn.Name, &pkmn.Slug, &pkmn.PokedexId, &pkmn.Sprite, &pkmn.HP, &pkmn.Attack, &pkmn.Defense, &pkmn.SpecialAttack, &pkmn.SpecialDefense, &pkmn.Speed, &pkmn.IsBaby, &pkmn.IsLegendary, &pkmn.IsMythical, &pkmn.Description, &typeId)
			if err != nil {
				pokemonList := make([]*model.PokemonList, len(typeIds))
				emptyPokemonList := model.NewEmptyPokemonList()
				errors := make([]error, len(typeIds))
				for i := range typeIds {
					pokemonList[i] = &emptyPokemonList
					errors[i] = fmt.Errorf("error scanning result pokemon for type in PokemonByTypeIdDataLoader: %w", err)
				}
				return pokemonList, errors
			}

			_, ok := pokemonByTypeId[typeId]
			if !ok {
				pl := model.NewEmptyPokemonList()
				pokemonByTypeId[typeId] = &pl
			}

			pokemonByTypeId[typeId].AddPokemon(&pkmn)
		}

		pokemonList := make([]*model.PokemonList, len(typeIds))
		for i, id := range typeIds {
			pokemonList[i] = pokemonByTypeId[id]
			i++
		}

		err = rows.Err()
		if err != nil {
			errors := make([]error, len(typeIds))
			for i := range typeIds {
				errors[i] = fmt.Errorf("error after fetching pokemon for type in PokemonByTypeIdDataLoader: %w", err)
			}
			return pokemonList, errors
		}

		return pokemonList, nil
	}
}
