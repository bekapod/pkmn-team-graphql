package model

import "bekapod/pkmn-team-graphql/data/db"

func NewEggGroupEdgeFromDb(dbEggGroup db.EggGroupModel) EggGroupEdge {
	e := EggGroupEdge{
		Cursor: dbEggGroup.ID,
		Node: &EggGroup{
			ID:   dbEggGroup.ID,
			Slug: dbEggGroup.Slug,
			Name: dbEggGroup.Name,
		},
	}

	return e
}

func NewEmptyEggGroupConnection() EggGroupConnection {
	return EggGroupConnection{
		PageInfo: &PageInfo{
			HasNextPage:     false,
			HasPreviousPage: false,
		},
		Edges: []*EggGroupEdge{},
	}
}

func (c *EggGroupConnection) AddEdge(e *EggGroupEdge) {
	if c.PageInfo.StartCursor == nil {
		c.PageInfo.StartCursor = &e.Cursor
	}
	c.PageInfo.EndCursor = &e.Cursor
	c.Edges = append(c.Edges, e)
}
