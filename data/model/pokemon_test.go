package model

import (
	"bekapod/pkmn-team-graphql/data/db"
	"reflect"
	"testing"

	"github.com/go-test/deep"
)

func TestNewPokemonFromDb_WithNulls(t *testing.T) {
	eggGroup := db.EggGroupModel{
		InnerEggGroup: db.InnerEggGroup{
			ID:   "123",
			Slug: "some-egg-group",
			Name: "Some Egg Group",
		},
	}
	pokemon := db.PokemonModel{
		InnerPokemon: db.InnerPokemon{
			ID:               "123",
			Slug:             "some-pokemon",
			Name:             "Some Pokemon",
			Hp:               243,
			Attack:           432,
			Defense:          423,
			SpecialAttack:    65,
			SpecialDefense:   32,
			Speed:            43,
			IsBaby:           true,
			IsLegendary:      false,
			IsMythical:       true,
			Color:            db.ColorPINK,
			Shape:            db.ShapeBLOB,
			IsDefaultVariant: true,
			Genus:            "Some genus",
			Height:           43,
			Weight:           543,
		},
		RelationsPokemon: db.RelationsPokemon{
			EggGroups: []db.EggGroupModel{eggGroup},
		},
	}
	expEggGroup := NewEggGroupEdgeFromDb(eggGroup)
	expEggGroups := NewEmptyEggGroupConnection()
	expEggGroups.AddEdge(&expEggGroup)
	exp := Pokemon{
		ID:               pokemon.ID,
		Slug:             pokemon.Slug,
		Name:             pokemon.Name,
		Hp:               pokemon.Hp,
		Attack:           pokemon.Attack,
		Defense:          pokemon.Defense,
		SpecialAttack:    pokemon.SpecialAttack,
		SpecialDefense:   pokemon.SpecialDefense,
		Speed:            pokemon.Speed,
		IsBaby:           pokemon.IsBaby,
		IsLegendary:      pokemon.IsLegendary,
		IsMythical:       pokemon.IsMythical,
		Color:            ColorPink,
		Shape:            ShapeBlob,
		IsDefaultVariant: true,
		Genus:            "Some genus",
		Height:           43,
		Weight:           543,
		EggGroups:        &expEggGroups,
	}

	got := NewPokemonFromDb(pokemon)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestNewPokemonFromDb_WithFullData(t *testing.T) {
	eggGroup := db.EggGroupModel{
		InnerEggGroup: db.InnerEggGroup{
			ID:   "123",
			Slug: "some-egg-group",
			Name: "Some Egg Group",
		},
	}
	sprite := "pokemon.png"
	description := "Some pokemon description"
	dbHabitat := db.HabitatWATERSEDGE
	habitat := HabitatWatersEdge
	pokemon := db.PokemonModel{
		InnerPokemon: db.InnerPokemon{
			ID:               "123",
			Slug:             "some-pokemon",
			Name:             "Some Pokemon",
			Sprite:           &sprite,
			Hp:               243,
			Attack:           432,
			Defense:          423,
			SpecialAttack:    65,
			SpecialDefense:   32,
			Speed:            43,
			IsBaby:           true,
			IsLegendary:      false,
			IsMythical:       true,
			Description:      &description,
			Color:            db.ColorPINK,
			Shape:            db.ShapeBLOB,
			Habitat:          &dbHabitat,
			IsDefaultVariant: true,
			Genus:            "Some genus",
			Height:           43,
			Weight:           543,
		},
		RelationsPokemon: db.RelationsPokemon{
			EggGroups: []db.EggGroupModel{eggGroup},
		},
	}
	expEggGroup := NewEggGroupEdgeFromDb(eggGroup)
	expEggGroups := NewEmptyEggGroupConnection()
	expEggGroups.AddEdge(&expEggGroup)
	exp := Pokemon{
		ID:               pokemon.ID,
		Slug:             pokemon.Slug,
		Name:             pokemon.Name,
		Sprite:           &sprite,
		Hp:               pokemon.Hp,
		Attack:           pokemon.Attack,
		Defense:          pokemon.Defense,
		SpecialAttack:    pokemon.SpecialAttack,
		SpecialDefense:   pokemon.SpecialDefense,
		Speed:            pokemon.Speed,
		IsBaby:           pokemon.IsBaby,
		IsLegendary:      pokemon.IsLegendary,
		IsMythical:       pokemon.IsMythical,
		Description:      &description,
		Color:            ColorPink,
		Shape:            ShapeBlob,
		Habitat:          &habitat,
		IsDefaultVariant: true,
		Genus:            "Some genus",
		Height:           43,
		Weight:           543,
		EggGroups:        &expEggGroups,
	}

	got := NewPokemonFromDb(pokemon)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestNewPokemonEdgeFromDb(t *testing.T) {
	eggGroup := db.EggGroupModel{
		InnerEggGroup: db.InnerEggGroup{
			ID:   "123",
			Slug: "some-egg-group",
			Name: "Some Egg Group",
		},
	}
	pokemon := db.PokemonModel{
		InnerPokemon: db.InnerPokemon{
			ID:               "123",
			Slug:             "some-pokemon",
			Name:             "Some Pokemon",
			Hp:               243,
			Attack:           432,
			Defense:          423,
			SpecialAttack:    65,
			SpecialDefense:   32,
			Speed:            43,
			IsBaby:           true,
			IsLegendary:      false,
			IsMythical:       true,
			Color:            db.ColorPINK,
			Shape:            db.ShapeBLOB,
			IsDefaultVariant: true,
			Genus:            "Some genus",
			Height:           43,
			Weight:           543,
		},
		RelationsPokemon: db.RelationsPokemon{
			EggGroups: []db.EggGroupModel{eggGroup},
		},
	}
	expEggGroup := NewEggGroupEdgeFromDb(eggGroup)
	expEggGroups := NewEmptyEggGroupConnection()
	expEggGroups.AddEdge(&expEggGroup)
	exp := PokemonEdge{
		Cursor: pokemon.ID,
		Node: &Pokemon{
			ID:               pokemon.ID,
			Slug:             pokemon.Slug,
			Name:             pokemon.Name,
			Hp:               pokemon.Hp,
			Attack:           pokemon.Attack,
			Defense:          pokemon.Defense,
			SpecialAttack:    pokemon.SpecialAttack,
			SpecialDefense:   pokemon.SpecialDefense,
			Speed:            pokemon.Speed,
			IsBaby:           pokemon.IsBaby,
			IsLegendary:      pokemon.IsLegendary,
			IsMythical:       pokemon.IsMythical,
			Color:            ColorPink,
			Shape:            ShapeBlob,
			IsDefaultVariant: true,
			Genus:            "Some genus",
			Height:           43,
			Weight:           543,
			EggGroups:        &expEggGroups,
		},
	}

	got := NewPokemonEdgeFromDb(pokemon)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestNewEmptyPokemonConnection(t *testing.T) {
	exp := PokemonConnection{
		PageInfo: &PageInfo{
			HasNextPage:     false,
			HasPreviousPage: false,
		},
		Edges: []*PokemonEdge{},
	}

	got := NewEmptyPokemonConnection()
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestPokemonConnection_AddEdge(t *testing.T) {
	pokemon := NewEmptyPokemonConnection()
	pokemon1 := &PokemonEdge{}
	pokemon2 := &PokemonEdge{}
	pokemon.AddEdge(pokemon1)
	pokemon.AddEdge(pokemon2)

	if !reflect.DeepEqual([]*PokemonEdge{pokemon1, pokemon2}, pokemon.Edges) {
		t.Errorf("the pokemon added do not match")
	}
}
