package model

import (
	"reflect"
	"testing"

	"github.com/go-test/deep"
)

func TestNewTeamList(t *testing.T) {
	teams := []*Team{
		{
			ID: "123-456",
		},
		{
			ID: "456-789",
		},
	}

	exp := TeamList{
		Total: 2,
		Teams: teams,
	}

	got := NewTeamList(teams)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestNewEmptyTeamList(t *testing.T) {
	exp := TeamList{
		Total: 0,
		Teams: []*Team{},
	}

	got := NewEmptyTeamList()
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestTeamList_AddTeam(t *testing.T) {
	teams := TeamList{}
	team1 := &Team{}
	team2 := &Team{}
	teams.AddTeam(team1)
	teams.AddTeam(team2)

	if teams.Total != 2 {
		t.Errorf("expected Total of 2, but got %d instead", teams.Total)
	}

	if !reflect.DeepEqual([]*Team{team1, team2}, teams.Teams) {
		t.Errorf("the teams added do not match")
	}
}
