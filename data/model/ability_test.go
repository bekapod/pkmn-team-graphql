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

func TestNewAbilityList(t *testing.T) {
	abilities := []*Ability{
		{
			ID: "123-456",
		},
		{
			ID: "456-789",
		},
	}

	exp := AbilityList{
		Total:     2,
		Abilities: abilities,
	}

	got := NewAbilityList(abilities)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestNewEmptyAbilityList(t *testing.T) {
	exp := AbilityList{
		Total:     0,
		Abilities: []*Ability{},
	}

	got := NewEmptyAbilityList()
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestAbilityList_AddAbility(t *testing.T) {
	abilities := AbilityList{}
	ability1 := &Ability{}
	ability2 := &Ability{}
	abilities.AddAbility(ability1)
	abilities.AddAbility(ability2)

	if abilities.Total != 2 {
		t.Errorf("expected Total of 2, but got %d instead", abilities.Total)
	}

	if !reflect.DeepEqual([]*Ability{ability1, ability2}, abilities.Abilities) {
		t.Errorf("the abilities added do not match")
	}
}
