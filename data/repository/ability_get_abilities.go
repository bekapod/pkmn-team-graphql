package repository

import (
	"bekapod/pkmn-team-graphql/data/model"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

var (
	ErrNoAbility = errors.New("no ability found")
)

var AbilityColumns = "abilities.id, abilities.name, abilities.slug, abilities.effect"

func (r Ability) GetAbilities(ctx context.Context) (*model.AbilityList, error) {
	abilities := model.NewEmptyAbilityList()

	rows, err := r.db.QueryContext(
		ctx,
		"SELECT "+AbilityColumns+" FROM abilities ORDER BY slug ASC",
	)
	if err != nil {
		return &abilities, fmt.Errorf("error fetching all abilities in GetAllAbilities: %w", err)
	}

	defer rows.Close()
	for rows.Next() {
		var a model.Ability
		err := rows.Scan(&a.ID, &a.Name, &a.Slug, &a.Effect)
		if err != nil {
			return &abilities, fmt.Errorf("error scanning result in GetAllAbilities: %w", err)
		}
		abilities.AddAbility(a)
	}
	err = rows.Err()
	if err != nil {
		return &abilities, fmt.Errorf("error after fetching all abilities in GetAllAbilities: %w", err)
	}

	return &abilities, nil
}

func (r Ability) GetAbilityById(ctx context.Context, id string) (*model.Ability, error) {
	a := model.Ability{}

	err := r.db.QueryRowContext(
		ctx,
		"SELECT  "+AbilityColumns+" FROM abilities WHERE id = $1",
		id,
	).Scan(&a.ID, &a.Name, &a.Slug, &a.Effect)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNoAbility
		}
		return nil, fmt.Errorf("error scanning result in GetAbilityById %s: %w", id, err)
	}

	return &a, nil
}

func (r Ability) AbilitiesByPokemonIdDataLoader(ctx context.Context) func(pokemonIds []string) ([]*model.AbilityList, []error) {
	return func(pokemonIds []string) ([]*model.AbilityList, []error) {
		abilitiesByPokemonId := map[string]*model.AbilityList{}
		placeholders := make([]string, len(pokemonIds))
		args := make([]interface{}, len(pokemonIds))
		for i := 0; i < len(pokemonIds); i++ {
			placeholders[i] = fmt.Sprintf("$%d", i+1)
			args[i] = pokemonIds[i]
		}

		query := `
			SELECT ` + AbilityColumns + `, pokemon_ability.pokemon_id
			FROM abilities
				LEFT JOIN pokemon_ability ON abilities.id = pokemon_ability.ability_id
			WHERE pokemon_ability.pokemon_id IN (` + strings.Join(placeholders, ",") + `)`

		rows, err := r.db.QueryContext(ctx,
			query,
			args...,
		)
		if err != nil {
			abilityList := make([]*model.AbilityList, len(pokemonIds))
			emptyAbilityList := model.NewEmptyAbilityList()
			errors := make([]error, len(pokemonIds))
			for i := range pokemonIds {
				abilityList[i] = &emptyAbilityList
				errors[i] = fmt.Errorf("error fetching abilities for pokemon in AbilitiesByPokemonIdDataLoader: %w", err)
			}
			return abilityList, errors
		}

		defer rows.Close()
		for rows.Next() {
			var a model.Ability
			var pokemonId string
			err := rows.Scan(&a.ID, &a.Name, &a.Slug, &a.Effect, &pokemonId)
			if err != nil {
				abilityList := make([]*model.AbilityList, len(pokemonIds))
				emptyAbilityList := model.NewEmptyAbilityList()
				errors := make([]error, len(pokemonIds))
				for i := range pokemonIds {
					abilityList[i] = &emptyAbilityList
					errors[i] = fmt.Errorf("error scanning result abilities for pokemon in AbilitiesByPokemonIdDataLoader: %w", err)
				}
				return abilityList, errors
			}

			_, ok := abilitiesByPokemonId[pokemonId]
			if !ok {
				al := model.NewEmptyAbilityList()
				abilitiesByPokemonId[pokemonId] = &al
			}

			abilitiesByPokemonId[pokemonId].AddAbility(a)
		}

		abilityList := make([]*model.AbilityList, len(pokemonIds))
		for i, id := range pokemonIds {
			abilityList[i] = abilitiesByPokemonId[id]
			i++
		}

		err = rows.Err()
		if err != nil {
			errors := make([]error, len(pokemonIds))
			for i := range pokemonIds {
				errors[i] = fmt.Errorf("error after fetching abilities for pokemon in AbilitiesByPokemonIdDataLoader: %w", err)
			}
			return abilityList, errors
		}

		return abilityList, nil
	}
}
