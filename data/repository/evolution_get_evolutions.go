package repository

import (
	"bekapod/pkmn-team-graphql/data/model"
	"context"
	"fmt"
	"strings"
)

func getItemObject(tableName string) string {
	return fmt.Sprintf("jsonb_build_object('id', %[1]s.id, 'name', %[1]s.name, 'slug', %[1]s.slug)", tableName)
}

func getLocationObject(tableName string) string {
	return fmt.Sprintf(`
		jsonb_build_object(
			'id', %[1]s.id,
			'name', %[1]s.name,
			'slug', %[1]s.slug,
			'region', (
				SELECT jsonb_build_object('id', region.id, 'name', region.name, 'slug', region.slug)
				FROM regions region
				WHERE region.id = location.region_id
			)
		)
	`, tableName)
}

var EvolutionColumns = "trigger_enum, gender_enum, min_level, min_happiness, min_beauty, min_affection, needs_overworld_rain, relative_physical_stats, time_of_day_enum, turn_upside_down, spin, take_damage, critical_hits, " + getItemObject("item") + " as item, " + getItemObject("held_item") + " as held_item, " + getLocationObject("location") + " as location, from_pokemon_id, to_pokemon_id, party_species_pokemon_id, trade_species_pokemon_id, known_move_id, known_move_type_id, party_type_id"

func (r Evolution) EvolutionsByPokemonIdDataLoader(ctx context.Context) func(pokemonIds []string) ([]*model.EvolutionList, []error) {
	return func(pokemonIds []string) ([]*model.EvolutionList, []error) {
		evolutionsByPokemonId := map[string]*model.EvolutionList{}
		placeholders := make([]string, len(pokemonIds))
		args := make([]interface{}, len(pokemonIds))
		for i := 0; i < len(pokemonIds); i++ {
			placeholders[i] = fmt.Sprintf("$%d", i+1)
			args[i] = pokemonIds[i]
		}

		query := `
			SELECT ` + EvolutionColumns + `
			FROM pokemon_evolutions
				LEFT JOIN items item ON item.id = item_id
				LEFT JOIN items held_item on held_item.id = held_item_id
				LEFT JOIN locations location ON location.id = location_id
			WHERE pokemon_evolutions.from_pokemon_id IN (` + strings.Join(placeholders, ",") + `)`

		rows, err := r.db.QueryContext(ctx,
			query,
			args...,
		)
		if err != nil {
			evolutionList := make([]*model.EvolutionList, len(pokemonIds))
			emptyEvolutionList := model.NewEmptyEvolutionList()
			errors := make([]error, len(pokemonIds))
			for i := range pokemonIds {
				evolutionList[i] = &emptyEvolutionList
				errors[i] = fmt.Errorf("error fetching evolutions for pokemon in EvolutionsByPokemonIdDataLoader: %w", err)
			}
			return evolutionList, errors
		}

		defer rows.Close()
		for rows.Next() {
			var e model.Evolution
			err := rows.Scan(&e.Trigger, &e.Gender, &e.MinLevel, &e.MinHappiness, &e.MinBeauty, &e.MinAffection, &e.NeedsOverworldRain, &e.RelativePhysicalStats, &e.TimeOfDay, &e.TurnUpsideDown, &e.Spin, &e.TakeDamage, &e.CriticalHits, &e.Item, &e.HeldItem, &e.Location, &e.FromPokemonID, &e.ToPokemonID, &e.PartySpeciesPokemonID, &e.TradeSpeciesPokemonID, &e.KnownMoveID, &e.KnownMoveTypeID, &e.PartyTypeID)
			if err != nil {
				evolutionList := make([]*model.EvolutionList, len(pokemonIds))
				emptyEvolutionList := model.NewEmptyEvolutionList()
				errors := make([]error, len(pokemonIds))
				for i := range pokemonIds {
					evolutionList[i] = &emptyEvolutionList
					errors[i] = fmt.Errorf("error scanning result evolutions for pokemon in EvolutionsByPokemonIdDataLoader: %w", err)
				}
				return evolutionList, errors
			}

			if e.Item.ID == "" {
				e.Item = nil
			}

			if e.HeldItem.ID == "" {
				e.HeldItem = nil
			}

			if e.Location.ID == "" {
				e.Location = nil
			}

			_, ok := evolutionsByPokemonId[*e.FromPokemonID]
			if !ok {
				el := model.NewEmptyEvolutionList()
				evolutionsByPokemonId[*e.FromPokemonID] = &el
			}

			evolutionsByPokemonId[*e.FromPokemonID].AddEvolution(e)
		}

		evolutionList := make([]*model.EvolutionList, len(pokemonIds))
		for i, id := range pokemonIds {
			evolutionList[i] = evolutionsByPokemonId[id]
			i++
		}

		err = rows.Err()
		if err != nil {
			errors := make([]error, len(pokemonIds))
			for i := range pokemonIds {
				errors[i] = fmt.Errorf("error after fetching evolutions for pokemon in EvolutionByPokemonIdDataLoader: %w", err)
			}
			return evolutionList, errors
		}

		return evolutionList, nil
	}
}
