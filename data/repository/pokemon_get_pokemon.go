package repository

import (
	"bekapod/pkmn-team-graphql/data/model"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/lib/pq"
)

var (
	ErrNoPokemon   = errors.New("no pokemon found")
	PokemonColumns = `
		pokemon.id, pokemon.pokedex_id, pokemon.slug, pokemon.name, pokemon.sprite, pokemon.hp, pokemon.attack, pokemon.defense, pokemon.special_attack, pokemon.special_defense, pokemon.speed, pokemon.is_baby, pokemon.is_legendary, pokemon.is_mythical, pokemon.description, pokemon.color_enum, pokemon.shape_enum, pokemon.habitat_enum, pokemon.is_default_variant, pokemon.genus, pokemon.height, pokemon.weight,
		array_agg(
			DISTINCT jsonb_build_object(
				'type', jsonb_build_object(
					'id', types.id,
					'name', types.name,
					'slug', types.slug,
					'noDamageTo', jsonb_build_object('types', ` + fmt.Sprintf(TemplatedTypeRelation, "no-damage-to") + `),
					'halfDamageTo', jsonb_build_object('types', ` + fmt.Sprintf(TemplatedTypeRelation, "half-damage-to") + `),
					'doubleDamageTo', jsonb_build_object('types', ` + fmt.Sprintf(TemplatedTypeRelation, "double-damage-to") + `),
					'noDamageFrom', jsonb_build_object('types', ` + fmt.Sprintf(TemplatedTypeRelation, "no-damage-from") + `),
					'halfDamageFrom', jsonb_build_object('types', ` + fmt.Sprintf(TemplatedTypeRelation, "half-damage-from") + `),
					'doubleDamageFrom', jsonb_build_object('types', ` + fmt.Sprintf(TemplatedTypeRelation, "double-damage-from") + `)
				),
				'slot', pokemon_type.slot
			)
		) as types,
		array_agg(
			DISTINCT jsonb_build_object('id', egg_groups.id, 'name', egg_groups.name, 'slug', egg_groups.slug)
		) as egg_groups,
		array_agg(
			DISTINCT jsonb_build_object(
				'ability', jsonb_build_object('id', abilities.id, 'name', abilities.name, 'slug', abilities.slug, 'effect', abilities.effect), 
				'slot', pokemon_ability.slot,
				'isHidden', pokemon_ability.is_hidden
			)
		) as abilities
	`
	PokemonJoins = `
		LEFT JOIN pokemon_type ON pokemon.id = pokemon_type.pokemon_id
		LEFT JOIN types ON pokemon_type.type_id = types.id
		LEFT JOIN pokemon_ability ON pokemon.id = pokemon_ability.pokemon_id
		LEFT JOIN abilities on pokemon_ability.ability_id = abilities.id
		LEFT JOIN pokemon_egg_group ON pokemon.id = pokemon_egg_group.pokemon_id
		LEFT JOIN egg_groups on pokemon_egg_group.egg_group_id = egg_groups.id
	`
)

