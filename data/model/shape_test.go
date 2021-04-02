package model

import (
	"os"
	"testing"
)

type TestDataShape struct {
	input Shape
	exp   interface{}
}

type TestDataShapeUnmarshal struct {
	input    Shape
	hasError bool
}

func TestShape_IsValid(t *testing.T) {
	shapes := []TestDataShape{
		{Squiggle, true},
		{Ball, true},
		{Heads, true},
		{"Something else", false},
	}

	for _, item := range shapes {
		if got := item.input.IsValid(); item.exp != got {
			t.Errorf("expected '%t' but got '%t' instead", item.exp, got)
		}
	}
}

func TestShape_String(t *testing.T) {
	shapes := []TestDataShape{
		{Squiggle, "squiggle"},
		{Ball, "ball"},
		{Heads, "heads"},
	}

	for _, item := range shapes {
		if got := item.input.String(); item.exp != got {
			t.Errorf("expected '%s' but got '%s' instead", item.exp, got)
		}
	}
}

func TestShape_UnmarshalGQL(t *testing.T) {
	shapes := []TestDataShapeUnmarshal{
		{Squiggle, false},
		{Ball, false},
		{Heads, false},
		{"Something else", true},
	}

	for _, item := range shapes {
		if got := (&item.input).UnmarshalGQL(item.input.String()); item.hasError && got == nil {
			t.Error("expected an error but didn't get one")
		}
	}
}

func TestShape_UnmarshalGQL_Error(t *testing.T) {
	shape := Heads
	got := (&shape).UnmarshalGQL(5)
	if got == nil {
		t.Error("expected an error but didn't get one")
	}
}

func ExampleShape_MarshalGQL() {
	shape := Ball
	(&shape).MarshalGQL(os.Stdout)
	// Output: "ball"
}
