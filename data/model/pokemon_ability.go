package model

import "bekapod/pkmn-team-graphql/data/db"

type PokemonAbility struct {
	AbilityID string `json:"typeId"`
	PokemonID string `json:"pokemonId"`
	Slot      int    `json:"slot"`
	IsHidden  bool   `json:"isHidden"`
}

func NewPokemonAbilityFromDb(dbPokemonAbility db.PokemonAbilityModel) PokemonAbility {
	pa := PokemonAbility{
		AbilityID: dbPokemonAbility.AbilityID,
		PokemonID: dbPokemonAbility.PokemonID,
		Slot:      dbPokemonAbility.Slot,
		IsHidden:  dbPokemonAbility.IsHidden,
	}

	return pa
}

func NewPokemonAbilityList(pokemonAbilities []*PokemonAbility) PokemonAbilityList {
	return PokemonAbilityList{
		Total:            len(pokemonAbilities),
		PokemonAbilities: pokemonAbilities,
	}
}

func NewEmptyPokemonAbilityList() PokemonAbilityList {
	return PokemonAbilityList{
		Total:            0,
		PokemonAbilities: []*PokemonAbility{},
	}
}

func (l *PokemonAbilityList) AddPokemonAbility(pt *PokemonAbility) {
	l.Total++
	l.PokemonAbilities = append(l.PokemonAbilities, pt)
}
