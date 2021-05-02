package model

import "bekapod/pkmn-team-graphql/data/db"

func NewEggGroupFromDb(dbEggGroup db.EggGroupModel) EggGroup {
	e := EggGroup{
		ID:   dbEggGroup.ID,
		Slug: dbEggGroup.Slug,
		Name: dbEggGroup.Name,
	}

	return e
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

func (l *EggGroupList) AddEggGroup(eg *EggGroup) {
	l.Total++
	l.EggGroups = append(l.EggGroups, eg)
}
