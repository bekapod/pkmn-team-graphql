package model

import (
	"bekapod/pkmn-team-graphql/data/db"
	"reflect"
	"testing"

	"github.com/go-test/deep"
)

func TestNewPokemonAbilityFromDb(t *testing.T) {
	pokemonAbility := db.PokemonAbilityModel{
		InnerPokemonAbility: db.InnerPokemonAbility{
			AbilityID: "ability-1",
			PokemonID: "pokemon-1",
			Slot:      2,
			IsHidden:  true,
		},
	}
	exp := PokemonAbility{
		AbilityID: pokemonAbility.AbilityID,
		PokemonID: pokemonAbility.PokemonID,
		Slot:      pokemonAbility.Slot,
		IsHidden:  pokemonAbility.IsHidden,
	}

	got := NewPokemonAbilityFromDb(pokemonAbility)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestNewPokemonAbilityList(t *testing.T) {
	pokemonAbilities := []*PokemonAbility{
		{
			AbilityID: "ability-1",
			PokemonID: "pokemon-1",
			Slot:      1,
		},
		{
			AbilityID: "ability-2",
			PokemonID: "pokemon-1",
			Slot:      2,
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
		PokemonAbilities: []*PokemonAbility{},
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

	if !reflect.DeepEqual([]*PokemonAbility{pokemonAbility1, pokemonAbility2}, pokemonAbilities.PokemonAbilities) {
		t.Errorf("the pokemon abilities added do not match")
	}
}
