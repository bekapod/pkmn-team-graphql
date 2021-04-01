package repository

import (
	"bekapod/pkmn-team-graphql/data/model"
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-test/deep"
)

func TestPokemonType_PokemonTypesByPokemonIdDataLoader(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM pokemon_type WHERE pokemon_type.pokemon_id IN (.*) ORDER BY slot").
		WithArgs("49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a").
		WillReturnRows(mockRowsForPokemonTypesByPokemonIdDataLoader(false, false, false, []string{"49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a"}))

	got, err := NewPokemonType(db).PokemonTypesByPokemonIdDataLoader(context.Background())([]string{"49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a"})
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	exp := []*model.PokemonTypeList{
		{
			Total: 1,
			PokemonTypes: []*model.PokemonType{
				{
					TypeID:    "5179f383-b765-4cc7-b9f9-8b1a3ba93019",
					PokemonID: "49653637-1d35-4138-98eb-14305a2741a0",
					Slot:      1,
				},
			},
		},
		nil,
		{
			Total: 2,
			PokemonTypes: []*model.PokemonType{
				{
					TypeID:    "5179f383-b765-4cc7-b9f9-8b1a3ba93019",
					PokemonID: "49de1627-e7b3-4a54-8d42-0ed7c795f28a",
					Slot:      1,
				},
				{
					TypeID:    "fd22f390-1745-4af2-adbe-9d9eca2db086",
					PokemonID: "49de1627-e7b3-4a54-8d42-0ed7c795f28a",
					Slot:      2,
				},
			},
		},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestPokemonType_PokemonTypesByPokemonIdDataLoader_WithQueryError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM pokemon_type WHERE pokemon_type.pokemon_id IN (.*) ORDER BY slot").
		WithArgs("49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a").
		WillReturnRows(mockRowsForPokemonTypesByPokemonIdDataLoader(false, false, false, []string{"49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a"})).
		WillReturnError(errors.New("i am Error"))

	got, err := NewPokemonType(db).PokemonTypesByPokemonIdDataLoader(context.Background())([]string{"49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a"})
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := []*model.PokemonTypeList{
		{
			Total:        0,
			PokemonTypes: []*model.PokemonType{},
		},
		{
			Total:        0,
			PokemonTypes: []*model.PokemonType{},
		},
		{
			Total:        0,
			PokemonTypes: []*model.PokemonType{},
		},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestPokemonType_PokemonTypesByPokemonIdDataLoader_WithScanError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM pokemon_type WHERE pokemon_type.pokemon_id IN (.*) ORDER BY slot").
		WithArgs("49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a").
		WillReturnRows(mockRowsForPokemonTypesByPokemonIdDataLoader(false, false, true, []string{"49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a"}))

	got, err := NewPokemonType(db).PokemonTypesByPokemonIdDataLoader(context.Background())([]string{"49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a"})
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := []*model.PokemonTypeList{
		{
			Total:        0,
			PokemonTypes: []*model.PokemonType{},
		},
		{
			Total:        0,
			PokemonTypes: []*model.PokemonType{},
		},
		{
			Total:        0,
			PokemonTypes: []*model.PokemonType{},
		},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestPokemonType_PokemonTypesByPokemonIdDataLoader_WithRowError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM pokemon_type WHERE pokemon_type.pokemon_id IN (.*) ORDER BY slot").
		WithArgs("49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a").
		WillReturnRows(mockRowsForPokemonTypesByPokemonIdDataLoader(false, true, false, []string{"49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a"}))

	got, err := NewPokemonType(db).PokemonTypesByPokemonIdDataLoader(context.Background())([]string{"49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a"})
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := []*model.PokemonTypeList{
		{
			Total: 1,
			PokemonTypes: []*model.PokemonType{
				{
					TypeID:    "5179f383-b765-4cc7-b9f9-8b1a3ba93019",
					PokemonID: "49653637-1d35-4138-98eb-14305a2741a0",
					Slot:      1,
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

func mockRowsForPokemonTypesByPokemonIdDataLoader(empty bool, hasRowError bool, hasScanError bool, ids []string) *sqlmock.Rows {
	if hasScanError {
		rows := sqlmock.NewRows([]string{"id"})
		rows.AddRow("9f61694f-34f0-4531-b5e4-aff9a3d9edde")
		return rows
	}
	rows := sqlmock.NewRows([]string{"pokemon_id", "type_id", "slot"})
	if !empty {
		rows.AddRow(ids[0], "5179f383-b765-4cc7-b9f9-8b1a3ba93019", 1).
			AddRow(ids[2], "5179f383-b765-4cc7-b9f9-8b1a3ba93019", 1).
			AddRow(ids[2], "fd22f390-1745-4af2-adbe-9d9eca2db086", 2)
	}
	if hasRowError {
		rows.RowError(1, errors.New("scan error"))
	}
	return rows
}
