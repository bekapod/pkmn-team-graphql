package repository

import (
	"bekapod/pkmn-team-graphql/data/model"
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-test/deep"
)

func TestEvolution_EvolutionsByPokemonIdDataLoader(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM pokemon_evolutions .* pokemon_evolutions.from_pokemon_id IN (.*)").
		WithArgs("49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a").
		WillReturnRows(mockRowsForEvolutionsByPokemonIdDataLoader(false, false, false, []string{"49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a"}, EvolvesTo))

	got, err := NewEvolution(db).EvolutionsByPokemonIdDataLoader(context.Background(), EvolvesTo)([]string{"49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a"})
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	pokemonId := "e0ccb29a-f73d-4478-a78e-25043420a1d1"
	pokemonId2 := "241b3038-a984-4104-a974-d13bee16905d"
	exp := []*model.EvolutionList{
		{
			Total: 1,
			Evolutions: []model.Evolution{
				{
					Trigger: model.Trade,
					Gender:  model.Unknown,
					HeldItem: &model.Item{
						ID:         "0d56a358-3e32-424f-9271-c2d9e73fe549",
						Slug:       "metal-coat",
						Name:       "Metal Coat",
						Cost:       2000,
						Effect:     "Held: Steel-Type moves from holder do 20% more damage.",
						Sprite:     "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/items/metal-coat.png",
						Category:   model.TypeEnhancement,
						Attributes: []model.ItemAttribute{model.Holdable, model.HoldableActive},
					},
					TimeOfDay: "any",
					PokemonID: &pokemonId,
				},
			},
		},
		nil,
		{
			Total: 2,
			Evolutions: []model.Evolution{
				{
					Trigger:   model.LevelUp,
					Gender:    model.Unknown,
					MinLevel:  22,
					TimeOfDay: model.Any,
					PokemonID: &pokemonId2,
				},
				{
					Trigger:   model.LevelUp,
					Gender:    model.Unknown,
					MinLevel:  22,
					TimeOfDay: model.Any,
					PokemonID: &pokemonId2,
				},
			},
		},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestEvolution_EvolutionsByPokemonIdDataLoader_WithQueryError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM pokemon_evolutions .* pokemon_evolutions.to_pokemon_id IN (.*)").
		WithArgs("49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a").
		WillReturnRows(mockRowsForEvolutionsByPokemonIdDataLoader(false, false, false, []string{"49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a"}, EvolvesFrom)).
		WillReturnError(errors.New("I am Error."))

	got, err := NewEvolution(db).EvolutionsByPokemonIdDataLoader(context.Background(), EvolvesFrom)([]string{"49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a"})
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := []*model.EvolutionList{
		{
			Total:      0,
			Evolutions: []model.Evolution{},
		},
		{
			Total:      0,
			Evolutions: []model.Evolution{},
		},
		{
			Total:      0,
			Evolutions: []model.Evolution{},
		},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestEvolution_EvolutionsByPokemonIdDataLoader_WithScanError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM pokemon_evolutions .* pokemon_evolutions.from_pokemon_id IN (.*)").
		WithArgs("49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a").
		WillReturnRows(mockRowsForEvolutionsByPokemonIdDataLoader(false, false, true, []string{"49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a"}, EvolvesTo))

	got, err := NewEvolution(db).EvolutionsByPokemonIdDataLoader(context.Background(), EvolvesTo)([]string{"49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a"})
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := []*model.EvolutionList{
		{
			Total:      0,
			Evolutions: []model.Evolution{},
		},
		{
			Total:      0,
			Evolutions: []model.Evolution{},
		},
		{
			Total:      0,
			Evolutions: []model.Evolution{},
		},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestEvolution_EvolutionsByPokemonIdDataLoader_WithRowError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM pokemon_evolutions .* pokemon_evolutions.to_pokemon_id IN (.*)").
		WithArgs("49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a").
		WillReturnRows(mockRowsForEvolutionsByPokemonIdDataLoader(false, true, false, []string{"49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a"}, EvolvesFrom))

	got, err := NewEvolution(db).EvolutionsByPokemonIdDataLoader(context.Background(), EvolvesFrom)([]string{"49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a"})
	if err == nil {
		t.Error("expected an error but got nil")
	}

	pokemonId := "e0ccb29a-f73d-4478-a78e-25043420a1d1"
	exp := []*model.EvolutionList{
		{
			Total: 1,
			Evolutions: []model.Evolution{
				{
					Trigger: model.Trade,
					Gender:  model.Unknown,
					HeldItem: &model.Item{
						ID:         "0d56a358-3e32-424f-9271-c2d9e73fe549",
						Slug:       "metal-coat",
						Name:       "Metal Coat",
						Cost:       2000,
						Effect:     "Held: Steel-Type moves from holder do 20% more damage.",
						Sprite:     "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/items/metal-coat.png",
						Category:   model.TypeEnhancement,
						Attributes: []model.ItemAttribute{model.Holdable, model.HoldableActive},
					},
					TimeOfDay: "any",
					PokemonID: &pokemonId,
				},
			},
		},
		nil,
		nil,
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func mockRowsForEvolutionsByPokemonIdDataLoader(empty bool, hasRowError bool, hasScanError bool, ids []string, direction EvolutionDirection) *sqlmock.Rows {
	if hasScanError {
		rows := sqlmock.NewRows([]string{"id"})
		rows.AddRow("9f61694f-34f0-4531-b5e4-aff9a3d9edde")
		return rows
	}
	rows := sqlmock.NewRows([]string{"trigger_enum", "gender_enum", "min_level", "min_happiness", "min_beauty", "min_affection", "needs_overworld_rain", "relative_physical_stats", "time_of_day_enum", "turn_upside_down", "spin", "take_damage", "critical_hits", "item", "held_item", "location", "party_species_pokemon_id", "trade_species_pokemon_id", "known_move_id", "known_move_type_id", "party_type_id", "to_pokemon_id", "from_pokemon_id"})

	if direction == EvolvesFrom {
		rows = sqlmock.NewRows([]string{"trigger_enum", "gender_enum", "min_level", "min_happiness", "min_beauty", "min_affection", "needs_overworld_rain", "relative_physical_stats", "time_of_day_enum", "turn_upside_down", "spin", "take_damage", "critical_hits", "item", "held_item", "location", "party_species_pokemon_id", "trade_species_pokemon_id", "known_move_id", "known_move_type_id", "party_type_id", "from_pokemon_id", "to_pokemon_id"})
	}

	if !empty {
		rows.AddRow("trade", "unknown", 0, 0, 0, 0, false, 0, "any", false, false, 0, 0, "{\"id\": null,\"cost\": null,\"name\": null,\"slug\": null,\"effect\": null,\"sprite\": null,\"category\": null,\"attributes\": null,\"flingPower\": null,\"flingEffect\": null}", "{\"id\": \"0d56a358-3e32-424f-9271-c2d9e73fe549\",\"cost\": 2000,\"name\": \"Metal Coat\",\"slug\": \"metal-coat\",\"effect\": \"Held: Steel-Type moves from holder do 20% more damage.\",\"sprite\": \"https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/items/metal-coat.png\",\"category\": \"type-enhancement\",\"attributes\": [\"holdable\",\"holdable-active\"],\"flingPower\": 30,\"flingEffect\": null}", "{\"id\": null,\"name\": null,\"slug\": null,\"region\": null}", nil, nil, nil, nil, nil, "e0ccb29a-f73d-4478-a78e-25043420a1d1", ids[0]).
			AddRow("level-up", "unknown", 22, 0, 0, 0, false, 0, "any", false, false, 0, 0, "{\"id\": null,\"cost\": null,\"name\": null,\"slug\": null,\"effect\": null,\"sprite\": null,\"category\": null,\"attributes\": null,\"flingPower\": null,\"flingEffect\": null}", "{\"id\": null,\"cost\": null,\"name\": null,\"slug\": null,\"effect\": null,\"sprite\": null,\"category\": null,\"attributes\": null,\"flingPower\": null,\"flingEffect\": null}", "{\"id\": null,\"name\": null,\"slug\": null,\"region\": null}", nil, nil, nil, nil, nil, "241b3038-a984-4104-a974-d13bee16905d", ids[2]).
			AddRow("level-up", "unknown", 22, 0, 0, 0, false, 0, "any", false, false, 0, 0, "{\"id\": null,\"cost\": null,\"name\": null,\"slug\": null,\"effect\": null,\"sprite\": null,\"category\": null,\"attributes\": null,\"flingPower\": null,\"flingEffect\": null}", "{\"id\": null,\"cost\": null,\"name\": null,\"slug\": null,\"effect\": null,\"sprite\": null,\"category\": null,\"attributes\": null,\"flingPower\": null,\"flingEffect\": null}", "{\"id\": null,\"name\": null,\"slug\": null,\"region\": null}", nil, nil, nil, nil, nil, "241b3038-a984-4104-a974-d13bee16905d", ids[2])
	}
	if hasRowError {
		rows.RowError(1, errors.New("row error"))
	}
	return rows
}
