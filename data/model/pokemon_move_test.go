package model

import (
	"bekapod/pkmn-team-graphql/data/db"
	"reflect"
	"testing"

	"github.com/go-test/deep"
)

func TestNewPokemonMoveEdgeFromDb(t *testing.T) {
	pokemonMove := db.PokemonMoveModel{
		InnerPokemonMove: db.InnerPokemonMove{
			ID:             "123",
			MoveID:         "move-1",
			PokemonID:      "pokemon-1",
			LearnMethod:    db.MoveLearnMethodSTADIUMSURFINGPIKACHU,
			LevelLearnedAt: 30,
		},
	}
	exp := PokemonMoveEdge{
		Cursor:         pokemonMove.ID,
		NodeID:         pokemonMove.MoveID,
		LearnMethod:    MoveLearnMethodStadiumSurfingPikachu,
		LevelLearnedAt: pokemonMove.LevelLearnedAt,
	}

	got := NewPokemonMoveEdgeFromDb(pokemonMove)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestNewEmptyPokemonMoveConnection(t *testing.T) {
	exp := PokemonMoveConnection{
		PageInfo: &PageInfo{
			HasNextPage:     false,
			HasPreviousPage: false,
		},
		Edges: []*PokemonMoveEdge{},
	}

	got := NewEmptyPokemonMoveConnection()
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestPokemonMoveConnection_AddEdge(t *testing.T) {
	pokemonMoves := NewEmptyPokemonMoveConnection()
	pokemonMove1 := &PokemonMoveEdge{}
	pokemonMove2 := &PokemonMoveEdge{}
	pokemonMoves.AddEdge(pokemonMove1)
	pokemonMoves.AddEdge(pokemonMove2)

	if !reflect.DeepEqual([]*PokemonMoveEdge{pokemonMove1, pokemonMove2}, pokemonMoves.Edges) {
		t.Errorf("the pokemon moves added do not match")
	}
}

func TestNewPokemonWithMoveEdgeFromDb(t *testing.T) {
	pokemonMove := db.PokemonMoveModel{
		InnerPokemonMove: db.InnerPokemonMove{
			ID:             "123",
			MoveID:         "move-1",
			PokemonID:      "pokemon-1",
			LearnMethod:    db.MoveLearnMethodSTADIUMSURFINGPIKACHU,
			LevelLearnedAt: 30,
		},
	}
	exp := PokemonWithMoveEdge{
		Cursor:         pokemonMove.ID,
		NodeID:         pokemonMove.PokemonID,
		LearnMethod:    MoveLearnMethodStadiumSurfingPikachu,
		LevelLearnedAt: pokemonMove.LevelLearnedAt,
	}

	got := NewPokemonWithMoveEdgeFromDb(pokemonMove)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestNewEmptyPokemonWithMoveConnection(t *testing.T) {
	exp := PokemonWithMoveConnection{
		PageInfo: &PageInfo{
			HasNextPage:     false,
			HasPreviousPage: false,
		},
		Edges: []*PokemonWithMoveEdge{},
	}

	got := NewEmptyPokemonWithMoveConnection()
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestPokemonWithMoveConnection_AddEdge(t *testing.T) {
	pokemonMoves := NewEmptyPokemonWithMoveConnection()
	pokemonMove1 := &PokemonWithMoveEdge{}
	pokemonMove2 := &PokemonWithMoveEdge{}
	pokemonMoves.AddEdge(pokemonMove1)
	pokemonMoves.AddEdge(pokemonMove2)

	if !reflect.DeepEqual([]*PokemonWithMoveEdge{pokemonMove1, pokemonMove2}, pokemonMoves.Edges) {
		t.Errorf("the pokemon moves added do not match")
	}
}
