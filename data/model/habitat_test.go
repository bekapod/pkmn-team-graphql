package model

import (
	"os"
	"testing"

	"github.com/go-test/deep"
)

type TestDataHabitat struct {
	input Habitat
	exp   interface{}
}

type TestDataHabitatUnmarshal struct {
	input    Habitat
	hasError bool
}

func TestHabitat_IsValid(t *testing.T) {
	habitats := []TestDataHabitat{
		{Cave, true},
		{WatersEdge, true},
		{Rare, true},
		{"Something else", false},
	}

	for _, item := range habitats {
		if got := item.input.IsValid(); item.exp != got {
			t.Errorf("expected '%t' but got '%t' instead", item.exp, got)
		}
	}
}

func TestHabitat_String(t *testing.T) {
	habitats := []TestDataHabitat{
		{Cave, "cave"},
		{WatersEdge, "waters-edge"},
		{Rare, "rare"},
	}

	for _, item := range habitats {
		if got := item.input.String(); item.exp != got {
			t.Errorf("expected '%s' but got '%s' instead", item.exp, got)
		}
	}
}

func TestHabitat_UnmarshalGQL(t *testing.T) {
	habitats := []TestDataHabitatUnmarshal{
		{Cave, false},
		{WatersEdge, false},
		{Rare, false},
		{"Something else", true},
	}

	for _, item := range habitats {
		if got := (&item.input).UnmarshalGQL(item.input.String()); item.hasError && got == nil {
			t.Error("expected an error but didn't get one")
		}
	}
}

func TestHabitat_UnmarshalGQL_Error(t *testing.T) {
	habitat := Cave
	got := (&habitat).UnmarshalGQL(5)
	if got == nil {
		t.Error("expected an error but didn't get one")
	}
}

func TestHabitat_Scan(t *testing.T) {
	exp := Grassland
	got := Grassland
	err := (&got).Scan("grassland")

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestHabitat_Scan_Nil(t *testing.T) {
	var got Habitat
	err := (&got).Scan(nil)

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestHabitat_Scan_Invalid(t *testing.T) {
	got := Grassland
	err := (&got).Scan("Sky")

	if err == nil {
		t.Error("expected an error but got nil")
	}
}

func TestHabitat_Scan_TypeError(t *testing.T) {
	got := Grassland
	err := (&got).Scan(5)

	if err == nil {
		t.Error("expected an error but got nil")
	}
}

func ExampleHabitat_MarshalGQL() {
	habitat := Cave
	(&habitat).MarshalGQL(os.Stdout)
	// Output: "cave"
}
