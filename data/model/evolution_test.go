package model

import (
	"testing"

	"github.com/go-test/deep"
)

func TestNewEvolutionList(t *testing.T) {
	evolutions := []Evolution{
		{
			Trigger: LevelUp,
		},
		{
			Trigger: Shed,
		},
	}

	exp := EvolutionList{
		Total:      2,
		Evolutions: evolutions,
	}

	got := NewEvolutionList(evolutions)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestNewEmptyEvolutionList(t *testing.T) {
	exp := EvolutionList{
		Total:      0,
		Evolutions: []Evolution{},
	}

	got := NewEmptyEvolutionList()
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestEvolutionList_AddEvolution(t *testing.T) {
	evolutions := EvolutionList{}
	evolution1 := Evolution{}
	evolution2 := Evolution{}
	evolutions.AddEvolution(evolution1)
	evolutions.AddEvolution(evolution2)

	if evolutions.Total != 2 {
		t.Errorf("expected Total of 2, but got %d instead", evolutions.Total)
	}
}
