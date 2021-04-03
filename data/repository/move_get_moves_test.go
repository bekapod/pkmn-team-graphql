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

	mock.ExpectQuery("SELECT .* FROM moves LEFT JOIN types ON moves.type_id = types.id ORDER BY moves.slug ASC").
		WillReturnRows(mockRowsForGetMoves(false, false, false))

	moves := []*model.Move{
		&accelerock,
		&infestation,
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

	mock.ExpectQuery("SELECT .* FROM moves ORDER BY moves.slug ASC").
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

	mock.ExpectQuery("SELECT .* FROM moves LEFT JOIN types ON moves.type_id = types.id ORDER BY moves.slug ASC").
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

	mock.ExpectQuery("SELECT .* FROM moves LEFT JOIN types ON moves.type_id = types.id ORDER BY moves.slug ASC").
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

	mock.ExpectQuery("SELECT .* FROM moves LEFT JOIN types ON moves.type_id = types.id ORDER BY moves.slug ASC").
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

	mock.ExpectQuery("SELECT .* FROM moves LEFT JOIN types ON moves.type_id = types.id WHERE moves.id.*").
		WithArgs("9f61694f-34f0-4531-b5e4-aff9a3d9edde").
		WillReturnRows(mockRowsForGetMoveById(false))

	exp := accelerock
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

	mock.ExpectQuery("SELECT .* FROM moves LEFT JOIN types ON moves.type_id = types.id WHERE moves.id.*").
		WithArgs("9f61694f-34f0-4531-b5e4-aff9a3d9edde").
		WillReturnError(errors.New("I am Error."))

	_, err := NewMove(db).GetMoveById(context.Background(), "9f61694f-34f0-4531-b5e4-aff9a3d9edde")
	if err == nil {
		t.Error("expected an error but got nil")
	}
}

