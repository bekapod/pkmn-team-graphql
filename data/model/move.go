package model

type Move struct {
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	Slug         string      `json:"slug"`
	Accuracy     int         `json:"accuracy"`
	PP           int         `json:"pp"`
	Power        int         `json:"power"`
	DamageClass  DamageClass `json:"damageClass"`
	Effect       string      `json:"effect"`
	EffectChance int         `json:"effectChance"`
	Target       string      `json:"target"`
	Type         Type        `json:"type"`
}

type MoveList struct {
	Total int    `json:"total"`
	Moves []Move `json:"moves"`
}

func NewMoveList(m []Move) MoveList {
	return MoveList{
		Total: len(m),
		Moves: m,
	}
}

func NewEmptyMoveList() MoveList {
	return MoveList{
		Total: 0,
		Moves: []Move{},
	}
}

func (l *MoveList) AddMove(m Move) {
	l.Total++
	l.Moves = append(l.Moves, m)
}

func (Move) IsEntity() {}
