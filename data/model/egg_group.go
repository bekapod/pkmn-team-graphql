package model

type EggGroup struct {
	ID   string `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
}

type EggGroupList struct {
	Total     int         `json:"total"`
	EggGroups []*EggGroup `json:"egg_groups"`
}

func NewEggGroupList(eggGroups []*EggGroup) EggGroupList {
	return EggGroupList{
		Total:     len(eggGroups),
		EggGroups: eggGroups,
	}
}

func NewEmptyEggGroupList() EggGroupList {
	return EggGroupList{
		Total:     0,
		EggGroups: []*EggGroup{},
	}
}

func (l *EggGroupList) AddEggGroup(e *EggGroup) {
	l.Total++
	l.EggGroups = append(l.EggGroups, e)
}

func (EggGroup) IsEntity() {}
