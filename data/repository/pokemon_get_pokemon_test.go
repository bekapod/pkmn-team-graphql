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

	mock.ExpectQuery("SELECT .* FROM pokemon ORDER BY pokedex_id, slug ASC").
		WillReturnRows(mockRowsForGetPokemon(false, false, false))

	pokemon := []*model.Pokemon{
		{
			ID:             "3ab43625-a18d-4b11-98a3-86d7d959fbe1",
			Slug:           "castform-snowy",
			Name:           "Castform",
			PokedexId:      351,
			Sprite:         "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/10015.png",
			HP:             70,
			Attack:         70,
			Defense:        70,
			SpecialAttack:  70,
			SpecialDefense: 70,
			Speed:          70,
			IsBaby:         false,
			IsLegendary:    false,
			IsMythical:     false,
			Description:    "Its form changes depending on the weather.\nThe rougher conditions get, the rougher\nCastform’s disposition!",
		},
		{
			ID:             "51948cca-743a-4e6d-9c00-579140daccc5",
			Slug:           "snorunt",
			Name:           "Snorunt",
			PokedexId:      361,
			Sprite:         "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/361.png",
			HP:             50,
			Attack:         50,
			Defense:        50,
			SpecialAttack:  50,
			SpecialDefense: 50,
			Speed:          50,
			IsBaby:         false,
			IsLegendary:    false,
			IsMythical:     false,
			Description:    "Rich people from cold areas all share childhood\nmemories of playing with Snorunt.",
		},
		{
			ID:             "85da4120-96bb-42b1-8e8f-9f8bded11a31",
			Slug:           "bronzong",
			Name:           "Bronzong",
			PokedexId:      437,
			Sprite:         "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/437.png",
			HP:             67,
			Attack:         89,
			Defense:        116,
			SpecialAttack:  79,
			SpecialDefense: 116,
			Speed:          33,
			IsBaby:         false,
			IsLegendary:    false,
			IsMythical:     false,
			Description:    "",
		},
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

	mock.ExpectQuery("SELECT .* FROM pokemon ORDER BY pokedex_id, slug ASC").
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

	mock.ExpectQuery("SELECT .* FROM pokemon ORDER BY pokedex_id, slug ASC").
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

	mock.ExpectQuery("SELECT .* FROM pokemon ORDER BY pokedex_id, slug ASC").
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

	mock.ExpectQuery("SELECT .* FROM pokemon ORDER BY pokedex_id, slug ASC").
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

	mock.ExpectQuery("SELECT .* FROM pokemon WHERE id.*").
		WithArgs("3ab43625-a18d-4b11-98a3-86d7d959fbe1").
		WillReturnRows(mockRowsForGetPokemonById(false))

	exp := model.Pokemon{
		ID:             "3ab43625-a18d-4b11-98a3-86d7d959fbe1",
		Slug:           "castform-snowy",
		Name:           "Castform",
		PokedexId:      351,
		Sprite:         "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/10015.png",
		HP:             70,
		Attack:         70,
		Defense:        70,
		SpecialAttack:  70,
		SpecialDefense: 70,
		Speed:          70,
		IsBaby:         false,
		IsLegendary:    false,
		IsMythical:     false,
		Description:    "Its form changes depending on the weather.\nThe rougher conditions get, the rougher\nCastform’s disposition!",
	}
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

	mock.ExpectQuery("SELECT .* FROM pokemon WHERE id.*").
		WithArgs("3ab43625-a18d-4b11-98a3-86d7d959fbe1").
		WillReturnError(errors.New("I am Error."))

	_, err := NewPokemon(db).GetPokemonById(context.Background(), "3ab43625-a18d-4b11-98a3-86d7d959fbe1")
	if err == nil {
		t.Error("expected an error but got nil")
	}
}

