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

	moves := []model.Move{
		accelerock,
		infestation,
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
		Moves: []model.Move{},
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
		Moves: []model.Move{},
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
		Moves: []model.Move{},
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
		Moves: []model.Move{},
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
			Moves: []model.Move{
				accelerock,
			},
		},
		nil,
		{
			Total: 2,
			Moves: []model.Move{
				accelerock,
				infestation,
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
			Moves: []model.Move{},
		},
		{
			Total: 0,
			Moves: []model.Move{},
		},
		{
			Total: 0,
			Moves: []model.Move{},
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
			Moves: []model.Move{},
		},
		{
			Total: 0,
			Moves: []model.Move{},
		},
		{
			Total: 0,
			Moves: []model.Move{},
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
			Moves: []model.Move{
				accelerock,
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
			Moves: []model.Move{
				accelerock,
				accelerock,
			},
		},
		nil,
		{
			Total: 1,
			Moves: []model.Move{
				infestation,
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
			Moves: []model.Move{},
		},
		{
			Total: 0,
			Moves: []model.Move{},
		},
		{
			Total: 0,
			Moves: []model.Move{},
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
			Moves: []model.Move{},
		},
		{
			Total: 0,
			Moves: []model.Move{},
		},
		{
			Total: 0,
			Moves: []model.Move{},
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
			Moves: []model.Move{
				accelerock,
			},
		},
		nil,
		nil,
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestMove_MovesByIdDataLoader(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM moves LEFT JOIN .* WHERE moves.id IN (.*)").
		WithArgs(accelerock.ID, "a248c127-8e9c-4f87-8513-c5dbc3385011", infestation.ID).
		WillReturnRows(mockRowsForMovesByIdDataLoader(false, false, false, []string{accelerock.ID, "a248c127-8e9c-4f87-8513-c5dbc3385011", infestation.ID}))

	got, err := NewMove(db).MovesByIdDataLoader(context.Background())([]string{accelerock.ID, "a248c127-8e9c-4f87-8513-c5dbc3385011", infestation.ID})
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	exp := []*model.Move{
		&accelerock,
		nil,
		&infestation,
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestMove_MovesByIdDataLoader_WithQueryError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM moves LEFT JOIN .* WHERE moves.id IN (.*)").
		WithArgs(accelerock.ID, "a248c127-8e9c-4f87-8513-c5dbc3385011", infestation.ID).
		WillReturnRows(mockRowsForMovesByIdDataLoader(false, false, false, []string{accelerock.ID, "a248c127-8e9c-4f87-8513-c5dbc3385011", infestation.ID})).
		WillReturnError(errors.New("I am Error."))

	got, err := NewMove(db).MovesByIdDataLoader(context.Background())([]string{accelerock.ID, "a248c127-8e9c-4f87-8513-c5dbc3385011", infestation.ID})
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := []*model.Move{
		nil,
		nil,
		nil,
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestMove_MovesByIdDataLoader_WithScanError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM moves LEFT JOIN .* WHERE moves.id IN (.*)").
		WithArgs(accelerock.ID, "a248c127-8e9c-4f87-8513-c5dbc3385011", infestation.ID).
		WillReturnRows(mockRowsForMovesByIdDataLoader(false, false, true, []string{accelerock.ID, "a248c127-8e9c-4f87-8513-c5dbc3385011", infestation.ID}))

	got, err := NewMove(db).MovesByIdDataLoader(context.Background())([]string{accelerock.ID, "a248c127-8e9c-4f87-8513-c5dbc3385011", infestation.ID})
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := []*model.Move{
		nil,
		nil,
		nil,
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestMove_MovesByIdDataLoader_WithRowError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM moves LEFT JOIN .* WHERE moves.id IN (.*)").
		WithArgs(accelerock.ID, "a248c127-8e9c-4f87-8513-c5dbc3385011", infestation.ID).
		WillReturnRows(mockRowsForMovesByIdDataLoader(false, true, false, []string{accelerock.ID, "a248c127-8e9c-4f87-8513-c5dbc3385011", infestation.ID}))

	got, err := NewMove(db).MovesByIdDataLoader(context.Background())([]string{accelerock.ID, "a248c127-8e9c-4f87-8513-c5dbc3385011", infestation.ID})
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := []*model.Move{
		&accelerock,
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
		ID:   Rock.ID,
		Name: Rock.Name,
		Slug: Rock.Slug,
		NoDamageTo: model.TypeList{
			Total: 0,
			Types: nil,
		},
		HalfDamageTo: model.TypeList{
			Total: 3,
			Types: []model.Type{
				Steel,
				Ground,
				Fighting,
			},
		},
		DoubleDamageTo: model.TypeList{
			Total: 4,
			Types: []model.Type{
				Ice,
				Flying,
				Bug,
				Fire,
			},
		},
		NoDamageFrom: model.TypeList{
			Total: 0,
			Types: nil,
		},
		HalfDamageFrom: model.TypeList{
			Total: 4,
			Types: []model.Type{
				Normal,
				Poison,
				Flying,
				Fire,
			},
		},
		DoubleDamageFrom: model.TypeList{
			Total: 5,
			Types: []model.Type{
				Steel,
				Grass,
				Ground,
				Fighting,
				Water,
			},
		},
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
		ID:   Bug.ID,
		Name: Bug.Name,
		Slug: Bug.Slug,
		NoDamageTo: model.TypeList{
			Total: 0,
			Types: nil,
		},
		HalfDamageTo: model.TypeList{
			Total: 7,
			Types: []model.Type{
				Ghost,
				Steel,
				Poison,
				Flying,
				Fighting,
				Fairy,
				Fire,
			},
		},
		DoubleDamageTo: model.TypeList{
			Total: 3,
			Types: []model.Type{
				Grass,
				Psychic,
				Dark,
			},
		},
		NoDamageFrom: model.TypeList{
			Total: 0,
			Types: nil,
		},
		HalfDamageFrom: model.TypeList{
			Total: 3,
			Types: []model.Type{
				Grass,
				Ground,
				Fighting,
			},
		},
		DoubleDamageFrom: model.TypeList{
			Total: 3,
			Types: []model.Type{
				Flying,
				Rock,
				Fire,
			},
		},
	},
}

func mockRowsForGetMoves(empty bool, hasRowError bool, hasScanError bool) *sqlmock.Rows {
	if hasScanError {
		rows := sqlmock.NewRows([]string{"id"})
		rows.AddRow("9f61694f-34f0-4531-b5e4-aff9a3d9edde")
		return rows
	}
	rows := sqlmock.NewRows([]string{"id", "name", "slug", "accuracy", "pp", "power", "damage_class_enum", "effect", "effect_chance", "target", "types.id", "types.name", "types.slug", "no_damage_to", "half_damage_to", "double_damage_to", "no_damage_from", "half_damage_from", "double_damage_from"})
	if !empty {
		rows.AddRow(accelerock.ID, accelerock.Name, accelerock.Slug, accelerock.Accuracy, accelerock.PP, accelerock.Power, accelerock.DamageClass, accelerock.Effect, accelerock.EffectChance, accelerock.Target, accelerock.Type.ID, accelerock.Type.Name, accelerock.Type.Slug, nil, `{"{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}","{\"id\": \"1b7d7950-305a-48fa-a771-01f7bc4dad8d\", \"name\": \"Ground\", \"slug\": \"ground\"}","{\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}"}`, `{"{\"id\": \"1dcc9d3c-55d4-4d33-809a-d1580c6e6542\", \"name\": \"Ice\", \"slug\": \"ice\"}","{\"id\": \"4f09ea3c-2d93-4908-aabc-bc6e04ff24bb\", \"name\": \"Flying\", \"slug\": \"flying\"}","{\"id\": \"56dddb9a-3623-43c5-8228-ea24d598afe7\", \"name\": \"Bug\", \"slug\": \"bug\"}","{\"id\": \"d43f57ab-e5b3-4912-a667-9f237d21d391\", \"name\": \"Fire\", \"slug\": \"fire\"}"}`, nil, `{"{\"id\": \"366a8621-9fa7-419b-b710-9100bcbb98d8\", \"name\": \"Normal\", \"slug\": \"normal\"}","{\"id\": \"42b31825-de68-4c1c-bea1-b32a290f1fef\", \"name\": \"Poison\", \"slug\": \"poison\"}","{\"id\": \"4f09ea3c-2d93-4908-aabc-bc6e04ff24bb\", \"name\": \"Flying\", \"slug\": \"flying\"}","{\"id\": \"d43f57ab-e5b3-4912-a667-9f237d21d391\", \"name\": \"Fire\", \"slug\": \"fire\"}"}`, `{"{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}","{\"id\": \"07b9eb0f-e676-4649-bf2e-0e5ef2c2c2e3\", \"name\": \"Grass\", \"slug\": \"grass\"}","{\"id\": \"1b7d7950-305a-48fa-a771-01f7bc4dad8d\", \"name\": \"Ground\", \"slug\": \"ground\"}","{\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}","{\"id\": \"de384e1c-89aa-44de-88fe-5e914e468f2b\", \"name\": \"Water\", \"slug\": \"water\"}"}`).
			AddRow(infestation.ID, infestation.Name, infestation.Slug, infestation.Accuracy, infestation.PP, infestation.Power, infestation.DamageClass, infestation.Effect, infestation.EffectChance, infestation.Target, infestation.Type.ID, infestation.Type.Name, infestation.Type.Slug, nil, `{"{\"id\": \"027f1455-8e6f-4891-8c62-d75bb6c49dae\", \"name\": \"Ghost\", \"slug\": \"ghost\"}","{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}","{\"id\": \"42b31825-de68-4c1c-bea1-b32a290f1fef\", \"name\": \"Poison\", \"slug\": \"poison\"}","{\"id\": \"4f09ea3c-2d93-4908-aabc-bc6e04ff24bb\", \"name\": \"Flying\", \"slug\": \"flying\"}","{\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}","{\"id\": \"a248c127-8e9c-4f87-8513-c5dbc3385011\", \"name\": \"Fairy\", \"slug\": \"fairy\"}","{\"id\": \"d43f57ab-e5b3-4912-a667-9f237d21d391\", \"name\": \"Fire\", \"slug\": \"fire\"}"}`, `{"{\"id\": \"07b9eb0f-e676-4649-bf2e-0e5ef2c2c2e3\", \"name\": \"Grass\", \"slug\": \"grass\"}","{\"id\": \"2222c839-3c6e-4727-b6b5-a946bb8af5fa\", \"name\": \"Psychic\", \"slug\": \"psychic\"}","{\"id\": \"9ca47516-fff8-4f5e-8eb5-582c1f7c05af\", \"name\": \"Dark\", \"slug\": \"dark\"}"}`, nil, `{"{\"id\": \"07b9eb0f-e676-4649-bf2e-0e5ef2c2c2e3\", \"name\": \"Grass\", \"slug\": \"grass\"}","{\"id\": \"1b7d7950-305a-48fa-a771-01f7bc4dad8d\", \"name\": \"Ground\", \"slug\": \"ground\"}","{\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}"}`, `{"{\"id\": \"4f09ea3c-2d93-4908-aabc-bc6e04ff24bb\", \"name\": \"Flying\", \"slug\": \"flying\"}","{\"id\": \"5179f383-b765-4cc7-b9f9-8b1a3ba93019\", \"name\": \"Rock\", \"slug\": \"rock\"}","{\"id\": \"d43f57ab-e5b3-4912-a667-9f237d21d391\", \"name\": \"Fire\", \"slug\": \"fire\"}"}`)
	}
	if hasRowError {
		rows.RowError(0, errors.New("row error"))
	}
	return rows
}

func mockRowsForGetMoveById(empty bool) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"id", "name", "slug", "accuracy", "pp", "power", "damage_class_enum", "effect", "effect_chance", "target", "types.id", "types.name", "types.slug", "no_damage_to", "half_damage_to", "double_damage_to", "no_damage_from", "half_damage_from", "double_damage_from"})
	if !empty {
		rows.AddRow(accelerock.ID, accelerock.Name, accelerock.Slug, accelerock.Accuracy, accelerock.PP, accelerock.Power, accelerock.DamageClass, accelerock.Effect, accelerock.EffectChance, accelerock.Target, accelerock.Type.ID, accelerock.Type.Name, accelerock.Type.Slug, nil, `{"{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}","{\"id\": \"1b7d7950-305a-48fa-a771-01f7bc4dad8d\", \"name\": \"Ground\", \"slug\": \"ground\"}","{\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}"}`, `{"{\"id\": \"1dcc9d3c-55d4-4d33-809a-d1580c6e6542\", \"name\": \"Ice\", \"slug\": \"ice\"}","{\"id\": \"4f09ea3c-2d93-4908-aabc-bc6e04ff24bb\", \"name\": \"Flying\", \"slug\": \"flying\"}","{\"id\": \"56dddb9a-3623-43c5-8228-ea24d598afe7\", \"name\": \"Bug\", \"slug\": \"bug\"}","{\"id\": \"d43f57ab-e5b3-4912-a667-9f237d21d391\", \"name\": \"Fire\", \"slug\": \"fire\"}"}`, nil, `{"{\"id\": \"366a8621-9fa7-419b-b710-9100bcbb98d8\", \"name\": \"Normal\", \"slug\": \"normal\"}","{\"id\": \"42b31825-de68-4c1c-bea1-b32a290f1fef\", \"name\": \"Poison\", \"slug\": \"poison\"}","{\"id\": \"4f09ea3c-2d93-4908-aabc-bc6e04ff24bb\", \"name\": \"Flying\", \"slug\": \"flying\"}","{\"id\": \"d43f57ab-e5b3-4912-a667-9f237d21d391\", \"name\": \"Fire\", \"slug\": \"fire\"}"}`, `{"{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}","{\"id\": \"07b9eb0f-e676-4649-bf2e-0e5ef2c2c2e3\", \"name\": \"Grass\", \"slug\": \"grass\"}","{\"id\": \"1b7d7950-305a-48fa-a771-01f7bc4dad8d\", \"name\": \"Ground\", \"slug\": \"ground\"}","{\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}","{\"id\": \"de384e1c-89aa-44de-88fe-5e914e468f2b\", \"name\": \"Water\", \"slug\": \"water\"}"}`)
	}
	return rows
}

func mockRowsForMovesByPokemonIdDataLoader(empty bool, hasRowError bool, hasScanError bool, ids []string) *sqlmock.Rows {
	if hasScanError {
		rows := sqlmock.NewRows([]string{"id"})
		rows.AddRow("9f61694f-34f0-4531-b5e4-aff9a3d9edde")
		return rows
	}
	rows := sqlmock.NewRows([]string{"id", "name", "slug", "accuracy", "pp", "power", "damage_class_enum", "effect", "effect_chance", "target", "types.id", "types.name", "types.slug", "no_damage_to", "half_damage_to", "double_damage_to", "no_damage_from", "half_damage_from", "double_damage_from", "pokemon_id"})
	if !empty {
		rows.AddRow(accelerock.ID, accelerock.Name, accelerock.Slug, accelerock.Accuracy, accelerock.PP, accelerock.Power, accelerock.DamageClass, accelerock.Effect, accelerock.EffectChance, accelerock.Target, accelerock.Type.ID, accelerock.Type.Name, accelerock.Type.Slug, nil, `{"{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}","{\"id\": \"1b7d7950-305a-48fa-a771-01f7bc4dad8d\", \"name\": \"Ground\", \"slug\": \"ground\"}","{\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}"}`, `{"{\"id\": \"1dcc9d3c-55d4-4d33-809a-d1580c6e6542\", \"name\": \"Ice\", \"slug\": \"ice\"}","{\"id\": \"4f09ea3c-2d93-4908-aabc-bc6e04ff24bb\", \"name\": \"Flying\", \"slug\": \"flying\"}","{\"id\": \"56dddb9a-3623-43c5-8228-ea24d598afe7\", \"name\": \"Bug\", \"slug\": \"bug\"}","{\"id\": \"d43f57ab-e5b3-4912-a667-9f237d21d391\", \"name\": \"Fire\", \"slug\": \"fire\"}"}`, nil, `{"{\"id\": \"366a8621-9fa7-419b-b710-9100bcbb98d8\", \"name\": \"Normal\", \"slug\": \"normal\"}","{\"id\": \"42b31825-de68-4c1c-bea1-b32a290f1fef\", \"name\": \"Poison\", \"slug\": \"poison\"}","{\"id\": \"4f09ea3c-2d93-4908-aabc-bc6e04ff24bb\", \"name\": \"Flying\", \"slug\": \"flying\"}","{\"id\": \"d43f57ab-e5b3-4912-a667-9f237d21d391\", \"name\": \"Fire\", \"slug\": \"fire\"}"}`, `{"{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}","{\"id\": \"07b9eb0f-e676-4649-bf2e-0e5ef2c2c2e3\", \"name\": \"Grass\", \"slug\": \"grass\"}","{\"id\": \"1b7d7950-305a-48fa-a771-01f7bc4dad8d\", \"name\": \"Ground\", \"slug\": \"ground\"}","{\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}","{\"id\": \"de384e1c-89aa-44de-88fe-5e914e468f2b\", \"name\": \"Water\", \"slug\": \"water\"}"}`, ids[0]).
			AddRow(accelerock.ID, accelerock.Name, accelerock.Slug, accelerock.Accuracy, accelerock.PP, accelerock.Power, accelerock.DamageClass, accelerock.Effect, accelerock.EffectChance, accelerock.Target, accelerock.Type.ID, accelerock.Type.Name, accelerock.Type.Slug, nil, `{"{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}","{\"id\": \"1b7d7950-305a-48fa-a771-01f7bc4dad8d\", \"name\": \"Ground\", \"slug\": \"ground\"}","{\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}"}`, `{"{\"id\": \"1dcc9d3c-55d4-4d33-809a-d1580c6e6542\", \"name\": \"Ice\", \"slug\": \"ice\"}","{\"id\": \"4f09ea3c-2d93-4908-aabc-bc6e04ff24bb\", \"name\": \"Flying\", \"slug\": \"flying\"}","{\"id\": \"56dddb9a-3623-43c5-8228-ea24d598afe7\", \"name\": \"Bug\", \"slug\": \"bug\"}","{\"id\": \"d43f57ab-e5b3-4912-a667-9f237d21d391\", \"name\": \"Fire\", \"slug\": \"fire\"}"}`, nil, `{"{\"id\": \"366a8621-9fa7-419b-b710-9100bcbb98d8\", \"name\": \"Normal\", \"slug\": \"normal\"}","{\"id\": \"42b31825-de68-4c1c-bea1-b32a290f1fef\", \"name\": \"Poison\", \"slug\": \"poison\"}","{\"id\": \"4f09ea3c-2d93-4908-aabc-bc6e04ff24bb\", \"name\": \"Flying\", \"slug\": \"flying\"}","{\"id\": \"d43f57ab-e5b3-4912-a667-9f237d21d391\", \"name\": \"Fire\", \"slug\": \"fire\"}"}`, `{"{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}","{\"id\": \"07b9eb0f-e676-4649-bf2e-0e5ef2c2c2e3\", \"name\": \"Grass\", \"slug\": \"grass\"}","{\"id\": \"1b7d7950-305a-48fa-a771-01f7bc4dad8d\", \"name\": \"Ground\", \"slug\": \"ground\"}","{\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}","{\"id\": \"de384e1c-89aa-44de-88fe-5e914e468f2b\", \"name\": \"Water\", \"slug\": \"water\"}"}`, ids[2]).
			AddRow(infestation.ID, infestation.Name, infestation.Slug, infestation.Accuracy, infestation.PP, infestation.Power, infestation.DamageClass, infestation.Effect, infestation.EffectChance, infestation.Target, infestation.Type.ID, infestation.Type.Name, infestation.Type.Slug, nil, `{"{\"id\": \"027f1455-8e6f-4891-8c62-d75bb6c49dae\", \"name\": \"Ghost\", \"slug\": \"ghost\"}","{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}","{\"id\": \"42b31825-de68-4c1c-bea1-b32a290f1fef\", \"name\": \"Poison\", \"slug\": \"poison\"}","{\"id\": \"4f09ea3c-2d93-4908-aabc-bc6e04ff24bb\", \"name\": \"Flying\", \"slug\": \"flying\"}","{\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}","{\"id\": \"a248c127-8e9c-4f87-8513-c5dbc3385011\", \"name\": \"Fairy\", \"slug\": \"fairy\"}","{\"id\": \"d43f57ab-e5b3-4912-a667-9f237d21d391\", \"name\": \"Fire\", \"slug\": \"fire\"}"}`, `{"{\"id\": \"07b9eb0f-e676-4649-bf2e-0e5ef2c2c2e3\", \"name\": \"Grass\", \"slug\": \"grass\"}","{\"id\": \"2222c839-3c6e-4727-b6b5-a946bb8af5fa\", \"name\": \"Psychic\", \"slug\": \"psychic\"}","{\"id\": \"9ca47516-fff8-4f5e-8eb5-582c1f7c05af\", \"name\": \"Dark\", \"slug\": \"dark\"}"}`, nil, `{"{\"id\": \"07b9eb0f-e676-4649-bf2e-0e5ef2c2c2e3\", \"name\": \"Grass\", \"slug\": \"grass\"}","{\"id\": \"1b7d7950-305a-48fa-a771-01f7bc4dad8d\", \"name\": \"Ground\", \"slug\": \"ground\"}","{\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}"}`, `{"{\"id\": \"4f09ea3c-2d93-4908-aabc-bc6e04ff24bb\", \"name\": \"Flying\", \"slug\": \"flying\"}","{\"id\": \"5179f383-b765-4cc7-b9f9-8b1a3ba93019\", \"name\": \"Rock\", \"slug\": \"rock\"}","{\"id\": \"d43f57ab-e5b3-4912-a667-9f237d21d391\", \"name\": \"Fire\", \"slug\": \"fire\"}"}`, ids[2])
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
	rows := sqlmock.NewRows([]string{"id", "name", "slug", "accuracy", "pp", "power", "damage_class_enum", "effect", "effect_chance", "target", "types.id", "types.name", "types.slug", "no_damage_to", "half_damage_to", "double_damage_to", "no_damage_from", "half_damage_from", "double_damage_from"})
	if !empty {
		rows.AddRow(accelerock.ID, accelerock.Name, accelerock.Slug, accelerock.Accuracy, accelerock.PP, accelerock.Power, accelerock.DamageClass, accelerock.Effect, accelerock.EffectChance, accelerock.Target, accelerock.Type.ID, accelerock.Type.Name, accelerock.Type.Slug, nil, `{"{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}","{\"id\": \"1b7d7950-305a-48fa-a771-01f7bc4dad8d\", \"name\": \"Ground\", \"slug\": \"ground\"}","{\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}"}`, `{"{\"id\": \"1dcc9d3c-55d4-4d33-809a-d1580c6e6542\", \"name\": \"Ice\", \"slug\": \"ice\"}","{\"id\": \"4f09ea3c-2d93-4908-aabc-bc6e04ff24bb\", \"name\": \"Flying\", \"slug\": \"flying\"}","{\"id\": \"56dddb9a-3623-43c5-8228-ea24d598afe7\", \"name\": \"Bug\", \"slug\": \"bug\"}","{\"id\": \"d43f57ab-e5b3-4912-a667-9f237d21d391\", \"name\": \"Fire\", \"slug\": \"fire\"}"}`, nil, `{"{\"id\": \"366a8621-9fa7-419b-b710-9100bcbb98d8\", \"name\": \"Normal\", \"slug\": \"normal\"}","{\"id\": \"42b31825-de68-4c1c-bea1-b32a290f1fef\", \"name\": \"Poison\", \"slug\": \"poison\"}","{\"id\": \"4f09ea3c-2d93-4908-aabc-bc6e04ff24bb\", \"name\": \"Flying\", \"slug\": \"flying\"}","{\"id\": \"d43f57ab-e5b3-4912-a667-9f237d21d391\", \"name\": \"Fire\", \"slug\": \"fire\"}"}`, `{"{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}","{\"id\": \"07b9eb0f-e676-4649-bf2e-0e5ef2c2c2e3\", \"name\": \"Grass\", \"slug\": \"grass\"}","{\"id\": \"1b7d7950-305a-48fa-a771-01f7bc4dad8d\", \"name\": \"Ground\", \"slug\": \"ground\"}","{\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}","{\"id\": \"de384e1c-89aa-44de-88fe-5e914e468f2b\", \"name\": \"Water\", \"slug\": \"water\"}"}`).
			AddRow(accelerock.ID, accelerock.Name, accelerock.Slug, accelerock.Accuracy, accelerock.PP, accelerock.Power, accelerock.DamageClass, accelerock.Effect, accelerock.EffectChance, accelerock.Target, accelerock.Type.ID, accelerock.Type.Name, accelerock.Type.Slug, nil, `{"{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}","{\"id\": \"1b7d7950-305a-48fa-a771-01f7bc4dad8d\", \"name\": \"Ground\", \"slug\": \"ground\"}","{\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}"}`, `{"{\"id\": \"1dcc9d3c-55d4-4d33-809a-d1580c6e6542\", \"name\": \"Ice\", \"slug\": \"ice\"}","{\"id\": \"4f09ea3c-2d93-4908-aabc-bc6e04ff24bb\", \"name\": \"Flying\", \"slug\": \"flying\"}","{\"id\": \"56dddb9a-3623-43c5-8228-ea24d598afe7\", \"name\": \"Bug\", \"slug\": \"bug\"}","{\"id\": \"d43f57ab-e5b3-4912-a667-9f237d21d391\", \"name\": \"Fire\", \"slug\": \"fire\"}"}`, nil, `{"{\"id\": \"366a8621-9fa7-419b-b710-9100bcbb98d8\", \"name\": \"Normal\", \"slug\": \"normal\"}","{\"id\": \"42b31825-de68-4c1c-bea1-b32a290f1fef\", \"name\": \"Poison\", \"slug\": \"poison\"}","{\"id\": \"4f09ea3c-2d93-4908-aabc-bc6e04ff24bb\", \"name\": \"Flying\", \"slug\": \"flying\"}","{\"id\": \"d43f57ab-e5b3-4912-a667-9f237d21d391\", \"name\": \"Fire\", \"slug\": \"fire\"}"}`, `{"{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}","{\"id\": \"07b9eb0f-e676-4649-bf2e-0e5ef2c2c2e3\", \"name\": \"Grass\", \"slug\": \"grass\"}","{\"id\": \"1b7d7950-305a-48fa-a771-01f7bc4dad8d\", \"name\": \"Ground\", \"slug\": \"ground\"}","{\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}","{\"id\": \"de384e1c-89aa-44de-88fe-5e914e468f2b\", \"name\": \"Water\", \"slug\": \"water\"}"}`).
			AddRow(infestation.ID, infestation.Name, infestation.Slug, infestation.Accuracy, infestation.PP, infestation.Power, infestation.DamageClass, infestation.Effect, infestation.EffectChance, infestation.Target, infestation.Type.ID, infestation.Type.Name, infestation.Type.Slug, nil, `{"{\"id\": \"027f1455-8e6f-4891-8c62-d75bb6c49dae\", \"name\": \"Ghost\", \"slug\": \"ghost\"}","{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}","{\"id\": \"42b31825-de68-4c1c-bea1-b32a290f1fef\", \"name\": \"Poison\", \"slug\": \"poison\"}","{\"id\": \"4f09ea3c-2d93-4908-aabc-bc6e04ff24bb\", \"name\": \"Flying\", \"slug\": \"flying\"}","{\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}","{\"id\": \"a248c127-8e9c-4f87-8513-c5dbc3385011\", \"name\": \"Fairy\", \"slug\": \"fairy\"}","{\"id\": \"d43f57ab-e5b3-4912-a667-9f237d21d391\", \"name\": \"Fire\", \"slug\": \"fire\"}"}`, `{"{\"id\": \"07b9eb0f-e676-4649-bf2e-0e5ef2c2c2e3\", \"name\": \"Grass\", \"slug\": \"grass\"}","{\"id\": \"2222c839-3c6e-4727-b6b5-a946bb8af5fa\", \"name\": \"Psychic\", \"slug\": \"psychic\"}","{\"id\": \"9ca47516-fff8-4f5e-8eb5-582c1f7c05af\", \"name\": \"Dark\", \"slug\": \"dark\"}"}`, nil, `{"{\"id\": \"07b9eb0f-e676-4649-bf2e-0e5ef2c2c2e3\", \"name\": \"Grass\", \"slug\": \"grass\"}","{\"id\": \"1b7d7950-305a-48fa-a771-01f7bc4dad8d\", \"name\": \"Ground\", \"slug\": \"ground\"}","{\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}"}`, `{"{\"id\": \"4f09ea3c-2d93-4908-aabc-bc6e04ff24bb\", \"name\": \"Flying\", \"slug\": \"flying\"}","{\"id\": \"5179f383-b765-4cc7-b9f9-8b1a3ba93019\", \"name\": \"Rock\", \"slug\": \"rock\"}","{\"id\": \"d43f57ab-e5b3-4912-a667-9f237d21d391\", \"name\": \"Fire\", \"slug\": \"fire\"}"}`)
	}
	if hasRowError {
		rows.RowError(1, errors.New("row error"))
	}
	return rows
}

func mockRowsForMovesByIdDataLoader(empty bool, hasRowError bool, hasScanError bool, ids []string) *sqlmock.Rows {
	if hasScanError {
		rows := sqlmock.NewRows([]string{"id"})
		rows.AddRow("9f61694f-34f0-4531-b5e4-aff9a3d9edde")
		return rows
	}
	rows := sqlmock.NewRows([]string{"id", "name", "slug", "accuracy", "pp", "power", "damage_class_enum", "effect", "effect_chance", "target", "types.id", "types.name", "types.slug", "no_damage_to", "half_damage_to", "double_damage_to", "no_damage_from", "half_damage_from", "double_damage_from"})
	if !empty {
		rows.AddRow(ids[0], accelerock.Name, accelerock.Slug, accelerock.Accuracy, accelerock.PP, accelerock.Power, accelerock.DamageClass, accelerock.Effect, accelerock.EffectChance, accelerock.Target, accelerock.Type.ID, accelerock.Type.Name, accelerock.Type.Slug, nil, `{"{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}","{\"id\": \"1b7d7950-305a-48fa-a771-01f7bc4dad8d\", \"name\": \"Ground\", \"slug\": \"ground\"}","{\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}"}`, `{"{\"id\": \"1dcc9d3c-55d4-4d33-809a-d1580c6e6542\", \"name\": \"Ice\", \"slug\": \"ice\"}","{\"id\": \"4f09ea3c-2d93-4908-aabc-bc6e04ff24bb\", \"name\": \"Flying\", \"slug\": \"flying\"}","{\"id\": \"56dddb9a-3623-43c5-8228-ea24d598afe7\", \"name\": \"Bug\", \"slug\": \"bug\"}","{\"id\": \"d43f57ab-e5b3-4912-a667-9f237d21d391\", \"name\": \"Fire\", \"slug\": \"fire\"}"}`, nil, `{"{\"id\": \"366a8621-9fa7-419b-b710-9100bcbb98d8\", \"name\": \"Normal\", \"slug\": \"normal\"}","{\"id\": \"42b31825-de68-4c1c-bea1-b32a290f1fef\", \"name\": \"Poison\", \"slug\": \"poison\"}","{\"id\": \"4f09ea3c-2d93-4908-aabc-bc6e04ff24bb\", \"name\": \"Flying\", \"slug\": \"flying\"}","{\"id\": \"d43f57ab-e5b3-4912-a667-9f237d21d391\", \"name\": \"Fire\", \"slug\": \"fire\"}"}`, `{"{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}","{\"id\": \"07b9eb0f-e676-4649-bf2e-0e5ef2c2c2e3\", \"name\": \"Grass\", \"slug\": \"grass\"}","{\"id\": \"1b7d7950-305a-48fa-a771-01f7bc4dad8d\", \"name\": \"Ground\", \"slug\": \"ground\"}","{\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}","{\"id\": \"de384e1c-89aa-44de-88fe-5e914e468f2b\", \"name\": \"Water\", \"slug\": \"water\"}"}`).
			AddRow(ids[0], accelerock.Name, accelerock.Slug, accelerock.Accuracy, accelerock.PP, accelerock.Power, accelerock.DamageClass, accelerock.Effect, accelerock.EffectChance, accelerock.Target, accelerock.Type.ID, accelerock.Type.Name, accelerock.Type.Slug, nil, `{"{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}","{\"id\": \"1b7d7950-305a-48fa-a771-01f7bc4dad8d\", \"name\": \"Ground\", \"slug\": \"ground\"}","{\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}"}`, `{"{\"id\": \"1dcc9d3c-55d4-4d33-809a-d1580c6e6542\", \"name\": \"Ice\", \"slug\": \"ice\"}","{\"id\": \"4f09ea3c-2d93-4908-aabc-bc6e04ff24bb\", \"name\": \"Flying\", \"slug\": \"flying\"}","{\"id\": \"56dddb9a-3623-43c5-8228-ea24d598afe7\", \"name\": \"Bug\", \"slug\": \"bug\"}","{\"id\": \"d43f57ab-e5b3-4912-a667-9f237d21d391\", \"name\": \"Fire\", \"slug\": \"fire\"}"}`, nil, `{"{\"id\": \"366a8621-9fa7-419b-b710-9100bcbb98d8\", \"name\": \"Normal\", \"slug\": \"normal\"}","{\"id\": \"42b31825-de68-4c1c-bea1-b32a290f1fef\", \"name\": \"Poison\", \"slug\": \"poison\"}","{\"id\": \"4f09ea3c-2d93-4908-aabc-bc6e04ff24bb\", \"name\": \"Flying\", \"slug\": \"flying\"}","{\"id\": \"d43f57ab-e5b3-4912-a667-9f237d21d391\", \"name\": \"Fire\", \"slug\": \"fire\"}"}`, `{"{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}","{\"id\": \"07b9eb0f-e676-4649-bf2e-0e5ef2c2c2e3\", \"name\": \"Grass\", \"slug\": \"grass\"}","{\"id\": \"1b7d7950-305a-48fa-a771-01f7bc4dad8d\", \"name\": \"Ground\", \"slug\": \"ground\"}","{\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}","{\"id\": \"de384e1c-89aa-44de-88fe-5e914e468f2b\", \"name\": \"Water\", \"slug\": \"water\"}"}`).
			AddRow(ids[2], infestation.Name, infestation.Slug, infestation.Accuracy, infestation.PP, infestation.Power, infestation.DamageClass, infestation.Effect, infestation.EffectChance, infestation.Target, infestation.Type.ID, infestation.Type.Name, infestation.Type.Slug, nil, `{"{\"id\": \"027f1455-8e6f-4891-8c62-d75bb6c49dae\", \"name\": \"Ghost\", \"slug\": \"ghost\"}","{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}","{\"id\": \"42b31825-de68-4c1c-bea1-b32a290f1fef\", \"name\": \"Poison\", \"slug\": \"poison\"}","{\"id\": \"4f09ea3c-2d93-4908-aabc-bc6e04ff24bb\", \"name\": \"Flying\", \"slug\": \"flying\"}","{\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}","{\"id\": \"a248c127-8e9c-4f87-8513-c5dbc3385011\", \"name\": \"Fairy\", \"slug\": \"fairy\"}","{\"id\": \"d43f57ab-e5b3-4912-a667-9f237d21d391\", \"name\": \"Fire\", \"slug\": \"fire\"}"}`, `{"{\"id\": \"07b9eb0f-e676-4649-bf2e-0e5ef2c2c2e3\", \"name\": \"Grass\", \"slug\": \"grass\"}","{\"id\": \"2222c839-3c6e-4727-b6b5-a946bb8af5fa\", \"name\": \"Psychic\", \"slug\": \"psychic\"}","{\"id\": \"9ca47516-fff8-4f5e-8eb5-582c1f7c05af\", \"name\": \"Dark\", \"slug\": \"dark\"}"}`, nil, `{"{\"id\": \"07b9eb0f-e676-4649-bf2e-0e5ef2c2c2e3\", \"name\": \"Grass\", \"slug\": \"grass\"}","{\"id\": \"1b7d7950-305a-48fa-a771-01f7bc4dad8d\", \"name\": \"Ground\", \"slug\": \"ground\"}","{\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}"}`, `{"{\"id\": \"4f09ea3c-2d93-4908-aabc-bc6e04ff24bb\", \"name\": \"Flying\", \"slug\": \"flying\"}","{\"id\": \"5179f383-b765-4cc7-b9f9-8b1a3ba93019\", \"name\": \"Rock\", \"slug\": \"rock\"}","{\"id\": \"d43f57ab-e5b3-4912-a667-9f237d21d391\", \"name\": \"Fire\", \"slug\": \"fire\"}"}`)
	}
	if hasRowError {
		rows.RowError(1, errors.New("row error"))
	}
	return rows
}
