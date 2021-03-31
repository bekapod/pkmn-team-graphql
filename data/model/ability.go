package model

type Ability struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Slug   string `json:"slug"`
	Effect string `json:"effect"`
}

type AbilityList struct {
	Total     int        `json:"total"`
	Abilities []*Ability `json:"abilities"`
}

func NewAbilityList(a []*Ability) AbilityList {
	return AbilityList{
		Total:     len(a),
		Abilities: a,
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

func (Ability) IsEntity() {}