func TestPokemon_GetPokemonById_WithNoRows(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM pokemon WHERE id.*").
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

	mock.ExpectQuery("SELECT .* FROM pokemon LEFT JOIN pokemon_move ON pokemon.id = pokemon_move.pokemon_id WHERE pokemon_move.move_id IN (.*)").
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
				{
					ID:             "3ab43625-a18d-4b11-98a3-86d7d959fbe1",
					Slug:           "castform-snowy",
					Name:           "Castform",
					PokedexId:      351,
					Sprite:         "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/10015.png",
					HP:             70,
					Attack:         70,
					Defense:        70,
					SpecialAttack:  70,
					SpecialDefense: 70,
					Speed:          70,
					IsBaby:         false,
					IsLegendary:    false,
					IsMythical:     false,
					Description:    "Its form changes depending on the weather.\nThe rougher conditions get, the rougher\nCastform’s disposition!",
				},
			},
		},
		nil,
		{
			Total: 2,
			Pokemon: []*model.Pokemon{
				{
					ID:             "51948cca-743a-4e6d-9c00-579140daccc5",
					Slug:           "snorunt",
					Name:           "Snorunt",
					PokedexId:      361,
					Sprite:         "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/361.png",
					HP:             50,
					Attack:         50,
					Defense:        50,
					SpecialAttack:  50,
					SpecialDefense: 50,
					Speed:          50,
					IsBaby:         false,
					IsLegendary:    false,
					IsMythical:     false,
					Description:    "Rich people from cold areas all share childhood\nmemories of playing with Snorunt.",
				},
				{
					ID:             "85da4120-96bb-42b1-8e8f-9f8bded11a31",
					Slug:           "bronzong",
					Name:           "Bronzong",
					PokedexId:      437,
					Sprite:         "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/437.png",
					HP:             67,
					Attack:         89,
					Defense:        116,
					SpecialAttack:  79,
					SpecialDefense: 116,
					Speed:          33,
					IsBaby:         false,
					IsLegendary:    false,
					IsMythical:     false,
					Description:    "",
				},
			},
		},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestPokemon_PokemonByMoveIdDataLoader_WithQueryError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM pokemon LEFT JOIN pokemon_move ON pokemon.id = pokemon_move.pokemon_id WHERE pokemon_move.move_id IN (.*)").
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

	mock.ExpectQuery("SELECT .* FROM pokemon LEFT JOIN pokemon_move ON pokemon.id = pokemon_move.pokemon_id WHERE pokemon_move.move_id IN (.*)").
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

	mock.ExpectQuery("SELECT .* FROM pokemon LEFT JOIN pokemon_move ON pokemon.id = pokemon_move.pokemon_id WHERE pokemon_move.move_id IN (.*)").
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
				{
					ID:             "3ab43625-a18d-4b11-98a3-86d7d959fbe1",
					Slug:           "castform-snowy",
					Name:           "Castform",
					PokedexId:      351,
					Sprite:         "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/10015.png",
					HP:             70,
					Attack:         70,
					Defense:        70,
					SpecialAttack:  70,
					SpecialDefense: 70,
					Speed:          70,
					IsBaby:         false,
					IsLegendary:    false,
					IsMythical:     false,
					Description:    "Its form changes depending on the weather.\nThe rougher conditions get, the rougher\nCastform’s disposition!",
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

func TestPokemon_PokemonByTypeIdDataLoader(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM pokemon LEFT JOIN pokemon_type ON pokemon.id = pokemon_type.pokemon_id WHERE pokemon_type.type_id IN (.*)").
		WithArgs("7b230001-b57e-4163-ac0a-3d157fc172e8", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "30e3e45e-9cc6-4fc2-a69a-fc91cd4908fb").
		WillReturnRows(mockRowsForPokemonByTypeIdDataLoader(false, false, false, []string{"7b230001-b57e-4163-ac0a-3d157fc172e8", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "30e3e45e-9cc6-4fc2-a69a-fc91cd4908fb"}))

	got, err := NewPokemon(db).PokemonByTypeIdDataLoader(context.Background())([]string{"7b230001-b57e-4163-ac0a-3d157fc172e8", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "30e3e45e-9cc6-4fc2-a69a-fc91cd4908fb"})
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	exp := []*model.PokemonList{
		{
			Total: 1,
			Pokemon: []*model.Pokemon{
				{
					ID:             "3ab43625-a18d-4b11-98a3-86d7d959fbe1",
					Slug:           "castform-snowy",
					Name:           "Castform",
					PokedexId:      351,
					Sprite:         "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/10015.png",
					HP:             70,
					Attack:         70,
					Defense:        70,
					SpecialAttack:  70,
					SpecialDefense: 70,
					Speed:          70,
					IsBaby:         false,
					IsLegendary:    false,
					IsMythical:     false,
					Description:    "Its form changes depending on the weather.\nThe rougher conditions get, the rougher\nCastform’s disposition!",
				},
			},
		},
		nil,
		{
			Total: 2,
			Pokemon: []*model.Pokemon{
				{
					ID:             "51948cca-743a-4e6d-9c00-579140daccc5",
					Slug:           "snorunt",
					Name:           "Snorunt",
					PokedexId:      361,
					Sprite:         "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/361.png",
					HP:             50,
					Attack:         50,
					Defense:        50,
					SpecialAttack:  50,
					SpecialDefense: 50,
					Speed:          50,
					IsBaby:         false,
					IsLegendary:    false,
					IsMythical:     false,
					Description:    "Rich people from cold areas all share childhood\nmemories of playing with Snorunt.",
				},
				{
					ID:             "85da4120-96bb-42b1-8e8f-9f8bded11a31",
					Slug:           "bronzong",
					Name:           "Bronzong",
					PokedexId:      437,
					Sprite:         "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/437.png",
					HP:             67,
					Attack:         89,
					Defense:        116,
					SpecialAttack:  79,
					SpecialDefense: 116,
					Speed:          33,
					IsBaby:         false,
					IsLegendary:    false,
					IsMythical:     false,
					Description:    "",
				},
			},
		},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestPokemon_PokemonByTypeIdDataLoader_WithQueryError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM pokemon LEFT JOIN pokemon_type ON pokemon.id = pokemon_type.pokemon_id WHERE pokemon_type.type_id IN (.*)").
		WithArgs("7b230001-b57e-4163-ac0a-3d157fc172e8", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "30e3e45e-9cc6-4fc2-a69a-fc91cd4908fb").
		WillReturnRows(mockRowsForPokemonByTypeIdDataLoader(false, false, false, []string{"7b230001-b57e-4163-ac0a-3d157fc172e8", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "30e3e45e-9cc6-4fc2-a69a-fc91cd4908fb"})).
		WillReturnError(errors.New("I am Error."))

	got, err := NewPokemon(db).PokemonByTypeIdDataLoader(context.Background())([]string{"7b230001-b57e-4163-ac0a-3d157fc172e8", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "30e3e45e-9cc6-4fc2-a69a-fc91cd4908fb"})
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

	mock.ExpectQuery("SELECT .* FROM pokemon LEFT JOIN pokemon_type ON pokemon.id = pokemon_type.pokemon_id WHERE pokemon_type.type_id IN (.*)").
		WithArgs("7b230001-b57e-4163-ac0a-3d157fc172e8", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "30e3e45e-9cc6-4fc2-a69a-fc91cd4908fb").
		WillReturnRows(mockRowsForPokemonByTypeIdDataLoader(false, false, true, []string{"7b230001-b57e-4163-ac0a-3d157fc172e8", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "30e3e45e-9cc6-4fc2-a69a-fc91cd4908fb"}))

	got, err := NewPokemon(db).PokemonByTypeIdDataLoader(context.Background())([]string{"7b230001-b57e-4163-ac0a-3d157fc172e8", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "30e3e45e-9cc6-4fc2-a69a-fc91cd4908fb"})
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

	mock.ExpectQuery("SELECT .* FROM pokemon LEFT JOIN pokemon_type ON pokemon.id = pokemon_type.pokemon_id WHERE pokemon_type.type_id IN (.*)").
		WithArgs("7b230001-b57e-4163-ac0a-3d157fc172e8", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "30e3e45e-9cc6-4fc2-a69a-fc91cd4908fb").
		WillReturnRows(mockRowsForPokemonByTypeIdDataLoader(false, true, false, []string{"7b230001-b57e-4163-ac0a-3d157fc172e8", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "30e3e45e-9cc6-4fc2-a69a-fc91cd4908fb"}))

	got, err := NewPokemon(db).PokemonByTypeIdDataLoader(context.Background())([]string{"7b230001-b57e-4163-ac0a-3d157fc172e8", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "30e3e45e-9cc6-4fc2-a69a-fc91cd4908fb"})
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := []*model.PokemonList{
		{
			Total: 1,
			Pokemon: []*model.Pokemon{
				{
					ID:             "3ab43625-a18d-4b11-98a3-86d7d959fbe1",
					Slug:           "castform-snowy",
					Name:           "Castform",
					PokedexId:      351,
					Sprite:         "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/10015.png",
					HP:             70,
					Attack:         70,
					Defense:        70,
					SpecialAttack:  70,
					SpecialDefense: 70,
					Speed:          70,
					IsBaby:         false,
					IsLegendary:    false,
					IsMythical:     false,
					Description:    "Its form changes depending on the weather.\nThe rougher conditions get, the rougher\nCastform’s disposition!",
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

func TestPokemon_PokemonByAbilityIdDataLoader(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM pokemon LEFT JOIN pokemon_ability ON pokemon.id = pokemon_ability.pokemon_id WHERE pokemon_ability.ability_id IN (.*)").
		WithArgs("7b230001-b57e-4163-ac0a-3d157fc172e8", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "30e3e45e-9cc6-4fc2-a69a-fc91cd4908fb").
		WillReturnRows(mockRowsForPokemonByAbilityIdDataLoader(false, false, false, []string{"7b230001-b57e-4163-ac0a-3d157fc172e8", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "30e3e45e-9cc6-4fc2-a69a-fc91cd4908fb"}))

	got, err := NewPokemon(db).PokemonByAbilityIdDataLoader(context.Background())([]string{"7b230001-b57e-4163-ac0a-3d157fc172e8", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "30e3e45e-9cc6-4fc2-a69a-fc91cd4908fb"})
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	exp := []*model.PokemonList{
		{
			Total: 1,
			Pokemon: []*model.Pokemon{
				{
					ID:             "3ab43625-a18d-4b11-98a3-86d7d959fbe1",
					Slug:           "castform-snowy",
					Name:           "Castform",
					PokedexId:      351,
					Sprite:         "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/10015.png",
					HP:             70,
					Attack:         70,
					Defense:        70,
					SpecialAttack:  70,
					SpecialDefense: 70,
					Speed:          70,
					IsBaby:         false,
					IsLegendary:    false,
					IsMythical:     false,
					Description:    "Its form changes depending on the weather.\nThe rougher conditions get, the rougher\nCastform’s disposition!",
				},
			},
		},
		nil,
		{
			Total: 2,
			Pokemon: []*model.Pokemon{
				{
					ID:             "51948cca-743a-4e6d-9c00-579140daccc5",
					Slug:           "snorunt",
					Name:           "Snorunt",
					PokedexId:      361,
					Sprite:         "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/361.png",
					HP:             50,
					Attack:         50,
					Defense:        50,
					SpecialAttack:  50,
					SpecialDefense: 50,
					Speed:          50,
					IsBaby:         false,
					IsLegendary:    false,
					IsMythical:     false,
					Description:    "Rich people from cold areas all share childhood\nmemories of playing with Snorunt.",
				},
				{
					ID:             "85da4120-96bb-42b1-8e8f-9f8bded11a31",
					Slug:           "bronzong",
					Name:           "Bronzong",
					PokedexId:      437,
					Sprite:         "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/437.png",
					HP:             67,
					Attack:         89,
					Defense:        116,
					SpecialAttack:  79,
					SpecialDefense: 116,
					Speed:          33,
					IsBaby:         false,
					IsLegendary:    false,
					IsMythical:     false,
					Description:    "",
				},
			},
		},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestPokemon_PokemonByAbilityIdDataLoader_WithQueryError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM pokemon LEFT JOIN pokemon_ability ON pokemon.id = pokemon_type.pokemon_id WHERE pokemon_ability.ability_id IN (.*)").
		WithArgs("7b230001-b57e-4163-ac0a-3d157fc172e8", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "30e3e45e-9cc6-4fc2-a69a-fc91cd4908fb").
		WillReturnRows(mockRowsForPokemonByAbilityIdDataLoader(false, false, false, []string{"7b230001-b57e-4163-ac0a-3d157fc172e8", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "30e3e45e-9cc6-4fc2-a69a-fc91cd4908fb"})).
		WillReturnError(errors.New("I am Error."))

	got, err := NewPokemon(db).PokemonByAbilityIdDataLoader(context.Background())([]string{"7b230001-b57e-4163-ac0a-3d157fc172e8", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "30e3e45e-9cc6-4fc2-a69a-fc91cd4908fb"})
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

	mock.ExpectQuery("SELECT .* FROM pokemon LEFT JOIN pokemon_ability ON pokemon.id = pokemon_ability.pokemon_id WHERE pokemon_ability.ability_id IN (.*)").
		WithArgs("7b230001-b57e-4163-ac0a-3d157fc172e8", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "30e3e45e-9cc6-4fc2-a69a-fc91cd4908fb").
		WillReturnRows(mockRowsForPokemonByAbilityIdDataLoader(false, false, true, []string{"7b230001-b57e-4163-ac0a-3d157fc172e8", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "30e3e45e-9cc6-4fc2-a69a-fc91cd4908fb"}))

	got, err := NewPokemon(db).PokemonByAbilityIdDataLoader(context.Background())([]string{"7b230001-b57e-4163-ac0a-3d157fc172e8", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "30e3e45e-9cc6-4fc2-a69a-fc91cd4908fb"})
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

	mock.ExpectQuery("SELECT .* FROM pokemon LEFT JOIN pokemon_ability ON pokemon.id = pokemon_ability.pokemon_id WHERE pokemon_ability.ability_id IN (.*)").
		WithArgs("7b230001-b57e-4163-ac0a-3d157fc172e8", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "30e3e45e-9cc6-4fc2-a69a-fc91cd4908fb").
		WillReturnRows(mockRowsForPokemonByAbilityIdDataLoader(false, true, false, []string{"7b230001-b57e-4163-ac0a-3d157fc172e8", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "30e3e45e-9cc6-4fc2-a69a-fc91cd4908fb"}))

	got, err := NewPokemon(db).PokemonByAbilityIdDataLoader(context.Background())([]string{"7b230001-b57e-4163-ac0a-3d157fc172e8", "b58c8651-90a1-4349-b3ae-b9c2caecc09a", "30e3e45e-9cc6-4fc2-a69a-fc91cd4908fb"})
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := []*model.PokemonList{
		{
			Total: 1,
			Pokemon: []*model.Pokemon{
				{
					ID:             "3ab43625-a18d-4b11-98a3-86d7d959fbe1",
					Slug:           "castform-snowy",
					Name:           "Castform",
					PokedexId:      351,
					Sprite:         "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/10015.png",
					HP:             70,
					Attack:         70,
					Defense:        70,
					SpecialAttack:  70,
					SpecialDefense: 70,
					Speed:          70,
					IsBaby:         false,
					IsLegendary:    false,
					IsMythical:     false,
					Description:    "Its form changes depending on the weather.\nThe rougher conditions get, the rougher\nCastform’s disposition!",
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

func TestPokemon_PokemonByPokemonIdDataLoader(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM pokemon WHERE id IN (.*)").
		WithArgs("56dddb9a-3623-43c5-8228-ea24d598afe7", "a248c127-8e9c-4f87-8513-c5dbc3385011", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b").
		WillReturnRows(mockRowsForPokemonByPokemonIdDataLoader(false, false, false, []string{"56dddb9a-3623-43c5-8228-ea24d598afe7", "a248c127-8e9c-4f87-8513-c5dbc3385011", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b"}))

	got, err := NewPokemon(db).PokemonByPokemonIdDataLoader(context.Background())([]string{"56dddb9a-3623-43c5-8228-ea24d598afe7", "a248c127-8e9c-4f87-8513-c5dbc3385011", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b"})
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	exp := []*model.Pokemon{
		{
			ID:             "56dddb9a-3623-43c5-8228-ea24d598afe7",
			Slug:           "castform-snowy",
			Name:           "Castform",
			PokedexId:      351,
			Sprite:         "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/10015.png",
			HP:             70,
			Attack:         70,
			Defense:        70,
			SpecialAttack:  70,
			SpecialDefense: 70,
			Speed:          70,
			IsBaby:         false,
			IsLegendary:    false,
			IsMythical:     false,
			Description:    "Its form changes depending on the weather.\nThe rougher conditions get, the rougher\nCastform’s disposition!",
		},
		nil,
		{
			ID:             "05cd51bd-23ca-4736-b8ec-aa93aca68a8b",
			Slug:           "snorunt",
			Name:           "Snorunt",
			PokedexId:      361,
			Sprite:         "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/361.png",
			HP:             50,
			Attack:         50,
			Defense:        50,
			SpecialAttack:  50,
			SpecialDefense: 50,
			Speed:          50,
			IsBaby:         false,
			IsLegendary:    false,
			IsMythical:     false,
			Description:    "Rich people from cold areas all share childhood\nmemories of playing with Snorunt.",
		},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestPokemon_PokemonByPokemonIdDataLoader_WithQueryError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM pokemon WHERE id IN (.*)").
		WithArgs("56dddb9a-3623-43c5-8228-ea24d598afe7", "a248c127-8e9c-4f87-8513-c5dbc3385011", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b").
		WillReturnRows(mockRowsForPokemonByPokemonIdDataLoader(false, false, false, []string{"56dddb9a-3623-43c5-8228-ea24d598afe7", "a248c127-8e9c-4f87-8513-c5dbc3385011", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b"})).
		WillReturnError(errors.New("I am Error."))

	got, err := NewPokemon(db).PokemonByPokemonIdDataLoader(context.Background())([]string{"56dddb9a-3623-43c5-8228-ea24d598afe7", "a248c127-8e9c-4f87-8513-c5dbc3385011", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b"})
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := []*model.Pokemon{
		nil,
		nil,
		nil,
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestPokemon_PokemonsByPokemonIdDataLoader_WithScanError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM pokemon WHERE id IN (.*)").
		WithArgs("56dddb9a-3623-43c5-8228-ea24d598afe7", "a248c127-8e9c-4f87-8513-c5dbc3385011", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b").
		WillReturnRows(mockRowsForPokemonByPokemonIdDataLoader(false, false, true, []string{"56dddb9a-3623-43c5-8228-ea24d598afe7", "a248c127-8e9c-4f87-8513-c5dbc3385011", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b"}))

	got, err := NewPokemon(db).PokemonByPokemonIdDataLoader(context.Background())([]string{"56dddb9a-3623-43c5-8228-ea24d598afe7", "a248c127-8e9c-4f87-8513-c5dbc3385011", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b"})
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := []*model.Pokemon{
		nil,
		nil,
		nil,
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestPokemon_PokemonByPokemonIdDataLoader_WithRowError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM pokemon WHERE id IN (.*)").
		WithArgs("56dddb9a-3623-43c5-8228-ea24d598afe7", "a248c127-8e9c-4f87-8513-c5dbc3385011", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b").
		WillReturnRows(mockRowsForPokemonByPokemonIdDataLoader(false, true, false, []string{"56dddb9a-3623-43c5-8228-ea24d598afe7", "a248c127-8e9c-4f87-8513-c5dbc3385011", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b"}))

	got, err := NewPokemon(db).PokemonByPokemonIdDataLoader(context.Background())([]string{"56dddb9a-3623-43c5-8228-ea24d598afe7", "a248c127-8e9c-4f87-8513-c5dbc3385011", "05cd51bd-23ca-4736-b8ec-aa93aca68a8b"})
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := []*model.Pokemon{
		{
			ID:             "56dddb9a-3623-43c5-8228-ea24d598afe7",
			Slug:           "castform-snowy",
			Name:           "Castform",
			PokedexId:      351,
			Sprite:         "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/10015.png",
			HP:             70,
			Attack:         70,
			Defense:        70,
			SpecialAttack:  70,
			SpecialDefense: 70,
			Speed:          70,
			IsBaby:         false,
			IsLegendary:    false,
			IsMythical:     false,
			Description:    "Its form changes depending on the weather.\nThe rougher conditions get, the rougher\nCastform’s disposition!",
		},
		nil,
		nil,
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func mockRowsForGetPokemon(empty bool, hasRowError bool, hasScanError bool) *sqlmock.Rows {
	if hasScanError {
		rows := sqlmock.NewRows([]string{"id"})
		rows.AddRow("3ab43625-a18d-4b11-98a3-86d7d959fbe1")
		return rows
	}
	rows := sqlmock.NewRows([]string{"id", "name", "slug", "pokedex_id", "sprite", "hp", "attack", "defense", "special_attack", "special_defense", "speed", "is_baby", "is_legendary", "is_mythical", "description"})
	if !empty {
		rows.AddRow("3ab43625-a18d-4b11-98a3-86d7d959fbe1", "Castform", "castform-snowy", 351, "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/10015.png", 70, 70, 70, 70, 70, 70, false, false, false, "Its form changes depending on the weather.\nThe rougher conditions get, the rougher\nCastform’s disposition!").
			AddRow("51948cca-743a-4e6d-9c00-579140daccc5", "Snorunt", "snorunt", 361, "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/361.png", 50, 50, 50, 50, 50, 50, false, false, false, "Rich people from cold areas all share childhood\nmemories of playing with Snorunt.").
			AddRow("85da4120-96bb-42b1-8e8f-9f8bded11a31", "Bronzong", "bronzong", 437, "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/437.png", 67, 89, 116, 79, 116, 33, false, false, false, "")
	}

	if hasRowError {
		rows.RowError(0, errors.New("scan error"))
	}
	return rows
}

func mockRowsForGetPokemonById(empty bool) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"id", "name", "slug", "pokedex_id", "sprite", "hp", "attack", "defense", "special_attack", "special_defense", "speed", "is_baby", "is_legendary", "is_mythical", "description"})
	if !empty {
		rows.AddRow("3ab43625-a18d-4b11-98a3-86d7d959fbe1", "Castform", "castform-snowy", 351, "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/10015.png", 70, 70, 70, 70, 70, 70, false, false, false, "Its form changes depending on the weather.\nThe rougher conditions get, the rougher\nCastform’s disposition!")
	}
	return rows
}

func mockRowsForPokemonByMoveIdDataLoader(empty bool, hasRowError bool, hasScanError bool, ids []string) *sqlmock.Rows {
	if hasScanError {
		rows := sqlmock.NewRows([]string{"id"})
		rows.AddRow("3ab43625-a18d-4b11-98a3-86d7d959fbe1")
		return rows
	}
	rows := sqlmock.NewRows([]string{"id", "name", "slug", "pokedex_id", "sprite", "hp", "attack", "defense", "special_attack", "special_defense", "speed", "is_baby", "is_legendary", "is_mythical", "description", "pokemon_move.move_id"})
	if !empty {
		rows.AddRow("3ab43625-a18d-4b11-98a3-86d7d959fbe1", "Castform", "castform-snowy", 351, "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/10015.png", 70, 70, 70, 70, 70, 70, false, false, false, "Its form changes depending on the weather.\nThe rougher conditions get, the rougher\nCastform’s disposition!", ids[0]).
			AddRow("51948cca-743a-4e6d-9c00-579140daccc5", "Snorunt", "snorunt", 361, "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/361.png", 50, 50, 50, 50, 50, 50, false, false, false, "Rich people from cold areas all share childhood\nmemories of playing with Snorunt.", ids[2]).
			AddRow("85da4120-96bb-42b1-8e8f-9f8bded11a31", "Bronzong", "bronzong", 437, "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/437.png", 67, 89, 116, 79, 116, 33, false, false, false, "", ids[2])
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
	rows := sqlmock.NewRows([]string{"id", "name", "slug", "pokedex_id", "sprite", "hp", "attack", "defense", "special_attack", "special_defense", "speed", "is_baby", "is_legendary", "is_mythical", "description", "pokemon_type.type_id"})
	if !empty {
		rows.AddRow("3ab43625-a18d-4b11-98a3-86d7d959fbe1", "Castform", "castform-snowy", 351, "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/10015.png", 70, 70, 70, 70, 70, 70, false, false, false, "Its form changes depending on the weather.\nThe rougher conditions get, the rougher\nCastform’s disposition!", ids[0]).
			AddRow("51948cca-743a-4e6d-9c00-579140daccc5", "Snorunt", "snorunt", 361, "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/361.png", 50, 50, 50, 50, 50, 50, false, false, false, "Rich people from cold areas all share childhood\nmemories of playing with Snorunt.", ids[2]).
			AddRow("85da4120-96bb-42b1-8e8f-9f8bded11a31", "Bronzong", "bronzong", 437, "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/437.png", 67, 89, 116, 79, 116, 33, false, false, false, "", ids[2])
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
	rows := sqlmock.NewRows([]string{"id", "name", "slug", "pokedex_id", "sprite", "hp", "attack", "defense", "special_attack", "special_defense", "speed", "is_baby", "is_legendary", "is_mythical", "description", "pokemon_type.type_id"})
	if !empty {
		rows.AddRow("3ab43625-a18d-4b11-98a3-86d7d959fbe1", "Castform", "castform-snowy", 351, "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/10015.png", 70, 70, 70, 70, 70, 70, false, false, false, "Its form changes depending on the weather.\nThe rougher conditions get, the rougher\nCastform’s disposition!", ids[0]).
			AddRow("51948cca-743a-4e6d-9c00-579140daccc5", "Snorunt", "snorunt", 361, "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/361.png", 50, 50, 50, 50, 50, 50, false, false, false, "Rich people from cold areas all share childhood\nmemories of playing with Snorunt.", ids[2]).
			AddRow("85da4120-96bb-42b1-8e8f-9f8bded11a31", "Bronzong", "bronzong", 437, "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/437.png", 67, 89, 116, 79, 116, 33, false, false, false, "", ids[2])
	}
	if hasRowError {
		rows.RowError(1, errors.New("scan error"))
	}
	return rows
}

func mockRowsForPokemonByPokemonIdDataLoader(empty bool, hasRowError bool, hasScanError bool, ids []string) *sqlmock.Rows {
	if hasScanError {
		rows := sqlmock.NewRows([]string{"id"})
		rows.AddRow("3ab43625-a18d-4b11-98a3-86d7d959fbe1")
		return rows
	}
	rows := sqlmock.NewRows([]string{"id", "name", "slug", "pokedex_id", "sprite", "hp", "attack", "defense", "special_attack", "special_defense", "speed", "is_baby", "is_legendary", "is_mythical", "description"})
	if !empty {
		rows.AddRow(ids[0], "Castform", "castform-snowy", 351, "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/10015.png", 70, 70, 70, 70, 70, 70, false, false, false, "Its form changes depending on the weather.\nThe rougher conditions get, the rougher\nCastform’s disposition!").
			AddRow(ids[2], "Snorunt", "snorunt", 361, "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/361.png", 50, 50, 50, 50, 50, 50, false, false, false, "Rich people from cold areas all share childhood\nmemories of playing with Snorunt.")
	}
	if hasRowError {
		rows.RowError(1, errors.New("scan error"))
	}
	return rows
}
