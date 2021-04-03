package model

import (
	"testing"

	"github.com/go-test/deep"
)

func TestNewPokemonTypeList(t *testing.T) {
	pokemonTypes := []PokemonType{
		{
			Type: Type{ID: "123-456"},
			Slot: 1,
		},
		{
			Type: Type{ID: "456-789"},
			Slot: 2,
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
		PokemonTypes: []PokemonType{},
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
}

func TestPokemonType_Scan(t *testing.T) {
	exp := PokemonType{
		Slot: 1,
		Type: Type{
			ID:   "07b9eb0f-e676-4649-bf2e-0e5ef2c2c2e3",
			Name: "Grass",
			Slug: "grass",
		},
	}
	got := PokemonType{}
	err := (&got).Scan([]uint8{123, 34, 115, 108, 111, 116, 34, 58, 32, 49, 44, 32, 34, 116, 121, 112, 101, 34, 58, 32, 123, 34, 105, 100, 34, 58, 32, 34, 48, 55, 98, 57, 101, 98, 48, 102, 45, 101, 54, 55, 54, 45, 52, 54, 52, 57, 45, 98, 102, 50, 101, 45, 48, 101, 53, 101, 102, 50, 99, 50, 99, 50, 101, 51, 34, 44, 32, 34, 110, 97, 109, 101, 34, 58, 32, 34, 71, 114, 97, 115, 115, 34, 44, 32, 34, 115, 108, 117, 103, 34, 58, 32, 34, 103, 114, 97, 115, 115, 34, 125, 125})

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestPokemonTypeList_Scan_Error(t *testing.T) {
	got := PokemonType{}
	err := (&got).Scan([]uint8{123, 34, 115, 108, 111, 116, 34, 58, 32, 49, 44, 32, 34, 116, 121, 112, 101, 34, 58, 32, 123, 34, 105, 100, 34, 58, 32, 34, 48, 55, 98, 57, 101, 98, 125})

	if err == nil {
		t.Error("expected an error but got nil")
	}
}

func TestPokemonTypeList_Scan_TypeError(t *testing.T) {
	got := PokemonType{}
	err := (&got).Scan(5)

	if err == nil {
		t.Error("expected an error but got nil")
	}
}
