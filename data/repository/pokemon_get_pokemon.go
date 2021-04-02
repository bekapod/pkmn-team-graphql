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
		"SELECT id, name, slug, pokedex_id, sprite, hp, attack, defense, special_attack, special_defense, speed, is_baby, is_legendary, is_mythical, description, color_enum, habitat_enum, shape_enum, height, weight, is_default_variant, genus FROM pokemon ORDER BY pokedex_id, slug ASC",
	)
	if err != nil {
		return &pokemon, fmt.Errorf("error fetching all pokemon in GetAllPokemon: %w", err)
	}

	defer rows.Close()
	for rows.Next() {
		var pkmn model.Pokemon
		var habitat sql.NullString
		err := rows.Scan(&pkmn.ID, &pkmn.Name, &pkmn.Slug, &pkmn.PokedexId, &pkmn.Sprite, &pkmn.HP, &pkmn.Attack, &pkmn.Defense, &pkmn.SpecialAttack, &pkmn.SpecialDefense, &pkmn.Speed, &pkmn.IsBaby, &pkmn.IsLegendary, &pkmn.IsMythical, &pkmn.Description, &pkmn.Color, &habitat, &pkmn.Shape, &pkmn.Height, &pkmn.Weight, &pkmn.IsDefaultVariant, &pkmn.Genus)
		if err != nil {
			return &pokemon, fmt.Errorf("error scanning result in GetAllPokemon: %w", err)
		}
		pkmn.Habitat = model.Habitat(habitat.String)
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
	var habitat sql.NullString

	err := r.db.QueryRowContext(
		ctx,
		"SELECT id, name, slug, pokedex_id, sprite, hp, attack, defense, special_attack, special_defense, speed, is_baby, is_legendary, is_mythical, description, color_enum, habitat_enum, shape_enum, height, weight, is_default_variant, genus FROM pokemon WHERE id = $1",
		id,
	).Scan(&pkmn.ID, &pkmn.Name, &pkmn.Slug, &pkmn.PokedexId, &pkmn.Sprite, &pkmn.HP, &pkmn.Attack, &pkmn.Defense, &pkmn.SpecialAttack, &pkmn.SpecialDefense, &pkmn.Speed, &pkmn.IsBaby, &pkmn.IsLegendary, &pkmn.IsMythical, &pkmn.Description, &pkmn.Color, &habitat, &pkmn.Shape, &pkmn.Height, &pkmn.Weight, &pkmn.IsDefaultVariant, &pkmn.Genus)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNoPokemon
		}
		return nil, fmt.Errorf("error scanning result in GetPokemonById %s: %w", id, err)
	}

	pkmn.Habitat = model.Habitat(habitat.String)
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

		query := "SELECT id, name, slug, pokedex_id, sprite, hp, attack, defense, special_attack, special_defense, speed, is_baby, is_legendary, is_mythical, description, color_enum, habitat_enum, shape_enum, height, weight, is_default_variant, genus, pokemon_move.move_id FROM pokemon LEFT JOIN pokemon_move ON pokemon.id = pokemon_move.pokemon_id WHERE pokemon_move.move_id IN (" + strings.Join(placeholders, ",") + ")"

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
			var habitat sql.NullString
			err := rows.Scan(&pkmn.ID, &pkmn.Name, &pkmn.Slug, &pkmn.PokedexId, &pkmn.Sprite, &pkmn.HP, &pkmn.Attack, &pkmn.Defense, &pkmn.SpecialAttack, &pkmn.SpecialDefense, &pkmn.Speed, &pkmn.IsBaby, &pkmn.IsLegendary, &pkmn.IsMythical, &pkmn.Description, &pkmn.Color, &habitat, &pkmn.Shape, &pkmn.Height, &pkmn.Weight, &pkmn.IsDefaultVariant, &pkmn.Genus, &moveId)
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

			pkmn.Habitat = model.Habitat(habitat.String)
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

		query := "SELECT id, name, slug, pokedex_id, sprite, hp, attack, defense, special_attack, special_defense, speed, is_baby, is_legendary, is_mythical, description, color_enum, habitat_enum, shape_enum, height, weight, is_default_variant, genus, pokemon_type.type_id FROM pokemon LEFT JOIN pokemon_type ON pokemon.id = pokemon_type.pokemon_id WHERE pokemon_type.type_id IN (" + strings.Join(placeholders, ",") + ")"

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
			var habitat sql.NullString
			err := rows.Scan(&pkmn.ID, &pkmn.Name, &pkmn.Slug, &pkmn.PokedexId, &pkmn.Sprite, &pkmn.HP, &pkmn.Attack, &pkmn.Defense, &pkmn.SpecialAttack, &pkmn.SpecialDefense, &pkmn.Speed, &pkmn.IsBaby, &pkmn.IsLegendary, &pkmn.IsMythical, &pkmn.Description, &pkmn.Color, &habitat, &pkmn.Shape, &pkmn.Height, &pkmn.Weight, &pkmn.IsDefaultVariant, &pkmn.Genus, &typeId)
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

			pkmn.Habitat = model.Habitat(habitat.String)
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

