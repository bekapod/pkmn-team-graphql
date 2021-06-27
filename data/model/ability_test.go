package model

import (
	"bekapod/pkmn-team-graphql/data/db"
	"reflect"
	"testing"

	"github.com/go-test/deep"
)

func TestNewAbilityFromDb_WithEffect(t *testing.T) {
	effect := "Some description of the ability"
	ability := db.AbilityModel{
		InnerAbility: db.InnerAbility{
			ID:     "123",
			Slug:   "some-ability",
			Name:   "Some Ability",
			Effect: &effect,
		},
	}
	exp := Ability{
		ID:     ability.ID,
		Slug:   ability.Slug,
		Name:   ability.Name,
		Effect: &effect,
	}

	got := NewAbilityFromDb(ability)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestNewAbilityFromDb_WithoutEffect(t *testing.T) {
	ability := db.AbilityModel{
		InnerAbility: db.InnerAbility{
			ID:   "123",
			Slug: "some-ability",
			Name: "Some Ability",
		},
	}
	exp := Ability{
		ID:     ability.ID,
		Slug:   ability.Slug,
		Name:   ability.Name,
		Effect: nil,
	}

	got := NewAbilityFromDb(ability)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestNewAbilityEdgeFromDb(t *testing.T) {
	ability := db.AbilityModel{
		InnerAbility: db.InnerAbility{
			ID:   "123",
			Slug: "some-ability",
			Name: "Some Ability",
		},
	}
	exp := AbilityEdge{
		Cursor: ability.ID,
		Node: &Ability{
			ID:     ability.ID,
			Slug:   ability.Slug,
			Name:   ability.Name,
			Effect: nil,
		},
	}

	got := NewAbilityEdgeFromDb(ability)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestNewEmptyAbilityConnection(t *testing.T) {
	exp := AbilityConnection{
		PageInfo: &PageInfo{
			HasNextPage:     false,
			HasPreviousPage: false,
		},
		Edges: []*AbilityEdge{},
	}

	got := NewEmptyAbilityConnection()
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestAbilityConnection_AddEdge(t *testing.T) {
	abilities := NewEmptyAbilityConnection()
	ability1 := &AbilityEdge{}
	ability2 := &AbilityEdge{}
	abilities.AddEdge(ability1)
	abilities.AddEdge(ability2)

	if !reflect.DeepEqual([]*AbilityEdge{ability1, ability2}, abilities.Edges) {
		t.Errorf("the edges added do not match")
	}
}
