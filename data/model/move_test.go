package model

import (
	"bekapod/pkmn-team-graphql/data/db"
	"reflect"
	"testing"

	"github.com/go-test/deep"
)

func TestNewMoveFromDb_WithNulls(t *testing.T) {
	move := db.MoveModel{
		InnerMove: db.InnerMove{
			ID:          "123",
			Slug:        "some-move",
			Name:        "Some Move",
			DamageClass: db.DamageClassPHYSICAL,
			Target:      db.MoveTargetALLOTHERPOKEMON,
			TypeID:      "some-type-id",
		},
	}
	exp := Move{
		ID:           move.ID,
		Slug:         move.Slug,
		Name:         move.Name,
		Accuracy:     nil,
		Pp:           nil,
		Power:        nil,
		DamageClass:  DamageClassPhysical,
		Effect:       nil,
		EffectChance: nil,
		Target:       MoveTargetAllOtherPokemon,
		TypeID:       move.TypeID,
	}

	got := NewMoveFromDb(move)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestNewMoveFromDb_WithFullData(t *testing.T) {
	accuracy := 70
	pp := 10
	power := 120
	effect := "Some move effect"
	effectChance := 10
	move := db.MoveModel{
		InnerMove: db.InnerMove{
			ID:           "123",
			Slug:         "some-move",
			Name:         "Some Move",
			Accuracy:     &accuracy,
			Pp:           &pp,
			Power:        &power,
			DamageClass:  db.DamageClassPHYSICAL,
			Effect:       &effect,
			EffectChance: &effectChance,
			Target:       db.MoveTargetALLOTHERPOKEMON,
			TypeID:       "some-type-id",
		},
	}
	exp := Move{
		ID:           move.ID,
		Slug:         move.Slug,
		Name:         move.Name,
		Accuracy:     &accuracy,
		Pp:           &pp,
		Power:        &power,
		DamageClass:  DamageClassPhysical,
		Effect:       &effect,
		EffectChance: &effectChance,
		Target:       MoveTargetAllOtherPokemon,
		TypeID:       move.TypeID,
	}

	got := NewMoveFromDb(move)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestNewMoveList(t *testing.T) {
	moves := []*Move{
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
		Moves: []*Move{},
	}

	got := NewEmptyMoveList()
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestMoveList_AddMove(t *testing.T) {
	moves := MoveList{}
	move1 := &Move{}
	move2 := &Move{}
	moves.AddMove(move1)
	moves.AddMove(move2)

	if moves.Total != 2 {
		t.Errorf("expected Total of 2, but got %d instead", moves.Total)
	}

	if !reflect.DeepEqual([]*Move{move1, move2}, moves.Moves) {
		t.Errorf("the moves added do not match")
	}
}
