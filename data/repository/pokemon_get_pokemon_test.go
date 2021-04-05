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

	pokemon := []model.Pokemon{
		castform,
		snorunt,
		bronzong,
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
		Pokemon: []model.Pokemon{},
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
		Pokemon: []model.Pokemon{},
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
		Pokemon: []model.Pokemon{},
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
		Pokemon: []model.Pokemon{},
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
			Pokemon: []model.Pokemon{
				castform,
			},
		},
		nil,
		{
			Total: 2,
			Pokemon: []model.Pokemon{
				snorunt,
				bronzong,
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
			Pokemon: []model.Pokemon{},
		},
		{
			Total:   0,
			Pokemon: []model.Pokemon{},
		},
		{
			Total:   0,
			Pokemon: []model.Pokemon{},
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
			Pokemon: []model.Pokemon{},
		},
		{
			Total:   0,
			Pokemon: []model.Pokemon{},
		},
		{
			Total:   0,
			Pokemon: []model.Pokemon{},
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
			Pokemon: []model.Pokemon{
				castform,
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
		WithArgs("366a8621-9fa7-419b-b710-9100bcbb98d8", "5179f383-b765-4cc7-b9f9-8b1a3ba93019", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b").
		WillReturnRows(mockRowsForPokemonByTypeIdDataLoader(false, false, false, []string{"366a8621-9fa7-419b-b710-9100bcbb98d8", "5179f383-b765-4cc7-b9f9-8b1a3ba93019", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b"}))

	got, err := NewPokemon(db).PokemonByTypeIdDataLoader(context.Background())([]string{"366a8621-9fa7-419b-b710-9100bcbb98d8", "5179f383-b765-4cc7-b9f9-8b1a3ba93019", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b"})
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	exp := []*model.PokemonList{
		{
			Total: 1,
			Pokemon: []model.Pokemon{
				castform,
			},
		},
		nil,
		{
			Total: 1,
			Pokemon: []model.Pokemon{
				bronzong,
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
		WithArgs("366a8621-9fa7-419b-b710-9100bcbb98d8", "5179f383-b765-4cc7-b9f9-8b1a3ba93019", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b").
		WillReturnRows(mockRowsForPokemonByTypeIdDataLoader(false, false, false, []string{"366a8621-9fa7-419b-b710-9100bcbb98d8", "5179f383-b765-4cc7-b9f9-8b1a3ba93019", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b"})).
		WillReturnError(errors.New("I am Error."))

	got, err := NewPokemon(db).PokemonByTypeIdDataLoader(context.Background())([]string{"366a8621-9fa7-419b-b710-9100bcbb98d8", "5179f383-b765-4cc7-b9f9-8b1a3ba93019", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b"})
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := []*model.PokemonList{
		{
			Total:   0,
			Pokemon: []model.Pokemon{},
		},
		{
			Total:   0,
			Pokemon: []model.Pokemon{},
		},
		{
			Total:   0,
			Pokemon: []model.Pokemon{},
		},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestPokemon_PokemonByTypeIdDataLoader_WithScanError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM pokemon LEFT JOIN .* WHERE pokemon_type.type_id IN (.*)").
		WithArgs("366a8621-9fa7-419b-b710-9100bcbb98d8", "5179f383-b765-4cc7-b9f9-8b1a3ba93019", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b").
		WillReturnRows(mockRowsForPokemonByTypeIdDataLoader(false, false, true, []string{"366a8621-9fa7-419b-b710-9100bcbb98d8", "5179f383-b765-4cc7-b9f9-8b1a3ba93019", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b"}))

	got, err := NewPokemon(db).PokemonByTypeIdDataLoader(context.Background())([]string{"366a8621-9fa7-419b-b710-9100bcbb98d8", "5179f383-b765-4cc7-b9f9-8b1a3ba93019", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b"})
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := []*model.PokemonList{
		{
			Total:   0,
			Pokemon: []model.Pokemon{},
		},
		{
			Total:   0,
			Pokemon: []model.Pokemon{},
		},
		{
			Total:   0,
			Pokemon: []model.Pokemon{},
		},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestPokemon_PokemonByTypeIdDataLoader_WithRowError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM pokemon LEFT JOIN .* WHERE pokemon_type.type_id IN (.*)").
		WithArgs("366a8621-9fa7-419b-b710-9100bcbb98d8", "5179f383-b765-4cc7-b9f9-8b1a3ba93019", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b").
		WillReturnRows(mockRowsForPokemonByTypeIdDataLoader(false, true, false, []string{"366a8621-9fa7-419b-b710-9100bcbb98d8", "5179f383-b765-4cc7-b9f9-8b1a3ba93019", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b"}))

	got, err := NewPokemon(db).PokemonByTypeIdDataLoader(context.Background())([]string{"366a8621-9fa7-419b-b710-9100bcbb98d8", "5179f383-b765-4cc7-b9f9-8b1a3ba93019", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b"})
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := []*model.PokemonList{
		{
			Total: 1,
			Pokemon: []model.Pokemon{
				castform,
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
			Pokemon: []model.Pokemon{
				castform,
			},
		},
		{
			Total: 1,
			Pokemon: []model.Pokemon{
				snorunt,
			},
		},
		{
			Total: 1,
			Pokemon: []model.Pokemon{
				bronzong,
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
			Pokemon: []model.Pokemon{},
		},
		{
			Total:   0,
			Pokemon: []model.Pokemon{},
		},
		{
			Total:   0,
			Pokemon: []model.Pokemon{},
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
			Pokemon: []model.Pokemon{},
		},
		{
			Total:   0,
			Pokemon: []model.Pokemon{},
		},
		{
			Total:   0,
			Pokemon: []model.Pokemon{},
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
			Pokemon: []model.Pokemon{
				castform,
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
					ID:   Normal.ID,
					Name: Normal.Name,
					Slug: Normal.Slug,
					NoDamageTo: model.TypeList{
						Total: 1,
						Types: []model.Type{
							Ghost,
						},
					},
					HalfDamageTo: model.TypeList{
						Total: 2,
						Types: []model.Type{
							Steel,
							Rock,
						},
					},
					DoubleDamageTo: model.TypeList{
						Total: 0,
						Types: nil,
					},
					NoDamageFrom: model.TypeList{
						Total: 1,
						Types: []model.Type{
							Ghost,
						},
					},
					HalfDamageFrom: model.TypeList{
						Total: 0,
						Types: nil,
					},
					DoubleDamageFrom: model.TypeList{
						Total: 1,
						Types: []model.Type{
							Fighting,
						},
					},
				},
			},
		},
	},
	Abilities: model.PokemonAbilityList{
		Total: 1,
		PokemonAbilities: []model.PokemonAbility{
			{
				Slot:     1,
				IsHidden: false,
				Ability: model.Ability{
					ID:     "0efe4eb9-537c-4b4c-92f6-d184a95b4923",
					Name:   "Forecast",
					Slug:   "forecast",
					Effect: "Changes castform's type and form to match the weather.",
				},
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
					ID:   Ice.ID,
					Name: Ice.Name,
					Slug: Ice.Slug,
					NoDamageTo: model.TypeList{
						Total: 0,
						Types: nil,
					},
					HalfDamageTo: model.TypeList{
						Total: 4,
						Types: []model.Type{
							Steel,
							Ice,
							Fire,
							Water,
						},
					},
					DoubleDamageTo: model.TypeList{
						Total: 4,
						Types: []model.Type{
							Grass,
							Ground,
							Flying,
							Dragon,
						},
					},
					NoDamageFrom: model.TypeList{
						Total: 0,
						Types: nil,
					},
					HalfDamageFrom: model.TypeList{
						Total: 1,
						Types: []model.Type{
							Ice,
						},
					},
					DoubleDamageFrom: model.TypeList{
						Total: 4,
						Types: []model.Type{
							Steel,
							Rock,
							Fighting,
							Fire,
						},
					},
				},
			},
		},
	},
	Abilities: model.PokemonAbilityList{
		Total: 3,
		PokemonAbilities: []model.PokemonAbility{
			{
				Slot:     1,
				IsHidden: false,
				Ability: model.Ability{
					ID:     "3eb38751-a341-457e-a211-1fc4641eac53",
					Name:   "Inner Focus",
					Slug:   "inner-focus",
					Effect: "Prevents flinching.",
				},
			},
			{
				Slot:     2,
				IsHidden: false,
				Ability: model.Ability{
					ID:     "673dd8ad-1494-49e1-86cd-9572df34540b",
					Name:   "Ice Body",
					Slug:   "ice-body",
					Effect: "Heals for 1/16 max HP after each turn during hail.  Protects against hail damage.",
				},
			},
			{
				Slot:     3,
				IsHidden: true,
				Ability: model.Ability{
					ID:     "ba77aff4-9bab-4dc7-acdc-e0bbba9b5c88",
					Name:   "Moody",
					Slug:   "moody",
					Effect: "Raises a random stat two stages and lowers another one stage after each turn.",
				},
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
					ID:   Steel.ID,
					Name: Steel.Name,
					Slug: Steel.Slug,
					NoDamageTo: model.TypeList{
						Total: 0,
						Types: nil,
					},
					HalfDamageTo: model.TypeList{
						Total: 4,
						Types: []model.Type{
							Steel,
							Electric,
							Fire,
							Water,
						},
					},
					DoubleDamageTo: model.TypeList{
						Total: 3,
						Types: []model.Type{
							Ice,
							Rock,
							Fairy,
						},
					},
					NoDamageFrom: model.TypeList{
						Total: 1,
						Types: []model.Type{
							Poison,
						},
					},
					HalfDamageFrom: model.TypeList{
						Total: 10,
						Types: []model.Type{
							Steel,
							Grass,
							Ice,
							Psychic,
							Normal,
							Flying,
							Rock,
							Bug,
							Fairy,
							Dragon,
						},
					},
					DoubleDamageFrom: model.TypeList{
						Total: 3,
						Types: []model.Type{
							Ground,
							Fighting,
							Fire,
						},
					},
				},
			},
			{
				Slot: 2,
				Type: model.Type{
					ID:   Psychic.ID,
					Name: Psychic.Name,
					Slug: Psychic.Slug,
					NoDamageTo: model.TypeList{
						Total: 1,
						Types: []model.Type{
							Dark,
						},
					},
					HalfDamageTo: model.TypeList{
						Total: 2,
						Types: []model.Type{
							Steel,
							Psychic,
						},
					},
					DoubleDamageTo: model.TypeList{
						Total: 2,
						Types: []model.Type{
							Poison,
							Fighting,
						},
					},
					NoDamageFrom: model.TypeList{
						Total: 0,
						Types: nil,
					},
					HalfDamageFrom: model.TypeList{
						Total: 2,
						Types: []model.Type{
							Psychic,
							Fighting,
						},
					},
					DoubleDamageFrom: model.TypeList{
						Total: 3,
						Types: []model.Type{
							Ghost,
							Bug,
							Dark,
						},
					},
				},
			},
		},
	},
	Abilities: model.PokemonAbilityList{
		Total: 3,
		PokemonAbilities: []model.PokemonAbility{
			{
				Slot:     1,
				IsHidden: false,
				Ability: model.Ability{
					ID:     "e55c279d-4554-4d5e-8120-7bf3a0477181",
					Name:   "Levitate",
					Slug:   "levitate",
					Effect: "Evades ground moves.",
				},
			},
			{
				Slot:     2,
				IsHidden: false,
				Ability: model.Ability{
					ID:     "9f0d876d-7e98-40d5-bfb3-2c0f079e2b26",
					Name:   "Heatproof",
					Slug:   "heatproof",
					Effect: "Halves damage from fire moves and burns.",
				},
			},
			{
				Slot:     3,
				IsHidden: true,
				Ability: model.Ability{
					ID:     "4acbc86c-4a5f-4f6e-99e8-4feab6337ad6",
					Name:   "Heavy Metal",
					Slug:   "heavy-metal",
					Effect: "Doubles the Pokémon's weight.",
				},
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
		rows.AddRow(castform.ID, castform.PokedexId, castform.Slug, castform.Name, castform.Sprite, castform.HP, castform.Attack, castform.Defense, castform.SpecialAttack, castform.SpecialDefense, castform.Speed, castform.IsBaby, castform.IsLegendary, castform.IsMythical, castform.Description, castform.Color.String(), castform.Shape.String(), castform.Habitat.String(), castform.IsDefaultVariant, castform.Genus, castform.Height, castform.Weight, `{"{\"slot\": 1, \"type\": {\"id\": \"366a8621-9fa7-419b-b710-9100bcbb98d8\", \"name\": \"Normal\", \"slug\": \"normal\", \"noDamageTo\": {\"types\": [{\"id\": \"027f1455-8e6f-4891-8c62-d75bb6c49dae\", \"name\": \"Ghost\", \"slug\": \"ghost\"}]}, \"halfDamageTo\": {\"types\": [{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}, {\"id\": \"5179f383-b765-4cc7-b9f9-8b1a3ba93019\", \"name\": \"Rock\", \"slug\": \"rock\"}]}, \"noDamageFrom\": {\"types\": [{\"id\": \"027f1455-8e6f-4891-8c62-d75bb6c49dae\", \"name\": \"Ghost\", \"slug\": \"ghost\"}]}, \"doubleDamageTo\": {\"types\": null}, \"halfDamageFrom\": {\"types\": null}, \"doubleDamageFrom\": {\"types\": [{\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}]}}}"}`, `{"{\"id\": \"1f0958a0-48ca-4160-9f18-7e5f06d96d27\", \"name\": \"Fairy\", \"slug\": \"fairy\"}","{\"id\": \"465ed2fa-0ff8-4cad-89af-e9db971026df\", \"name\": \"Amorphous\", \"slug\": \"indeterminate\"}"}`, `{"{\"slot\": 1, \"ability\": {\"id\": \"0efe4eb9-537c-4b4c-92f6-d184a95b4923\", \"name\": \"Forecast\", \"slug\": \"forecast\", \"effect\": \"Changes castform's type and form to match the weather.\"}, \"isHidden\": false}"}`).
			AddRow(snorunt.ID, snorunt.PokedexId, snorunt.Slug, snorunt.Name, snorunt.Sprite, snorunt.HP, snorunt.Attack, snorunt.Defense, snorunt.SpecialAttack, snorunt.SpecialDefense, snorunt.Speed, snorunt.IsBaby, snorunt.IsLegendary, snorunt.IsMythical, snorunt.Description, snorunt.Color.String(), snorunt.Shape.String(), snorunt.Habitat.String(), snorunt.IsDefaultVariant, snorunt.Genus, snorunt.Height, snorunt.Weight, `{"{\"slot\": 1, \"type\": {\"id\": \"1dcc9d3c-55d4-4d33-809a-d1580c6e6542\", \"name\": \"Ice\", \"slug\": \"ice\", \"noDamageTo\": {\"types\": null}, \"halfDamageTo\": {\"types\": [{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}, {\"id\": \"1dcc9d3c-55d4-4d33-809a-d1580c6e6542\", \"name\": \"Ice\", \"slug\": \"ice\"}, {\"id\": \"d43f57ab-e5b3-4912-a667-9f237d21d391\", \"name\": \"Fire\", \"slug\": \"fire\"}, {\"id\": \"de384e1c-89aa-44de-88fe-5e914e468f2b\", \"name\": \"Water\", \"slug\": \"water\"}]}, \"noDamageFrom\": {\"types\": null}, \"doubleDamageTo\": {\"types\": [{\"id\": \"07b9eb0f-e676-4649-bf2e-0e5ef2c2c2e3\", \"name\": \"Grass\", \"slug\": \"grass\"}, {\"id\": \"1b7d7950-305a-48fa-a771-01f7bc4dad8d\", \"name\": \"Ground\", \"slug\": \"ground\"}, {\"id\": \"4f09ea3c-2d93-4908-aabc-bc6e04ff24bb\", \"name\": \"Flying\", \"slug\": \"flying\"}, {\"id\": \"a82aa044-d8fd-43b3-9dd6-0ce0bfb29fb1\", \"name\": \"Dragon\", \"slug\": \"dragon\"}]}, \"halfDamageFrom\": {\"types\": [{\"id\": \"1dcc9d3c-55d4-4d33-809a-d1580c6e6542\", \"name\": \"Ice\", \"slug\": \"ice\"}]}, \"doubleDamageFrom\": {\"types\": [{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}, {\"id\": \"5179f383-b765-4cc7-b9f9-8b1a3ba93019\", \"name\": \"Rock\", \"slug\": \"rock\"}, {\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}, {\"id\": \"d43f57ab-e5b3-4912-a667-9f237d21d391\", \"name\": \"Fire\", \"slug\": \"fire\"}]}}}"}`, `{"{\"id\": \"1f0958a0-48ca-4160-9f18-7e5f06d96d27\", \"name\": \"Fairy\", \"slug\": \"fairy\"}","{\"id\": \"b140921f-74c9-4537-9a08-996277d4fcb4\", \"name\": \"Mineral\", \"slug\": \"mineral\"}"}`, `{"{\"slot\": 1, \"ability\": {\"id\": \"3eb38751-a341-457e-a211-1fc4641eac53\", \"name\": \"Inner Focus\", \"slug\": \"inner-focus\", \"effect\": \"Prevents flinching.\"}, \"isHidden\": false}","{\"slot\": 2, \"ability\": {\"id\": \"673dd8ad-1494-49e1-86cd-9572df34540b\", \"name\": \"Ice Body\", \"slug\": \"ice-body\", \"effect\": \"Heals for 1/16 max HP after each turn during hail.  Protects against hail damage.\"}, \"isHidden\": false}","{\"slot\": 3, \"ability\": {\"id\": \"ba77aff4-9bab-4dc7-acdc-e0bbba9b5c88\", \"name\": \"Moody\", \"slug\": \"moody\", \"effect\": \"Raises a random stat two stages and lowers another one stage after each turn.\"}, \"isHidden\": true}"}`).
			AddRow(bronzong.ID, bronzong.PokedexId, bronzong.Slug, bronzong.Name, bronzong.Sprite, bronzong.HP, bronzong.Attack, bronzong.Defense, bronzong.SpecialAttack, bronzong.SpecialDefense, bronzong.Speed, bronzong.IsBaby, bronzong.IsLegendary, bronzong.IsMythical, bronzong.Description, bronzong.Color.String(), bronzong.Shape.String(), nil, bronzong.IsDefaultVariant, bronzong.Genus, bronzong.Height, bronzong.Weight, `{"{\"slot\": 1, \"type\": {\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\", \"noDamageTo\": {\"types\": null}, \"halfDamageTo\": {\"types\": [{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}, {\"id\": \"b3468930-5d60-418f-aaaf-f16cbc93f08d\", \"name\": \"Electric\", \"slug\": \"electric\"}, {\"id\": \"d43f57ab-e5b3-4912-a667-9f237d21d391\", \"name\": \"Fire\", \"slug\": \"fire\"}, {\"id\": \"de384e1c-89aa-44de-88fe-5e914e468f2b\", \"name\": \"Water\", \"slug\": \"water\"}]}, \"noDamageFrom\": {\"types\": [{\"id\": \"42b31825-de68-4c1c-bea1-b32a290f1fef\", \"name\": \"Poison\", \"slug\": \"poison\"}]}, \"doubleDamageTo\": {\"types\": [{\"id\": \"1dcc9d3c-55d4-4d33-809a-d1580c6e6542\", \"name\": \"Ice\", \"slug\": \"ice\"}, {\"id\": \"5179f383-b765-4cc7-b9f9-8b1a3ba93019\", \"name\": \"Rock\", \"slug\": \"rock\"}, {\"id\": \"a248c127-8e9c-4f87-8513-c5dbc3385011\", \"name\": \"Fairy\", \"slug\": \"fairy\"}]}, \"halfDamageFrom\": {\"types\": [{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}, {\"id\": \"07b9eb0f-e676-4649-bf2e-0e5ef2c2c2e3\", \"name\": \"Grass\", \"slug\": \"grass\"}, {\"id\": \"1dcc9d3c-55d4-4d33-809a-d1580c6e6542\", \"name\": \"Ice\", \"slug\": \"ice\"}, {\"id\": \"2222c839-3c6e-4727-b6b5-a946bb8af5fa\", \"name\": \"Psychic\", \"slug\": \"psychic\"}, {\"id\": \"366a8621-9fa7-419b-b710-9100bcbb98d8\", \"name\": \"Normal\", \"slug\": \"normal\"}, {\"id\": \"4f09ea3c-2d93-4908-aabc-bc6e04ff24bb\", \"name\": \"Flying\", \"slug\": \"flying\"}, {\"id\": \"5179f383-b765-4cc7-b9f9-8b1a3ba93019\", \"name\": \"Rock\", \"slug\": \"rock\"}, {\"id\": \"56dddb9a-3623-43c5-8228-ea24d598afe7\", \"name\": \"Bug\", \"slug\": \"bug\"}, {\"id\": \"a248c127-8e9c-4f87-8513-c5dbc3385011\", \"name\": \"Fairy\", \"slug\": \"fairy\"}, {\"id\": \"a82aa044-d8fd-43b3-9dd6-0ce0bfb29fb1\", \"name\": \"Dragon\", \"slug\": \"dragon\"}]}, \"doubleDamageFrom\": {\"types\": [{\"id\": \"1b7d7950-305a-48fa-a771-01f7bc4dad8d\", \"name\": \"Ground\", \"slug\": \"ground\"}, {\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}, {\"id\": \"d43f57ab-e5b3-4912-a667-9f237d21d391\", \"name\": \"Fire\", \"slug\": \"fire\"}]}}}","{\"slot\": 2, \"type\": {\"id\": \"2222c839-3c6e-4727-b6b5-a946bb8af5fa\", \"name\": \"Psychic\", \"slug\": \"psychic\", \"noDamageTo\": {\"types\": [{\"id\": \"9ca47516-fff8-4f5e-8eb5-582c1f7c05af\", \"name\": \"Dark\", \"slug\": \"dark\"}]}, \"halfDamageTo\": {\"types\": [{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}, {\"id\": \"2222c839-3c6e-4727-b6b5-a946bb8af5fa\", \"name\": \"Psychic\", \"slug\": \"psychic\"}]}, \"noDamageFrom\": {\"types\": null}, \"doubleDamageTo\": {\"types\": [{\"id\": \"42b31825-de68-4c1c-bea1-b32a290f1fef\", \"name\": \"Poison\", \"slug\": \"poison\"}, {\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}]}, \"halfDamageFrom\": {\"types\": [{\"id\": \"2222c839-3c6e-4727-b6b5-a946bb8af5fa\", \"name\": \"Psychic\", \"slug\": \"psychic\"}, {\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}]}, \"doubleDamageFrom\": {\"types\": [{\"id\": \"027f1455-8e6f-4891-8c62-d75bb6c49dae\", \"name\": \"Ghost\", \"slug\": \"ghost\"}, {\"id\": \"56dddb9a-3623-43c5-8228-ea24d598afe7\", \"name\": \"Bug\", \"slug\": \"bug\"}, {\"id\": \"9ca47516-fff8-4f5e-8eb5-582c1f7c05af\", \"name\": \"Dark\", \"slug\": \"dark\"}]}}}"}`, `{"{\"id\": \"b140921f-74c9-4537-9a08-996277d4fcb4\", \"name\": \"Mineral\", \"slug\": \"mineral\"}"}`, `{"{\"slot\": 1, \"ability\": {\"id\": \"e55c279d-4554-4d5e-8120-7bf3a0477181\", \"name\": \"Levitate\", \"slug\": \"levitate\", \"effect\": \"Evades ground moves.\"}, \"isHidden\": false}","{\"slot\": 2, \"ability\": {\"id\": \"9f0d876d-7e98-40d5-bfb3-2c0f079e2b26\", \"name\": \"Heatproof\", \"slug\": \"heatproof\", \"effect\": \"Halves damage from fire moves and burns.\"}, \"isHidden\": false}","{\"slot\": 3, \"ability\": {\"id\": \"4acbc86c-4a5f-4f6e-99e8-4feab6337ad6\", \"name\": \"Heavy Metal\", \"slug\": \"heavy-metal\", \"effect\": \"Doubles the Pokémon's weight.\"}, \"isHidden\": true}"}`)
	}

	if hasRowError {
		rows.RowError(0, errors.New("scan error"))
	}
	return rows
}

func mockRowsForGetPokemonById(empty bool) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"id", "pokedex_id", "slug", "name", "sprite", "hp", "attack", "defense", "special_attack", "special_defense", "speed", "is_baby", "is_legendary", "is_mythical", "description", "color_enum", "shape_enum", "habitat_enum", "is_default_variant", "genus", "height", "weight", "types", "egg_groups", "abilities"})
	if !empty {
		rows.AddRow(castform.ID, castform.PokedexId, castform.Slug, castform.Name, castform.Sprite, castform.HP, castform.Attack, castform.Defense, castform.SpecialAttack, castform.SpecialDefense, castform.Speed, castform.IsBaby, castform.IsLegendary, castform.IsMythical, castform.Description, castform.Color.String(), castform.Shape.String(), castform.Habitat.String(), castform.IsDefaultVariant, castform.Genus, castform.Height, castform.Weight, `{"{\"slot\": 1, \"type\": {\"id\": \"366a8621-9fa7-419b-b710-9100bcbb98d8\", \"name\": \"Normal\", \"slug\": \"normal\", \"noDamageTo\": {\"types\": [{\"id\": \"027f1455-8e6f-4891-8c62-d75bb6c49dae\", \"name\": \"Ghost\", \"slug\": \"ghost\"}]}, \"halfDamageTo\": {\"types\": [{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}, {\"id\": \"5179f383-b765-4cc7-b9f9-8b1a3ba93019\", \"name\": \"Rock\", \"slug\": \"rock\"}]}, \"noDamageFrom\": {\"types\": [{\"id\": \"027f1455-8e6f-4891-8c62-d75bb6c49dae\", \"name\": \"Ghost\", \"slug\": \"ghost\"}]}, \"doubleDamageTo\": {\"types\": null}, \"halfDamageFrom\": {\"types\": null}, \"doubleDamageFrom\": {\"types\": [{\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}]}}}"}`, `{"{\"id\": \"1f0958a0-48ca-4160-9f18-7e5f06d96d27\", \"name\": \"Fairy\", \"slug\": \"fairy\"}","{\"id\": \"465ed2fa-0ff8-4cad-89af-e9db971026df\", \"name\": \"Amorphous\", \"slug\": \"indeterminate\"}"}`, `{"{\"slot\": 1, \"ability\": {\"id\": \"0efe4eb9-537c-4b4c-92f6-d184a95b4923\", \"name\": \"Forecast\", \"slug\": \"forecast\", \"effect\": \"Changes castform's type and form to match the weather.\"}, \"isHidden\": false}"}`)
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
		rows.AddRow(castform.ID, castform.PokedexId, castform.Slug, castform.Name, castform.Sprite, castform.HP, castform.Attack, castform.Defense, castform.SpecialAttack, castform.SpecialDefense, castform.Speed, castform.IsBaby, castform.IsLegendary, castform.IsMythical, castform.Description, castform.Color.String(), castform.Shape.String(), castform.Habitat.String(), castform.IsDefaultVariant, castform.Genus, castform.Height, castform.Weight, `{"{\"slot\": 1, \"type\": {\"id\": \"366a8621-9fa7-419b-b710-9100bcbb98d8\", \"name\": \"Normal\", \"slug\": \"normal\", \"noDamageTo\": {\"types\": [{\"id\": \"027f1455-8e6f-4891-8c62-d75bb6c49dae\", \"name\": \"Ghost\", \"slug\": \"ghost\"}]}, \"halfDamageTo\": {\"types\": [{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}, {\"id\": \"5179f383-b765-4cc7-b9f9-8b1a3ba93019\", \"name\": \"Rock\", \"slug\": \"rock\"}]}, \"noDamageFrom\": {\"types\": [{\"id\": \"027f1455-8e6f-4891-8c62-d75bb6c49dae\", \"name\": \"Ghost\", \"slug\": \"ghost\"}]}, \"doubleDamageTo\": {\"types\": null}, \"halfDamageFrom\": {\"types\": null}, \"doubleDamageFrom\": {\"types\": [{\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}]}}}"}`, `{"{\"id\": \"1f0958a0-48ca-4160-9f18-7e5f06d96d27\", \"name\": \"Fairy\", \"slug\": \"fairy\"}","{\"id\": \"465ed2fa-0ff8-4cad-89af-e9db971026df\", \"name\": \"Amorphous\", \"slug\": \"indeterminate\"}"}`, `{"{\"slot\": 1, \"ability\": {\"id\": \"0efe4eb9-537c-4b4c-92f6-d184a95b4923\", \"name\": \"Forecast\", \"slug\": \"forecast\", \"effect\": \"Changes castform's type and form to match the weather.\"}, \"isHidden\": false}"}`, ids[0]).
			AddRow(snorunt.ID, snorunt.PokedexId, snorunt.Slug, snorunt.Name, snorunt.Sprite, snorunt.HP, snorunt.Attack, snorunt.Defense, snorunt.SpecialAttack, snorunt.SpecialDefense, snorunt.Speed, snorunt.IsBaby, snorunt.IsLegendary, snorunt.IsMythical, snorunt.Description, snorunt.Color.String(), snorunt.Shape.String(), snorunt.Habitat.String(), snorunt.IsDefaultVariant, snorunt.Genus, snorunt.Height, snorunt.Weight, `{"{\"slot\": 1, \"type\": {\"id\": \"1dcc9d3c-55d4-4d33-809a-d1580c6e6542\", \"name\": \"Ice\", \"slug\": \"ice\", \"noDamageTo\": {\"types\": null}, \"halfDamageTo\": {\"types\": [{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}, {\"id\": \"1dcc9d3c-55d4-4d33-809a-d1580c6e6542\", \"name\": \"Ice\", \"slug\": \"ice\"}, {\"id\": \"d43f57ab-e5b3-4912-a667-9f237d21d391\", \"name\": \"Fire\", \"slug\": \"fire\"}, {\"id\": \"de384e1c-89aa-44de-88fe-5e914e468f2b\", \"name\": \"Water\", \"slug\": \"water\"}]}, \"noDamageFrom\": {\"types\": null}, \"doubleDamageTo\": {\"types\": [{\"id\": \"07b9eb0f-e676-4649-bf2e-0e5ef2c2c2e3\", \"name\": \"Grass\", \"slug\": \"grass\"}, {\"id\": \"1b7d7950-305a-48fa-a771-01f7bc4dad8d\", \"name\": \"Ground\", \"slug\": \"ground\"}, {\"id\": \"4f09ea3c-2d93-4908-aabc-bc6e04ff24bb\", \"name\": \"Flying\", \"slug\": \"flying\"}, {\"id\": \"a82aa044-d8fd-43b3-9dd6-0ce0bfb29fb1\", \"name\": \"Dragon\", \"slug\": \"dragon\"}]}, \"halfDamageFrom\": {\"types\": [{\"id\": \"1dcc9d3c-55d4-4d33-809a-d1580c6e6542\", \"name\": \"Ice\", \"slug\": \"ice\"}]}, \"doubleDamageFrom\": {\"types\": [{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}, {\"id\": \"5179f383-b765-4cc7-b9f9-8b1a3ba93019\", \"name\": \"Rock\", \"slug\": \"rock\"}, {\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}, {\"id\": \"d43f57ab-e5b3-4912-a667-9f237d21d391\", \"name\": \"Fire\", \"slug\": \"fire\"}]}}}"}`, `{"{\"id\": \"1f0958a0-48ca-4160-9f18-7e5f06d96d27\", \"name\": \"Fairy\", \"slug\": \"fairy\"}","{\"id\": \"b140921f-74c9-4537-9a08-996277d4fcb4\", \"name\": \"Mineral\", \"slug\": \"mineral\"}"}`, `{"{\"slot\": 1, \"ability\": {\"id\": \"3eb38751-a341-457e-a211-1fc4641eac53\", \"name\": \"Inner Focus\", \"slug\": \"inner-focus\", \"effect\": \"Prevents flinching.\"}, \"isHidden\": false}","{\"slot\": 2, \"ability\": {\"id\": \"673dd8ad-1494-49e1-86cd-9572df34540b\", \"name\": \"Ice Body\", \"slug\": \"ice-body\", \"effect\": \"Heals for 1/16 max HP after each turn during hail.  Protects against hail damage.\"}, \"isHidden\": false}","{\"slot\": 3, \"ability\": {\"id\": \"ba77aff4-9bab-4dc7-acdc-e0bbba9b5c88\", \"name\": \"Moody\", \"slug\": \"moody\", \"effect\": \"Raises a random stat two stages and lowers another one stage after each turn.\"}, \"isHidden\": true}"}`, ids[2]).
			AddRow(bronzong.ID, bronzong.PokedexId, bronzong.Slug, bronzong.Name, bronzong.Sprite, bronzong.HP, bronzong.Attack, bronzong.Defense, bronzong.SpecialAttack, bronzong.SpecialDefense, bronzong.Speed, bronzong.IsBaby, bronzong.IsLegendary, bronzong.IsMythical, bronzong.Description, bronzong.Color.String(), bronzong.Shape.String(), nil, bronzong.IsDefaultVariant, bronzong.Genus, bronzong.Height, bronzong.Weight, `{"{\"slot\": 1, \"type\": {\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\", \"noDamageTo\": {\"types\": null}, \"halfDamageTo\": {\"types\": [{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}, {\"id\": \"b3468930-5d60-418f-aaaf-f16cbc93f08d\", \"name\": \"Electric\", \"slug\": \"electric\"}, {\"id\": \"d43f57ab-e5b3-4912-a667-9f237d21d391\", \"name\": \"Fire\", \"slug\": \"fire\"}, {\"id\": \"de384e1c-89aa-44de-88fe-5e914e468f2b\", \"name\": \"Water\", \"slug\": \"water\"}]}, \"noDamageFrom\": {\"types\": [{\"id\": \"42b31825-de68-4c1c-bea1-b32a290f1fef\", \"name\": \"Poison\", \"slug\": \"poison\"}]}, \"doubleDamageTo\": {\"types\": [{\"id\": \"1dcc9d3c-55d4-4d33-809a-d1580c6e6542\", \"name\": \"Ice\", \"slug\": \"ice\"}, {\"id\": \"5179f383-b765-4cc7-b9f9-8b1a3ba93019\", \"name\": \"Rock\", \"slug\": \"rock\"}, {\"id\": \"a248c127-8e9c-4f87-8513-c5dbc3385011\", \"name\": \"Fairy\", \"slug\": \"fairy\"}]}, \"halfDamageFrom\": {\"types\": [{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}, {\"id\": \"07b9eb0f-e676-4649-bf2e-0e5ef2c2c2e3\", \"name\": \"Grass\", \"slug\": \"grass\"}, {\"id\": \"1dcc9d3c-55d4-4d33-809a-d1580c6e6542\", \"name\": \"Ice\", \"slug\": \"ice\"}, {\"id\": \"2222c839-3c6e-4727-b6b5-a946bb8af5fa\", \"name\": \"Psychic\", \"slug\": \"psychic\"}, {\"id\": \"366a8621-9fa7-419b-b710-9100bcbb98d8\", \"name\": \"Normal\", \"slug\": \"normal\"}, {\"id\": \"4f09ea3c-2d93-4908-aabc-bc6e04ff24bb\", \"name\": \"Flying\", \"slug\": \"flying\"}, {\"id\": \"5179f383-b765-4cc7-b9f9-8b1a3ba93019\", \"name\": \"Rock\", \"slug\": \"rock\"}, {\"id\": \"56dddb9a-3623-43c5-8228-ea24d598afe7\", \"name\": \"Bug\", \"slug\": \"bug\"}, {\"id\": \"a248c127-8e9c-4f87-8513-c5dbc3385011\", \"name\": \"Fairy\", \"slug\": \"fairy\"}, {\"id\": \"a82aa044-d8fd-43b3-9dd6-0ce0bfb29fb1\", \"name\": \"Dragon\", \"slug\": \"dragon\"}]}, \"doubleDamageFrom\": {\"types\": [{\"id\": \"1b7d7950-305a-48fa-a771-01f7bc4dad8d\", \"name\": \"Ground\", \"slug\": \"ground\"}, {\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}, {\"id\": \"d43f57ab-e5b3-4912-a667-9f237d21d391\", \"name\": \"Fire\", \"slug\": \"fire\"}]}}}","{\"slot\": 2, \"type\": {\"id\": \"2222c839-3c6e-4727-b6b5-a946bb8af5fa\", \"name\": \"Psychic\", \"slug\": \"psychic\", \"noDamageTo\": {\"types\": [{\"id\": \"9ca47516-fff8-4f5e-8eb5-582c1f7c05af\", \"name\": \"Dark\", \"slug\": \"dark\"}]}, \"halfDamageTo\": {\"types\": [{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}, {\"id\": \"2222c839-3c6e-4727-b6b5-a946bb8af5fa\", \"name\": \"Psychic\", \"slug\": \"psychic\"}]}, \"noDamageFrom\": {\"types\": null}, \"doubleDamageTo\": {\"types\": [{\"id\": \"42b31825-de68-4c1c-bea1-b32a290f1fef\", \"name\": \"Poison\", \"slug\": \"poison\"}, {\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}]}, \"halfDamageFrom\": {\"types\": [{\"id\": \"2222c839-3c6e-4727-b6b5-a946bb8af5fa\", \"name\": \"Psychic\", \"slug\": \"psychic\"}, {\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}]}, \"doubleDamageFrom\": {\"types\": [{\"id\": \"027f1455-8e6f-4891-8c62-d75bb6c49dae\", \"name\": \"Ghost\", \"slug\": \"ghost\"}, {\"id\": \"56dddb9a-3623-43c5-8228-ea24d598afe7\", \"name\": \"Bug\", \"slug\": \"bug\"}, {\"id\": \"9ca47516-fff8-4f5e-8eb5-582c1f7c05af\", \"name\": \"Dark\", \"slug\": \"dark\"}]}}}"}`, `{"{\"id\": \"b140921f-74c9-4537-9a08-996277d4fcb4\", \"name\": \"Mineral\", \"slug\": \"mineral\"}"}`, `{"{\"slot\": 1, \"ability\": {\"id\": \"e55c279d-4554-4d5e-8120-7bf3a0477181\", \"name\": \"Levitate\", \"slug\": \"levitate\", \"effect\": \"Evades ground moves.\"}, \"isHidden\": false}","{\"slot\": 2, \"ability\": {\"id\": \"9f0d876d-7e98-40d5-bfb3-2c0f079e2b26\", \"name\": \"Heatproof\", \"slug\": \"heatproof\", \"effect\": \"Halves damage from fire moves and burns.\"}, \"isHidden\": false}","{\"slot\": 3, \"ability\": {\"id\": \"4acbc86c-4a5f-4f6e-99e8-4feab6337ad6\", \"name\": \"Heavy Metal\", \"slug\": \"heavy-metal\", \"effect\": \"Doubles the Pokémon's weight.\"}, \"isHidden\": true}"}`, ids[2])
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
		rows.AddRow(castform.ID, castform.PokedexId, castform.Slug, castform.Name, castform.Sprite, castform.HP, castform.Attack, castform.Defense, castform.SpecialAttack, castform.SpecialDefense, castform.Speed, castform.IsBaby, castform.IsLegendary, castform.IsMythical, castform.Description, castform.Color.String(), castform.Shape.String(), castform.Habitat.String(), castform.IsDefaultVariant, castform.Genus, castform.Height, castform.Weight, `{"{\"slot\": 1, \"type\": {\"id\": \"366a8621-9fa7-419b-b710-9100bcbb98d8\", \"name\": \"Normal\", \"slug\": \"normal\", \"noDamageTo\": {\"types\": [{\"id\": \"027f1455-8e6f-4891-8c62-d75bb6c49dae\", \"name\": \"Ghost\", \"slug\": \"ghost\"}]}, \"halfDamageTo\": {\"types\": [{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}, {\"id\": \"5179f383-b765-4cc7-b9f9-8b1a3ba93019\", \"name\": \"Rock\", \"slug\": \"rock\"}]}, \"noDamageFrom\": {\"types\": [{\"id\": \"027f1455-8e6f-4891-8c62-d75bb6c49dae\", \"name\": \"Ghost\", \"slug\": \"ghost\"}]}, \"doubleDamageTo\": {\"types\": null}, \"halfDamageFrom\": {\"types\": null}, \"doubleDamageFrom\": {\"types\": [{\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}]}}}"}`, `{"{\"id\": \"1f0958a0-48ca-4160-9f18-7e5f06d96d27\", \"name\": \"Fairy\", \"slug\": \"fairy\"}","{\"id\": \"465ed2fa-0ff8-4cad-89af-e9db971026df\", \"name\": \"Amorphous\", \"slug\": \"indeterminate\"}"}`, `{"{\"slot\": 1, \"ability\": {\"id\": \"0efe4eb9-537c-4b4c-92f6-d184a95b4923\", \"name\": \"Forecast\", \"slug\": \"forecast\", \"effect\": \"Changes castform's type and form to match the weather.\"}, \"isHidden\": false}"}`).
			AddRow(snorunt.ID, snorunt.PokedexId, snorunt.Slug, snorunt.Name, snorunt.Sprite, snorunt.HP, snorunt.Attack, snorunt.Defense, snorunt.SpecialAttack, snorunt.SpecialDefense, snorunt.Speed, snorunt.IsBaby, snorunt.IsLegendary, snorunt.IsMythical, snorunt.Description, snorunt.Color.String(), snorunt.Shape.String(), snorunt.Habitat.String(), snorunt.IsDefaultVariant, snorunt.Genus, snorunt.Height, snorunt.Weight, `{"{\"slot\": 1, \"type\": {\"id\": \"1dcc9d3c-55d4-4d33-809a-d1580c6e6542\", \"name\": \"Ice\", \"slug\": \"ice\", \"noDamageTo\": {\"types\": null}, \"halfDamageTo\": {\"types\": [{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}, {\"id\": \"1dcc9d3c-55d4-4d33-809a-d1580c6e6542\", \"name\": \"Ice\", \"slug\": \"ice\"}, {\"id\": \"d43f57ab-e5b3-4912-a667-9f237d21d391\", \"name\": \"Fire\", \"slug\": \"fire\"}, {\"id\": \"de384e1c-89aa-44de-88fe-5e914e468f2b\", \"name\": \"Water\", \"slug\": \"water\"}]}, \"noDamageFrom\": {\"types\": null}, \"doubleDamageTo\": {\"types\": [{\"id\": \"07b9eb0f-e676-4649-bf2e-0e5ef2c2c2e3\", \"name\": \"Grass\", \"slug\": \"grass\"}, {\"id\": \"1b7d7950-305a-48fa-a771-01f7bc4dad8d\", \"name\": \"Ground\", \"slug\": \"ground\"}, {\"id\": \"4f09ea3c-2d93-4908-aabc-bc6e04ff24bb\", \"name\": \"Flying\", \"slug\": \"flying\"}, {\"id\": \"a82aa044-d8fd-43b3-9dd6-0ce0bfb29fb1\", \"name\": \"Dragon\", \"slug\": \"dragon\"}]}, \"halfDamageFrom\": {\"types\": [{\"id\": \"1dcc9d3c-55d4-4d33-809a-d1580c6e6542\", \"name\": \"Ice\", \"slug\": \"ice\"}]}, \"doubleDamageFrom\": {\"types\": [{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}, {\"id\": \"5179f383-b765-4cc7-b9f9-8b1a3ba93019\", \"name\": \"Rock\", \"slug\": \"rock\"}, {\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}, {\"id\": \"d43f57ab-e5b3-4912-a667-9f237d21d391\", \"name\": \"Fire\", \"slug\": \"fire\"}]}}}"}`, `{"{\"id\": \"1f0958a0-48ca-4160-9f18-7e5f06d96d27\", \"name\": \"Fairy\", \"slug\": \"fairy\"}","{\"id\": \"b140921f-74c9-4537-9a08-996277d4fcb4\", \"name\": \"Mineral\", \"slug\": \"mineral\"}"}`, `{"{\"slot\": 1, \"ability\": {\"id\": \"3eb38751-a341-457e-a211-1fc4641eac53\", \"name\": \"Inner Focus\", \"slug\": \"inner-focus\", \"effect\": \"Prevents flinching.\"}, \"isHidden\": false}","{\"slot\": 2, \"ability\": {\"id\": \"673dd8ad-1494-49e1-86cd-9572df34540b\", \"name\": \"Ice Body\", \"slug\": \"ice-body\", \"effect\": \"Heals for 1/16 max HP after each turn during hail.  Protects against hail damage.\"}, \"isHidden\": false}","{\"slot\": 3, \"ability\": {\"id\": \"ba77aff4-9bab-4dc7-acdc-e0bbba9b5c88\", \"name\": \"Moody\", \"slug\": \"moody\", \"effect\": \"Raises a random stat two stages and lowers another one stage after each turn.\"}, \"isHidden\": true}"}`).
			AddRow(bronzong.ID, bronzong.PokedexId, bronzong.Slug, bronzong.Name, bronzong.Sprite, bronzong.HP, bronzong.Attack, bronzong.Defense, bronzong.SpecialAttack, bronzong.SpecialDefense, bronzong.Speed, bronzong.IsBaby, bronzong.IsLegendary, bronzong.IsMythical, bronzong.Description, bronzong.Color.String(), bronzong.Shape.String(), nil, bronzong.IsDefaultVariant, bronzong.Genus, bronzong.Height, bronzong.Weight, `{"{\"slot\": 1, \"type\": {\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\", \"noDamageTo\": {\"types\": null}, \"halfDamageTo\": {\"types\": [{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}, {\"id\": \"b3468930-5d60-418f-aaaf-f16cbc93f08d\", \"name\": \"Electric\", \"slug\": \"electric\"}, {\"id\": \"d43f57ab-e5b3-4912-a667-9f237d21d391\", \"name\": \"Fire\", \"slug\": \"fire\"}, {\"id\": \"de384e1c-89aa-44de-88fe-5e914e468f2b\", \"name\": \"Water\", \"slug\": \"water\"}]}, \"noDamageFrom\": {\"types\": [{\"id\": \"42b31825-de68-4c1c-bea1-b32a290f1fef\", \"name\": \"Poison\", \"slug\": \"poison\"}]}, \"doubleDamageTo\": {\"types\": [{\"id\": \"1dcc9d3c-55d4-4d33-809a-d1580c6e6542\", \"name\": \"Ice\", \"slug\": \"ice\"}, {\"id\": \"5179f383-b765-4cc7-b9f9-8b1a3ba93019\", \"name\": \"Rock\", \"slug\": \"rock\"}, {\"id\": \"a248c127-8e9c-4f87-8513-c5dbc3385011\", \"name\": \"Fairy\", \"slug\": \"fairy\"}]}, \"halfDamageFrom\": {\"types\": [{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}, {\"id\": \"07b9eb0f-e676-4649-bf2e-0e5ef2c2c2e3\", \"name\": \"Grass\", \"slug\": \"grass\"}, {\"id\": \"1dcc9d3c-55d4-4d33-809a-d1580c6e6542\", \"name\": \"Ice\", \"slug\": \"ice\"}, {\"id\": \"2222c839-3c6e-4727-b6b5-a946bb8af5fa\", \"name\": \"Psychic\", \"slug\": \"psychic\"}, {\"id\": \"366a8621-9fa7-419b-b710-9100bcbb98d8\", \"name\": \"Normal\", \"slug\": \"normal\"}, {\"id\": \"4f09ea3c-2d93-4908-aabc-bc6e04ff24bb\", \"name\": \"Flying\", \"slug\": \"flying\"}, {\"id\": \"5179f383-b765-4cc7-b9f9-8b1a3ba93019\", \"name\": \"Rock\", \"slug\": \"rock\"}, {\"id\": \"56dddb9a-3623-43c5-8228-ea24d598afe7\", \"name\": \"Bug\", \"slug\": \"bug\"}, {\"id\": \"a248c127-8e9c-4f87-8513-c5dbc3385011\", \"name\": \"Fairy\", \"slug\": \"fairy\"}, {\"id\": \"a82aa044-d8fd-43b3-9dd6-0ce0bfb29fb1\", \"name\": \"Dragon\", \"slug\": \"dragon\"}]}, \"doubleDamageFrom\": {\"types\": [{\"id\": \"1b7d7950-305a-48fa-a771-01f7bc4dad8d\", \"name\": \"Ground\", \"slug\": \"ground\"}, {\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}, {\"id\": \"d43f57ab-e5b3-4912-a667-9f237d21d391\", \"name\": \"Fire\", \"slug\": \"fire\"}]}}}","{\"slot\": 2, \"type\": {\"id\": \"2222c839-3c6e-4727-b6b5-a946bb8af5fa\", \"name\": \"Psychic\", \"slug\": \"psychic\", \"noDamageTo\": {\"types\": [{\"id\": \"9ca47516-fff8-4f5e-8eb5-582c1f7c05af\", \"name\": \"Dark\", \"slug\": \"dark\"}]}, \"halfDamageTo\": {\"types\": [{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}, {\"id\": \"2222c839-3c6e-4727-b6b5-a946bb8af5fa\", \"name\": \"Psychic\", \"slug\": \"psychic\"}]}, \"noDamageFrom\": {\"types\": null}, \"doubleDamageTo\": {\"types\": [{\"id\": \"42b31825-de68-4c1c-bea1-b32a290f1fef\", \"name\": \"Poison\", \"slug\": \"poison\"}, {\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}]}, \"halfDamageFrom\": {\"types\": [{\"id\": \"2222c839-3c6e-4727-b6b5-a946bb8af5fa\", \"name\": \"Psychic\", \"slug\": \"psychic\"}, {\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}]}, \"doubleDamageFrom\": {\"types\": [{\"id\": \"027f1455-8e6f-4891-8c62-d75bb6c49dae\", \"name\": \"Ghost\", \"slug\": \"ghost\"}, {\"id\": \"56dddb9a-3623-43c5-8228-ea24d598afe7\", \"name\": \"Bug\", \"slug\": \"bug\"}, {\"id\": \"9ca47516-fff8-4f5e-8eb5-582c1f7c05af\", \"name\": \"Dark\", \"slug\": \"dark\"}]}}}"}`, `{"{\"id\": \"b140921f-74c9-4537-9a08-996277d4fcb4\", \"name\": \"Mineral\", \"slug\": \"mineral\"}"}`, `{"{\"slot\": 1, \"ability\": {\"id\": \"e55c279d-4554-4d5e-8120-7bf3a0477181\", \"name\": \"Levitate\", \"slug\": \"levitate\", \"effect\": \"Evades ground moves.\"}, \"isHidden\": false}","{\"slot\": 2, \"ability\": {\"id\": \"9f0d876d-7e98-40d5-bfb3-2c0f079e2b26\", \"name\": \"Heatproof\", \"slug\": \"heatproof\", \"effect\": \"Halves damage from fire moves and burns.\"}, \"isHidden\": false}","{\"slot\": 3, \"ability\": {\"id\": \"4acbc86c-4a5f-4f6e-99e8-4feab6337ad6\", \"name\": \"Heavy Metal\", \"slug\": \"heavy-metal\", \"effect\": \"Doubles the Pokémon's weight.\"}, \"isHidden\": true}"}`)
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
		rows.AddRow(castform.ID, castform.PokedexId, castform.Slug, castform.Name, castform.Sprite, castform.HP, castform.Attack, castform.Defense, castform.SpecialAttack, castform.SpecialDefense, castform.Speed, castform.IsBaby, castform.IsLegendary, castform.IsMythical, castform.Description, castform.Color.String(), castform.Shape.String(), castform.Habitat.String(), castform.IsDefaultVariant, castform.Genus, castform.Height, castform.Weight, `{"{\"slot\": 1, \"type\": {\"id\": \"366a8621-9fa7-419b-b710-9100bcbb98d8\", \"name\": \"Normal\", \"slug\": \"normal\", \"noDamageTo\": {\"types\": [{\"id\": \"027f1455-8e6f-4891-8c62-d75bb6c49dae\", \"name\": \"Ghost\", \"slug\": \"ghost\"}]}, \"halfDamageTo\": {\"types\": [{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}, {\"id\": \"5179f383-b765-4cc7-b9f9-8b1a3ba93019\", \"name\": \"Rock\", \"slug\": \"rock\"}]}, \"noDamageFrom\": {\"types\": [{\"id\": \"027f1455-8e6f-4891-8c62-d75bb6c49dae\", \"name\": \"Ghost\", \"slug\": \"ghost\"}]}, \"doubleDamageTo\": {\"types\": null}, \"halfDamageFrom\": {\"types\": null}, \"doubleDamageFrom\": {\"types\": [{\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}]}}}"}`, `{"{\"id\": \"1f0958a0-48ca-4160-9f18-7e5f06d96d27\", \"name\": \"Fairy\", \"slug\": \"fairy\"}","{\"id\": \"465ed2fa-0ff8-4cad-89af-e9db971026df\", \"name\": \"Amorphous\", \"slug\": \"indeterminate\"}"}`, `{"{\"slot\": 1, \"ability\": {\"id\": \"`+ids[0]+`\", \"name\": \"Forecast\", \"slug\": \"forecast\", \"effect\": \"Changes castform's type and form to match the weather.\"}, \"isHidden\": false}"}`).
			AddRow(snorunt.ID, snorunt.PokedexId, snorunt.Slug, snorunt.Name, snorunt.Sprite, snorunt.HP, snorunt.Attack, snorunt.Defense, snorunt.SpecialAttack, snorunt.SpecialDefense, snorunt.Speed, snorunt.IsBaby, snorunt.IsLegendary, snorunt.IsMythical, snorunt.Description, snorunt.Color.String(), snorunt.Shape.String(), snorunt.Habitat.String(), snorunt.IsDefaultVariant, snorunt.Genus, snorunt.Height, snorunt.Weight, `{"{\"slot\": 1, \"type\": {\"id\": \"1dcc9d3c-55d4-4d33-809a-d1580c6e6542\", \"name\": \"Ice\", \"slug\": \"ice\", \"noDamageTo\": {\"types\": null}, \"halfDamageTo\": {\"types\": [{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}, {\"id\": \"1dcc9d3c-55d4-4d33-809a-d1580c6e6542\", \"name\": \"Ice\", \"slug\": \"ice\"}, {\"id\": \"d43f57ab-e5b3-4912-a667-9f237d21d391\", \"name\": \"Fire\", \"slug\": \"fire\"}, {\"id\": \"de384e1c-89aa-44de-88fe-5e914e468f2b\", \"name\": \"Water\", \"slug\": \"water\"}]}, \"noDamageFrom\": {\"types\": null}, \"doubleDamageTo\": {\"types\": [{\"id\": \"07b9eb0f-e676-4649-bf2e-0e5ef2c2c2e3\", \"name\": \"Grass\", \"slug\": \"grass\"}, {\"id\": \"1b7d7950-305a-48fa-a771-01f7bc4dad8d\", \"name\": \"Ground\", \"slug\": \"ground\"}, {\"id\": \"4f09ea3c-2d93-4908-aabc-bc6e04ff24bb\", \"name\": \"Flying\", \"slug\": \"flying\"}, {\"id\": \"a82aa044-d8fd-43b3-9dd6-0ce0bfb29fb1\", \"name\": \"Dragon\", \"slug\": \"dragon\"}]}, \"halfDamageFrom\": {\"types\": [{\"id\": \"1dcc9d3c-55d4-4d33-809a-d1580c6e6542\", \"name\": \"Ice\", \"slug\": \"ice\"}]}, \"doubleDamageFrom\": {\"types\": [{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}, {\"id\": \"5179f383-b765-4cc7-b9f9-8b1a3ba93019\", \"name\": \"Rock\", \"slug\": \"rock\"}, {\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}, {\"id\": \"d43f57ab-e5b3-4912-a667-9f237d21d391\", \"name\": \"Fire\", \"slug\": \"fire\"}]}}}"}`, `{"{\"id\": \"1f0958a0-48ca-4160-9f18-7e5f06d96d27\", \"name\": \"Fairy\", \"slug\": \"fairy\"}","{\"id\": \"b140921f-74c9-4537-9a08-996277d4fcb4\", \"name\": \"Mineral\", \"slug\": \"mineral\"}"}`, `{"{\"slot\": 1, \"ability\": {\"id\": \"`+ids[1]+`\", \"name\": \"Inner Focus\", \"slug\": \"inner-focus\", \"effect\": \"Prevents flinching.\"}, \"isHidden\": false}","{\"slot\": 2, \"ability\": {\"id\": \"673dd8ad-1494-49e1-86cd-9572df34540b\", \"name\": \"Ice Body\", \"slug\": \"ice-body\", \"effect\": \"Heals for 1/16 max HP after each turn during hail.  Protects against hail damage.\"}, \"isHidden\": false}","{\"slot\": 3, \"ability\": {\"id\": \"ba77aff4-9bab-4dc7-acdc-e0bbba9b5c88\", \"name\": \"Moody\", \"slug\": \"moody\", \"effect\": \"Raises a random stat two stages and lowers another one stage after each turn.\"}, \"isHidden\": true}"}`).
			AddRow(bronzong.ID, bronzong.PokedexId, bronzong.Slug, bronzong.Name, bronzong.Sprite, bronzong.HP, bronzong.Attack, bronzong.Defense, bronzong.SpecialAttack, bronzong.SpecialDefense, bronzong.Speed, bronzong.IsBaby, bronzong.IsLegendary, bronzong.IsMythical, bronzong.Description, bronzong.Color.String(), bronzong.Shape.String(), nil, bronzong.IsDefaultVariant, bronzong.Genus, bronzong.Height, bronzong.Weight, `{"{\"slot\": 1, \"type\": {\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\", \"noDamageTo\": {\"types\": null}, \"halfDamageTo\": {\"types\": [{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}, {\"id\": \"b3468930-5d60-418f-aaaf-f16cbc93f08d\", \"name\": \"Electric\", \"slug\": \"electric\"}, {\"id\": \"d43f57ab-e5b3-4912-a667-9f237d21d391\", \"name\": \"Fire\", \"slug\": \"fire\"}, {\"id\": \"de384e1c-89aa-44de-88fe-5e914e468f2b\", \"name\": \"Water\", \"slug\": \"water\"}]}, \"noDamageFrom\": {\"types\": [{\"id\": \"42b31825-de68-4c1c-bea1-b32a290f1fef\", \"name\": \"Poison\", \"slug\": \"poison\"}]}, \"doubleDamageTo\": {\"types\": [{\"id\": \"1dcc9d3c-55d4-4d33-809a-d1580c6e6542\", \"name\": \"Ice\", \"slug\": \"ice\"}, {\"id\": \"5179f383-b765-4cc7-b9f9-8b1a3ba93019\", \"name\": \"Rock\", \"slug\": \"rock\"}, {\"id\": \"a248c127-8e9c-4f87-8513-c5dbc3385011\", \"name\": \"Fairy\", \"slug\": \"fairy\"}]}, \"halfDamageFrom\": {\"types\": [{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}, {\"id\": \"07b9eb0f-e676-4649-bf2e-0e5ef2c2c2e3\", \"name\": \"Grass\", \"slug\": \"grass\"}, {\"id\": \"1dcc9d3c-55d4-4d33-809a-d1580c6e6542\", \"name\": \"Ice\", \"slug\": \"ice\"}, {\"id\": \"2222c839-3c6e-4727-b6b5-a946bb8af5fa\", \"name\": \"Psychic\", \"slug\": \"psychic\"}, {\"id\": \"366a8621-9fa7-419b-b710-9100bcbb98d8\", \"name\": \"Normal\", \"slug\": \"normal\"}, {\"id\": \"4f09ea3c-2d93-4908-aabc-bc6e04ff24bb\", \"name\": \"Flying\", \"slug\": \"flying\"}, {\"id\": \"5179f383-b765-4cc7-b9f9-8b1a3ba93019\", \"name\": \"Rock\", \"slug\": \"rock\"}, {\"id\": \"56dddb9a-3623-43c5-8228-ea24d598afe7\", \"name\": \"Bug\", \"slug\": \"bug\"}, {\"id\": \"a248c127-8e9c-4f87-8513-c5dbc3385011\", \"name\": \"Fairy\", \"slug\": \"fairy\"}, {\"id\": \"a82aa044-d8fd-43b3-9dd6-0ce0bfb29fb1\", \"name\": \"Dragon\", \"slug\": \"dragon\"}]}, \"doubleDamageFrom\": {\"types\": [{\"id\": \"1b7d7950-305a-48fa-a771-01f7bc4dad8d\", \"name\": \"Ground\", \"slug\": \"ground\"}, {\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}, {\"id\": \"d43f57ab-e5b3-4912-a667-9f237d21d391\", \"name\": \"Fire\", \"slug\": \"fire\"}]}}}","{\"slot\": 2, \"type\": {\"id\": \"2222c839-3c6e-4727-b6b5-a946bb8af5fa\", \"name\": \"Psychic\", \"slug\": \"psychic\", \"noDamageTo\": {\"types\": [{\"id\": \"9ca47516-fff8-4f5e-8eb5-582c1f7c05af\", \"name\": \"Dark\", \"slug\": \"dark\"}]}, \"halfDamageTo\": {\"types\": [{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}, {\"id\": \"2222c839-3c6e-4727-b6b5-a946bb8af5fa\", \"name\": \"Psychic\", \"slug\": \"psychic\"}]}, \"noDamageFrom\": {\"types\": null}, \"doubleDamageTo\": {\"types\": [{\"id\": \"42b31825-de68-4c1c-bea1-b32a290f1fef\", \"name\": \"Poison\", \"slug\": \"poison\"}, {\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}]}, \"halfDamageFrom\": {\"types\": [{\"id\": \"2222c839-3c6e-4727-b6b5-a946bb8af5fa\", \"name\": \"Psychic\", \"slug\": \"psychic\"}, {\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}]}, \"doubleDamageFrom\": {\"types\": [{\"id\": \"027f1455-8e6f-4891-8c62-d75bb6c49dae\", \"name\": \"Ghost\", \"slug\": \"ghost\"}, {\"id\": \"56dddb9a-3623-43c5-8228-ea24d598afe7\", \"name\": \"Bug\", \"slug\": \"bug\"}, {\"id\": \"9ca47516-fff8-4f5e-8eb5-582c1f7c05af\", \"name\": \"Dark\", \"slug\": \"dark\"}]}}}"}`, `{"{\"id\": \"b140921f-74c9-4537-9a08-996277d4fcb4\", \"name\": \"Mineral\", \"slug\": \"mineral\"}"}`, `{"{\"slot\": 1, \"ability\": {\"id\": \"e55c279d-4554-4d5e-8120-7bf3a0477181\", \"name\": \"Levitate\", \"slug\": \"levitate\", \"effect\": \"Evades ground moves.\"}, \"isHidden\": false}","{\"slot\": 2, \"ability\": {\"id\": \"`+ids[2]+`\", \"name\": \"Heatproof\", \"slug\": \"heatproof\", \"effect\": \"Halves damage from fire moves and burns.\"}, \"isHidden\": false}","{\"slot\": 3, \"ability\": {\"id\": \"4acbc86c-4a5f-4f6e-99e8-4feab6337ad6\", \"name\": \"Heavy Metal\", \"slug\": \"heavy-metal\", \"effect\": \"Doubles the Pokémon's weight.\"}, \"isHidden\": true}"}`)
	}
	if hasRowError {
		rows.RowError(1, errors.New("scan error"))
	}
	return rows
}
