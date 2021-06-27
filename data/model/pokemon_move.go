package model

import "bekapod/pkmn-team-graphql/data/db"

type PokemonMoveEdge struct {
	Cursor         string          `json:"cursor"`
	NodeID         string          `json:"nodeId"`
	LearnMethod    MoveLearnMethod `json:"learnMethod"`
	LevelLearnedAt int             `json:"levelLearnedAt"`
}

type PokemonWithMoveEdge struct {
	Cursor         string          `json:"cursor"`
	NodeID         string          `json:"nodeId"`
	LearnMethod    MoveLearnMethod `json:"learnMethod"`
	LevelLearnedAt int             `json:"levelLearnedAt"`
}

func NewPokemonMoveEdgeFromDb(dbPokemonMove db.PokemonMoveModel) PokemonMoveEdge {
	pm := PokemonMoveEdge{
		Cursor:         dbPokemonMove.ID,
		NodeID:         dbPokemonMove.MoveID,
		LearnMethod:    MoveLearnMethod(dbPokemonMove.LearnMethod),
		LevelLearnedAt: dbPokemonMove.LevelLearnedAt,
	}

	return pm
}

func NewEmptyPokemonMoveConnection() PokemonMoveConnection {
	return PokemonMoveConnection{
		PageInfo: &PageInfo{
			HasNextPage:     false,
			HasPreviousPage: false,
		},
		Edges: []*PokemonMoveEdge{},
	}
}

func (c *PokemonMoveConnection) AddEdge(e *PokemonMoveEdge) {
	if c.PageInfo.StartCursor == nil {
		c.PageInfo.StartCursor = &e.Cursor
	}
	c.PageInfo.EndCursor = &e.Cursor
	c.Edges = append(c.Edges, e)
}

func NewPokemonWithMoveEdgeFromDb(dbPokemonMove db.PokemonMoveModel) PokemonWithMoveEdge {
	pm := PokemonWithMoveEdge{
		Cursor:         dbPokemonMove.ID,
		NodeID:         dbPokemonMove.PokemonID,
		LearnMethod:    MoveLearnMethod(dbPokemonMove.LearnMethod),
		LevelLearnedAt: dbPokemonMove.LevelLearnedAt,
	}

	return pm
}

func NewEmptyPokemonWithMoveConnection() PokemonWithMoveConnection {
	return PokemonWithMoveConnection{
		PageInfo: &PageInfo{
			HasNextPage:     false,
			HasPreviousPage: false,
		},
		Edges: []*PokemonWithMoveEdge{},
	}
}

func (c *PokemonWithMoveConnection) AddEdge(e *PokemonWithMoveEdge) {
	if c.PageInfo.StartCursor == nil {
		c.PageInfo.StartCursor = &e.Cursor
	}
	c.PageInfo.EndCursor = &e.Cursor
	c.Edges = append(c.Edges, e)
}
