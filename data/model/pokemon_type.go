package model

import (
	"encoding/json"
	"fmt"
)

type PokemonType struct {
	Type Type `json:"type"`
	Slot int  `json:"slot"`
}

type PokemonTypeList struct {
	Total        int           `json:"total"`
	PokemonTypes []PokemonType `json:"pokemonTypes"`
}

func NewPokemonTypeList(t []PokemonType) PokemonTypeList {
	return PokemonTypeList{
		Total:        len(t),
		PokemonTypes: t,
	}
}

func NewEmptyPokemonTypeList() PokemonTypeList {
	return PokemonTypeList{
		Total:        0,
		PokemonTypes: []PokemonType{},
	}
}

func (l *PokemonTypeList) AddPokemonType(t *PokemonType) {
	l.Total++
	l.PokemonTypes = append(l.PokemonTypes, *t)
}

func (t *PokemonType) Scan(src interface{}) error {
	switch v := src.(type) {
	case []uint8:
		err := json.Unmarshal([]byte(v), &t)
		return err
	}

	return fmt.Errorf("failed to scan pokemon type")
}

func (PokemonType) IsEntity() {}
