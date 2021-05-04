package model

import (
	"bekapod/pkmn-team-graphql/data/db"
	"reflect"
	"testing"

	"github.com/go-test/deep"
)

func TestNewTeamMemberFromDb(t *testing.T) {
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
	teamMember := db.TeamMemberModel{
		InnerTeamMember: db.InnerTeamMember{
			ID:        "team-member-1",
			Slot:      2,
			PokemonID: "pokemon-2",
		},
		RelationsTeamMember: db.RelationsTeamMember{
			Team: &team,
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
	}
	expTeam := Team{
		ID:      team.ID,
		Name:    team.Name,
		Members: &expTeamMembers,
	}
	exp := TeamMember{
		ID:        teamMember.ID,
		Slot:      teamMember.Slot,
		PokemonID: teamMember.PokemonID,
		Moves:     &expTeamMemberMoves,
		Team:      &expTeam,
	}

	got := NewTeamMemberFromDb(teamMember)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

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

func TestTeamMemberList_RemoveTeamMember(t *testing.T) {
	teamMembers := TeamMemberList{}
	teamMember1 := &TeamMember{ID: "team-member-1"}
	teamMember2 := &TeamMember{ID: "team-member-2"}
	teamMembers.AddTeamMember(teamMember1)
	teamMembers.AddTeamMember(teamMember2)
	teamMembers.RemoveTeamMember(teamMember2.ID)

	if teamMembers.Total != 1 {
		t.Errorf("expected Total of 1, but got %d instead", teamMembers.Total)
	}

	if !reflect.DeepEqual([]*TeamMember{teamMember1}, teamMembers.TeamMembers) {
		t.Errorf("the teamMembers do not match")
	}
}
