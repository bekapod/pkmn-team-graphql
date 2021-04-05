package model

import (
	"testing"

	"github.com/go-test/deep"
)

func TestNewPokemonAbilityList(t *testing.T) {
	pokemonAbilities := []PokemonAbility{
		{
			Ability: Ability{ID: "123-456"},
			Slot:    1,
		},
		{
			Ability: Ability{ID: "456-789"},
			Slot:    2,
		},
	}

	exp := PokemonAbilityList{
		Total:            2,
		PokemonAbilities: pokemonAbilities,
	}

	got := NewPokemonAbilityList(pokemonAbilities)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestNewEmptyPokemonAbilityList(t *testing.T) {
	exp := PokemonAbilityList{
		Total:            0,
		PokemonAbilities: []PokemonAbility{},
	}

	got := NewEmptyPokemonAbilityList()
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestPokemonAbilityList_AddPokemonAbility(t *testing.T) {
	pokemonAbilities := PokemonAbilityList{}
	pokemonAbility1 := &PokemonAbility{}
	pokemonAbility2 := &PokemonAbility{}
	pokemonAbilities.AddPokemonAbility(pokemonAbility1)
	pokemonAbilities.AddPokemonAbility(pokemonAbility2)

	if pokemonAbilities.Total != 2 {
		t.Errorf("expected Total of 2, but got %d instead", pokemonAbilities.Total)
	}
}

func TestPokemonAbility_Scan(t *testing.T) {
	exp := PokemonAbility{
		Slot: 1,
		Ability: Ability{
			ID:     "5d654881-83d1-49db-81aa-e5ca84e04f95",
			Name:   "Static",
			Slug:   "static",
			Effect: "Has a 30% chance of paralyzing attacking Pokémon on contact.",
		},
	}
	got := PokemonAbility{}
	err := (&got).Scan([]uint8{123, 34, 115, 108, 111, 116, 34, 58, 32, 49, 44, 32, 34, 97, 98, 105, 108, 105, 116, 121, 34, 58, 32, 123, 34, 105, 100, 34, 58, 32, 34, 53, 100, 54, 53, 52, 56, 56, 49, 45, 56, 51, 100, 49, 45, 52, 57, 100, 98, 45, 56, 49, 97, 97, 45, 101, 53, 99, 97, 56, 52, 101, 48, 52, 102, 57, 53, 34, 44, 32, 34, 110, 97, 109, 101, 34, 58, 32, 34, 83, 116, 97, 116, 105, 99, 34, 44, 32, 34, 115, 108, 117, 103, 34, 58, 32, 34, 115, 116, 97, 116, 105, 99, 34, 44, 32, 34, 101, 102, 102, 101, 99, 116, 34, 58, 32, 34, 72, 97, 115, 32, 97, 32, 51, 48, 37, 32, 99, 104, 97, 110, 99, 101, 32, 111, 102, 32, 112, 97, 114, 97, 108, 121, 122, 105, 110, 103, 32, 97, 116, 116, 97, 99, 107, 105, 110, 103, 32, 80, 111, 107, 195, 169, 109, 111, 110, 32, 111, 110, 32, 99, 111, 110, 116, 97, 99, 116, 46, 34, 125, 44, 32, 34, 105, 115, 72, 105, 100, 100, 101, 110, 34, 58, 32, 102, 97, 108, 115, 101, 125})

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestPokemonAbilityList_Scan_Error(t *testing.T) {
	got := PokemonAbility{}
	err := (&got).Scan([]uint8{123, 34, 115, 108, 111, 116, 34, 58, 32, 49, 44, 32, 34, 116, 121, 112, 101, 34, 58, 32, 123, 34, 105, 100, 34, 58, 32, 34, 48, 55, 98, 57, 101, 98, 125})

	if err == nil {
		t.Error("expected an error but got nil")
	}
}

func TestPokemonAbilityList_Scan_AbilityError(t *testing.T) {
	got := PokemonAbility{}
	err := (&got).Scan(5)

	if err == nil {
		t.Error("expected an error but got nil")
	}
}
