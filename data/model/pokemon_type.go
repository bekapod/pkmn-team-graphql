package model

import "bekapod/pkmn-team-graphql/data/db"

type PokemonTypeEdge struct {
	Cursor string `json:"cursor"`
	NodeID string `json:"nodeId"`
	Slot   int    `json:"slot"`
}

type PokemonWithTypeEdge struct {
	Cursor string `json:"cursor"`
	NodeID string `json:"nodeId"`
	Slot   int    `json:"slot"`
}

func NewPokemonTypeFromDb(dbPokemonType db.PokemonTypeModel) PokemonTypeEdge {
	pt := PokemonTypeEdge{
		Cursor: dbPokemonType.ID,
		NodeID: dbPokemonType.TypeID,
		Slot:   dbPokemonType.Slot,
	}

	return pt
}

func NewEmptyPokemonTypeConnection() PokemonTypeConnection {
	return PokemonTypeConnection{
		PageInfo: &PageInfo{
			HasNextPage:     false,
			HasPreviousPage: false,
		},
		Edges: []*PokemonTypeEdge{},
	}
}

func (c *PokemonTypeConnection) AddEdge(e *PokemonTypeEdge) {
	if c.PageInfo.StartCursor == nil {
		c.PageInfo.StartCursor = &e.Cursor
	}
	c.PageInfo.EndCursor = &e.Cursor
	c.Edges = append(c.Edges, e)
}

func NewPokemonWithTypeFromDb(dbPokemonType db.PokemonTypeModel) PokemonWithTypeEdge {
	pt := PokemonWithTypeEdge{
		Cursor: dbPokemonType.ID,
		NodeID: dbPokemonType.PokemonID,
		Slot:   dbPokemonType.Slot,
	}

	return pt
}

func NewEmptyPokemonWithTypeConnection() PokemonWithTypeConnection {
	return PokemonWithTypeConnection{
		PageInfo: &PageInfo{
			HasNextPage:     false,
			HasPreviousPage: false,
		},
		Edges: []*PokemonWithTypeEdge{},
	}
}

func (c *PokemonWithTypeConnection) AddEdge(e *PokemonWithTypeEdge) {
	if c.PageInfo.StartCursor == nil {
		c.PageInfo.StartCursor = &e.Cursor
	}
	c.PageInfo.EndCursor = &e.Cursor
	c.Edges = append(c.Edges, e)
}
