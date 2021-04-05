package model

import (
	"encoding/json"
	"fmt"
)

type PokemonAbility struct {
	Ability  Ability `json:"ability"`
	Slot     int     `json:"slot"`
	IsHidden bool    `json:"isHidden"`
}

type PokemonAbilityList struct {
	Total            int              `json:"total"`
	PokemonAbilities []PokemonAbility `json:"pokemonAbilities"`
}

func NewPokemonAbilityList(a []PokemonAbility) PokemonAbilityList {
	return PokemonAbilityList{
		Total:            len(a),
		PokemonAbilities: a,
	}
}

func NewEmptyPokemonAbilityList() PokemonAbilityList {
	return PokemonAbilityList{
		Total:            0,
		PokemonAbilities: []PokemonAbility{},
	}
}

func (l *PokemonAbilityList) AddPokemonAbility(a *PokemonAbility) {
	l.Total++
	l.PokemonAbilities = append(l.PokemonAbilities, *a)
}

func (a *PokemonAbility) Scan(src interface{}) error {
	switch v := src.(type) {
	case []uint8:
		err := json.Unmarshal([]byte(v), &a)
		return err
	}

	return fmt.Errorf("failed to scan pokemon ability")
}

func (PokemonAbility) IsEntity() {}
