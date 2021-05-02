package model

import "bekapod/pkmn-team-graphql/data/db"

type Ability struct {
	ID     string  `json:"id"`
	Slug   string  `json:"slug"`
	Name   string  `json:"name"`
	Effect *string `json:"effect"`
}

func NewAbilityFromDb(dbAbility db.AbilityModel) Ability {
	a := Ability{
		ID:   dbAbility.ID,
		Slug: dbAbility.Slug,
		Name: dbAbility.Name,
	}

	if value, ok := dbAbility.Effect(); ok {
		a.Effect = &value
	} else {
		a.Effect = nil
	}

	return a
}

func NewAbilityList(abilities []*Ability) AbilityList {
	return AbilityList{
		Total:     len(abilities),
		Abilities: abilities,
	}
}

func NewEmptyAbilityList() AbilityList {
	return AbilityList{
		Total:     0,
		Abilities: []*Ability{},
	}
}

func (l *AbilityList) AddAbility(a *Ability) {
	l.Total++
	l.Abilities = append(l.Abilities, a)
}
