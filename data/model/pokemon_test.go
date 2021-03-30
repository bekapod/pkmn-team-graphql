package model

import (
	"reflect"
	"testing"

	"github.com/go-test/deep"
)

func TestNewPokemonList(t *testing.T) {
	pokemon := []*Pokemon{
		{
			ID: "123-456",
		},
		{
			ID: "456-789",
		},
	}

	exp := PokemonList{
		Total:   2,
		Pokemon: pokemon,
	}

	got := NewPokemonList(pokemon)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestNewEmptyPokemonList(t *testing.T) {
	exp := PokemonList{
		Total:   0,
		Pokemon: []*Pokemon{},
	}

	got := NewEmptyPokemonList()
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestPokemonList_AddPokemon(t *testing.T) {
	pokemon := PokemonList{}
	pokemon1 := &Pokemon{}
	pokemon2 := &Pokemon{}
	pokemon.AddPokemon(pokemon1)
	pokemon.AddPokemon(pokemon2)

	if pokemon.Total != 2 {
		t.Errorf("expected Total of 2, but got %d instead", pokemon.Total)
	}

	if !reflect.DeepEqual([]*Pokemon{pokemon1, pokemon2}, pokemon.Pokemon) {
		t.Errorf("the pokemon added do not match")
	}
}
