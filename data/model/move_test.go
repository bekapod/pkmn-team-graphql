package model

import (
	"testing"

	"github.com/go-test/deep"
)

func TestNewMoveList(t *testing.T) {
	moves := []Move{
		{
			ID: "123-456",
		},
		{
			ID: "456-789",
		},
	}

	exp := MoveList{
		Total: 2,
		Moves: moves,
	}

	got := NewMoveList(moves)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestNewEmptyMoveList(t *testing.T) {
	exp := MoveList{
		Total: 0,
		Moves: []Move{},
	}

	got := NewEmptyMoveList()
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestMoveList_AddMove(t *testing.T) {
	moves := MoveList{}
	move1 := Move{}
	move2 := Move{}
	moves.AddMove(move1)
	moves.AddMove(move2)

	if moves.Total != 2 {
		t.Errorf("expected Total of 2, but got %d instead", moves.Total)
	}
}
