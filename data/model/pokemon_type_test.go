package model

import (
	"reflect"
	"testing"

	"github.com/go-test/deep"
)

func TestNewPokemonTypeList(t *testing.T) {
	pokemonTypes := []*PokemonType{
		{
			TypeID:    "123-456",
			PokemonID: "456-789",
			Slot:      1,
		},
		{
			TypeID:    "456-789",
			PokemonID: "123-456",
			Slot:      2,
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
		t.Errorf("the pokemonTypes added do not match")
	}
}
