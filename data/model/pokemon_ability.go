package model

import "bekapod/pkmn-team-graphql/data/db"

type PokemonAbilityEdge struct {
	Cursor   string `json:"cursor"`
	Slot     int    `json:"slot"`
	IsHidden bool   `json:"isHidden"`
	NodeID   string `json:"nodeId"`
}

type PokemonWithAbilityEdge struct {
	Cursor   string `json:"cursor"`
	Slot     int    `json:"slot"`
	IsHidden bool   `json:"isHidden"`
	NodeID   string `json:"nodeId"`
}

func NewPokemonAbilityEdgeFromDb(dbPokemonAbility db.PokemonAbilityModel) PokemonAbilityEdge {
	pa := PokemonAbilityEdge{
		Cursor:   dbPokemonAbility.ID,
		NodeID:   dbPokemonAbility.AbilityID,
		Slot:     dbPokemonAbility.Slot,
		IsHidden: dbPokemonAbility.IsHidden,
	}

	return pa
}

func NewEmptyPokemonAbilityConnection() PokemonAbilityConnection {
	return PokemonAbilityConnection{
		PageInfo: &PageInfo{
			HasNextPage:     false,
			HasPreviousPage: false,
		},
		Edges: []*PokemonAbilityEdge{},
	}
}

func (c *PokemonAbilityConnection) AddEdge(e *PokemonAbilityEdge) {
	if c.PageInfo.StartCursor == nil {
		c.PageInfo.StartCursor = &e.Cursor
	}
	c.PageInfo.EndCursor = &e.Cursor
	c.Edges = append(c.Edges, e)
}

func NewPokemonWithAbilityEdgeFromDb(dbPokemonAbility db.PokemonAbilityModel) PokemonWithAbilityEdge {
	pa := PokemonWithAbilityEdge{
		Cursor:   dbPokemonAbility.ID,
		NodeID:   dbPokemonAbility.PokemonID,
		Slot:     dbPokemonAbility.Slot,
		IsHidden: dbPokemonAbility.IsHidden,
	}

	return pa
}

func NewEmptyPokemonWithAbilityConnection() PokemonWithAbilityConnection {
	return PokemonWithAbilityConnection{
		PageInfo: &PageInfo{
			HasNextPage:     false,
			HasPreviousPage: false,
		},
		Edges: []*PokemonWithAbilityEdge{},
	}
}

func (c *PokemonWithAbilityConnection) AddEdge(e *PokemonWithAbilityEdge) {
	if c.PageInfo.StartCursor == nil {
		c.PageInfo.StartCursor = &e.Cursor
	}
	c.PageInfo.EndCursor = &e.Cursor
	c.Edges = append(c.Edges, e)
}
