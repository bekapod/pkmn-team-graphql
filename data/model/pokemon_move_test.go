package model

import (
	"bekapod/pkmn-team-graphql/data/db"
	"reflect"
	"testing"

	"github.com/go-test/deep"
)

func TestNewPokemonMoveFromDb(t *testing.T) {
	pokemonMove := db.PokemonMoveModel{
		InnerPokemonMove: db.InnerPokemonMove{
			MoveID:         "move-1",
			PokemonID:      "pokemon-1",
			LearnMethod:    db.MoveLearnMethodSTADIUMSURFINGPIKACHU,
			LevelLearnedAt: 30,
		},
	}
	exp := PokemonMove{
		MoveID:         pokemonMove.MoveID,
		PokemonID:      pokemonMove.PokemonID,
		LearnMethod:    MoveLearnMethodStadiumSurfingPikachu,
		LevelLearnedAt: pokemonMove.LevelLearnedAt,
	}

	got := NewPokemonMoveFromDb(pokemonMove)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestNewPokemonMoveList(t *testing.T) {
	pokemonMoves := []*PokemonMove{
		{
			MoveID:    "move-1",
			PokemonID: "pokemon-1",
		},
		{
			MoveID:    "move-2",
			PokemonID: "pokemon-1",
		},
	}

	exp := PokemonMoveList{
		Total:        2,
		PokemonMoves: pokemonMoves,
	}

	got := NewPokemonMoveList(pokemonMoves)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestNewEmptyPokemonMoveList(t *testing.T) {
	exp := PokemonMoveList{
		Total:        0,
		PokemonMoves: []*PokemonMove{},
	}

	got := NewEmptyPokemonMoveList()
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestPokemonMoveList_AddPokemonMove(t *testing.T) {
	pokemonMoves := PokemonMoveList{}
	pokemonMove1 := &PokemonMove{}
	pokemonMove2 := &PokemonMove{}
	pokemonMoves.AddPokemonMove(pokemonMove1)
	pokemonMoves.AddPokemonMove(pokemonMove2)

	if pokemonMoves.Total != 2 {
		t.Errorf("expected Total of 2, but got %d instead", pokemonMoves.Total)
	}

	if !reflect.DeepEqual([]*PokemonMove{pokemonMove1, pokemonMove2}, pokemonMoves.PokemonMoves) {
		t.Errorf("the pokemon moves added do not match")
	}
}