func TestMove_GetMoveById_WithNoRows(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM moves LEFT JOIN types ON moves.type_id = types.id WHERE moves.id.*").
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

	mock.ExpectQuery("SELECT .* FROM moves LEFT JOIN types ON moves.type_id = types.id LEFT JOIN pokemon_move ON moves.id = pokemon_move.move_id WHERE pokemon_move.pokemon_id IN (.*)").
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
				&accelerock,
			},
		},
		nil,
		{
			Total: 2,
			Moves: []*model.Move{
				&accelerock,
				&infestation,
			},
		},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestMove_MovesByPokemonIdDataLoader_WithQueryError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM moves LEFT JOIN types ON moves.type_id = types.id LEFT JOIN pokemon_move ON moves.id = pokemon_move.move_id WHERE pokemon_move.pokemon_id IN (.*)").
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

	mock.ExpectQuery("SELECT .* FROM moves LEFT JOIN types ON moves.type_id = types.id LEFT JOIN pokemon_move ON moves.id = pokemon_move.move_id WHERE pokemon_move.pokemon_id IN (.*)").
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

	mock.ExpectQuery("SELECT .* FROM moves LEFT JOIN types ON moves.type_id = types.id LEFT JOIN pokemon_move ON moves.id = pokemon_move.move_id WHERE pokemon_move.pokemon_id IN (.*)").
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
				&accelerock,
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

	mock.ExpectQuery("SELECT .* FROM moves  LEFT JOIN types ON moves.type_id = types.id WHERE type_id IN (.*)").
		WithArgs(accelerock.Type.ID, "a248c127-8e9c-4f87-8513-c5dbc3385011", infestation.Type.ID).
		WillReturnRows(mockRowsForMovesByTypeIdDataLoader(false, false, false, []string{accelerock.Type.ID, "a248c127-8e9c-4f87-8513-c5dbc3385011", infestation.Type.ID}))

	got, err := NewMove(db).MovesByTypeIdDataLoader(context.Background())([]string{accelerock.Type.ID, "a248c127-8e9c-4f87-8513-c5dbc3385011", infestation.Type.ID})
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	exp := []*model.MoveList{
		{
			Total: 2,
			Moves: []*model.Move{
				&accelerock,
				&accelerock,
			},
		},
		nil,
		{
			Total: 1,
			Moves: []*model.Move{
				&infestation,
			},
		},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestMove_MovesByTypeIdDataLoader_WithQueryError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM moves LEFT JOIN types ON moves.type_id = types.id WHERE type_id IN (.*)").
		WithArgs(accelerock.Type.ID, "a248c127-8e9c-4f87-8513-c5dbc3385011", infestation.Type.ID).
		WillReturnRows(mockRowsForMovesByTypeIdDataLoader(false, false, false, []string{accelerock.Type.ID, "a248c127-8e9c-4f87-8513-c5dbc3385011", infestation.Type.ID})).
		WillReturnError(errors.New("I am Error."))

	got, err := NewMove(db).MovesByTypeIdDataLoader(context.Background())([]string{accelerock.Type.ID, "a248c127-8e9c-4f87-8513-c5dbc3385011", infestation.Type.ID})
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

	mock.ExpectQuery("SELECT .* FROM moves LEFT JOIN types ON moves.type_id = types.id WHERE type_id IN (.*)").
		WithArgs(accelerock.Type.ID, "a248c127-8e9c-4f87-8513-c5dbc3385011", infestation.Type.ID).
		WillReturnRows(mockRowsForMovesByTypeIdDataLoader(false, false, true, []string{accelerock.Type.ID, "a248c127-8e9c-4f87-8513-c5dbc3385011", infestation.Type.ID}))

	got, err := NewMove(db).MovesByTypeIdDataLoader(context.Background())([]string{accelerock.Type.ID, "a248c127-8e9c-4f87-8513-c5dbc3385011", infestation.Type.ID})
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

	mock.ExpectQuery("SELECT .* FROM moves LEFT JOIN types ON moves.type_id = types.id WHERE type_id IN (.*)").
		WithArgs(accelerock.Type.ID, "a248c127-8e9c-4f87-8513-c5dbc3385011", infestation.Type.ID).
		WillReturnRows(mockRowsForMovesByTypeIdDataLoader(false, true, false, []string{accelerock.Type.ID, "a248c127-8e9c-4f87-8513-c5dbc3385011", infestation.Type.ID}))

	got, err := NewMove(db).MovesByTypeIdDataLoader(context.Background())([]string{accelerock.Type.ID, "a248c127-8e9c-4f87-8513-c5dbc3385011", infestation.Type.ID})
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := []*model.MoveList{
		{
			Total: 1,
			Moves: []*model.Move{
				&accelerock,
			},
		},
		nil,
		nil,
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

var accelerock = model.Move{
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
	Type: model.Type{
		ID:   "5179f383-b765-4cc7-b9f9-8b1a3ba93019",
		Name: "Rock",
		Slug: "rock",
	},
}

var infestation = model.Move{
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
	Type: model.Type{
		ID:   "56dddb9a-3623-43c5-8228-ea24d598afe7",
		Name: "Bug",
		Slug: "bug",
	},
}

func mockRowsForGetMoves(empty bool, hasRowError bool, hasScanError bool) *sqlmock.Rows {
	if hasScanError {
		rows := sqlmock.NewRows([]string{"id"})
		rows.AddRow("9f61694f-34f0-4531-b5e4-aff9a3d9edde")
		return rows
	}
	rows := sqlmock.NewRows([]string{"id", "name", "slug", "accuracy", "pp", "power", "damage_class_enum", "effect", "effect_chance", "target", "types.id", "types.name", "types.slug"})
	if !empty {
		rows.AddRow(accelerock.ID, accelerock.Name, accelerock.Slug, accelerock.Accuracy, accelerock.PP, accelerock.Power, accelerock.DamageClass, accelerock.Effect, accelerock.EffectChance, accelerock.Target, accelerock.Type.ID, accelerock.Type.Name, accelerock.Type.Slug).
			AddRow(infestation.ID, infestation.Name, infestation.Slug, infestation.Accuracy, infestation.PP, infestation.Power, infestation.DamageClass, infestation.Effect, infestation.EffectChance, infestation.Target, infestation.Type.ID, infestation.Type.Name, infestation.Type.Slug)
	}
	if hasRowError {
		rows.RowError(0, errors.New("row error"))
	}
	return rows
}

func mockRowsForGetMoveById(empty bool) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"id", "name", "slug", "accuracy", "pp", "power", "damage_class_enum", "effect", "effect_chance", "target", "type_id", "types.name", "types.slug"})
	if !empty {
		rows.AddRow(accelerock.ID, accelerock.Name, accelerock.Slug, accelerock.Accuracy, accelerock.PP, accelerock.Power, accelerock.DamageClass, accelerock.Effect, accelerock.EffectChance, accelerock.Target, accelerock.Type.ID, accelerock.Type.Name, accelerock.Type.Slug)
	}
	return rows
}

