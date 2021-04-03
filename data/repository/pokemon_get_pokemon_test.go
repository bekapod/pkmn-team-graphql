package repository

import (
	"bekapod/pkmn-team-graphql/data/model"
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-test/deep"
)

func TestPokemon_GetPokemon(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM pokemon LEFT JOIN .* GROUP BY pokemon.id ORDER BY pokedex_id, pokemon.slug ASC").
		WillReturnRows(mockRowsForGetPokemon(false, false, false))

	pokemon := []*model.Pokemon{
		&castform,
		&snorunt,
		&bronzong,
	}

	exp := model.NewPokemonList(pokemon)
	got, err := NewPokemon(db).GetPokemon(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if diff := deep.Equal(&exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestPokemon_GetPokemon_WithQueryError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM pokemon LEFT JOIN .* GROUP BY pokemon.id ORDER BY pokedex_id, pokemon.slug ASC").
		WillReturnError(errors.New("I am Error."))

	got, err := NewPokemon(db).GetPokemon(context.Background())
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := &model.PokemonList{
		Total:   0,
		Pokemon: []*model.Pokemon{},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestMove_GetPokemon_WithScanError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM pokemon LEFT JOIN .* GROUP BY pokemon.id ORDER BY pokedex_id, pokemon.slug ASC").
		WillReturnRows(mockRowsForGetPokemon(false, false, true))

	got, err := NewPokemon(db).GetPokemon(context.Background())
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := &model.PokemonList{
		Total:   0,
		Pokemon: []*model.Pokemon{},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestMove_GetPokemon_WithRowError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM pokemon LEFT JOIN .* GROUP BY pokemon.id ORDER BY pokedex_id, pokemon.slug ASC").
		WillReturnRows(mockRowsForGetPokemon(false, true, false))

	got, err := NewPokemon(db).GetPokemon(context.Background())
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := &model.PokemonList{
		Total:   0,
		Pokemon: []*model.Pokemon{},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestPokemon_GetPokemon_WithNoRows(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM pokemon LEFT JOIN .* GROUP BY pokemon.id ORDER BY pokedex_id, pokemon.slug ASC").
		WillReturnRows(mockRowsForGetMoves(true, false, false))

	got, err := NewPokemon(db).GetPokemon(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	exp := &model.PokemonList{
		Total:   0,
		Pokemon: []*model.Pokemon{},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestPokemon_GetPokemonById(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM pokemon LEFT JOIN .* WHERE pokemon.id.* GROUP BY pokemon.id").
		WithArgs("3ab43625-a18d-4b11-98a3-86d7d959fbe1").
		WillReturnRows(mockRowsForGetPokemonById(false))

	exp := castform
	got, err := NewPokemon(db).GetPokemonById(context.Background(), "3ab43625-a18d-4b11-98a3-86d7d959fbe1")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if diff := deep.Equal(&exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestPokemon_GetPokemonById_WithQueryError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM pokemon LEFT JOIN .* WHERE pokemon.id.* GROUP BY pokemon.id").
		WithArgs("3ab43625-a18d-4b11-98a3-86d7d959fbe1").
		WillReturnError(errors.New("I am Error."))

	_, err := NewPokemon(db).GetPokemonById(context.Background(), "3ab43625-a18d-4b11-98a3-86d7d959fbe1")
	if err == nil {
		t.Error("expected an error but got nil")
	}
}

func TestPokemon_GetPokemonById_WithNoRows(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM pokemon LEFT JOIN .* WHERE pokemon.id.* GROUP BY pokemon.id").
		WithArgs("3ab43625-a18d-4b11-98a3-86d7d959fbe1").
		WillReturnRows(mockRowsForGetPokemonById(true))

	got, err := NewPokemon(db).GetPokemonById(context.Background(), "3ab43625-a18d-4b11-98a3-86d7d959fbe1")
	if err == nil {
		t.Error("expected an error but got nil")
	}

	var exp *model.Pokemon = nil

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestPokemon_PokemonByMoveIdDataLoader(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM pokemon LEFT JOIN .* WHERE pokemon_move.move_id IN (.*)").
		WithArgs("7b230001-b57e-4163-ac0a-3d157fc172e8", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "30e3e45e-9cc6-4fc2-a69a-fc91cd4908fb").
		WillReturnRows(mockRowsForPokemonByMoveIdDataLoader(false, false, false, []string{"7b230001-b57e-4163-ac0a-3d157fc172e8", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "30e3e45e-9cc6-4fc2-a69a-fc91cd4908fb"}))

	got, err := NewPokemon(db).PokemonByMoveIdDataLoader(context.Background())([]string{"7b230001-b57e-4163-ac0a-3d157fc172e8", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "30e3e45e-9cc6-4fc2-a69a-fc91cd4908fb"})
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	exp := []*model.PokemonList{
		{
			Total: 1,
			Pokemon: []*model.Pokemon{
				&castform,
			},
		},
		nil,
		{
			Total: 2,
			Pokemon: []*model.Pokemon{
				&snorunt,
				&bronzong,
			},
		},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestPokemon_PokemonByMoveIdDataLoader_WithQueryError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM pokemon LEFT JOIN .* WHERE pokemon_move.move_id IN (.*)").
		WithArgs("7b230001-b57e-4163-ac0a-3d157fc172e8", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "30e3e45e-9cc6-4fc2-a69a-fc91cd4908fb").
		WillReturnRows(mockRowsForPokemonByMoveIdDataLoader(false, false, false, []string{"7b230001-b57e-4163-ac0a-3d157fc172e8", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "30e3e45e-9cc6-4fc2-a69a-fc91cd4908fb"})).
		WillReturnError(errors.New("I am Error."))

	got, err := NewPokemon(db).PokemonByMoveIdDataLoader(context.Background())([]string{"7b230001-b57e-4163-ac0a-3d157fc172e8", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "30e3e45e-9cc6-4fc2-a69a-fc91cd4908fb"})
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := []*model.PokemonList{
		{
			Total:   0,
			Pokemon: []*model.Pokemon{},
		},
		{
			Total:   0,
			Pokemon: []*model.Pokemon{},
		},
		{
			Total:   0,
			Pokemon: []*model.Pokemon{},
		},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestPokemon_PokemonByMoveIdDataLoader_WithScanError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM pokemon LEFT JOIN .* WHERE pokemon_move.move_id IN (.*)").
		WithArgs("7b230001-b57e-4163-ac0a-3d157fc172e8", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "30e3e45e-9cc6-4fc2-a69a-fc91cd4908fb").
		WillReturnRows(mockRowsForPokemonByMoveIdDataLoader(false, false, true, []string{"7b230001-b57e-4163-ac0a-3d157fc172e8", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "30e3e45e-9cc6-4fc2-a69a-fc91cd4908fb"}))

	got, err := NewPokemon(db).PokemonByMoveIdDataLoader(context.Background())([]string{"7b230001-b57e-4163-ac0a-3d157fc172e8", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "30e3e45e-9cc6-4fc2-a69a-fc91cd4908fb"})
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := []*model.PokemonList{
		{
			Total:   0,
			Pokemon: []*model.Pokemon{},
		},
		{
			Total:   0,
			Pokemon: []*model.Pokemon{},
		},
		{
			Total:   0,
			Pokemon: []*model.Pokemon{},
		},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestPokemon_PokemonByMoveDataLoader_WithRowError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM pokemon LEFT JOIN .* WHERE pokemon_move.move_id IN (.*)").
		WithArgs("7b230001-b57e-4163-ac0a-3d157fc172e8", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "30e3e45e-9cc6-4fc2-a69a-fc91cd4908fb").
		WillReturnRows(mockRowsForPokemonByMoveIdDataLoader(false, true, false, []string{"7b230001-b57e-4163-ac0a-3d157fc172e8", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "30e3e45e-9cc6-4fc2-a69a-fc91cd4908fb"}))

	got, err := NewPokemon(db).PokemonByMoveIdDataLoader(context.Background())([]string{"7b230001-b57e-4163-ac0a-3d157fc172e8", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "30e3e45e-9cc6-4fc2-a69a-fc91cd4908fb"})
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := []*model.PokemonList{
		{
			Total: 1,
			Pokemon: []*model.Pokemon{
				&castform,
			},
		},
		nil,
		nil,
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestPokemon_PokemonByTypeIdDataLoader(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM pokemon LEFT JOIN .* WHERE pokemon_type.type_id IN (.*)").
		WithArgs("1dcc9d3c-55d4-4d33-809a-d1580c6e6542", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b").
		WillReturnRows(mockRowsForPokemonByTypeIdDataLoader(false, false, false, []string{"1dcc9d3c-55d4-4d33-809a-d1580c6e6542", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b"}))

	got, err := NewPokemon(db).PokemonByTypeIdDataLoader(context.Background())([]string{"1dcc9d3c-55d4-4d33-809a-d1580c6e6542", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b"})
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	exp := []*model.PokemonList{
		{
			Total: 2,
			Pokemon: []*model.Pokemon{
				&castform,
				&snorunt,
			},
		},
		nil,
		{
			Total: 1,
			Pokemon: []*model.Pokemon{
				&bronzong,
			},
		},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestPokemon_PokemonByTypeIdDataLoader_WithQueryError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM pokemon LEFT JOIN .* WHERE pokemon_type.type_id IN (.*)").
		WithArgs("1dcc9d3c-55d4-4d33-809a-d1580c6e6542", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b").
		WillReturnRows(mockRowsForPokemonByTypeIdDataLoader(false, false, false, []string{"1dcc9d3c-55d4-4d33-809a-d1580c6e6542", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b"})).
		WillReturnError(errors.New("I am Error."))

	got, err := NewPokemon(db).PokemonByTypeIdDataLoader(context.Background())([]string{"1dcc9d3c-55d4-4d33-809a-d1580c6e6542", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b"})
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := []*model.PokemonList{
		{
			Total:   0,
			Pokemon: []*model.Pokemon{},
		},
		{
			Total:   0,
			Pokemon: []*model.Pokemon{},
		},
		{
			Total:   0,
			Pokemon: []*model.Pokemon{},
		},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestPokemon_PokemonByTypeIdDataLoader_WithScanError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM pokemon LEFT JOIN .* WHERE pokemon_type.type_id IN (.*)").
		WithArgs("1dcc9d3c-55d4-4d33-809a-d1580c6e6542", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b").
		WillReturnRows(mockRowsForPokemonByTypeIdDataLoader(false, false, true, []string{"1dcc9d3c-55d4-4d33-809a-d1580c6e6542", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b"}))

	got, err := NewPokemon(db).PokemonByTypeIdDataLoader(context.Background())([]string{"1dcc9d3c-55d4-4d33-809a-d1580c6e6542", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b"})
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := []*model.PokemonList{
		{
			Total:   0,
			Pokemon: []*model.Pokemon{},
		},
		{
			Total:   0,
			Pokemon: []*model.Pokemon{},
		},
		{
			Total:   0,
			Pokemon: []*model.Pokemon{},
		},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestPokemon_PokemonByTypeIdDataLoader_WithRowError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM pokemon LEFT JOIN .* WHERE pokemon_type.type_id IN (.*)").
		WithArgs("1dcc9d3c-55d4-4d33-809a-d1580c6e6542", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b").
		WillReturnRows(mockRowsForPokemonByTypeIdDataLoader(false, true, false, []string{"1dcc9d3c-55d4-4d33-809a-d1580c6e6542", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b"}))

	got, err := NewPokemon(db).PokemonByTypeIdDataLoader(context.Background())([]string{"1dcc9d3c-55d4-4d33-809a-d1580c6e6542", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b"})
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := []*model.PokemonList{
		{
			Total: 1,
			Pokemon: []*model.Pokemon{
				&castform,
			},
		},
		nil,
		nil,
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestPokemon_PokemonByAbilityIdDataLoader(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM pokemon LEFT JOIN .* WHERE pokemon_ability.ability_id IN (.*)").
		WithArgs("0efe4eb9-537c-4b4c-92f6-d184a95b4923", "3eb38751-a341-457e-a211-1fc4641eac53", "9f0d876d-7e98-40d5-bfb3-2c0f079e2b26").
		WillReturnRows(mockRowsForPokemonByAbilityIdDataLoader(false, false, false, []string{"0efe4eb9-537c-4b4c-92f6-d184a95b4923", "3eb38751-a341-457e-a211-1fc4641eac53", "9f0d876d-7e98-40d5-bfb3-2c0f079e2b26"}))

	got, err := NewPokemon(db).PokemonByAbilityIdDataLoader(context.Background())([]string{"0efe4eb9-537c-4b4c-92f6-d184a95b4923", "3eb38751-a341-457e-a211-1fc4641eac53", "9f0d876d-7e98-40d5-bfb3-2c0f079e2b26"})
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	exp := []*model.PokemonList{
		{
			Total: 1,
			Pokemon: []*model.Pokemon{
				&castform,
			},
		},
		{
			Total: 1,
			Pokemon: []*model.Pokemon{
				&snorunt,
			},
		},
		{
			Total: 1,
			Pokemon: []*model.Pokemon{
				&bronzong,
			},
		},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestPokemon_PokemonByAbilityIdDataLoader_WithQueryError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM pokemon LEFT JOIN .* WHERE pokemon_ability.ability_id IN (.*)").
		WithArgs("0efe4eb9-537c-4b4c-92f6-d184a95b4923", "3eb38751-a341-457e-a211-1fc4641eac53", "9f0d876d-7e98-40d5-bfb3-2c0f079e2b26").
		WillReturnRows(mockRowsForPokemonByAbilityIdDataLoader(false, false, false, []string{"0efe4eb9-537c-4b4c-92f6-d184a95b4923", "3eb38751-a341-457e-a211-1fc4641eac53", "9f0d876d-7e98-40d5-bfb3-2c0f079e2b26"})).
		WillReturnError(errors.New("I am Error."))

	got, err := NewPokemon(db).PokemonByAbilityIdDataLoader(context.Background())([]string{"0efe4eb9-537c-4b4c-92f6-d184a95b4923", "3eb38751-a341-457e-a211-1fc4641eac53", "9f0d876d-7e98-40d5-bfb3-2c0f079e2b26"})
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := []*model.PokemonList{
		{
			Total:   0,
			Pokemon: []*model.Pokemon{},
		},
		{
			Total:   0,
			Pokemon: []*model.Pokemon{},
		},
		{
			Total:   0,
			Pokemon: []*model.Pokemon{},
		},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestPokemon_PokemonByAbilityIdDataLoader_WithScanError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM pokemon LEFT JOIN .* WHERE pokemon_ability.ability_id IN (.*)").
		WithArgs("0efe4eb9-537c-4b4c-92f6-d184a95b4923", "3eb38751-a341-457e-a211-1fc4641eac53", "9f0d876d-7e98-40d5-bfb3-2c0f079e2b26").
		WillReturnRows(mockRowsForPokemonByAbilityIdDataLoader(false, false, true, []string{"0efe4eb9-537c-4b4c-92f6-d184a95b4923", "3eb38751-a341-457e-a211-1fc4641eac53", "9f0d876d-7e98-40d5-bfb3-2c0f079e2b26"}))

	got, err := NewPokemon(db).PokemonByAbilityIdDataLoader(context.Background())([]string{"0efe4eb9-537c-4b4c-92f6-d184a95b4923", "3eb38751-a341-457e-a211-1fc4641eac53", "9f0d876d-7e98-40d5-bfb3-2c0f079e2b26"})
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := []*model.PokemonList{
		{
			Total:   0,
			Pokemon: []*model.Pokemon{},
		},
		{
			Total:   0,
			Pokemon: []*model.Pokemon{},
		},
		{
			Total:   0,
			Pokemon: []*model.Pokemon{},
		},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestPokemon_PokemonByAbilityIdDataLoader_WithRowError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM pokemon LEFT JOIN .* WHERE pokemon_ability.ability_id IN (.*)").
		WithArgs("0efe4eb9-537c-4b4c-92f6-d184a95b4923", "3eb38751-a341-457e-a211-1fc4641eac53", "9f0d876d-7e98-40d5-bfb3-2c0f079e2b26").
		WillReturnRows(mockRowsForPokemonByAbilityIdDataLoader(false, true, false, []string{"0efe4eb9-537c-4b4c-92f6-d184a95b4923", "3eb38751-a341-457e-a211-1fc4641eac53", "9f0d876d-7e98-40d5-bfb3-2c0f079e2b26"}))

	got, err := NewPokemon(db).PokemonByAbilityIdDataLoader(context.Background())([]string{"0efe4eb9-537c-4b4c-92f6-d184a95b4923", "3eb38751-a341-457e-a211-1fc4641eac53", "9f0d876d-7e98-40d5-bfb3-2c0f079e2b26"})
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := []*model.PokemonList{
		{
			Total: 1,
			Pokemon: []*model.Pokemon{
				&castform,
			},
		},
		nil,
		nil,
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

var castform = model.Pokemon{
	ID:               "3ab43625-a18d-4b11-98a3-86d7d959fbe1",
	Slug:             "castform-snowy",
	Name:             "Castform",
	PokedexId:        351,
	Sprite:           "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/10015.png",
	HP:               70,
	Attack:           70,
	Defense:          70,
	SpecialAttack:    70,
	SpecialDefense:   70,
	Speed:            70,
	IsBaby:           false,
	IsLegendary:      false,
	IsMythical:       false,
	Description:      "Its form changes depending on the weather.\nThe rougher conditions get, the rougher\nCastform’s disposition!",
	Color:            model.Gray,
	Shape:            model.Ball,
	Habitat:          model.Grassland,
	IsDefaultVariant: false,
	Genus:            "Weather Pokémon",
	Height:           3,
	Weight:           8,
	Types: model.PokemonTypeList{
		Total: 1,
		PokemonTypes: []model.PokemonType{
			{
				Slot: 1,
				Type: model.Type{
					ID:   "1dcc9d3c-55d4-4d33-809a-d1580c6e6542",
					Name: "Ice",
					Slug: "ice",
				},
			},
		},
	},
	Abilities: model.AbilityList{
		Total: 1,
		Abilities: []model.Ability{
			{
				ID:     "0efe4eb9-537c-4b4c-92f6-d184a95b4923",
				Name:   "Forecast",
				Slug:   "forecast",
				Effect: "Changes castform's type and form to match the weather.",
			},
		},
	},
	EggGroups: model.EggGroupList{
		Total: 2,
		EggGroups: []model.EggGroup{
			{
				ID:   "1f0958a0-48ca-4160-9f18-7e5f06d96d27",
				Name: "Fairy",
				Slug: "fairy",
			},
			{
				ID:   "465ed2fa-0ff8-4cad-89af-e9db971026df",
				Name: "Amorphous",
				Slug: "indeterminate",
			},
		},
	},
}

var snorunt = model.Pokemon{
	ID:               "51948cca-743a-4e6d-9c00-579140daccc5",
	Slug:             "snorunt",
	Name:             "Snorunt",
	PokedexId:        361,
	Sprite:           "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/361.png",
	HP:               50,
	Attack:           50,
	Defense:          50,
	SpecialAttack:    50,
	SpecialDefense:   50,
	Speed:            50,
	IsBaby:           false,
	IsLegendary:      false,
	IsMythical:       false,
	Description:      "Rich people from cold areas all share childhood\nmemories of playing with Snorunt.",
	Color:            model.Gray,
	Shape:            model.Humanoid,
	Habitat:          model.Cave,
	IsDefaultVariant: true,
	Genus:            "Snow Hat Pokémon",
	Height:           7,
	Weight:           168,
	Types: model.PokemonTypeList{
		Total: 1,
		PokemonTypes: []model.PokemonType{
			{
				Slot: 1,
				Type: model.Type{
					ID:   "1dcc9d3c-55d4-4d33-809a-d1580c6e6542",
					Name: "Ice",
					Slug: "ice",
				},
			},
		},
	},
	Abilities: model.AbilityList{
		Total: 3,
		Abilities: []model.Ability{
			{
				ID:     "3eb38751-a341-457e-a211-1fc4641eac53",
				Name:   "Inner Focus",
				Slug:   "inner-focus",
				Effect: "Prevents flinching.",
			},
			{
				ID:     "673dd8ad-1494-49e1-86cd-9572df34540b",
				Name:   "Ice Body",
				Slug:   "ice-body",
				Effect: "Heals for 1/16 max HP after each turn during hail.  Protects against hail damage.",
			},
			{
				ID:     "ba77aff4-9bab-4dc7-acdc-e0bbba9b5c88",
				Name:   "Moody",
				Slug:   "moody",
				Effect: "Raises a random stat two stages and lowers another one stage after each turn.",
			},
		},
	},
	EggGroups: model.EggGroupList{
		Total: 2,
		EggGroups: []model.EggGroup{
			{
				ID:   "1f0958a0-48ca-4160-9f18-7e5f06d96d27",
				Name: "Fairy",
				Slug: "fairy",
			},
			{
				ID:   "b140921f-74c9-4537-9a08-996277d4fcb4",
				Name: "Mineral",
				Slug: "mineral",
			},
		},
	},
}

var bronzong = model.Pokemon{
	ID:               "85da4120-96bb-42b1-8e8f-9f8bded11a31",
	Slug:             "bronzong",
	Name:             "Bronzong",
	PokedexId:        437,
	Sprite:           "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/437.png",
	HP:               67,
	Attack:           89,
	Defense:          116,
	SpecialAttack:    79,
	SpecialDefense:   116,
	Speed:            33,
	IsBaby:           false,
	IsLegendary:      false,
	IsMythical:       false,
	Description:      "",
	Color:            model.Green,
	Shape:            model.Arms,
	IsDefaultVariant: true,
	Genus:            "Bronze Bell Pokémon",
	Height:           13,
	Weight:           1870,
	Types: model.PokemonTypeList{
		Total: 2,
		PokemonTypes: []model.PokemonType{
			{
				Slot: 1,
				Type: model.Type{
					ID:   "05cd51bd-23ca-4736-b8ec-aa93aca68a8b",
					Name: "Steel",
					Slug: "steel",
				},
			},
			{
				Slot: 2,
				Type: model.Type{
					ID:   "2222c839-3c6e-4727-b6b5-a946bb8af5fa",
					Name: "Psychic",
					Slug: "psychic",
				},
			},
		},
	},
	Abilities: model.AbilityList{
		Total: 3,
		Abilities: []model.Ability{
			{
				ID:     "4acbc86c-4a5f-4f6e-99e8-4feab6337ad6",
				Name:   "Heavy Metal",
				Slug:   "heavy-metal",
				Effect: "Doubles the Pokémon's weight.",
			},
			{
				ID:     "9f0d876d-7e98-40d5-bfb3-2c0f079e2b26",
				Name:   "Heatproof",
				Slug:   "heatproof",
				Effect: "Halves damage from fire moves and burns.",
			},
			{
				ID:     "e55c279d-4554-4d5e-8120-7bf3a0477181",
				Name:   "Levitate",
				Slug:   "levitate",
				Effect: "Evades ground moves.",
			},
		},
	},
	EggGroups: model.EggGroupList{
		Total: 1,
		EggGroups: []model.EggGroup{
			{
				ID:   "b140921f-74c9-4537-9a08-996277d4fcb4",
				Name: "Mineral",
				Slug: "mineral",
			},
		},
	},
}

func mockRowsForGetPokemon(empty bool, hasRowError bool, hasScanError bool) *sqlmock.Rows {
	if hasScanError {
		rows := sqlmock.NewRows([]string{"id"})
		rows.AddRow("3ab43625-a18d-4b11-98a3-86d7d959fbe1")
		return rows
	}
	rows := sqlmock.NewRows([]string{"id", "pokedex_id", "slug", "name", "sprite", "hp", "attack", "defense", "special_attack", "special_defense", "speed", "is_baby", "is_legendary", "is_mythical", "description", "color_enum", "shape_enum", "habitat_enum", "is_default_variant", "genus", "height", "weight", "types", "egg_groups", "abilities"})
	if !empty {
		rows.AddRow(castform.ID, castform.PokedexId, castform.Slug, castform.Name, castform.Sprite, castform.HP, castform.Attack, castform.Defense, castform.SpecialAttack, castform.SpecialDefense, castform.Speed, castform.IsBaby, castform.IsLegendary, castform.IsMythical, castform.Description, castform.Color.String(), castform.Shape.String(), castform.Habitat.String(), castform.IsDefaultVariant, castform.Genus, castform.Height, castform.Weight, `{"{\"slot\": 1, \"type\": {\"id\": \"1dcc9d3c-55d4-4d33-809a-d1580c6e6542\", \"name\": \"Ice\", \"slug\": \"ice\"}}"}`, `{"{\"id\": \"1f0958a0-48ca-4160-9f18-7e5f06d96d27\", \"name\": \"Fairy\", \"slug\": \"fairy\"}","{\"id\": \"465ed2fa-0ff8-4cad-89af-e9db971026df\", \"name\": \"Amorphous\", \"slug\": \"indeterminate\"}"}`, `{"{\"id\": \"0efe4eb9-537c-4b4c-92f6-d184a95b4923\", \"name\": \"Forecast\", \"slug\": \"forecast\", \"effect\": \"Changes castform's type and form to match the weather.\"}"}`).
			AddRow(snorunt.ID, snorunt.PokedexId, snorunt.Slug, snorunt.Name, snorunt.Sprite, snorunt.HP, snorunt.Attack, snorunt.Defense, snorunt.SpecialAttack, snorunt.SpecialDefense, snorunt.Speed, snorunt.IsBaby, snorunt.IsLegendary, snorunt.IsMythical, snorunt.Description, snorunt.Color.String(), snorunt.Shape.String(), snorunt.Habitat.String(), snorunt.IsDefaultVariant, snorunt.Genus, snorunt.Height, snorunt.Weight, `{"{\"slot\": 1, \"type\": {\"id\": \"1dcc9d3c-55d4-4d33-809a-d1580c6e6542\", \"name\": \"Ice\", \"slug\": \"ice\"}}"}`, `{"{\"id\": \"1f0958a0-48ca-4160-9f18-7e5f06d96d27\", \"name\": \"Fairy\", \"slug\": \"fairy\"}","{\"id\": \"b140921f-74c9-4537-9a08-996277d4fcb4\", \"name\": \"Mineral\", \"slug\": \"mineral\"}"}`, `{"{\"id\": \"3eb38751-a341-457e-a211-1fc4641eac53\", \"name\": \"Inner Focus\", \"slug\": \"inner-focus\", \"effect\": \"Prevents flinching.\"}","{\"id\": \"673dd8ad-1494-49e1-86cd-9572df34540b\", \"name\": \"Ice Body\", \"slug\": \"ice-body\", \"effect\": \"Heals for 1/16 max HP after each turn during hail.  Protects against hail damage.\"}","{\"id\": \"ba77aff4-9bab-4dc7-acdc-e0bbba9b5c88\", \"name\": \"Moody\", \"slug\": \"moody\", \"effect\": \"Raises a random stat two stages and lowers another one stage after each turn.\"}"}`).
			AddRow(bronzong.ID, bronzong.PokedexId, bronzong.Slug, bronzong.Name, bronzong.Sprite, bronzong.HP, bronzong.Attack, bronzong.Defense, bronzong.SpecialAttack, bronzong.SpecialDefense, bronzong.Speed, bronzong.IsBaby, bronzong.IsLegendary, bronzong.IsMythical, bronzong.Description, bronzong.Color.String(), bronzong.Shape.String(), nil, bronzong.IsDefaultVariant, bronzong.Genus, bronzong.Height, bronzong.Weight, `{"{\"slot\": 1, \"type\": {\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}}","{\"slot\": 2, \"type\": {\"id\": \"2222c839-3c6e-4727-b6b5-a946bb8af5fa\", \"name\": \"Psychic\", \"slug\": \"psychic\"}}"}`, `{"{\"id\": \"b140921f-74c9-4537-9a08-996277d4fcb4\", \"name\": \"Mineral\", \"slug\": \"mineral\"}"}`, `{"{\"id\": \"4acbc86c-4a5f-4f6e-99e8-4feab6337ad6\", \"name\": \"Heavy Metal\", \"slug\": \"heavy-metal\", \"effect\": \"Doubles the Pokémon's weight.\"}","{\"id\": \"9f0d876d-7e98-40d5-bfb3-2c0f079e2b26\", \"name\": \"Heatproof\", \"slug\": \"heatproof\", \"effect\": \"Halves damage from fire moves and burns.\"}","{\"id\": \"e55c279d-4554-4d5e-8120-7bf3a0477181\", \"name\": \"Levitate\", \"slug\": \"levitate\", \"effect\": \"Evades ground moves.\"}"}`)
	}

	if hasRowError {
		rows.RowError(0, errors.New("scan error"))
	}
	return rows
}

func mockRowsForGetPokemonById(empty bool) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"id", "pokedex_id", "slug", "name", "sprite", "hp", "attack", "defense", "special_attack", "special_defense", "speed", "is_baby", "is_legendary", "is_mythical", "description", "color_enum", "shape_enum", "habitat_enum", "is_default_variant", "genus", "height", "weight", "types", "egg_groups", "abilities"})
	if !empty {
		rows.AddRow(castform.ID, castform.PokedexId, castform.Slug, castform.Name, castform.Sprite, castform.HP, castform.Attack, castform.Defense, castform.SpecialAttack, castform.SpecialDefense, castform.Speed, castform.IsBaby, castform.IsLegendary, castform.IsMythical, castform.Description, castform.Color.String(), castform.Shape.String(), castform.Habitat.String(), castform.IsDefaultVariant, castform.Genus, castform.Height, castform.Weight, `{"{\"slot\": 1, \"type\": {\"id\": \"1dcc9d3c-55d4-4d33-809a-d1580c6e6542\", \"name\": \"Ice\", \"slug\": \"ice\"}}"}`, `{"{\"id\": \"1f0958a0-48ca-4160-9f18-7e5f06d96d27\", \"name\": \"Fairy\", \"slug\": \"fairy\"}","{\"id\": \"465ed2fa-0ff8-4cad-89af-e9db971026df\", \"name\": \"Amorphous\", \"slug\": \"indeterminate\"}"}`, `{"{\"id\": \"0efe4eb9-537c-4b4c-92f6-d184a95b4923\", \"name\": \"Forecast\", \"slug\": \"forecast\", \"effect\": \"Changes castform's type and form to match the weather.\"}"}`)
	}
	return rows
}

func mockRowsForPokemonByMoveIdDataLoader(empty bool, hasRowError bool, hasScanError bool, ids []string) *sqlmock.Rows {
	if hasScanError {
		rows := sqlmock.NewRows([]string{"id"})
		rows.AddRow("3ab43625-a18d-4b11-98a3-86d7d959fbe1")
		return rows
	}
	rows := sqlmock.NewRows([]string{"id", "name", "slug", "pokedex_id", "sprite", "hp", "attack", "defense", "special_attack", "special_defense", "speed", "is_baby", "is_legendary", "is_mythical", "description", "color_enum", "habitat_enum", "shape_enum", "height", "weight", "is_default_variant", "genus", "types", "egg_groups", "abilities", "pokemon_move.move_id"})
	if !empty {
		rows.AddRow(castform.ID, castform.PokedexId, castform.Slug, castform.Name, castform.Sprite, castform.HP, castform.Attack, castform.Defense, castform.SpecialAttack, castform.SpecialDefense, castform.Speed, castform.IsBaby, castform.IsLegendary, castform.IsMythical, castform.Description, castform.Color.String(), castform.Shape.String(), castform.Habitat.String(), castform.IsDefaultVariant, castform.Genus, castform.Height, castform.Weight, `{"{\"slot\": 1, \"type\": {\"id\": \"1dcc9d3c-55d4-4d33-809a-d1580c6e6542\", \"name\": \"Ice\", \"slug\": \"ice\"}}"}`, `{"{\"id\": \"1f0958a0-48ca-4160-9f18-7e5f06d96d27\", \"name\": \"Fairy\", \"slug\": \"fairy\"}","{\"id\": \"465ed2fa-0ff8-4cad-89af-e9db971026df\", \"name\": \"Amorphous\", \"slug\": \"indeterminate\"}"}`, `{"{\"id\": \"0efe4eb9-537c-4b4c-92f6-d184a95b4923\", \"name\": \"Forecast\", \"slug\": \"forecast\", \"effect\": \"Changes castform's type and form to match the weather.\"}"}`, ids[0]).
			AddRow(snorunt.ID, snorunt.PokedexId, snorunt.Slug, snorunt.Name, snorunt.Sprite, snorunt.HP, snorunt.Attack, snorunt.Defense, snorunt.SpecialAttack, snorunt.SpecialDefense, snorunt.Speed, snorunt.IsBaby, snorunt.IsLegendary, snorunt.IsMythical, snorunt.Description, snorunt.Color.String(), snorunt.Shape.String(), snorunt.Habitat.String(), snorunt.IsDefaultVariant, snorunt.Genus, snorunt.Height, snorunt.Weight, `{"{\"slot\": 1, \"type\": {\"id\": \"1dcc9d3c-55d4-4d33-809a-d1580c6e6542\", \"name\": \"Ice\", \"slug\": \"ice\"}}"}`, `{"{\"id\": \"1f0958a0-48ca-4160-9f18-7e5f06d96d27\", \"name\": \"Fairy\", \"slug\": \"fairy\"}","{\"id\": \"b140921f-74c9-4537-9a08-996277d4fcb4\", \"name\": \"Mineral\", \"slug\": \"mineral\"}"}`, `{"{\"id\": \"3eb38751-a341-457e-a211-1fc4641eac53\", \"name\": \"Inner Focus\", \"slug\": \"inner-focus\", \"effect\": \"Prevents flinching.\"}","{\"id\": \"673dd8ad-1494-49e1-86cd-9572df34540b\", \"name\": \"Ice Body\", \"slug\": \"ice-body\", \"effect\": \"Heals for 1/16 max HP after each turn during hail.  Protects against hail damage.\"}","{\"id\": \"ba77aff4-9bab-4dc7-acdc-e0bbba9b5c88\", \"name\": \"Moody\", \"slug\": \"moody\", \"effect\": \"Raises a random stat two stages and lowers another one stage after each turn.\"}"}`, ids[2]).
			AddRow(bronzong.ID, bronzong.PokedexId, bronzong.Slug, bronzong.Name, bronzong.Sprite, bronzong.HP, bronzong.Attack, bronzong.Defense, bronzong.SpecialAttack, bronzong.SpecialDefense, bronzong.Speed, bronzong.IsBaby, bronzong.IsLegendary, bronzong.IsMythical, bronzong.Description, bronzong.Color.String(), bronzong.Shape.String(), nil, bronzong.IsDefaultVariant, bronzong.Genus, bronzong.Height, bronzong.Weight, `{"{\"slot\": 1, \"type\": {\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}}","{\"slot\": 2, \"type\": {\"id\": \"2222c839-3c6e-4727-b6b5-a946bb8af5fa\", \"name\": \"Psychic\", \"slug\": \"psychic\"}}"}`, `{"{\"id\": \"b140921f-74c9-4537-9a08-996277d4fcb4\", \"name\": \"Mineral\", \"slug\": \"mineral\"}"}`, `{"{\"id\": \"4acbc86c-4a5f-4f6e-99e8-4feab6337ad6\", \"name\": \"Heavy Metal\", \"slug\": \"heavy-metal\", \"effect\": \"Doubles the Pokémon's weight.\"}","{\"id\": \"9f0d876d-7e98-40d5-bfb3-2c0f079e2b26\", \"name\": \"Heatproof\", \"slug\": \"heatproof\", \"effect\": \"Halves damage from fire moves and burns.\"}","{\"id\": \"e55c279d-4554-4d5e-8120-7bf3a0477181\", \"name\": \"Levitate\", \"slug\": \"levitate\", \"effect\": \"Evades ground moves.\"}"}`, ids[2])
	}
	if hasRowError {
		rows.RowError(1, errors.New("scan error"))
	}
	return rows
}

func mockRowsForPokemonByTypeIdDataLoader(empty bool, hasRowError bool, hasScanError bool, ids []string) *sqlmock.Rows {
	if hasScanError {
		rows := sqlmock.NewRows([]string{"id"})
		rows.AddRow("3ab43625-a18d-4b11-98a3-86d7d959fbe1")
		return rows
	}
	rows := sqlmock.NewRows([]string{"id", "name", "slug", "pokedex_id", "sprite", "hp", "attack", "defense", "special_attack", "special_defense", "speed", "is_baby", "is_legendary", "is_mythical", "description", "color_enum", "habitat_enum", "shape_enum", "height", "weight", "is_default_variant", "genus", "types", "egg_groups", "abilities"})
	if !empty {
		rows.AddRow(castform.ID, castform.PokedexId, castform.Slug, castform.Name, castform.Sprite, castform.HP, castform.Attack, castform.Defense, castform.SpecialAttack, castform.SpecialDefense, castform.Speed, castform.IsBaby, castform.IsLegendary, castform.IsMythical, castform.Description, castform.Color.String(), castform.Shape.String(), castform.Habitat.String(), castform.IsDefaultVariant, castform.Genus, castform.Height, castform.Weight, `{"{\"slot\": 1, \"type\": {\"id\": \"`+ids[0]+`\", \"name\": \"Ice\", \"slug\": \"ice\"}}"}`, `{"{\"id\": \"1f0958a0-48ca-4160-9f18-7e5f06d96d27\", \"name\": \"Fairy\", \"slug\": \"fairy\"}","{\"id\": \"465ed2fa-0ff8-4cad-89af-e9db971026df\", \"name\": \"Amorphous\", \"slug\": \"indeterminate\"}"}`, `{"{\"id\": \"0efe4eb9-537c-4b4c-92f6-d184a95b4923\", \"name\": \"Forecast\", \"slug\": \"forecast\", \"effect\": \"Changes castform's type and form to match the weather.\"}"}`).
			AddRow(snorunt.ID, snorunt.PokedexId, snorunt.Slug, snorunt.Name, snorunt.Sprite, snorunt.HP, snorunt.Attack, snorunt.Defense, snorunt.SpecialAttack, snorunt.SpecialDefense, snorunt.Speed, snorunt.IsBaby, snorunt.IsLegendary, snorunt.IsMythical, snorunt.Description, snorunt.Color.String(), snorunt.Shape.String(), snorunt.Habitat.String(), snorunt.IsDefaultVariant, snorunt.Genus, snorunt.Height, snorunt.Weight, `{"{\"slot\": 1, \"type\": {\"id\": \"`+ids[0]+`\", \"name\": \"Ice\", \"slug\": \"ice\"}}"}`, `{"{\"id\": \"1f0958a0-48ca-4160-9f18-7e5f06d96d27\", \"name\": \"Fairy\", \"slug\": \"fairy\"}","{\"id\": \"b140921f-74c9-4537-9a08-996277d4fcb4\", \"name\": \"Mineral\", \"slug\": \"mineral\"}"}`, `{"{\"id\": \"3eb38751-a341-457e-a211-1fc4641eac53\", \"name\": \"Inner Focus\", \"slug\": \"inner-focus\", \"effect\": \"Prevents flinching.\"}","{\"id\": \"673dd8ad-1494-49e1-86cd-9572df34540b\", \"name\": \"Ice Body\", \"slug\": \"ice-body\", \"effect\": \"Heals for 1/16 max HP after each turn during hail.  Protects against hail damage.\"}","{\"id\": \"ba77aff4-9bab-4dc7-acdc-e0bbba9b5c88\", \"name\": \"Moody\", \"slug\": \"moody\", \"effect\": \"Raises a random stat two stages and lowers another one stage after each turn.\"}"}`).
			AddRow(bronzong.ID, bronzong.PokedexId, bronzong.Slug, bronzong.Name, bronzong.Sprite, bronzong.HP, bronzong.Attack, bronzong.Defense, bronzong.SpecialAttack, bronzong.SpecialDefense, bronzong.Speed, bronzong.IsBaby, bronzong.IsLegendary, bronzong.IsMythical, bronzong.Description, bronzong.Color.String(), bronzong.Shape.String(), nil, bronzong.IsDefaultVariant, bronzong.Genus, bronzong.Height, bronzong.Weight, `{"{\"slot\": 1, \"type\": {\"id\": \"`+ids[2]+`\", \"name\": \"Steel\", \"slug\": \"steel\"}}","{\"slot\": 2, \"type\": {\"id\": \"2222c839-3c6e-4727-b6b5-a946bb8af5fa\", \"name\": \"Psychic\", \"slug\": \"psychic\"}}"}`, `{"{\"id\": \"b140921f-74c9-4537-9a08-996277d4fcb4\", \"name\": \"Mineral\", \"slug\": \"mineral\"}"}`, `{"{\"id\": \"4acbc86c-4a5f-4f6e-99e8-4feab6337ad6\", \"name\": \"Heavy Metal\", \"slug\": \"heavy-metal\", \"effect\": \"Doubles the Pokémon's weight.\"}","{\"id\": \"9f0d876d-7e98-40d5-bfb3-2c0f079e2b26\", \"name\": \"Heatproof\", \"slug\": \"heatproof\", \"effect\": \"Halves damage from fire moves and burns.\"}","{\"id\": \"e55c279d-4554-4d5e-8120-7bf3a0477181\", \"name\": \"Levitate\", \"slug\": \"levitate\", \"effect\": \"Evades ground moves.\"}"}`)
	}
	if hasRowError {
		rows.RowError(1, errors.New("scan error"))
	}
	return rows
}

func mockRowsForPokemonByAbilityIdDataLoader(empty bool, hasRowError bool, hasScanError bool, ids []string) *sqlmock.Rows {
	if hasScanError {
		rows := sqlmock.NewRows([]string{"id"})
		rows.AddRow("3ab43625-a18d-4b11-98a3-86d7d959fbe1")
		return rows
	}
	rows := sqlmock.NewRows([]string{"id", "name", "slug", "pokedex_id", "sprite", "hp", "attack", "defense", "special_attack", "special_defense", "speed", "is_baby", "is_legendary", "is_mythical", "description", "color_enum", "habitat_enum", "shape_enum", "height", "weight", "is_default_variant", "genus", "types", "egg_groups", "abilities"})
	if !empty {
		rows.AddRow(castform.ID, castform.PokedexId, castform.Slug, castform.Name, castform.Sprite, castform.HP, castform.Attack, castform.Defense, castform.SpecialAttack, castform.SpecialDefense, castform.Speed, castform.IsBaby, castform.IsLegendary, castform.IsMythical, castform.Description, castform.Color.String(), castform.Shape.String(), castform.Habitat.String(), castform.IsDefaultVariant, castform.Genus, castform.Height, castform.Weight, `{"{\"slot\": 1, \"type\": {\"id\": \"1dcc9d3c-55d4-4d33-809a-d1580c6e6542\", \"name\": \"Ice\", \"slug\": \"ice\"}}"}`, `{"{\"id\": \"1f0958a0-48ca-4160-9f18-7e5f06d96d27\", \"name\": \"Fairy\", \"slug\": \"fairy\"}","{\"id\": \"465ed2fa-0ff8-4cad-89af-e9db971026df\", \"name\": \"Amorphous\", \"slug\": \"indeterminate\"}"}`, `{"{\"id\": \"`+ids[0]+`\", \"name\": \"Forecast\", \"slug\": \"forecast\", \"effect\": \"Changes castform's type and form to match the weather.\"}"}`).
			AddRow(snorunt.ID, snorunt.PokedexId, snorunt.Slug, snorunt.Name, snorunt.Sprite, snorunt.HP, snorunt.Attack, snorunt.Defense, snorunt.SpecialAttack, snorunt.SpecialDefense, snorunt.Speed, snorunt.IsBaby, snorunt.IsLegendary, snorunt.IsMythical, snorunt.Description, snorunt.Color.String(), snorunt.Shape.String(), snorunt.Habitat.String(), snorunt.IsDefaultVariant, snorunt.Genus, snorunt.Height, snorunt.Weight, `{"{\"slot\": 1, \"type\": {\"id\": \"1dcc9d3c-55d4-4d33-809a-d1580c6e6542\", \"name\": \"Ice\", \"slug\": \"ice\"}}"}`, `{"{\"id\": \"1f0958a0-48ca-4160-9f18-7e5f06d96d27\", \"name\": \"Fairy\", \"slug\": \"fairy\"}","{\"id\": \"b140921f-74c9-4537-9a08-996277d4fcb4\", \"name\": \"Mineral\", \"slug\": \"mineral\"}"}`, `{"{\"id\": \"`+ids[1]+`\", \"name\": \"Inner Focus\", \"slug\": \"inner-focus\", \"effect\": \"Prevents flinching.\"}","{\"id\": \"673dd8ad-1494-49e1-86cd-9572df34540b\", \"name\": \"Ice Body\", \"slug\": \"ice-body\", \"effect\": \"Heals for 1/16 max HP after each turn during hail.  Protects against hail damage.\"}","{\"id\": \"ba77aff4-9bab-4dc7-acdc-e0bbba9b5c88\", \"name\": \"Moody\", \"slug\": \"moody\", \"effect\": \"Raises a random stat two stages and lowers another one stage after each turn.\"}"}`).
			AddRow(bronzong.ID, bronzong.PokedexId, bronzong.Slug, bronzong.Name, bronzong.Sprite, bronzong.HP, bronzong.Attack, bronzong.Defense, bronzong.SpecialAttack, bronzong.SpecialDefense, bronzong.Speed, bronzong.IsBaby, bronzong.IsLegendary, bronzong.IsMythical, bronzong.Description, bronzong.Color.String(), bronzong.Shape.String(), nil, bronzong.IsDefaultVariant, bronzong.Genus, bronzong.Height, bronzong.Weight, `{"{\"slot\": 1, \"type\": {\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}}","{\"slot\": 2, \"type\": {\"id\": \"2222c839-3c6e-4727-b6b5-a946bb8af5fa\", \"name\": \"Psychic\", \"slug\": \"psychic\"}}"}`, `{"{\"id\": \"b140921f-74c9-4537-9a08-996277d4fcb4\", \"name\": \"Mineral\", \"slug\": \"mineral\"}"}`, `{"{\"id\": \"4acbc86c-4a5f-4f6e-99e8-4feab6337ad6\", \"name\": \"Heavy Metal\", \"slug\": \"heavy-metal\", \"effect\": \"Doubles the Pokémon's weight.\"}","{\"id\": \"`+ids[2]+`\", \"name\": \"Heatproof\", \"slug\": \"heatproof\", \"effect\": \"Halves damage from fire moves and burns.\"}","{\"id\": \"e55c279d-4554-4d5e-8120-7bf3a0477181\", \"name\": \"Levitate\", \"slug\": \"levitate\", \"effect\": \"Evades ground moves.\"}"}`)
	}
	if hasRowError {
		rows.RowError(1, errors.New("scan error"))
	}
	return rows
}
