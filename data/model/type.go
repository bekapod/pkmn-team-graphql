package model

type Type struct {
	ID   string `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
}

type TypeList struct {
	Total int     `json:"total"`
	Types []*Type `json:"types"`
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

func (Type) IsEntity() {}
