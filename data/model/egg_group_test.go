package model

import (
	"bekapod/pkmn-team-graphql/data/db"
	"reflect"
	"testing"

	"github.com/go-test/deep"
)

func TestNewEggGroupFromDb(t *testing.T) {
	eggGroup := db.EggGroupModel{
		InnerEggGroup: db.InnerEggGroup{
			ID:   "123",
			Slug: "some-egg-group",
			Name: "Some Egg Group",
		},
	}
	exp := EggGroup{
		ID:   eggGroup.ID,
		Slug: eggGroup.Slug,
		Name: eggGroup.Name,
	}

	got := NewEggGroupFromDb(eggGroup)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestNewEggGroupList(t *testing.T) {
	eggGroups := []*EggGroup{
		{
			ID: "123-456",
		},
		{
			ID: "456-789",
		},
	}

	exp := EggGroupList{
		Total:     2,
		EggGroups: eggGroups,
	}

	got := NewEggGroupList(eggGroups)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestNewEmptyEggGroupList(t *testing.T) {
	exp := EggGroupList{
		Total:     0,
		EggGroups: []*EggGroup{},
	}

	got := NewEmptyEggGroupList()
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestEggGroupList_AddEggGroup(t *testing.T) {
	eggGroups := EggGroupList{}
	eggGroup1 := EggGroup{}
	eggGroup2 := EggGroup{}
	eggGroups.AddEggGroup(&eggGroup1)
	eggGroups.AddEggGroup(&eggGroup2)

	if eggGroups.Total != 2 {
		t.Errorf("expected Total of 2, but got %d instead", eggGroups.Total)
	}

	if !reflect.DeepEqual([]*EggGroup{&eggGroup1, &eggGroup2}, eggGroups.EggGroups) {
		t.Errorf("the egg groups added do not match")
	}
}
