package model

import (
	"bekapod/pkmn-team-graphql/data/db"
	"reflect"
	"testing"

	"github.com/go-test/deep"
)

func TestNewEggGroupEdgeFromDb(t *testing.T) {
	eggGroup := db.EggGroupModel{
		InnerEggGroup: db.InnerEggGroup{
			ID:   "123",
			Slug: "some-egg-group",
			Name: "Some Egg Group",
		},
	}
	exp := EggGroupEdge{
		Cursor: eggGroup.ID,
		Node: &EggGroup{
			ID:   eggGroup.ID,
			Slug: eggGroup.Slug,
			Name: eggGroup.Name,
		},
	}

	got := NewEggGroupEdgeFromDb(eggGroup)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestNewEmptyEggGroupConnection(t *testing.T) {
	exp := EggGroupConnection{
		PageInfo: &PageInfo{
			HasNextPage:     false,
			HasPreviousPage: false,
		},
		Edges: []*EggGroupEdge{},
	}

	got := NewEmptyEggGroupConnection()
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestEggGroupConnection_AddEdge(t *testing.T) {
	eggGroups := NewEmptyEggGroupConnection()
	eggGroup1 := EggGroupEdge{}
	eggGroup2 := EggGroupEdge{}
	eggGroups.AddEdge(&eggGroup1)
	eggGroups.AddEdge(&eggGroup2)

	if !reflect.DeepEqual([]*EggGroupEdge{&eggGroup1, &eggGroup2}, eggGroups.Edges) {
		t.Errorf("the egg groups added do not match")
	}
}
