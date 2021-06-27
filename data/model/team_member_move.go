package model

import "bekapod/pkmn-team-graphql/data/db"

type TeamMemberMoveEdge struct {
	Cursor         string          `json:"cursor"`
	NodeID         string          `json:"nodeId"`
	LearnMethod    MoveLearnMethod `json:"learnMethod"`
	LevelLearnedAt int             `json:"levelLearnedAt"`
	Slot           int             `json:"slot"`
}

func NewTeamMemberMoveEdgeFromDb(dbTeamMemberMove db.TeamMemberMoveModel) TeamMemberMoveEdge {
	return TeamMemberMoveEdge{
		Cursor:         dbTeamMemberMove.ID,
		NodeID:         dbTeamMemberMove.PokemonMove().MoveID,
		LearnMethod:    MoveLearnMethod(dbTeamMemberMove.PokemonMove().LearnMethod),
		LevelLearnedAt: dbTeamMemberMove.PokemonMove().LevelLearnedAt,
		Slot:           dbTeamMemberMove.Slot,
	}
}

func NewEmptyTeamMemberMoveConnection() TeamMemberMoveConnection {
	return TeamMemberMoveConnection{
		PageInfo: &PageInfo{
			HasNextPage:     false,
			HasPreviousPage: false,
		},
		Edges: []*TeamMemberMoveEdge{},
	}
}

func (c *TeamMemberMoveConnection) AddEdge(e *TeamMemberMoveEdge) {
	if c.PageInfo.StartCursor == nil {
		c.PageInfo.StartCursor = &e.Cursor
	}
	c.PageInfo.EndCursor = &e.Cursor
	c.Edges = append(c.Edges, e)
}