func (r Pokemon) GetPokemon(ctx context.Context) (*model.PokemonList, error) {
	pokemon := model.NewEmptyPokemonList()

	rows, err := r.db.QueryContext(
		ctx,
		`SELECT `+PokemonColumns+`	
		FROM pokemon
			`+PokemonJoins+`
		GROUP BY pokemon.id
		ORDER BY pokedex_id, pokemon.slug ASC;`,
	)
	if err != nil {
		return &pokemon, fmt.Errorf("error fetching all pokemon in GetAllPokemon: %w", err)
	}

	defer rows.Close()
	for rows.Next() {
		var pkmn model.Pokemon
		err := rows.Scan(&pkmn.ID, &pkmn.PokedexId, &pkmn.Slug, &pkmn.Name, &pkmn.Sprite, &pkmn.HP, &pkmn.Attack, &pkmn.Defense, &pkmn.SpecialAttack, &pkmn.SpecialDefense, &pkmn.Speed, &pkmn.IsBaby, &pkmn.IsLegendary, &pkmn.IsMythical, &pkmn.Description, &pkmn.Color, &pkmn.Shape, &pkmn.Habitat, &pkmn.IsDefaultVariant, &pkmn.Genus, &pkmn.Height, &pkmn.Weight, pq.Array(&pkmn.Types.PokemonTypes), pq.Array(&pkmn.EggGroups.EggGroups), pq.Array(&pkmn.Abilities.PokemonAbilities))
		if err != nil {
			return &pokemon, fmt.Errorf("error scanning result in GetAllPokemon: %w", err)
		}
		pkmn.Types.Total = len(pkmn.Types.PokemonTypes)
		pkmn.EggGroups.Total = len(pkmn.EggGroups.EggGroups)
		pkmn.Abilities.Total = len(pkmn.Abilities.PokemonAbilities)

		for i := range pkmn.Types.PokemonTypes {
			pkmn.Types.PokemonTypes[i].Type.NoDamageTo.Total = len(pkmn.Types.PokemonTypes[i].Type.NoDamageTo.Types)
			pkmn.Types.PokemonTypes[i].Type.HalfDamageTo.Total = len(pkmn.Types.PokemonTypes[i].Type.HalfDamageTo.Types)
			pkmn.Types.PokemonTypes[i].Type.DoubleDamageTo.Total = len(pkmn.Types.PokemonTypes[i].Type.DoubleDamageTo.Types)
			pkmn.Types.PokemonTypes[i].Type.NoDamageFrom.Total = len(pkmn.Types.PokemonTypes[i].Type.NoDamageFrom.Types)
			pkmn.Types.PokemonTypes[i].Type.HalfDamageFrom.Total = len(pkmn.Types.PokemonTypes[i].Type.HalfDamageFrom.Types)
			pkmn.Types.PokemonTypes[i].Type.DoubleDamageFrom.Total = len(pkmn.Types.PokemonTypes[i].Type.DoubleDamageFrom.Types)
		}

		pokemon.AddPokemon(pkmn)
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
		`SELECT `+PokemonColumns+`
		FROM pokemon
			`+PokemonJoins+`
		WHERE pokemon.id = $1
		GROUP BY pokemon.id;`,
		id,
	).Scan(&pkmn.ID, &pkmn.PokedexId, &pkmn.Slug, &pkmn.Name, &pkmn.Sprite, &pkmn.HP, &pkmn.Attack, &pkmn.Defense, &pkmn.SpecialAttack, &pkmn.SpecialDefense, &pkmn.Speed, &pkmn.IsBaby, &pkmn.IsLegendary, &pkmn.IsMythical, &pkmn.Description, &pkmn.Color, &pkmn.Shape, &pkmn.Habitat, &pkmn.IsDefaultVariant, &pkmn.Genus, &pkmn.Height, &pkmn.Weight, pq.Array(&pkmn.Types.PokemonTypes), pq.Array(&pkmn.EggGroups.EggGroups), pq.Array(&pkmn.Abilities.PokemonAbilities))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNoPokemon
		}
		return nil, fmt.Errorf("error scanning result in GetPokemonById %s: %w", id, err)
	}
	pkmn.Types.Total = len(pkmn.Types.PokemonTypes)
	pkmn.EggGroups.Total = len(pkmn.EggGroups.EggGroups)
	pkmn.Abilities.Total = len(pkmn.Abilities.PokemonAbilities)

	for i := range pkmn.Types.PokemonTypes {
		pkmn.Types.PokemonTypes[i].Type.NoDamageTo.Total = len(pkmn.Types.PokemonTypes[i].Type.NoDamageTo.Types)
		pkmn.Types.PokemonTypes[i].Type.HalfDamageTo.Total = len(pkmn.Types.PokemonTypes[i].Type.HalfDamageTo.Types)
		pkmn.Types.PokemonTypes[i].Type.DoubleDamageTo.Total = len(pkmn.Types.PokemonTypes[i].Type.DoubleDamageTo.Types)
		pkmn.Types.PokemonTypes[i].Type.NoDamageFrom.Total = len(pkmn.Types.PokemonTypes[i].Type.NoDamageFrom.Types)
		pkmn.Types.PokemonTypes[i].Type.HalfDamageFrom.Total = len(pkmn.Types.PokemonTypes[i].Type.HalfDamageFrom.Types)
		pkmn.Types.PokemonTypes[i].Type.DoubleDamageFrom.Total = len(pkmn.Types.PokemonTypes[i].Type.DoubleDamageFrom.Types)
	}

	return &pkmn, nil
}

