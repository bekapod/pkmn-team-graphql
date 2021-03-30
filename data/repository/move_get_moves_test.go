package repository

import (
	"bekapod/pkmn-team-graphql/data/model"
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-test/deep"
)

func TestMove_GetMoves(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM moves ORDER BY slug ASC").
		WillReturnRows(mockRowsForGetMoves(false, false, false))

	moves := []*model.Move{
		{
			ID:           "9f61694f-34f0-4531-b5e4-aff9a3d9edde",
			Slug:         "accelerock",
			Name:         "Accelerock",
			Accuracy:     100,
			PP:           20,
			Power:        40,
			DamageClass:  model.Physical,
			Effect:       "Inflicts regular damage with no additional effect.",
			EffectChance: 0,
			Target:       "Selected Pokémon",
			TypeID:       "5179f383-b765-4cc7-b9f9-8b1a3ba93019",
		},
		{
			ID:           "fd22f390-1745-4af2-adbe-9d9eca2db086",
			Slug:         "infestation",
			Name:         "Infestation",
			Accuracy:     100,
			PP:           20,
			Power:        20,
			DamageClass:  model.Special,
			Effect:       "Prevents the target from fleeing and inflicts damage for 2-5 turns.",
			EffectChance: 100,
			Target:       "Selected Pokémon",
			TypeID:       "56dddb9a-3623-43c5-8228-ea24d598afe7",
		},
	}

	exp := model.NewMoveList(moves)
	got, err := NewMove(db).GetMoves(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if diff := deep.Equal(&exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestMove_GetMoves_WithQueryError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM moves ORDER BY slug ASC").
		WillReturnError(errors.New("I am Error."))

	got, err := NewMove(db).GetMoves(context.Background())
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := &model.MoveList{
		Total: 0,
		Moves: []*model.Move{},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestMove_GetMoves_WithScanError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM moves ORDER BY slug ASC").
		WillReturnRows(mockRowsForGetMoves(false, false, true))

	got, err := NewMove(db).GetMoves(context.Background())
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := &model.MoveList{
		Total: 0,
		Moves: []*model.Move{},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestMove_GetMoves_WithRowError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM moves ORDER BY slug ASC").
		WillReturnRows(mockRowsForGetMoves(false, true, false))

	got, err := NewMove(db).GetMoves(context.Background())
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := &model.MoveList{
		Total: 0,
		Moves: []*model.Move{},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestMove_GetMoves_WithNoRows(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM moves ORDER BY slug ASC").
		WillReturnRows(mockRowsForGetMoves(true, false, false))

	got, err := NewMove(db).GetMoves(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	exp := &model.MoveList{
		Total: 0,
		Moves: []*model.Move{},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestMove_GetMoveById(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM moves WHERE id.*").
		WithArgs("9f61694f-34f0-4531-b5e4-aff9a3d9edde").
		WillReturnRows(mockRowsForGetMoveById(false))

	exp := model.Move{
		ID:           "9f61694f-34f0-4531-b5e4-aff9a3d9edde",
		Slug:         "accelerock",
		Name:         "Accelerock",
		Accuracy:     100,
		PP:           20,
		Power:        40,
		DamageClass:  model.Physical,
		Effect:       "Inflicts regular damage with no additional effect.",
		EffectChance: 0,
		Target:       "Selected Pokémon",
		TypeID:       "5179f383-b765-4cc7-b9f9-8b1a3ba93019",
	}
	got, err := NewMove(db).GetMoveById(context.Background(), "9f61694f-34f0-4531-b5e4-aff9a3d9edde")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if diff := deep.Equal(&exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestMove_GetMoveById_WithQueryError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM moves WHERE id.*").
		WithArgs("9f61694f-34f0-4531-b5e4-aff9a3d9edde").
		WillReturnError(errors.New("I am Error."))

	_, err := NewMove(db).GetMoveById(context.Background(), "9f61694f-34f0-4531-b5e4-aff9a3d9edde")
	if err == nil {
		t.Error("expected an error but got nil")
	}
}

func TestMove_GetMoveById_WithNoRows(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM moves WHERE id.*").
		WithArgs("9f61694f-34f0-4531-b5e4-aff9a3d9edde").
		WillReturnRows(mockRowsForGetMoveById(true))

	got, err := NewMove(db).GetMoveById(context.Background(), "9f61694f-34f0-4531-b5e4-aff9a3d9edde")
	if err == nil {
		t.Error("expected an error but got nil")
	}

	var exp *model.Move = nil

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestMove_MovesByPokemonIdDataLoader(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM moves LEFT JOIN pokemon_move ON moves.id = pokemon_move.move_id WHERE pokemon_move.pokemon_id IN (.*)").
		WithArgs("49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a").
		WillReturnRows(mockRowsForMovesByPokemonIdDataLoader(false, false, false, []string{"49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a"}))

	got, err := NewMove(db).MovesByPokemonIdDataLoader(context.Background())([]string{"49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a"})
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	exp := []*model.MoveList{
		{
			Total: 1,
			Moves: []*model.Move{
				{
					ID:           "9f61694f-34f0-4531-b5e4-aff9a3d9edde",
					Slug:         "accelerock",
					Name:         "Accelerock",
					Accuracy:     100,
					PP:           20,
					Power:        40,
					DamageClass:  model.Physical,
					Effect:       "Inflicts regular damage with no additional effect.",
					EffectChance: 0,
					Target:       "Selected Pokémon",
					TypeID:       "5179f383-b765-4cc7-b9f9-8b1a3ba93019",
				},
			},
		},
		nil,
		{
			Total: 2,
			Moves: []*model.Move{
				{
					ID:           "9f61694f-34f0-4531-b5e4-aff9a3d9edde",
					Slug:         "accelerock",
					Name:         "Accelerock",
					Accuracy:     100,
					PP:           20,
					Power:        40,
					DamageClass:  model.Physical,
					Effect:       "Inflicts regular damage with no additional effect.",
					EffectChance: 0,
					Target:       "Selected Pokémon",
					TypeID:       "5179f383-b765-4cc7-b9f9-8b1a3ba93019",
				},
				{
					ID:           "fd22f390-1745-4af2-adbe-9d9eca2db086",
					Slug:         "infestation",
					Name:         "Infestation",
					Accuracy:     100,
					PP:           20,
					Power:        20,
					DamageClass:  model.Special,
					Effect:       "Prevents the target from fleeing and inflicts damage for 2-5 turns.",
					EffectChance: 100,
					Target:       "Selected Pokémon",
					TypeID:       "56dddb9a-3623-43c5-8228-ea24d598afe7",
				},
			},
		},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestMove_MovesByPokemonIdDataLoader_WithQueryError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM moves LEFT JOIN pokemon_move ON moves.id = pokemon_move.move_id WHERE pokemon_move.pokemon_id IN (.*)").
		WithArgs("49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a").
		WillReturnRows(mockRowsForMovesByPokemonIdDataLoader(false, false, false, []string{"49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a"})).
		WillReturnError(errors.New("I am Error."))

	got, err := NewMove(db).MovesByPokemonIdDataLoader(context.Background())([]string{"49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a"})
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := []*model.MoveList{
		{
			Total: 0,
			Moves: []*model.Move{},
		},
		{
			Total: 0,
			Moves: []*model.Move{},
		},
		{
			Total: 0,
			Moves: []*model.Move{},
		},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestMove_MovesByPokemonIdDataLoader_WithScanError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM moves LEFT JOIN pokemon_move ON moves.id = pokemon_move.move_id WHERE pokemon_move.pokemon_id IN (.*)").
		WithArgs("49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a").
		WillReturnRows(mockRowsForMovesByPokemonIdDataLoader(false, false, true, []string{"49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a"}))

	got, err := NewMove(db).MovesByPokemonIdDataLoader(context.Background())([]string{"49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a"})
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := []*model.MoveList{
		{
			Total: 0,
			Moves: []*model.Move{},
		},
		{
			Total: 0,
			Moves: []*model.Move{},
		},
		{
			Total: 0,
			Moves: []*model.Move{},
		},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestMove_MovesByPokemonIdDataLoader_WithRowError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM moves LEFT JOIN pokemon_move ON moves.id = pokemon_move.move_id WHERE pokemon_move.pokemon_id IN (.*)").
		WithArgs("49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a").
		WillReturnRows(mockRowsForMovesByPokemonIdDataLoader(false, true, false, []string{"49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a"}))

	got, err := NewMove(db).MovesByPokemonIdDataLoader(context.Background())([]string{"49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a"})
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := []*model.MoveList{
		{
			Total: 1,
			Moves: []*model.Move{
				{
					ID:           "9f61694f-34f0-4531-b5e4-aff9a3d9edde",
					Slug:         "accelerock",
					Name:         "Accelerock",
					Accuracy:     100,
					PP:           20,
					Power:        40,
					DamageClass:  model.Physical,
					Effect:       "Inflicts regular damage with no additional effect.",
					EffectChance: 0,
					Target:       "Selected Pokémon",
					TypeID:       "5179f383-b765-4cc7-b9f9-8b1a3ba93019",
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

func TestMove_MovesByTypeIdDataLoader(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM moves WHERE type_id IN (.*)").
		WithArgs("56dddb9a-3623-43c5-8228-ea24d598afe7", "a248c127-8e9c-4f87-8513-c5dbc3385011", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b").
		WillReturnRows(mockRowsForMovesByTypeIdDataLoader(false, false, false, []string{"56dddb9a-3623-43c5-8228-ea24d598afe7", "a248c127-8e9c-4f87-8513-c5dbc3385011", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b"}))

	got, err := NewMove(db).MovesByTypeIdDataLoader(context.Background())([]string{"56dddb9a-3623-43c5-8228-ea24d598afe7", "a248c127-8e9c-4f87-8513-c5dbc3385011", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b"})
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	exp := []*model.MoveList{
		{
			Total: 1,
			Moves: []*model.Move{
				{
					ID:           "9f61694f-34f0-4531-b5e4-aff9a3d9edde",
					Slug:         "accelerock",
					Name:         "Accelerock",
					Accuracy:     100,
					PP:           20,
					Power:        40,
					DamageClass:  model.Physical,
					Effect:       "Inflicts regular damage with no additional effect.",
					EffectChance: 0,
					Target:       "Selected Pokémon",
					TypeID:       "56dddb9a-3623-43c5-8228-ea24d598afe7",
				},
			},
		},
		nil,
		{
			Total: 2,
			Moves: []*model.Move{
				{
					ID:           "9f61694f-34f0-4531-b5e4-aff9a3d9edde",
					Slug:         "accelerock",
					Name:         "Accelerock",
					Accuracy:     100,
					PP:           20,
					Power:        40,
					DamageClass:  model.Physical,
					Effect:       "Inflicts regular damage with no additional effect.",
					EffectChance: 0,
					Target:       "Selected Pokémon",
					TypeID:       "05cd51bd-23ca-4736-b8ec-aa93aca68a8b",
				},
				{
					ID:           "fd22f390-1745-4af2-adbe-9d9eca2db086",
					Slug:         "infestation",
					Name:         "Infestation",
					Accuracy:     100,
					PP:           20,
					Power:        20,
					DamageClass:  model.Special,
					Effect:       "Prevents the target from fleeing and inflicts damage for 2-5 turns.",
					EffectChance: 100,
					Target:       "Selected Pokémon",
					TypeID:       "05cd51bd-23ca-4736-b8ec-aa93aca68a8b",
				},
			},
		},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestMove_MovesByTypeIdDataLoader_WithQueryError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM moves WHERE type_id IN (.*)").
		WithArgs("56dddb9a-3623-43c5-8228-ea24d598afe7", "a248c127-8e9c-4f87-8513-c5dbc3385011", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b").
		WillReturnRows(mockRowsForMovesByTypeIdDataLoader(false, false, false, []string{"56dddb9a-3623-43c5-8228-ea24d598afe7", "a248c127-8e9c-4f87-8513-c5dbc3385011", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b"})).
		WillReturnError(errors.New("I am Error."))

	got, err := NewMove(db).MovesByTypeIdDataLoader(context.Background())([]string{"56dddb9a-3623-43c5-8228-ea24d598afe7", "a248c127-8e9c-4f87-8513-c5dbc3385011", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b"})
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := []*model.MoveList{
		{
			Total: 0,
			Moves: []*model.Move{},
		},
		{
			Total: 0,
			Moves: []*model.Move{},
		},
		{
			Total: 0,
			Moves: []*model.Move{},
		},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestMove_MovesByTypeIdDataLoader_WithScanError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM moves WHERE type_id IN (.*)").
		WithArgs("56dddb9a-3623-43c5-8228-ea24d598afe7", "a248c127-8e9c-4f87-8513-c5dbc3385011", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b").
		WillReturnRows(mockRowsForMovesByTypeIdDataLoader(false, false, true, []string{"56dddb9a-3623-43c5-8228-ea24d598afe7", "a248c127-8e9c-4f87-8513-c5dbc3385011", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b"}))

	got, err := NewMove(db).MovesByTypeIdDataLoader(context.Background())([]string{"56dddb9a-3623-43c5-8228-ea24d598afe7", "a248c127-8e9c-4f87-8513-c5dbc3385011", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b"})
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := []*model.MoveList{
		{
			Total: 0,
			Moves: []*model.Move{},
		},
		{
			Total: 0,
			Moves: []*model.Move{},
		},
		{
			Total: 0,
			Moves: []*model.Move{},
		},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestMove_MovesByTypeIdDataLoader_WithRowError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM moves WHERE type_id IN (.*)").
		WithArgs("56dddb9a-3623-43c5-8228-ea24d598afe7", "a248c127-8e9c-4f87-8513-c5dbc3385011", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b").
		WillReturnRows(mockRowsForMovesByTypeIdDataLoader(false, true, false, []string{"56dddb9a-3623-43c5-8228-ea24d598afe7", "a248c127-8e9c-4f87-8513-c5dbc3385011", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b"}))

	got, err := NewMove(db).MovesByTypeIdDataLoader(context.Background())([]string{"56dddb9a-3623-43c5-8228-ea24d598afe7", "a248c127-8e9c-4f87-8513-c5dbc3385011", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b"})
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := []*model.MoveList{
		{
			Total: 1,
			Moves: []*model.Move{
				{
					ID:           "9f61694f-34f0-4531-b5e4-aff9a3d9edde",
					Slug:         "accelerock",
					Name:         "Accelerock",
					Accuracy:     100,
					PP:           20,
					Power:        40,
					DamageClass:  model.Physical,
					Effect:       "Inflicts regular damage with no additional effect.",
					EffectChance: 0,
					Target:       "Selected Pokémon",
					TypeID:       "56dddb9a-3623-43c5-8228-ea24d598afe7",
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

func mockRowsForGetMoves(empty bool, hasRowError bool, hasScanError bool) *sqlmock.Rows {
	if hasScanError {
		rows := sqlmock.NewRows([]string{"id"})
		rows.AddRow("9f61694f-34f0-4531-b5e4-aff9a3d9edde")
		return rows
	}
	rows := sqlmock.NewRows([]string{"id", "name", "slug", "accuracy", "pp", "power", "damage_class_enum", "effect", "effect_chance", "target", "type_id"})
	if !empty {
		rows.AddRow("9f61694f-34f0-4531-b5e4-aff9a3d9edde", "Accelerock", "accelerock", 100, 20, 40, "physical", "Inflicts regular damage with no additional effect.", 0, "Selected Pokémon", "5179f383-b765-4cc7-b9f9-8b1a3ba93019").
			AddRow("fd22f390-1745-4af2-adbe-9d9eca2db086", "Infestation", "infestation", 100, 20, 20, "special", "Prevents the target from fleeing and inflicts damage for 2-5 turns.", 100, "Selected Pokémon", "56dddb9a-3623-43c5-8228-ea24d598afe7")
	}
	if hasRowError {
		rows.RowError(0, errors.New("scan error"))
	}
	return rows
}

func mockRowsForGetMoveById(empty bool) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"id", "name", "slug", "accuracy", "pp", "power", "damage_class_enum", "effect", "effect_chance", "target", "type_id"})
	if !empty {
		rows.AddRow("9f61694f-34f0-4531-b5e4-aff9a3d9edde", "Accelerock", "accelerock", 100, 20, 40, "physical", "Inflicts regular damage with no additional effect.", 0, "Selected Pokémon", "5179f383-b765-4cc7-b9f9-8b1a3ba93019")
	}
	return rows
}

func mockRowsForMovesByPokemonIdDataLoader(empty bool, hasRowError bool, hasScanError bool, ids []string) *sqlmock.Rows {
	if hasScanError {
		rows := sqlmock.NewRows([]string{"id"})
		rows.AddRow("9f61694f-34f0-4531-b5e4-aff9a3d9edde")
		return rows
	}
	rows := sqlmock.NewRows([]string{"id", "name", "slug", "accuracy", "pp", "power", "damage_class_enum", "effect", "effect_chance", "target", "type_id", "pokemon_move.pokemon_id"})
	if !empty {
		rows.AddRow("9f61694f-34f0-4531-b5e4-aff9a3d9edde", "Accelerock", "accelerock", 100, 20, 40, "physical", "Inflicts regular damage with no additional effect.", 0, "Selected Pokémon", "5179f383-b765-4cc7-b9f9-8b1a3ba93019", ids[0]).
			AddRow("9f61694f-34f0-4531-b5e4-aff9a3d9edde", "Accelerock", "accelerock", 100, 20, 40, "physical", "Inflicts regular damage with no additional effect.", 0, "Selected Pokémon", "5179f383-b765-4cc7-b9f9-8b1a3ba93019", ids[2]).
			AddRow("fd22f390-1745-4af2-adbe-9d9eca2db086", "Infestation", "infestation", 100, 20, 20, "special", "Prevents the target from fleeing and inflicts damage for 2-5 turns.", 100, "Selected Pokémon", "56dddb9a-3623-43c5-8228-ea24d598afe7", ids[2])
	}
	if hasRowError {
		rows.RowError(1, errors.New("scan error"))
	}
	return rows
}

func mockRowsForMovesByTypeIdDataLoader(empty bool, hasRowError bool, hasScanError bool, ids []string) *sqlmock.Rows {
	if hasScanError {
		rows := sqlmock.NewRows([]string{"id"})
		rows.AddRow("9f61694f-34f0-4531-b5e4-aff9a3d9edde")
		return rows
	}
	rows := sqlmock.NewRows([]string{"id", "name", "slug", "accuracy", "pp", "power", "damage_class_enum", "effect", "effect_chance", "target", "type_id"})
	if !empty {
		rows.AddRow("9f61694f-34f0-4531-b5e4-aff9a3d9edde", "Accelerock", "accelerock", 100, 20, 40, "physical", "Inflicts regular damage with no additional effect.", 0, "Selected Pokémon", ids[0]).
			AddRow("9f61694f-34f0-4531-b5e4-aff9a3d9edde", "Accelerock", "accelerock", 100, 20, 40, "physical", "Inflicts regular damage with no additional effect.", 0, "Selected Pokémon", ids[2]).
			AddRow("fd22f390-1745-4af2-adbe-9d9eca2db086", "Infestation", "infestation", 100, 20, 20, "special", "Prevents the target from fleeing and inflicts damage for 2-5 turns.", 100, "Selected Pokémon", ids[2])
	}
	if hasRowError {
		rows.RowError(1, errors.New("scan error"))
	}
	return rows
}
