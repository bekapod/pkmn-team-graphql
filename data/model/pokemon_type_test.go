package model

import (
	"bekapod/pkmn-team-graphql/data/db"
	"reflect"
	"testing"

	"github.com/go-test/deep"
)

func TestNewPokemonTypeEdgeFromDb(t *testing.T) {
	pokemonType := db.PokemonTypeModel{
		InnerPokemonType: db.InnerPokemonType{
			ID:        "123",
			TypeID:    "type-1",
			PokemonID: "pokemon-1",
			Slot:      2,
		},
	}
	exp := PokemonTypeEdge{
		Cursor: pokemonType.ID,
		NodeID: pokemonType.TypeID,
		Slot:   pokemonType.Slot,
	}

	got := NewPokemonTypeFromDb(pokemonType)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestNewEmptyPokemonTypeConnection(t *testing.T) {
	exp := PokemonTypeConnection{
		PageInfo: &PageInfo{
			HasNextPage:     false,
			HasPreviousPage: false,
		},
		Edges: []*PokemonTypeEdge{},
	}

	got := NewEmptyPokemonTypeConnection()
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestPokemonTypeConnection_AddEdge(t *testing.T) {
	pokemonTypes := NewEmptyPokemonTypeConnection()
	pokemonType1 := &PokemonTypeEdge{Cursor: "1"}
	pokemonType2 := &PokemonTypeEdge{Cursor: "2"}
	pokemonTypes.AddEdge(pokemonType1)
	pokemonTypes.AddEdge(pokemonType2)

	if *pokemonTypes.PageInfo.StartCursor != pokemonType1.Cursor {
		t.Errorf("expected start cursor to be %s but got %s", pokemonType1.Cursor, *pokemonTypes.PageInfo.StartCursor)
	}

	if *pokemonTypes.PageInfo.EndCursor != pokemonType2.Cursor {
		t.Errorf("expected start cursor to be %s but got %s", pokemonType2.Cursor, *pokemonTypes.PageInfo.StartCursor)
	}

	if !reflect.DeepEqual([]*PokemonTypeEdge{pokemonType1, pokemonType2}, pokemonTypes.Edges) {
		t.Errorf("the pokemon types added do not match")
	}
}

func TestNewPokemonWithTypeEdgeFromDb(t *testing.T) {
	pokemonType := db.PokemonTypeModel{
		InnerPokemonType: db.InnerPokemonType{
			ID:        "123",
			TypeID:    "type-1",
			PokemonID: "pokemon-1",
			Slot:      2,
		},
	}
	exp := PokemonWithTypeEdge{
		Cursor: pokemonType.ID,
		NodeID: pokemonType.PokemonID,
		Slot:   pokemonType.Slot,
	}

	got := NewPokemonWithTypeFromDb(pokemonType)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestNewEmptyPokemonWithTypeConnection(t *testing.T) {
	exp := PokemonWithTypeConnection{
		PageInfo: &PageInfo{
			HasNextPage:     false,
			HasPreviousPage: false,
		},
		Edges: []*PokemonWithTypeEdge{},
	}

	got := NewEmptyPokemonWithTypeConnection()
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestPokemonWithTypeConnection_AddEdge(t *testing.T) {
	pokemonTypes := NewEmptyPokemonWithTypeConnection()
	pokemonType1 := &PokemonWithTypeEdge{Cursor: "1"}
	pokemonType2 := &PokemonWithTypeEdge{Cursor: "2"}
	pokemonTypes.AddEdge(pokemonType1)
	pokemonTypes.AddEdge(pokemonType2)

	if *pokemonTypes.PageInfo.StartCursor != pokemonType1.Cursor {
		t.Errorf("expected start cursor to be %s but got %s", pokemonType1.Cursor, *pokemonTypes.PageInfo.StartCursor)
	}

	if *pokemonTypes.PageInfo.EndCursor != pokemonType2.Cursor {
		t.Errorf("expected start cursor to be %s but got %s", pokemonType2.Cursor, *pokemonTypes.PageInfo.StartCursor)
	}

	if !reflect.DeepEqual([]*PokemonWithTypeEdge{pokemonType1, pokemonType2}, pokemonTypes.Edges) {
		t.Errorf("the pokemon types added do not match")
	}
}