func mockRowsForMovesByPokemonIdDataLoader(empty bool, hasRowError bool, hasScanError bool, ids []string) *sqlmock.Rows {
	if hasScanError {
		rows := sqlmock.NewRows([]string{"id"})
		rows.AddRow("9f61694f-34f0-4531-b5e4-aff9a3d9edde")
		return rows
	}
	rows := sqlmock.NewRows([]string{"id", "name", "slug", "accuracy", "pp", "power", "damage_class_enum", "effect", "effect_chance", "target", "type_id", "types.name", "types.slug", "pokemon_move.pokemon_id"})
	if !empty {
		rows.AddRow(accelerock.ID, accelerock.Name, accelerock.Slug, accelerock.Accuracy, accelerock.PP, accelerock.Power, accelerock.DamageClass, accelerock.Effect, accelerock.EffectChance, accelerock.Target, accelerock.Type.ID, accelerock.Type.Name, accelerock.Type.Slug, ids[0]).
			AddRow(accelerock.ID, accelerock.Name, accelerock.Slug, accelerock.Accuracy, accelerock.PP, accelerock.Power, accelerock.DamageClass, accelerock.Effect, accelerock.EffectChance, accelerock.Target, accelerock.Type.ID, accelerock.Type.Name, accelerock.Type.Slug, ids[2]).
			AddRow(infestation.ID, infestation.Name, infestation.Slug, infestation.Accuracy, infestation.PP, infestation.Power, infestation.DamageClass, infestation.Effect, infestation.EffectChance, infestation.Target, infestation.Type.ID, infestation.Type.Name, infestation.Type.Slug, ids[2])
	}
	if hasRowError {
		rows.RowError(1, errors.New("row error"))
	}
	return rows
}

func mockRowsForMovesByTypeIdDataLoader(empty bool, hasRowError bool, hasScanError bool, ids []string) *sqlmock.Rows {
	if hasScanError {
		rows := sqlmock.NewRows([]string{"id"})
		rows.AddRow(accelerock.ID)
		return rows
	}
	rows := sqlmock.NewRows([]string{"id", "name", "slug", "accuracy", "pp", "power", "damage_class_enum", "effect", "effect_chance", "target", "type_id", "types.name", "types.slug"})
	if !empty {
		rows.AddRow(accelerock.ID, accelerock.Name, accelerock.Slug, accelerock.Accuracy, accelerock.PP, accelerock.Power, accelerock.DamageClass, accelerock.Effect, accelerock.EffectChance, accelerock.Target, accelerock.Type.ID, accelerock.Type.Name, accelerock.Type.Slug).
			AddRow(accelerock.ID, accelerock.Name, accelerock.Slug, accelerock.Accuracy, accelerock.PP, accelerock.Power, accelerock.DamageClass, accelerock.Effect, accelerock.EffectChance, accelerock.Target, accelerock.Type.ID, accelerock.Type.Name, accelerock.Type.Slug).
			AddRow(infestation.ID, infestation.Name, infestation.Slug, infestation.Accuracy, infestation.PP, infestation.Power, infestation.DamageClass, infestation.Effect, infestation.EffectChance, infestation.Target, infestation.Type.ID, infestation.Type.Name, infestation.Type.Slug)
	}
	if hasRowError {
		rows.RowError(1, errors.New("row error"))
	}
	return rows
}
