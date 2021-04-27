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

func TestEvolution_Scan(t *testing.T) {
	exp := Evolution{
		Pokemon:               Pokemon{},
		Trigger:               LevelUp,
		Item:                  Item{},
		Gender:                Female,
		HeldItem:              Item{},
		KnownMove:             Move{},
		Location:              Location{},
		MinLevel:              32,
		MinHappiness:          5,
		MinBeauty:             4,
		MinAffection:          10,
		NeedsOverworldRain:    true,
		PartyPokemon:          Pokemon{},
		RelativePhysicalStats: 1,
		TimeOfDay:             Night,
		TradeWithPokemon:      Pokemon{},
		TurnUpsideDown:        true,
		Spin:                  false,
		TakeDamage:            20,
		CriticalHits:          3,
	}
	got := Evolution{}
	err := (&got).Scan([]uint8{123, 34, 105, 100, 34, 58, 32, 34, 49, 102, 48, 57, 53, 56, 97, 48, 45, 52, 56, 99, 97, 45, 52, 49, 54, 48, 45, 57, 102, 49, 56, 45, 55, 101, 53, 102, 48, 54, 100, 57, 54, 100, 50, 55, 34, 44, 32, 34, 110, 97, 109, 101, 34, 58, 32, 34, 70, 97, 105, 114, 121, 34, 44, 32, 34, 115, 108, 117, 103, 34, 58, 32, 34, 102, 97, 105, 114, 121, 34, 125})

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestEvolution_Scan_Error(t *testing.T) {
	got := Evolution{}
	err := (&got).Scan([]uint8{123, 34, 115, 108, 111, 116, 34, 58, 32, 49, 44, 32, 34, 116, 121, 112, 101, 34, 58, 32, 123, 34, 105, 100, 34, 58, 32, 34, 48, 55, 98, 57, 101, 98, 125})

	if err == nil {
		t.Error("expected an error but got nil")
	}
}

func TestEvolution_Scan_EvolutionError(t *testing.T) {
	got := Evolution{}
	err := (&got).Scan(5)

	if err == nil {
		t.Error("expected an error but got nil")
	}
}
