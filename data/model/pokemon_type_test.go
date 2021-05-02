package model

import (
	"bekapod/pkmn-team-graphql/data/db"
	"reflect"
	"testing"

	"github.com/go-test/deep"
)

func TestNewPokemonTypeFromDb(t *testing.T) {
	pokemonType := db.PokemonTypeModel{
		InnerPokemonType: db.InnerPokemonType{
			TypeID:    "type-1",
			PokemonID: "pokemon-1",
			Slot:      2,
		},
	}
	exp := PokemonType{
		TypeID:    pokemonType.TypeID,
		PokemonID: pokemonType.PokemonID,
		Slot:      pokemonType.Slot,
	}

	got := NewPokemonTypeFromDb(pokemonType)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestNewPokemonTypeList(t *testing.T) {
	pokemonTypes := []*PokemonType{
		{
			TypeID:    "type-1",
			PokemonID: "pokemon-1",
		},
		{
			TypeID:    "type-2",
			PokemonID: "pokemon-1",
		},
	}

	exp := PokemonTypeList{
		Total:        2,
		PokemonTypes: pokemonTypes,
	}

	got := NewPokemonTypeList(pokemonTypes)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestNewEmptyPokemonTypeList(t *testing.T) {
	exp := PokemonTypeList{
		Total:        0,
		PokemonTypes: []*PokemonType{},
	}

	got := NewEmptyPokemonTypeList()
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestPokemonTypeList_AddPokemonType(t *testing.T) {
	pokemonTypes := PokemonTypeList{}
	pokemonType1 := &PokemonType{}
	pokemonType2 := &PokemonType{}
	pokemonTypes.AddPokemonType(pokemonType1)
	pokemonTypes.AddPokemonType(pokemonType2)

	if pokemonTypes.Total != 2 {
		t.Errorf("expected Total of 2, but got %d instead", pokemonTypes.Total)
	}

	if !reflect.DeepEqual([]*PokemonType{pokemonType1, pokemonType2}, pokemonTypes.PokemonTypes) {
		t.Errorf("the pokemon types added do not match")
	}
}
