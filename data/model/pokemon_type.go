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

func (l *PokemonTypeList) Scan(src interface{}) error {
	ts := make([]string, 0)

	switch v := src.(type) {
	case string:
		err := json.Unmarshal([]byte(v), &ts)
		for _, val := range ts {
			var t struct {
				ID   string `json:"id"`
				Name string `json:"name"`
				Slug string `json:"slug"`
				Slot int    `json:"slot"`
			}
			err := json.Unmarshal([]byte(val), &t)
			if err != nil {
				return err
			}
			pokemonType := PokemonType{
				Slot: t.Slot,
				Type: Type{
					ID:   t.ID,
					Name: t.Name,
					Slug: t.Slug,
				},
			}
			l.AddPokemonType(&pokemonType)
		}
		if err != nil {
			return err
		}
		return nil
	}

	return fmt.Errorf("failed to scan pokemon type list")
}

func (PokemonType) IsEntity() {}