func (r Pokemon) PokemonByIdDataLoader(ctx context.Context) func(pokemonIds []string) ([]*model.Pokemon, []error) {
	return func(pokemonIds []string) ([]*model.Pokemon, []error) {
		pokemonById := map[string]*model.Pokemon{}
		placeholders := make([]string, len(pokemonIds))
		args := make([]interface{}, len(pokemonIds))
		for i := 0; i < len(pokemonIds); i++ {
			placeholders[i] = fmt.Sprintf("$%d", i+1)
			args[i] = pokemonIds[i]
		}

		query := `SELECT ` + PokemonColumns + `
		FROM pokemon
			` + PokemonJoins + `
		WHERE pokemon.id IN (` + strings.Join(placeholders, ",") + `)
		GROUP BY pokemon.id
		ORDER BY pokedex_id, pokemon.slug ASC;`

		rows, err := r.db.QueryContext(ctx,
			query,
			args...,
		)
		if err != nil {
			pokemon := make([]*model.Pokemon, len(pokemonIds))
			errors := make([]error, len(pokemonIds))
			for i := range pokemonIds {
				errors[i] = fmt.Errorf("error fetching pokemon by id in PokemonByIdDataLoader: %w", err)
			}
			return pokemon, errors
		}

		defer rows.Close()
		for rows.Next() {
			var pkmn model.Pokemon
			err := rows.Scan(&pkmn.ID, &pkmn.PokedexId, &pkmn.Slug, &pkmn.Name, &pkmn.Sprite, &pkmn.HP, &pkmn.Attack, &pkmn.Defense, &pkmn.SpecialAttack, &pkmn.SpecialDefense, &pkmn.Speed, &pkmn.IsBaby, &pkmn.IsLegendary, &pkmn.IsMythical, &pkmn.Description, &pkmn.Color, &pkmn.Shape, &pkmn.Habitat, &pkmn.IsDefaultVariant, &pkmn.Genus, &pkmn.Height, &pkmn.Weight, pq.Array(&pkmn.Types.PokemonTypes), pq.Array(&pkmn.EggGroups.EggGroups), pq.Array(&pkmn.Abilities.PokemonAbilities))
			if err != nil {
				pokemon := make([]*model.Pokemon, len(pokemonIds))
				errors := make([]error, len(pokemonIds))
				for i := range pokemonIds {
					errors[i] = fmt.Errorf("error scanning result pokemon by id in PokemonByIdDataLoader: %w", err)
				}
				return pokemon, errors
			}
			pkmn.Types.Total = len(pkmn.Types.PokemonTypes)
			pkmn.EggGroups.Total = len(pkmn.EggGroups.EggGroups)
			pkmn.Abilities.Total = len(pkmn.Abilities.PokemonAbilities)

			for i := range pkmn.Types.PokemonTypes {
				pkmn.Types.PokemonTypes[i].Type.NoDamageTo.Total = len(pkmn.Types.PokemonTypes[i].Type.NoDamageTo.Types)
				pkmn.Types.PokemonTypes[i].Type.HalfDamageTo.Total = len(pkmn.Types.PokemonTypes[i].Type.HalfDamageTo.Types)
				pkmn.Types.PokemonTypes[i].Type.DoubleDamageTo.Total = len(pkmn.Types.PokemonTypes[i].Type.DoubleDamageTo.Types)
				pkmn.Types.PokemonTypes[i].Type.NoDamageFrom.Total = len(pkmn.Types.PokemonTypes[i].Type.NoDamageFrom.Types)
				pkmn.Types.PokemonTypes[i].Type.HalfDamageFrom.Total = len(pkmn.Types.PokemonTypes[i].Type.HalfDamageFrom.Types)
				pkmn.Types.PokemonTypes[i].Type.DoubleDamageFrom.Total = len(pkmn.Types.PokemonTypes[i].Type.DoubleDamageFrom.Types)
			}

			pokemonById[pkmn.ID] = &pkmn
		}

		pokemon := make([]*model.Pokemon, len(pokemonIds))
		for i, id := range pokemonIds {
			pokemon[i] = pokemonById[id]
			i++
		}

		err = rows.Err()
		if err != nil {
			errors := make([]error, len(pokemonIds))
			for i := range pokemonIds {
				errors[i] = fmt.Errorf("error after fetching pokemon by id in PokemonByIdDataLoader: %w", err)
			}
			return pokemon, errors
		}

		return pokemon, nil
	}
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

		query := `SELECT ` + PokemonColumns + `, pokemon_move.move_id
		FROM pokemon
			` + PokemonJoins + `
			LEFT JOIN pokemon_move ON pokemon.id = pokemon_move.pokemon_id
		WHERE pokemon_move.move_id IN (` + strings.Join(placeholders, ",") + `)
		GROUP BY pokemon.id, pokemon_move.move_id
		ORDER BY pokedex_id, pokemon.slug ASC;`

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
			err := rows.Scan(&pkmn.ID, &pkmn.PokedexId, &pkmn.Slug, &pkmn.Name, &pkmn.Sprite, &pkmn.HP, &pkmn.Attack, &pkmn.Defense, &pkmn.SpecialAttack, &pkmn.SpecialDefense, &pkmn.Speed, &pkmn.IsBaby, &pkmn.IsLegendary, &pkmn.IsMythical, &pkmn.Description, &pkmn.Color, &pkmn.Shape, &pkmn.Habitat, &pkmn.IsDefaultVariant, &pkmn.Genus, &pkmn.Height, &pkmn.Weight, pq.Array(&pkmn.Types.PokemonTypes), pq.Array(&pkmn.EggGroups.EggGroups), pq.Array(&pkmn.Abilities.PokemonAbilities), &moveId)
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
			pkmn.EggGroups.Total = len(pkmn.EggGroups.EggGroups)
			pkmn.Abilities.Total = len(pkmn.Abilities.PokemonAbilities)

			for i := range pkmn.Types.PokemonTypes {
				pkmn.Types.PokemonTypes[i].Type.NoDamageTo.Total = len(pkmn.Types.PokemonTypes[i].Type.NoDamageTo.Types)
				pkmn.Types.PokemonTypes[i].Type.HalfDamageTo.Total = len(pkmn.Types.PokemonTypes[i].Type.HalfDamageTo.Types)
				pkmn.Types.PokemonTypes[i].Type.DoubleDamageTo.Total = len(pkmn.Types.PokemonTypes[i].Type.DoubleDamageTo.Types)
				pkmn.Types.PokemonTypes[i].Type.NoDamageFrom.Total = len(pkmn.Types.PokemonTypes[i].Type.NoDamageFrom.Types)
				pkmn.Types.PokemonTypes[i].Type.HalfDamageFrom.Total = len(pkmn.Types.PokemonTypes[i].Type.HalfDamageFrom.Types)
				pkmn.Types.PokemonTypes[i].Type.DoubleDamageFrom.Total = len(pkmn.Types.PokemonTypes[i].Type.DoubleDamageFrom.Types)
			}

			_, ok := pokemonByMoveId[moveId]
			if !ok {
				pl := model.NewEmptyPokemonList()
				pokemonByMoveId[moveId] = &pl
			}

			pokemonByMoveId[moveId].AddPokemon(pkmn)
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

		query := `SELECT ` + PokemonColumns + `
		FROM pokemon
			` + PokemonJoins + `
		WHERE pokemon_type.type_id IN (` + strings.Join(placeholders, ",") + `)
		GROUP BY pokemon.id
		ORDER BY pokedex_id, pokemon.slug ASC;`

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
			err := rows.Scan(&pkmn.ID, &pkmn.PokedexId, &pkmn.Slug, &pkmn.Name, &pkmn.Sprite, &pkmn.HP, &pkmn.Attack, &pkmn.Defense, &pkmn.SpecialAttack, &pkmn.SpecialDefense, &pkmn.Speed, &pkmn.IsBaby, &pkmn.IsLegendary, &pkmn.IsMythical, &pkmn.Description, &pkmn.Color, &pkmn.Shape, &pkmn.Habitat, &pkmn.IsDefaultVariant, &pkmn.Genus, &pkmn.Height, &pkmn.Weight, pq.Array(&pkmn.Types.PokemonTypes), pq.Array(&pkmn.EggGroups.EggGroups), pq.Array(&pkmn.Abilities.PokemonAbilities))
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
			pkmn.EggGroups.Total = len(pkmn.EggGroups.EggGroups)
			pkmn.Abilities.Total = len(pkmn.Abilities.PokemonAbilities)

			for i := range pkmn.Types.PokemonTypes {
				pkmn.Types.PokemonTypes[i].Type.NoDamageTo.Total = len(pkmn.Types.PokemonTypes[i].Type.NoDamageTo.Types)
				pkmn.Types.PokemonTypes[i].Type.HalfDamageTo.Total = len(pkmn.Types.PokemonTypes[i].Type.HalfDamageTo.Types)
				pkmn.Types.PokemonTypes[i].Type.DoubleDamageTo.Total = len(pkmn.Types.PokemonTypes[i].Type.DoubleDamageTo.Types)
				pkmn.Types.PokemonTypes[i].Type.NoDamageFrom.Total = len(pkmn.Types.PokemonTypes[i].Type.NoDamageFrom.Types)
				pkmn.Types.PokemonTypes[i].Type.HalfDamageFrom.Total = len(pkmn.Types.PokemonTypes[i].Type.HalfDamageFrom.Types)
				pkmn.Types.PokemonTypes[i].Type.DoubleDamageFrom.Total = len(pkmn.Types.PokemonTypes[i].Type.DoubleDamageFrom.Types)
			}

			for _, t := range pkmn.Types.PokemonTypes {
				_, ok := pokemonByTypeId[t.Type.ID]
				if !ok {
					pl := model.NewEmptyPokemonList()
					pokemonByTypeId[t.Type.ID] = &pl
				}

				pokemonByTypeId[t.Type.ID].AddPokemon(pkmn)
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

		query := `SELECT ` + PokemonColumns + `
		FROM pokemon
			` + PokemonJoins + `
		WHERE pokemon_ability.ability_id IN (` + strings.Join(placeholders, ",") + `)
		GROUP BY pokemon.id
		ORDER BY pokedex_id, pokemon.slug ASC;`

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
			err := rows.Scan(&pkmn.ID, &pkmn.PokedexId, &pkmn.Slug, &pkmn.Name, &pkmn.Sprite, &pkmn.HP, &pkmn.Attack, &pkmn.Defense, &pkmn.SpecialAttack, &pkmn.SpecialDefense, &pkmn.Speed, &pkmn.IsBaby, &pkmn.IsLegendary, &pkmn.IsMythical, &pkmn.Description, &pkmn.Color, &pkmn.Shape, &pkmn.Habitat, &pkmn.IsDefaultVariant, &pkmn.Genus, &pkmn.Height, &pkmn.Weight, pq.Array(&pkmn.Types.PokemonTypes), pq.Array(&pkmn.EggGroups.EggGroups), pq.Array(&pkmn.Abilities.PokemonAbilities))
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
			pkmn.EggGroups.Total = len(pkmn.EggGroups.EggGroups)
			pkmn.Abilities.Total = len(pkmn.Abilities.PokemonAbilities)

			for i := range pkmn.Types.PokemonTypes {
				pkmn.Types.PokemonTypes[i].Type.NoDamageTo.Total = len(pkmn.Types.PokemonTypes[i].Type.NoDamageTo.Types)
				pkmn.Types.PokemonTypes[i].Type.HalfDamageTo.Total = len(pkmn.Types.PokemonTypes[i].Type.HalfDamageTo.Types)
				pkmn.Types.PokemonTypes[i].Type.DoubleDamageTo.Total = len(pkmn.Types.PokemonTypes[i].Type.DoubleDamageTo.Types)
				pkmn.Types.PokemonTypes[i].Type.NoDamageFrom.Total = len(pkmn.Types.PokemonTypes[i].Type.NoDamageFrom.Types)
				pkmn.Types.PokemonTypes[i].Type.HalfDamageFrom.Total = len(pkmn.Types.PokemonTypes[i].Type.HalfDamageFrom.Types)
				pkmn.Types.PokemonTypes[i].Type.DoubleDamageFrom.Total = len(pkmn.Types.PokemonTypes[i].Type.DoubleDamageFrom.Types)
			}

			for _, a := range pkmn.Abilities.PokemonAbilities {
				_, ok := pokemonByAbilityId[a.Ability.ID]
				if !ok {
					pl := model.NewEmptyPokemonList()
					pokemonByAbilityId[a.Ability.ID] = &pl
				}

				pokemonByAbilityId[a.Ability.ID].AddPokemon(pkmn)
			}
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
