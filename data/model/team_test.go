package model

import (
	"bekapod/pkmn-team-graphql/data/db"
	"reflect"
	"testing"
	"time"

	"github.com/go-test/deep"
)

func TestNewTeamEdgeFromDb(t *testing.T) {
	team := db.TeamModel{
		InnerTeam: db.InnerTeam{
			ID:        "123",
			Name:      "Some team",
			CreatedAt: time.Date(2020, time.July, 11, 0, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2021, time.January, 10, 0, 0, 0, 0, time.UTC),
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
				},
				{
					InnerTeamMember: db.InnerTeamMember{
						ID:        "team-member-2",
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
				},
			},
		},
	}

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

	pokemon1Id := "pokemon-1"
	pokemon2Id := "pokemon-2"
	teamMembers := []*TeamMemberEdge{
		{
			Cursor: "team-member-1",
			Slot:   &team.TeamMembers()[0].Slot,
			Node: &TeamMember{
				ID:        "team-member-1",
				PokemonID: &pokemon1Id,
				Moves:     &expTeamMemberMoves,
			},
		},
		{
			Cursor: "team-member-2",
			Slot:   &team.TeamMembers()[1].Slot,
			Node: &TeamMember{
				ID:        "team-member-2",
				PokemonID: &pokemon2Id,
				Moves:     &expTeamMemberMoves,
			},
		},
	}
	expTeamMembers := NewEmptyTeamMemberConnection()
	expTeamMembers.AddEdge(teamMembers[0])
	expTeamMembers.AddEdge(teamMembers[1])
	exp := TeamEdge{
		Cursor: team.ID,
		Node: &Team{
			ID:        team.ID,
			Name:      team.Name,
			Members:   &expTeamMembers,
			CreatedAt: team.CreatedAt,
			UpdatedAt: team.UpdatedAt,
		},
	}

	got := NewTeamEdgeFromDb(team)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestNewEmptyTeamConnection(t *testing.T) {
	exp := TeamConnection{
		PageInfo: &PageInfo{
			HasNextPage:     false,
			HasPreviousPage: false,
		},
		Edges: []*TeamEdge{},
	}

	got := NewEmptyTeamConnection()
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestTeamConnection_AddEdge(t *testing.T) {
	teams := NewEmptyTeamConnection()
	team1 := &TeamEdge{Cursor: "1"}
	team2 := &TeamEdge{Cursor: "2"}
	teams.AddEdge(team1)
	teams.AddEdge(team2)

	if *teams.PageInfo.StartCursor != team1.Cursor {
		t.Errorf("expected start cursor to be %s but got %s", team1.Cursor, *teams.PageInfo.StartCursor)
	}

	if *teams.PageInfo.EndCursor != team2.Cursor {
		t.Errorf("expected start cursor to be %s but got %s", team2.Cursor, *teams.PageInfo.StartCursor)
	}

	if !reflect.DeepEqual([]*TeamEdge{team1, team2}, teams.Edges) {
		t.Errorf("the teams added do not match")
	}
}
