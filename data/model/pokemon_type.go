package model

type PokemonType struct {
	TypeID    string `json:"type"`
	PokemonID string `json:"pokemon"`
	Slot      int    `json:"slot"`
}

type PokemonTypeList struct {
	Total        int            `json:"total"`
	PokemonTypes []*PokemonType `json:"pokemonTypes"`
}

func NewPokemonTypeList(t []*PokemonType) PokemonTypeList {
	return PokemonTypeList{
		Total:        len(t),
		PokemonTypes: t,
	}
}

func NewEmptyPokemonTypeList() PokemonTypeList {
	return PokemonTypeList{
		Total:        0,
		PokemonTypes: []*PokemonType{},
	}
}

func (l *PokemonTypeList) AddPokemonType(t *PokemonType) {
	l.Total++
	l.PokemonTypes = append(l.PokemonTypes, t)
}

func (PokemonType) IsEntity() {}