func (r Pokemon) PokemonByAbilityIdDataLoader(ctx context.Context) func(abilityIds []string) ([]*model.PokemonList, []error) {
	return func(abilityIds []string) ([]*model.PokemonList, []error) {
		pokemonByAbilityId := map[string]*model.PokemonList{}
		placeholders := make([]string, len(abilityIds))
		args := make([]interface{}, len(abilityIds))
		for i := 0; i < len(abilityIds); i++ {
			placeholders[i] = fmt.Sprintf("$%d", i+1)
			args[i] = abilityIds[i]
		}

		query := "SELECT id, name, slug, pokedex_id, sprite, hp, attack, defense, special_attack, special_defense, speed, is_baby, is_legendary, is_mythical, description, color_enum, habitat_enum, shape_enum, height, weight, is_default_variant, genus, pokemon_ability.ability_id FROM pokemon LEFT JOIN pokemon_ability ON pokemon.id = pokemon_ability.pokemon_id WHERE pokemon_ability.ability_id IN (" + strings.Join(placeholders, ",") + ")"

		log.Logger.WithField("args", args).Debug(query)
		rows, err := r.db.QueryContext(ctx,
			query,
			args...,
		)
		if err != nil {
			pokemonList := make([]*model.PokemonList, len(abilityIds))
			emptyPokemonList := model.NewEmptyPokemonList()
			errors := make([]error, len(abilityIds))
			for i := range abilityIds {
				pokemonList[i] = &emptyPokemonList
				errors[i] = fmt.Errorf("error fetching pokemon for ability in PokemonByAbilityIdDataLoader: %w", err)
			}
			return pokemonList, errors
		}

		defer rows.Close()
		for rows.Next() {
			var pkmn model.Pokemon
			var abilityId string
			var habitat sql.NullString
			err := rows.Scan(&pkmn.ID, &pkmn.Name, &pkmn.Slug, &pkmn.PokedexId, &pkmn.Sprite, &pkmn.HP, &pkmn.Attack, &pkmn.Defense, &pkmn.SpecialAttack, &pkmn.SpecialDefense, &pkmn.Speed, &pkmn.IsBaby, &pkmn.IsLegendary, &pkmn.IsMythical, &pkmn.Description, &pkmn.Color, &habitat, &pkmn.Shape, &pkmn.Height, &pkmn.Weight, &pkmn.IsDefaultVariant, &pkmn.Genus, &abilityId)
			if err != nil {
				pokemonList := make([]*model.PokemonList, len(abilityIds))
				emptyPokemonList := model.NewEmptyPokemonList()
				errors := make([]error, len(abilityIds))
				for i := range abilityIds {
					pokemonList[i] = &emptyPokemonList
					errors[i] = fmt.Errorf("error scanning result pokemon for ability in PokemonByAbilityIdDataLoader: %w", err)
				}
				return pokemonList, errors
			}

			pkmn.Habitat = model.Habitat(habitat.String)
			_, ok := pokemonByAbilityId[abilityId]
			if !ok {
				pl := model.NewEmptyPokemonList()
				pokemonByAbilityId[abilityId] = &pl
			}

			pokemonByAbilityId[abilityId].AddPokemon(&pkmn)
		}

		pokemonList := make([]*model.PokemonList, len(abilityIds))
		for i, id := range abilityIds {
			pokemonList[i] = pokemonByAbilityId[id]
			i++
		}

		err = rows.Err()
		if err != nil {
			errors := make([]error, len(abilityIds))
			for i := range abilityIds {
				errors[i] = fmt.Errorf("error after fetching pokemon for ability in PokemonByAbilityIdDataLoader: %w", err)
			}
			return pokemonList, errors
		}

		return pokemonList, nil
	}
}

