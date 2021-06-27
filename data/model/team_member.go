package model

import "bekapod/pkmn-team-graphql/data/db"

type TeamMember struct {
	ID        string                    `json:"id"`
	PokemonID *string                   `json:"pokemonId"`
	Moves     *TeamMemberMoveConnection `json:"moves"`
}

func (TeamMember) IsNode() {}

func NewTeamMemberFromDb(dbTeamMember db.TeamMemberModel) TeamMember {
	teamMemberMoves := NewEmptyTeamMemberMoveConnection()

	for _, move := range dbTeamMember.Moves() {
		edge := NewTeamMemberMoveEdgeFromDb(move)
		teamMemberMoves.AddEdge(&edge)
	}

	return TeamMember{
		ID:        dbTeamMember.ID,
		PokemonID: &dbTeamMember.PokemonID,
		Moves:     &teamMemberMoves,
	}
}

func NewTeamMemberEdgeFromDb(dbTeamMember db.TeamMemberModel) TeamMemberEdge {
	node := NewTeamMemberFromDb(dbTeamMember)
	teamMember := TeamMemberEdge{
		Cursor: dbTeamMember.ID,
		Node:   &node,
		Slot:   &dbTeamMember.Slot,
	}

	return teamMember
}

func NewEmptyTeamMemberConnection() TeamMemberConnection {
	return TeamMemberConnection{
		PageInfo: &PageInfo{
			HasNextPage:     false,
			HasPreviousPage: false,
		},
		Edges: []*TeamMemberEdge{},
	}
}

func (c *TeamMemberConnection) AddEdge(e *TeamMemberEdge) {
	if c.PageInfo.StartCursor == nil {
		c.PageInfo.StartCursor = &e.Cursor
	}
	c.PageInfo.EndCursor = &e.Cursor
	c.Edges = append(c.Edges, e)
}
