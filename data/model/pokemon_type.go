package model

import "bekapod/pkmn-team-graphql/data/db"

type PokemonType struct {
	TypeID    string `json:"typeId"`
	PokemonID string `json:"pokemonId"`
	Slot      int    `json:"slot"`
}

func NewPokemonTypeFromDb(dbPokemonType db.PokemonTypeModel) PokemonType {
	pt := PokemonType{
		TypeID:    dbPokemonType.TypeID,
		PokemonID: dbPokemonType.PokemonID,
		Slot:      dbPokemonType.Slot,
	}

	return pt
}

func NewPokemonTypeList(pokemonTypes []*PokemonType) PokemonTypeList {
	return PokemonTypeList{
		Total:        len(pokemonTypes),
		PokemonTypes: pokemonTypes,
	}
}

func NewEmptyPokemonTypeList() PokemonTypeList {
	return PokemonTypeList{
		Total:        0,
		PokemonTypes: []*PokemonType{},
	}
}

func (l *PokemonTypeList) AddPokemonType(pt *PokemonType) {
	l.Total++
	l.PokemonTypes = append(l.PokemonTypes, pt)
}