func (r Pokemon) PokemonByPokemonIdDataLoader(ctx context.Context) func(pokemonIds []string) ([]*model.Pokemon, []error) {
	return func(pokemonIds []string) ([]*model.Pokemon, []error) {
		pokemonByPokemonId := map[string]*model.Pokemon{}
		placeholders := make([]string, len(pokemonIds))
		args := make([]interface{}, len(pokemonIds))
		for i := 0; i < len(pokemonIds); i++ {
			placeholders[i] = fmt.Sprintf("$%d", i+1)
			args[i] = pokemonIds[i]
		}

		query := "SELECT id, name, slug, pokedex_id, sprite, hp, attack, defense, special_attack, special_defense, speed, is_baby, is_legendary, is_mythical, description, color_enum, habitat_enum, shape_enum, height, weight, is_default_variant, genus FROM pokemon WHERE id IN (" + strings.Join(placeholders, ",") + ")"

		log.Logger.WithField("args", args).Debug(query)
		rows, err := r.db.QueryContext(ctx,
			query,
			args...,
		)
		if err != nil {
			pokemonList := make([]*model.Pokemon, len(pokemonIds))
			errors := make([]error, len(pokemonIds))
			for i := range pokemonIds {
				errors[i] = fmt.Errorf("error fetching pokemons for pokemon id in PokemonByPokemonIdDataLoader: %w", err)
			}
			return pokemonList, errors
		}

		defer rows.Close()
		for rows.Next() {
			var pkmn model.Pokemon
			var habitat sql.NullString
			err := rows.Scan(&pkmn.ID, &pkmn.Name, &pkmn.Slug, &pkmn.PokedexId, &pkmn.Sprite, &pkmn.HP, &pkmn.Attack, &pkmn.Defense, &pkmn.SpecialAttack, &pkmn.SpecialDefense, &pkmn.Speed, &pkmn.IsBaby, &pkmn.IsLegendary, &pkmn.IsMythical, &pkmn.Description, &pkmn.Color, &habitat, &pkmn.Shape, &pkmn.Height, &pkmn.Weight, &pkmn.IsDefaultVariant, &pkmn.Genus)
			if err != nil {
				pokemonList := make([]*model.Pokemon, len(pokemonIds))
				errors := make([]error, len(pokemonIds))
				for i := range pokemonIds {
					errors[i] = fmt.Errorf("error scanning result in PokemonByPokemonIdDataLoader: %w", err)
				}
				return pokemonList, errors
			}
			pkmn.Habitat = model.Habitat(habitat.String)
			pokemonByPokemonId[pkmn.ID] = &pkmn
		}

		pokemon := make([]*model.Pokemon, len(pokemonIds))
		for i, id := range pokemonIds {
			pokemon[i] = pokemonByPokemonId[id]
			i++
		}

		err = rows.Err()
		if err != nil {
			errors := make([]error, len(pokemonIds))
			for i := range pokemonIds {
				errors[i] = fmt.Errorf("error after fetching pokemon for pokemon id in PokemonByPokemonIdDataLoader: %w", err)
			}
			return pokemon, errors
		}

		return pokemon, nil
	}
}
