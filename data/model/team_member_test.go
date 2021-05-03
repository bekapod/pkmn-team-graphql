package model

import (
	"reflect"
	"testing"

	"github.com/go-test/deep"
)

func TestNewTeamMemberList(t *testing.T) {
	teamMembers := []*TeamMember{
		{
			ID: "123-456",
		},
		{
			ID: "456-789",
		},
	}

	exp := TeamMemberList{
		Total:       2,
		TeamMembers: teamMembers,
	}

	got := NewTeamMemberList(teamMembers)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestNewEmptyTeamMemberList(t *testing.T) {
	exp := TeamMemberList{
		Total:       0,
		TeamMembers: []*TeamMember{},
	}

	got := NewEmptyTeamMemberList()
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestTeamMemberList_AddTeamMember(t *testing.T) {
	teamMembers := TeamMemberList{}
	teamMember1 := &TeamMember{}
	teamMember2 := &TeamMember{}
	teamMembers.AddTeamMember(teamMember1)
	teamMembers.AddTeamMember(teamMember2)

	if teamMembers.Total != 2 {
		t.Errorf("expected Total of 2, but got %d instead", teamMembers.Total)
	}

	if !reflect.DeepEqual([]*TeamMember{teamMember1, teamMember2}, teamMembers.TeamMembers) {
		t.Errorf("the teamMembers added do not match")
	}
}
