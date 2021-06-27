package model

import (
	"bekapod/pkmn-team-graphql/data/db"
	"reflect"
	"testing"

	"github.com/go-test/deep"
)

func TestNewPokemonAbilityEdgeFromDb(t *testing.T) {
	pokemonAbility := db.PokemonAbilityModel{
		InnerPokemonAbility: db.InnerPokemonAbility{
			ID:        "123",
			AbilityID: "ability-1",
			PokemonID: "pokemon-1",
			Slot:      2,
			IsHidden:  true,
		},
	}
	exp := PokemonAbilityEdge{
		Cursor:   pokemonAbility.ID,
		Slot:     pokemonAbility.Slot,
		IsHidden: pokemonAbility.IsHidden,
		NodeID:   pokemonAbility.AbilityID,
	}

	got := NewPokemonAbilityEdgeFromDb(pokemonAbility)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestNewEmptyPokemonAbilityConnection(t *testing.T) {
	exp := PokemonAbilityConnection{
		PageInfo: &PageInfo{
			HasNextPage:     false,
			HasPreviousPage: false,
		},
		Edges: []*PokemonAbilityEdge{},
	}

	got := NewEmptyPokemonAbilityConnection()
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestPokemonAbilityConnection_AddEdge(t *testing.T) {
	pokemonAbilities := NewEmptyPokemonAbilityConnection()
	pokemonAbility1 := &PokemonAbilityEdge{}
	pokemonAbility2 := &PokemonAbilityEdge{}
	pokemonAbilities.AddEdge(pokemonAbility1)
	pokemonAbilities.AddEdge(pokemonAbility2)

	if !reflect.DeepEqual([]*PokemonAbilityEdge{pokemonAbility1, pokemonAbility2}, pokemonAbilities.Edges) {
		t.Errorf("the pokemon abilities added do not match")
	}
}

func TestNewPokemonWithAbilityEdgeFromDb(t *testing.T) {
	pokemonAbility := db.PokemonAbilityModel{
		InnerPokemonAbility: db.InnerPokemonAbility{
			ID:        "123",
			AbilityID: "ability-1",
			PokemonID: "pokemon-1",
			Slot:      2,
			IsHidden:  true,
		},
	}
	exp := PokemonWithAbilityEdge{
		Cursor:   pokemonAbility.ID,
		Slot:     pokemonAbility.Slot,
		IsHidden: pokemonAbility.IsHidden,
		NodeID:   pokemonAbility.PokemonID,
	}

	got := NewPokemonWithAbilityEdgeFromDb(pokemonAbility)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestNewEmptyPokemonWithAbilityConnection(t *testing.T) {
	exp := PokemonWithAbilityConnection{
		PageInfo: &PageInfo{
			HasNextPage:     false,
			HasPreviousPage: false,
		},
		Edges: []*PokemonWithAbilityEdge{},
	}

	got := NewEmptyPokemonWithAbilityConnection()
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestPokemonWithAbilityConnection_AddEdge(t *testing.T) {
	pokemonAbilities := NewEmptyPokemonWithAbilityConnection()
	pokemonAbility1 := &PokemonWithAbilityEdge{}
	pokemonAbility2 := &PokemonWithAbilityEdge{}
	pokemonAbilities.AddEdge(pokemonAbility1)
	pokemonAbilities.AddEdge(pokemonAbility2)

	if !reflect.DeepEqual([]*PokemonWithAbilityEdge{pokemonAbility1, pokemonAbility2}, pokemonAbilities.Edges) {
		t.Errorf("the pokemon abilities added do not match")
	}
}
