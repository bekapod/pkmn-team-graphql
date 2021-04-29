package model

import (
	"testing"

	"github.com/go-test/deep"
)

func TestItem_Scan(t *testing.T) {
	exp := Item{
		ID:          "b170853d-96d7-4682-b13f-9e8f9631af16",
		Name:        "Thunder Stone",
		Slug:        "thunder-stone",
		Cost:        3000,
		FlingPower:  0,
		FlingEffect: "",
		Effect:      "Evolves an Eelektrik into Eelektross, an Eevee into Jolteon, or a Pikachu into Raichu.",
		Sprite:      "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/items/thunder-stone.png",
		Category:    EvolutionItem,
		Attributes:  []ItemAttribute{Underground},
	}
	got := Item{}
	err := (&got).Scan([]uint8{123, 34, 105, 100, 34, 58, 32, 34, 98, 49, 55, 48, 56, 53, 51, 100, 45, 57, 54, 100, 55, 45, 52, 54, 56, 50, 45, 98, 49, 51, 102, 45, 57, 101, 56, 102, 57, 54, 51, 49, 97, 102, 49, 54, 34, 44, 32, 34, 99, 111, 115, 116, 34, 58, 32, 51, 48, 48, 48, 44, 32, 34, 110, 97, 109, 101, 34, 58, 32, 34, 84, 104, 117, 110, 100, 101, 114, 32, 83, 116, 111, 110, 101, 34, 44, 32, 34, 115, 108, 117, 103, 34, 58, 32, 34, 116, 104, 117, 110, 100, 101, 114, 45, 115, 116, 111, 110, 101, 34, 44, 32, 34, 101, 102, 102, 101, 99, 116, 34, 58, 32, 34, 69, 118, 111, 108, 118, 101, 115, 32, 97, 110, 32, 69, 101, 108, 101, 107, 116, 114, 105, 107, 32, 105, 110, 116, 111, 32, 69, 101, 108, 101, 107, 116, 114, 111, 115, 115, 44, 32, 97, 110, 32, 69, 101, 118, 101, 101, 32, 105, 110, 116, 111, 32, 74, 111, 108, 116, 101, 111, 110, 44, 32, 111, 114, 32, 97, 32, 80, 105, 107, 97, 99, 104, 117, 32, 105, 110, 116, 111, 32, 82, 97, 105, 99, 104, 117, 46, 34, 44, 32, 34, 115, 112, 114, 105, 116, 101, 34, 58, 32, 34, 104, 116, 116, 112, 115, 58, 47, 47, 114, 97, 119, 46, 103, 105, 116, 104, 117, 98, 117, 115, 101, 114, 99, 111, 110, 116, 101, 110, 116, 46, 99, 111, 109, 47, 80, 111, 107, 101, 65, 80, 73, 47, 115, 112, 114, 105, 116, 101, 115, 47, 109, 97, 115, 116, 101, 114, 47, 115, 112, 114, 105, 116, 101, 115, 47, 105, 116, 101, 109, 115, 47, 116, 104, 117, 110, 100, 101, 114, 45, 115, 116, 111, 110, 101, 46, 112, 110, 103, 34, 44, 32, 34, 99, 97, 116, 101, 103, 111, 114, 121, 34, 58, 32, 34, 101, 118, 111, 108, 117, 116, 105, 111, 110, 34, 44, 32, 34, 97, 116, 116, 114, 105, 98, 117, 116, 101, 115, 34, 58, 32, 91, 34, 117, 110, 100, 101, 114, 103, 114, 111, 117, 110, 100, 34, 93, 44, 32, 34, 102, 108, 105, 110, 103, 80, 111, 119, 101, 114, 34, 58, 32, 51, 48, 44, 32, 34, 102, 108, 105, 110, 103, 69, 102, 102, 101, 99, 116, 34, 58, 32, 110, 117, 108, 108, 125})

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestItem_ScanString(t *testing.T) {
	exp := Item{
		ID:          "b170853d-96d7-4682-b13f-9e8f9631af16",
		Name:        "Thunder Stone",
		Slug:        "thunder-stone",
		Cost:        3000,
		FlingPower:  0,
		FlingEffect: "",
		Effect:      "Evolves an Eelektrik into Eelektross, an Eevee into Jolteon, or a Pikachu into Raichu.",
		Sprite:      "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/items/thunder-stone.png",
		Category:    EvolutionItem,
		Attributes:  []ItemAttribute{Underground},
	}
	got := Item{}
	err := (&got).Scan("{\"id\": \"b170853d-96d7-4682-b13f-9e8f9631af16\",\"cost\": 3000,\"name\": \"Thunder Stone\",\"slug\": \"thunder-stone\",\"effect\": \"Evolves an Eelektrik into Eelektross, an Eevee into Jolteon, or a Pikachu into Raichu.\",\"sprite\": \"https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/items/thunder-stone.png\",\"category\": \"evolution\",\"attributes\": [\"underground\"],\"flingPower\": 30,\"flingEffect\": null}")

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestItem_Scan_Error(t *testing.T) {
	got := Item{}
	err := (&got).Scan([]uint8{123, 34, 115, 108, 111, 116, 34, 58, 32, 49, 44, 32, 34, 116, 121, 112, 101, 34, 58, 32, 123, 34, 105, 100, 34, 58, 32, 34, 48, 55, 98, 57, 101, 98, 125})

	if err == nil {
		t.Error("expected an error but got nil")
	}
}

func TestItem_Scan_ItemError(t *testing.T) {
	got := Item{}
	err := (&got).Scan(5)

	if err == nil {
		t.Error("expected an error but got nil")
	}
}
