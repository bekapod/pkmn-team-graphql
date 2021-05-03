package model

import (
	"bekapod/pkmn-team-graphql/data/db"
	"reflect"
	"testing"

	"github.com/go-test/deep"
)

func TestNewTeamFromDb(t *testing.T) {
	team := db.TeamModel{
		InnerTeam: db.InnerTeam{
			ID:   "123",
			Name: "Some team",
		},
		RelationsTeam: db.RelationsTeam{
			TeamMembers: []db.TeamMemberModel{
				{
					InnerTeamMember: db.InnerTeamMember{
						ID:        "team-member-1",
						Slot:      1,
						PokemonID: "pokemon-1",
					},
					RelationsTeamMember: db.RelationsTeamMember{
						Moves: []db.TeamMemberMoveModel{
							{
								InnerTeamMemberMove: db.InnerTeamMemberMove{
									ID:            "team-member-move-1",
									Slot:          1,
									PokemonMoveID: "pokemon-move-id-1",
								},
							},
							{
								InnerTeamMemberMove: db.InnerTeamMemberMove{
									ID:            "team-member-move-2",
									Slot:          2,
									PokemonMoveID: "pokemon-move-id-2",
								},
							},
						},
					},
				},
				{
					InnerTeamMember: db.InnerTeamMember{
						ID:        "team-member-2",
						Slot:      2,
						PokemonID: "pokemon-2",
					},
					RelationsTeamMember: db.RelationsTeamMember{
						Moves: []db.TeamMemberMoveModel{},
					},
				},
			},
		},
	}

	teamMemberMoves := []*TeamMemberMove{
		{
			ID:            "team-member-move-1",
			Slot:          1,
			PokemonMoveID: "pokemon-move-id-1",
		},
		{
			ID:            "team-member-move-2",
			Slot:          2,
			PokemonMoveID: "pokemon-move-id-2",
		},
	}
	expTeamMemberMoves := NewTeamMemberMoveList(teamMemberMoves)
	emptyTeamMemberMoves := NewEmptyTeamMemberMoveList()
	teamMembers := []*TeamMember{
		{
			ID:        "team-member-1",
			Slot:      1,
			PokemonID: "pokemon-1",
			Moves:     &expTeamMemberMoves,
		},
		{
			ID:        "team-member-2",
			Slot:      2,
			PokemonID: "pokemon-2",
			Moves:     &emptyTeamMemberMoves,
		},
	}
	expTeamMembers := NewTeamMemberList(teamMembers)
	exp := Team{
		ID:      team.ID,
		Name:    team.Name,
		Members: &expTeamMembers,
	}

	got := NewTeamFromDb(team)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

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
