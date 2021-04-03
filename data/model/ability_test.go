package model

import (
	"testing"

	"github.com/go-test/deep"
)

func TestNewAbilityList(t *testing.T) {
	abilities := []Ability{
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
		Abilities: []Ability{},
	}

	got := NewEmptyAbilityList()
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestAbilityList_AddAbility(t *testing.T) {
	abilities := AbilityList{}
	ability1 := Ability{}
	ability2 := Ability{}
	abilities.AddAbility(ability1)
	abilities.AddAbility(ability2)

	if abilities.Total != 2 {
		t.Errorf("expected Total of 2, but got %d instead", abilities.Total)
	}
}
