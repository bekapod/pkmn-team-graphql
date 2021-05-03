package model

import (
	"bekapod/pkmn-team-graphql/data/db"
)

type Type struct {
	ID   string `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
}

func NewTypeFromDb(dbType db.TypeModel) Type {
	return Type{
		ID:   dbType.ID,
		Slug: dbType.Slug,
		Name: dbType.Name,
	}
}

func NewTypeList(types []*Type) TypeList {
	return TypeList{
		Total: len(types),
		Types: types,
	}
}

func NewEmptyTypeList() TypeList {
	return TypeList{
		Total: 0,
		Types: []*Type{},
	}
}

func (l *TypeList) AddType(t *Type) {
	l.Total++
	l.Types = append(l.Types, t)
}
