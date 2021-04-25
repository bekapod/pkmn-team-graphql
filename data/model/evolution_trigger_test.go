package model

import (
	"os"
	"testing"

	"github.com/go-test/deep"
)

type TestDataEvolutionTrigger struct {
	input EvolutionTrigger
	exp   interface{}
}

type TestDataEvolutionTriggerUnmarshal struct {
	input    EvolutionTrigger
	hasError bool
}

func TestEvolutionTrigger_IsValid(t *testing.T) {
	evolutionTriggers := []TestDataEvolutionTrigger{
		{LevelUp, true},
		{Shed, true},
		{Trade, true},
		{"Something else", false},
	}

	for _, item := range evolutionTriggers {
		if got := item.input.IsValid(); item.exp != got {
			t.Errorf("expected '%t' but got '%t' instead", item.exp, got)
		}
	}
}

func TestEvolutionTrigger_String(t *testing.T) {
	evolutionTriggers := []TestDataEvolutionTrigger{
		{LevelUp, "level-up"},
		{Shed, "shed"},
		{Trade, "trade"},
	}

	for _, item := range evolutionTriggers {
		if got := item.input.String(); item.exp != got {
			t.Errorf("expected '%s' but got '%s' instead", item.exp, got)
		}
	}
}

func TestEvolutionTrigger_UnmarshalGQL(t *testing.T) {
	evolutionTriggers := []TestDataEvolutionTriggerUnmarshal{
		{LevelUp, false},
		{Shed, false},
		{Trade, false},
		{"Something else", true},
	}

	for _, item := range evolutionTriggers {
		if got := (&item.input).UnmarshalGQL(item.input.String()); item.hasError && got == nil {
			t.Error("expected an error but didn't get one")
		}
	}
}

func TestEvolutionTrigger_UnmarshalGQL_Error(t *testing.T) {
	evolutionTrigger := LevelUp
	got := (&evolutionTrigger).UnmarshalGQL(5)
	if got == nil {
		t.Error("expected an error but didn't get one")
	}
}

func TestEvolutionTrigger_Scan(t *testing.T) {
	exp := LevelUp
	got := LevelUp
	err := (&got).Scan("level-up")

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestEvolutionTrigger_Scan_Nil(t *testing.T) {
	var got EvolutionTrigger
	err := (&got).Scan(nil)

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestEvolutionTrigger_Scan_Invalid(t *testing.T) {
	got := LevelUp
	err := (&got).Scan("Sky")

	if err == nil {
		t.Error("expected an error but got nil")
	}
}

func TestEvolutionTrigger_Scan_TypeError(t *testing.T) {
	got := LevelUp
	err := (&got).Scan(5)

	if err == nil {
		t.Error("expected an error but got nil")
	}
}

func ExampleEvolutionTrigger_MarshalGQL() {
	evolutionTrigger := LevelUp
	(&evolutionTrigger).MarshalGQL(os.Stdout)
	// Output: "level-up"
}
