package model

import (
	"bekapod/pkmn-team-graphql/data/db"
)

type Type struct {
	ID   string `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
}

func (Type) IsNode() {}

func NewTypeFromDb(dbType db.TypeModel) Type {
	return Type{
		ID:   dbType.ID,
		Slug: dbType.Slug,
		Name: dbType.Name,
	}
}

func NewTypeEdgeFromDb(dbType db.TypeModel) TypeEdge {
	node := NewTypeFromDb(dbType)
	return TypeEdge{
		Cursor: dbType.ID,
		Node:   &node,
	}
}

func NewEmptyTypeConnection() TypeConnection {
	return TypeConnection{
		PageInfo: &PageInfo{
			HasNextPage:     false,
			HasPreviousPage: false,
		},
		Edges: []*TypeEdge{},
	}
}

func (c *TypeConnection) AddEdge(e *TypeEdge) {
	if c.PageInfo.StartCursor == nil {
		c.PageInfo.StartCursor = &e.Cursor
	}
	c.PageInfo.EndCursor = &e.Cursor
	c.Edges = append(c.Edges, e)
}
