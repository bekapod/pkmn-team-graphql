package model

import (
	"reflect"
	"testing"

	"github.com/go-test/deep"
)

func TestNewPokemonTypeList(t *testing.T) {
	pokemonTypes := []*PokemonType{
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

func TestPokemonTypeList_Scan(t *testing.T) {
	exp := PokemonTypeList{
		Total: 2,
		PokemonTypes: []*PokemonType{
			{
				Slot: 1,
				Type: Type{
					ID:   "05cd51bd-23ca-4736-b8ec-aa93aca68a8b",
					Name: "Steel",
					Slug: "steel",
				},
			},
			{
				Slot: 2,
				Type: Type{
					ID:   "2222c839-3c6e-4727-b6b5-a946bb8af5fa",
					Name: "Psychic",
					Slug: "psychic",
				},
			},
		},
	}
	got := PokemonTypeList{}
	err := (&got).Scan(`["{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slot\": 1, \"slug\": \"steel\"}","{\"id\": \"2222c839-3c6e-4727-b6b5-a946bb8af5fa\", \"name\": \"Psychic\", \"slot\": 2, \"slug\": \"psychic\"}"]`)

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestPokemonTypeList_Scan_TypeError(t *testing.T) {
	got := PokemonTypeList{}
	err := (&got).Scan(5)

	if err == nil {
		t.Error("expected an error but got nil")
	}
}

func TestPokemonTypeList_Scan_WrongJSONFormat(t *testing.T) {
	got := PokemonTypeList{}
	err := (&got).Scan(`5`)

	if err == nil {
		t.Error("expected an error but got nil")
	}
}

func TestPokemonTypeList_Scan_WrongTypeFormat(t *testing.T) {
	got := PokemonTypeList{}
	err := (&got).Scan(`["5","5"]`)

	if err == nil {
		t.Error("expected an error but got nil")
	}
}
