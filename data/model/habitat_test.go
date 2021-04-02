package model

import (
	"os"
	"testing"
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

func ExampleHabitat_MarshalGQL() {
	habitat := Cave
	(&habitat).MarshalGQL(os.Stdout)
	// Output: "cave"
}
