package model

import (
	"bekapod/pkmn-team-graphql/data/db"
	"reflect"
	"testing"

	"github.com/go-test/deep"
)

func TestNewTeamMemberMoveEdgeFromDb(t *testing.T) {
	teamMemberMove := db.TeamMemberMoveModel{
		InnerTeamMemberMove: db.InnerTeamMemberMove{
			ID:            "123",
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
					LearnMethod:    db.MoveLearnMethodSTADIUMSURFINGPIKACHU,
					LevelLearnedAt: 30,
				},
			},
		},
	}
	exp := TeamMemberMoveEdge{
		Cursor:         teamMemberMove.ID,
		NodeID:         "move-1",
		LearnMethod:    MoveLearnMethod(teamMemberMove.PokemonMove().LearnMethod),
		LevelLearnedAt: teamMemberMove.PokemonMove().LevelLearnedAt,
	}

	got := NewTeamMemberMoveEdgeFromDb(teamMemberMove)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestNewEmptyTeamMemberMoveConnection(t *testing.T) {
	exp := TeamMemberMoveConnection{
		PageInfo: &PageInfo{
			HasNextPage:     false,
			HasPreviousPage: false,
		},
		Edges: []*TeamMemberMoveEdge{},
	}

	got := NewEmptyTeamMemberMoveConnection()
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestTeamMemberMoveConnection_AddEdge(t *testing.T) {
	teamMemberMoves := NewEmptyTeamMemberMoveConnection()
	teamMemberMove1 := &TeamMemberMoveEdge{}
	teamMemberMove2 := &TeamMemberMoveEdge{}
	teamMemberMoves.AddEdge(teamMemberMove1)
	teamMemberMoves.AddEdge(teamMemberMove2)

	if !reflect.DeepEqual([]*TeamMemberMoveEdge{teamMemberMove1, teamMemberMove2}, teamMemberMoves.Edges) {
		t.Errorf("the teamMemberMoves added do not match")
	}
}
