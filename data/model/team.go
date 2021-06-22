package model

import "bekapod/pkmn-team-graphql/data/db"

func NewTeamFromDb(dbTeam db.TeamModel) Team {
	teamMembers := NewEmptyTeamMemberConnection()
	team := Team{
		ID:      dbTeam.ID,
		Name:    dbTeam.Name,
		Members: &teamMembers,
	}

	for _, tm := range dbTeam.TeamMembers() {
		teamMember := NewTeamMemberEdgeFromDb(tm)
		teamMembers.AddEdge(&teamMember)
	}

	return team
}

func NewTeamEdgeFromDb(dbTeam db.TeamModel) TeamEdge {
	node := NewTeamFromDb(dbTeam)
	return TeamEdge{
		Cursor: dbTeam.ID,
		Node:   &node,
	}
}

func NewEmptyTeamConnection() TeamConnection {
	return TeamConnection{
		PageInfo: &PageInfo{
			HasNextPage:     false,
			HasPreviousPage: false,
		},
		Edges: []*TeamEdge{},
	}
}

func (c *TeamConnection) AddEdge(e *TeamEdge) {
	if c.PageInfo.StartCursor == nil {
		c.PageInfo.StartCursor = &e.Cursor
	}
	c.PageInfo.EndCursor = &e.Cursor
	c.Edges = append(c.Edges, e)
}
