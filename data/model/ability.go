package model

import "bekapod/pkmn-team-graphql/data/db"

type Ability struct {
	ID     string  `json:"id"`
	Slug   string  `json:"slug"`
	Name   string  `json:"name"`
	Effect *string `json:"effect"`
}

func (Ability) IsNode() {}

func NewAbilityFromDb(dbAbility db.AbilityModel) Ability {
	a := Ability{
		ID:   dbAbility.ID,
		Slug: dbAbility.Slug,
		Name: dbAbility.Name,
	}

	if value, ok := dbAbility.Effect(); ok {
		a.Effect = &value
	} else {
		a.Effect = nil
	}

	return a
}

func NewAbilityEdgeFromDb(dbAbility db.AbilityModel) AbilityEdge {
	node := NewAbilityFromDb(dbAbility)
	return AbilityEdge{
		Cursor: dbAbility.ID,
		Node:   &node,
	}
}

func NewEmptyAbilityConnection() AbilityConnection {
	return AbilityConnection{
		PageInfo: &PageInfo{
			HasNextPage:     false,
			HasPreviousPage: false,
		},
		Edges: []*AbilityEdge{},
	}
}

func (c *AbilityConnection) AddEdge(e *AbilityEdge) {
	if c.PageInfo.StartCursor == nil {
		c.PageInfo.StartCursor = &e.Cursor
	}
	c.PageInfo.EndCursor = &e.Cursor
	c.Edges = append(c.Edges, e)
}
