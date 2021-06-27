package model

import (
	"bekapod/pkmn-team-graphql/data/db"
	"reflect"
	"testing"

	"github.com/go-test/deep"
)

func TestNewTeamMemberFromDb(t *testing.T) {
	teamMemberMoves := []*TeamMemberMoveEdge{
		{
			Cursor:         "team-member-move-1",
			NodeID:         "move-1",
			LearnMethod:    MoveLearnMethodEgg,
			LevelLearnedAt: 0,
		},
		{
			Cursor:         "team-member-move-2",
			NodeID:         "move-2",
			LearnMethod:    MoveLearnMethodLevelUp,
			LevelLearnedAt: 22,
		},
	}
	expTeamMemberMoves := NewEmptyTeamMemberMoveConnection()
	expTeamMemberMoves.AddEdge(teamMemberMoves[0])
	expTeamMemberMoves.AddEdge(teamMemberMoves[1])

	teamMember := db.TeamMemberModel{
		InnerTeamMember: db.InnerTeamMember{
			ID:        "team-member-1",
			Slot:      2,
			PokemonID: "pokemon-2",
		},
		RelationsTeamMember: db.RelationsTeamMember{
			Moves: []db.TeamMemberMoveModel{
				{
					InnerTeamMemberMove: db.InnerTeamMemberMove{
						ID:            "team-member-move-1",
						Slot:          2,
						TeamMemberID:  "team-member-id",
						PokemonMoveID: "pokemon-move-id",
					},
					RelationsTeamMemberMove: db.RelationsTeamMemberMove{
						PokemonMove: &db.PokemonMoveModel{
							InnerPokemonMove: db.InnerPokemonMove{
								ID:             "123",
								MoveID:         "move-1",
								PokemonID:      "pokemon-1",
								LearnMethod:    db.MoveLearnMethodEGG,
								LevelLearnedAt: 0,
							},
						},
					},
				},
				{
					InnerTeamMemberMove: db.InnerTeamMemberMove{
						ID:            "team-member-move-2",
						Slot:          2,
						TeamMemberID:  "team-member-id",
						PokemonMoveID: "pokemon-move-id",
					},
					RelationsTeamMemberMove: db.RelationsTeamMemberMove{
						PokemonMove: &db.PokemonMoveModel{
							InnerPokemonMove: db.InnerPokemonMove{
								ID:             "123",
								MoveID:         "move-2",
								PokemonID:      "pokemon-1",
								LearnMethod:    db.MoveLearnMethodLEVELUP,
								LevelLearnedAt: 22,
							},
						},
					},
				},
			},
		},
	}
	exp := TeamMember{
		ID:        teamMember.ID,
		PokemonID: &teamMember.PokemonID,
		Moves:     &expTeamMemberMoves,
	}

	got := NewTeamMemberFromDb(teamMember)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestNewTeamMemberEdgeFromDb(t *testing.T) {
	teamMemberMoves := []*TeamMemberMoveEdge{
		{
			Cursor:         "team-member-move-1",
			NodeID:         "move-1",
			LearnMethod:    MoveLearnMethodEgg,
			LevelLearnedAt: 0,
		},
		{
			Cursor:         "team-member-move-2",
			NodeID:         "move-2",
			LearnMethod:    MoveLearnMethodLevelUp,
			LevelLearnedAt: 22,
		},
	}
	expTeamMemberMoves := NewEmptyTeamMemberMoveConnection()
	expTeamMemberMoves.AddEdge(teamMemberMoves[0])
	expTeamMemberMoves.AddEdge(teamMemberMoves[1])
	teamMember := db.TeamMemberModel{
		InnerTeamMember: db.InnerTeamMember{
			ID:        "team-member-1",
			Slot:      2,
			PokemonID: "pokemon-2",
		},
		RelationsTeamMember: db.RelationsTeamMember{
			Moves: []db.TeamMemberMoveModel{
				{
					InnerTeamMemberMove: db.InnerTeamMemberMove{
						ID:            "team-member-move-1",
						Slot:          2,
						TeamMemberID:  "team-member-id",
						PokemonMoveID: "pokemon-move-1",
					},
					RelationsTeamMemberMove: db.RelationsTeamMemberMove{
						PokemonMove: &db.PokemonMoveModel{
							InnerPokemonMove: db.InnerPokemonMove{
								ID:             "pokemon-move-1",
								MoveID:         "move-1",
								PokemonID:      "pokemon-1",
								LearnMethod:    db.MoveLearnMethodEGG,
								LevelLearnedAt: 0,
							},
						},
					},
				},
				{
					InnerTeamMemberMove: db.InnerTeamMemberMove{
						ID:            "team-member-move-2",
						Slot:          2,
						TeamMemberID:  "team-member-id",
						PokemonMoveID: "pokemon-move-2",
					},
					RelationsTeamMemberMove: db.RelationsTeamMemberMove{
						PokemonMove: &db.PokemonMoveModel{
							InnerPokemonMove: db.InnerPokemonMove{
								ID:             "pokemon-move-2",
								MoveID:         "move-2",
								PokemonID:      "pokemon-1",
								LearnMethod:    db.MoveLearnMethodLEVELUP,
								LevelLearnedAt: 22,
							},
						},
					},
				},
			},
		},
	}
	exp := TeamMemberEdge{
		Cursor: teamMember.ID,
		Slot:   &teamMember.Slot,
		Node: &TeamMember{
			ID:        teamMember.ID,
			PokemonID: &teamMember.PokemonID,
			Moves:     &expTeamMemberMoves,
		},
	}

	got := NewTeamMemberEdgeFromDb(teamMember)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}
func TestNewEmptyTeamMemberConnection(t *testing.T) {
	exp := TeamMemberConnection{
		PageInfo: &PageInfo{
			HasNextPage:     false,
			HasPreviousPage: false,
		},
		Edges: []*TeamMemberEdge{},
	}

	got := NewEmptyTeamMemberConnection()
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestTeamMemberConnection_AddEdge(t *testing.T) {
	teamMembers := NewEmptyTeamMemberConnection()
	teamMember1 := &TeamMemberEdge{Cursor: "1"}
	teamMember2 := &TeamMemberEdge{Cursor: "2"}
	teamMembers.AddEdge(teamMember1)
	teamMembers.AddEdge(teamMember2)

	if *teamMembers.PageInfo.StartCursor != teamMember1.Cursor {
		t.Errorf("expected start cursor to be %s but got %s", teamMember1.Cursor, *teamMembers.PageInfo.StartCursor)
	}

	if *teamMembers.PageInfo.EndCursor != teamMember2.Cursor {
		t.Errorf("expected start cursor to be %s but got %s", teamMember2.Cursor, *teamMembers.PageInfo.StartCursor)
	}

	if !reflect.DeepEqual([]*TeamMemberEdge{teamMember1, teamMember2}, teamMembers.Edges) {
		t.Errorf("the teamMembers added do not match")
	}
}
