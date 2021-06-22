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

func TestNewMoveEdgeFromDb(t *testing.T) {
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
	exp := MoveEdge{
		Cursor: move.ID,
		Node: &Move{
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
		},
	}

	got := NewMoveEdgeFromDb(move)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestNewEmptyMoveConnection(t *testing.T) {
	exp := MoveConnection{
		PageInfo: &PageInfo{
			HasNextPage:     false,
			HasPreviousPage: false,
		},
		Edges: []*MoveEdge{},
	}

	got := NewEmptyMoveConnection()
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestMoveConnection_AddEdge(t *testing.T) {
	moves := NewEmptyMoveConnection()
	move1 := &MoveEdge{Cursor: "1"}
	move2 := &MoveEdge{Cursor: "2"}
	moves.AddEdge(move1)
	moves.AddEdge(move2)

	if *moves.PageInfo.StartCursor != move1.Cursor {
		t.Errorf("expected start cursor to be %s but got %s", move1.Cursor, *moves.PageInfo.StartCursor)
	}

	if *moves.PageInfo.EndCursor != move2.Cursor {
		t.Errorf("expected start cursor to be %s but got %s", move2.Cursor, *moves.PageInfo.StartCursor)
	}

	if !reflect.DeepEqual([]*MoveEdge{move1, move2}, moves.Edges) {
		t.Errorf("the moves added do not match")
	}
}
