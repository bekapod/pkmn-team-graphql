package model

import "bekapod/pkmn-team-graphql/data/db"

type Move struct {
	ID           string      `json:"id"`
	Slug         string      `json:"slug"`
	Name         string      `json:"name"`
	Accuracy     *int        `json:"accuracy"`
	Pp           *int        `json:"pp"`
	Power        *int        `json:"power"`
	DamageClass  DamageClass `json:"damageClass"`
	Effect       *string     `json:"effect"`
	EffectChance *int        `json:"effectChance"`
	Target       MoveTarget  `json:"target"`
	TypeID       string      `json:"typeId"`
}

func (Move) IsNode() {}

func NewMoveFromDb(dbMove db.MoveModel) Move {
	m := Move{
		ID:          dbMove.ID,
		Slug:        dbMove.Slug,
		Name:        dbMove.Name,
		DamageClass: DamageClass(dbMove.DamageClass),
		Target:      MoveTarget(dbMove.Target),
		TypeID:      dbMove.TypeID,
	}

	if accuracy, ok := dbMove.Accuracy(); ok {
		m.Accuracy = &accuracy
	} else {
		m.Accuracy = nil
	}

	if pp, ok := dbMove.Pp(); ok {
		m.Pp = &pp
	} else {
		m.Pp = nil
	}

	if power, ok := dbMove.Power(); ok {
		m.Power = &power
	} else {
		m.Power = nil
	}

	if effect, ok := dbMove.Effect(); ok {
		m.Effect = &effect
	} else {
		m.Effect = nil
	}

	if effectChance, ok := dbMove.EffectChance(); ok {
		m.EffectChance = &effectChance
	} else {
		m.EffectChance = nil
	}

	return m
}

func NewMoveEdgeFromDb(dbMove db.MoveModel) MoveEdge {
	node := NewMoveFromDb(dbMove)
	return MoveEdge{
		Cursor: dbMove.ID,
		Node:   &node,
	}
}

func NewEmptyMoveConnection() MoveConnection {
	return MoveConnection{
		PageInfo: &PageInfo{
			HasNextPage:     false,
			HasPreviousPage: false,
		},
		Edges: []*MoveEdge{},
	}
}

func (c *MoveConnection) AddEdge(e *MoveEdge) {
	if c.PageInfo.StartCursor == nil {
		c.PageInfo.StartCursor = &e.Cursor
	}
	c.PageInfo.EndCursor = &e.Cursor
	c.Edges = append(c.Edges, e)
}
