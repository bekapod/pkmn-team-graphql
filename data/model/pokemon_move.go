package model

import "bekapod/pkmn-team-graphql/data/db"

type PokemonMove struct {
	MoveID         string          `json:"typeId"`
	PokemonID      string          `json:"pokemonId"`
	LearnMethod    MoveLearnMethod `json:"learnMethod"`
	LevelLearnedAt int             `json:"levelLearnedAt"`
}

func NewPokemonMoveFromDb(dbPokemonMove db.PokemonMoveModel) PokemonMove {
	pm := PokemonMove{
		MoveID:         dbPokemonMove.MoveID,
		PokemonID:      dbPokemonMove.PokemonID,
		LearnMethod:    MoveLearnMethod(dbPokemonMove.LearnMethod),
		LevelLearnedAt: dbPokemonMove.LevelLearnedAt,
	}

	return pm
}

func NewPokemonMoveList(pokemonMoves []*PokemonMove) PokemonMoveList {
	return PokemonMoveList{
		Total:        len(pokemonMoves),
		PokemonMoves: pokemonMoves,
	}
}

func NewEmptyPokemonMoveList() PokemonMoveList {
	return PokemonMoveList{
		Total:        0,
		PokemonMoves: []*PokemonMove{},
	}
}

func (l *PokemonMoveList) AddPokemonMove(pm *PokemonMove) {
	l.Total++
	l.PokemonMoves = append(l.PokemonMoves, pm)
}
