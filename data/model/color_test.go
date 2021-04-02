package model

import (
	"os"
	"testing"
)

type TestDataColor struct {
	input Color
	exp   interface{}
}

type TestDataColorUnmarshal struct {
	input    Color
	hasError bool
}

func TestColor_IsValid(t *testing.T) {
	colors := []TestDataColor{
		{Red, true},
		{Yellow, true},
		{Blue, true},
		{"Something else", false},
	}

	for _, item := range colors {
		if got := item.input.IsValid(); item.exp != got {
			t.Errorf("expected '%t' but got '%t' instead", item.exp, got)
		}
	}
}

func TestColor_String(t *testing.T) {
	colors := []TestDataColor{
		{Red, "red"},
		{Yellow, "yellow"},
		{Blue, "blue"},
	}

	for _, item := range colors {
		if got := item.input.String(); item.exp != got {
			t.Errorf("expected '%s' but got '%s' instead", item.exp, got)
		}
	}
}

func TestColor_UnmarshalGQL(t *testing.T) {
	colors := []TestDataColorUnmarshal{
		{Red, false},
		{Yellow, false},
		{Blue, false},
		{"Something else", true},
	}

	for _, item := range colors {
		if got := (&item.input).UnmarshalGQL(item.input.String()); item.hasError && got == nil {
			t.Error("expected an error but didn't get one")
		}
	}
}

func TestColor_UnmarshalGQL_Error(t *testing.T) {
	color := Yellow
	got := (&color).UnmarshalGQL(5)
	if got == nil {
		t.Error("expected an error but didn't get one")
	}
}

func ExampleColor_MarshalGQL() {
	color := Yellow
	(&color).MarshalGQL(os.Stdout)
	// Output: "yellow"
}
