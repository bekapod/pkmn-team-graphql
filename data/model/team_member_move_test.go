package model

import (
	"reflect"
	"testing"

	"github.com/go-test/deep"
)

func TestNewTeamMemberMoveList(t *testing.T) {
	teamMemberMoves := []*TeamMemberMove{
		{
			ID: "123-456",
		},
		{
			ID: "456-789",
		},
	}

	exp := TeamMemberMoveList{
		Total:           2,
		TeamMemberMoves: teamMemberMoves,
	}

	got := NewTeamMemberMoveList(teamMemberMoves)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestNewEmptyTeamMemberMoveList(t *testing.T) {
	exp := TeamMemberMoveList{
		Total:           0,
		TeamMemberMoves: []*TeamMemberMove{},
	}

	got := NewEmptyTeamMemberMoveList()
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestTeamMemberMoveList_AddTeamMemberMove(t *testing.T) {
	teamMemberMoves := TeamMemberMoveList{}
	teamMemberMove1 := &TeamMemberMove{}
	teamMemberMove2 := &TeamMemberMove{}
	teamMemberMoves.AddTeamMemberMove(teamMemberMove1)
	teamMemberMoves.AddTeamMemberMove(teamMemberMove2)

	if teamMemberMoves.Total != 2 {
		t.Errorf("expected Total of 2, but got %d instead", teamMemberMoves.Total)
	}

	if !reflect.DeepEqual([]*TeamMemberMove{teamMemberMove1, teamMemberMove2}, teamMemberMoves.TeamMemberMoves) {
		t.Errorf("the teamMemberMoves added do not match")
	}
}
