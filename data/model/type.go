package model

import (
	"encoding/json"
	"fmt"
)

type Type struct {
	ID               string   `json:"id"`
	Slug             string   `json:"slug"`
	Name             string   `json:"name"`
	NoDamageTo       TypeList `json:"noDamageTo"`
	HalfDamageTo     TypeList `json:"halfDamageTo"`
	DoubleDamageTo   TypeList `json:"doubleDamageTo"`
	NoDamageFrom     TypeList `json:"noDamageFrom"`
	HalfDamageFrom   TypeList `json:"halfDamageFrom"`
	DoubleDamageFrom TypeList `json:"doubleDamageFrom"`
}

type TypeList struct {
	Total int    `json:"total"`
	Types []Type `json:"types"`
}

func NewTypeList(types []Type) TypeList {
	return TypeList{
		Total: len(types),
		Types: types,
	}
}

func NewEmptyTypeList() TypeList {
	return TypeList{
		Total: 0,
		Types: []Type{},
	}
}

func (l *TypeList) AddType(t *Type) {
	l.Total++
	l.Types = append(l.Types, *t)
}

func (t *Type) Scan(src interface{}) error {
	switch v := src.(type) {
	case []uint8:
		err := json.Unmarshal([]byte(v), &t)
		return err
	}

	return fmt.Errorf("failed to scan type")
}

func (Type) IsEntity() {}
