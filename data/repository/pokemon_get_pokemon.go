package repository

import (
	"bekapod/pkmn-team-graphql/data/model"
	"bekapod/pkmn-team-graphql/log"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/lib/pq"
)

var (
	ErrNoPokemon = errors.New("no pokemon found")
)

func (r Pokemon) GetPokemon(ctx context.Context) (*model.PokemonList, error) {
	pokemon := model.NewEmptyPokemonList()

	rows, err := r.db.QueryContext(
		ctx,
		`SELECT
			pokemon.*,
			array_agg(jsonb_build_object('type', jsonb_build_object('id', types.id, 'name', types.name, 'slug', types.slug), 'slot', pokemon_type.slot) ORDER BY pokemon_type.slot) as types
		FROM pokemon
			LEFT JOIN pokemon_type ON pokemon.id = pokemon_type.pokemon_id
			LEFT JOIN types ON pokemon_type.type_id = types.id
		GROUP BY pokemon.id
		ORDER BY pokedex_id, pokemon.slug ASC;`,
	)
	if err != nil {
		return &pokemon, fmt.Errorf("error fetching all pokemon in GetAllPokemon: %w", err)
	}

	defer rows.Close()
	for rows.Next() {
		var pkmn model.Pokemon
		err := rows.Scan(&pkmn.ID, &pkmn.PokedexId, &pkmn.Slug, &pkmn.Name, &pkmn.Sprite, &pkmn.HP, &pkmn.Attack, &pkmn.Defense, &pkmn.SpecialAttack, &pkmn.SpecialDefense, &pkmn.Speed, &pkmn.IsBaby, &pkmn.IsLegendary, &pkmn.IsMythical, &pkmn.Description, &pkmn.Color, &pkmn.Shape, &pkmn.Habitat, &pkmn.IsDefaultVariant, &pkmn.Genus, &pkmn.Height, &pkmn.Weight, pq.Array(&pkmn.Types.PokemonTypes))
		if err != nil {
			return &pokemon, fmt.Errorf("error scanning result in GetAllPokemon: %w", err)
		}
		pkmn.Types.Total = len(pkmn.Types.PokemonTypes)
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
		`SELECT
			pokemon.*,
			array_agg(jsonb_build_object('type', jsonb_build_object('id', types.id, 'name', types.name, 'slug', types.slug), 'slot', pokemon_type.slot) ORDER BY pokemon_type.slot) as types
		FROM pokemon
			LEFT JOIN pokemon_type ON pokemon.id = pokemon_type.pokemon_id
			LEFT JOIN types ON pokemon_type.type_id = types.id
		WHERE pokemon.id = $1
		GROUP BY pokemon.id;`,
		id,
	).Scan(&pkmn.ID, &pkmn.PokedexId, &pkmn.Slug, &pkmn.Name, &pkmn.Sprite, &pkmn.HP, &pkmn.Attack, &pkmn.Defense, &pkmn.SpecialAttack, &pkmn.SpecialDefense, &pkmn.Speed, &pkmn.IsBaby, &pkmn.IsLegendary, &pkmn.IsMythical, &pkmn.Description, &pkmn.Color, &pkmn.Shape, &pkmn.Habitat, &pkmn.IsDefaultVariant, &pkmn.Genus, &pkmn.Height, &pkmn.Weight, pq.Array(&pkmn.Types.PokemonTypes))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNoPokemon
		}
		return nil, fmt.Errorf("error scanning result in GetPokemonById %s: %w", id, err)
	}
	pkmn.Types.Total = len(pkmn.Types.PokemonTypes)
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

		query := `SELECT
			pokemon.*,
			array_agg(jsonb_build_object('type', jsonb_build_object('id', types.id, 'name', types.name, 'slug', types.slug), 'slot', pokemon_type.slot) ORDER BY pokemon_type.slot) as types
			pokemon_move.move_id
		FROM pokemon
			LEFT JOIN pokemon_type ON pokemon.id = pokemon_type.pokemon_id
			LEFT JOIN types ON pokemon_type.type_id = types.id
			LEFT JOIN pokemon_move ON pokemon.id = pokemon_move.pokemon_id
		WHERE pokemon_move.move_id IN (` + strings.Join(placeholders, ",") + `)
		GROUP BY pokemon.id
		ORDER BY pokedex_id, pokemon.slug ASC;`

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
			err := rows.Scan(&pkmn.ID, &pkmn.PokedexId, &pkmn.Slug, &pkmn.Name, &pkmn.Sprite, &pkmn.HP, &pkmn.Attack, &pkmn.Defense, &pkmn.SpecialAttack, &pkmn.SpecialDefense, &pkmn.Speed, &pkmn.IsBaby, &pkmn.IsLegendary, &pkmn.IsMythical, &pkmn.Description, &pkmn.Color, &pkmn.Shape, &pkmn.Habitat, &pkmn.IsDefaultVariant, &pkmn.Genus, &pkmn.Height, &pkmn.Weight, pq.Array(&pkmn.Types.PokemonTypes), &moveId)
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
			pkmn.Types.Total = len(pkmn.Types.PokemonTypes)

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

		query := `SELECT
			pokemon.*,
			array_agg(jsonb_build_object('type', jsonb_build_object('id', types.id, 'name', types.name, 'slug', types.slug), 'slot', pokemon_type.slot) ORDER BY pokemon_type.slot) as types
		FROM pokemon
			LEFT JOIN pokemon_type ON pokemon.id = pokemon_type.pokemon_id
			LEFT JOIN types ON pokemon_type.type_id = types.id
		WHERE pokemon_type.type_id IN (` + strings.Join(placeholders, ",") + `)
		GROUP BY pokemon.id
		ORDER BY pokedex_id, pokemon.slug ASC;`

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
			err := rows.Scan(&pkmn.ID, &pkmn.PokedexId, &pkmn.Slug, &pkmn.Name, &pkmn.Sprite, &pkmn.HP, &pkmn.Attack, &pkmn.Defense, &pkmn.SpecialAttack, &pkmn.SpecialDefense, &pkmn.Speed, &pkmn.IsBaby, &pkmn.IsLegendary, &pkmn.IsMythical, &pkmn.Description, &pkmn.Color, &pkmn.Shape, &pkmn.Habitat, &pkmn.IsDefaultVariant, &pkmn.Genus, &pkmn.Height, &pkmn.Weight, pq.Array(&pkmn.Types.PokemonTypes))
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
			pkmn.Types.Total = len(pkmn.Types.PokemonTypes)

			for _, t := range pkmn.Types.PokemonTypes {
				_, ok := pokemonByTypeId[t.Type.ID]
				if !ok {
					pl := model.NewEmptyPokemonList()
					pokemonByTypeId[t.Type.ID] = &pl
				}

				pokemonByTypeId[t.Type.ID].AddPokemon(&pkmn)
			}
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

		query := `SELECT
			pokemon.*,
			array_agg(jsonb_build_object('type', jsonb_build_object('id', types.id, 'name', types.name, 'slug', types.slug), 'slot', pokemon_type.slot) ORDER BY pokemon_type.slot) as types
			pokemon_ability.ability_id
		FROM pokemon
			LEFT JOIN pokemon_type ON pokemon.id = pokemon_type.pokemon_id
			LEFT JOIN types ON pokemon_type.type_id = types.id
			LEFT JOIN pokemon_ability ON pokemon.id = pokemon_ability.pokemon_id
		WHERE pokemon_ability.ability_id IN (` + strings.Join(placeholders, ",") + `)
		GROUP BY pokemon.id
		ORDER BY pokedex_id, pokemon.slug ASC;`

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
			err := rows.Scan(&pkmn.ID, &pkmn.PokedexId, &pkmn.Slug, &pkmn.Name, &pkmn.Sprite, &pkmn.HP, &pkmn.Attack, &pkmn.Defense, &pkmn.SpecialAttack, &pkmn.SpecialDefense, &pkmn.Speed, &pkmn.IsBaby, &pkmn.IsLegendary, &pkmn.IsMythical, &pkmn.Description, &pkmn.Color, &pkmn.Shape, &pkmn.Habitat, &pkmn.IsDefaultVariant, &pkmn.Genus, &pkmn.Height, &pkmn.Weight, pq.Array(&pkmn.Types.PokemonTypes), &abilityId)
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
			pkmn.Types.Total = len(pkmn.Types.PokemonTypes)

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
