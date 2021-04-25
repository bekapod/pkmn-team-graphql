package model

import (
	"os"
	"testing"

	"github.com/go-test/deep"
)

type TestDataTimeOfDay struct {
	input TimeOfDay
	exp   interface{}
}

type TestDataTimeOfDayUnmarshal struct {
	input    TimeOfDay
	hasError bool
}

func TestTimeOfDay_IsValid(t *testing.T) {
	timeOfDays := []TestDataTimeOfDay{
		{Any, true},
		{Day, true},
		{Night, true},
		{"Something else", false},
	}

	for _, item := range timeOfDays {
		if got := item.input.IsValid(); item.exp != got {
			t.Errorf("expected '%t' but got '%t' instead", item.exp, got)
		}
	}
}

func TestTimeOfDay_String(t *testing.T) {
	timeOfDays := []TestDataTimeOfDay{
		{Any, "any"},
		{Day, "day"},
		{Night, "night"},
	}

	for _, item := range timeOfDays {
		if got := item.input.String(); item.exp != got {
			t.Errorf("expected '%s' but got '%s' instead", item.exp, got)
		}
	}
}

func TestTimeOfDay_UnmarshalGQL(t *testing.T) {
	timeOfDays := []TestDataTimeOfDayUnmarshal{
		{Any, false},
		{Day, false},
		{Night, false},
		{"Something else", true},
	}

	for _, item := range timeOfDays {
		if got := (&item.input).UnmarshalGQL(item.input.String()); item.hasError && got == nil {
			t.Error("expected an error but didn't get one")
		}
	}
}

func TestTimeOfDay_UnmarshalGQL_Error(t *testing.T) {
	timeOfDay := Day
	got := (&timeOfDay).UnmarshalGQL(5)
	if got == nil {
		t.Error("expected an error but didn't get one")
	}
}

func TestTimeOfDay_Scan(t *testing.T) {
	exp := Night
	got := Night
	err := (&got).Scan("night")

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestTimeOfDay_Scan_Nil(t *testing.T) {
	var got TimeOfDay
	err := (&got).Scan(nil)

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestTimeOfDay_Scan_Invalid(t *testing.T) {
	got := Day
	err := (&got).Scan("Sky")

	if err == nil {
		t.Error("expected an error but got nil")
	}
}

func TestTimeOfDay_Scan_TypeError(t *testing.T) {
	got := Night
	err := (&got).Scan(5)

	if err == nil {
		t.Error("expected an error but got nil")
	}
}

func ExampleTimeOfDay_MarshalGQL() {
	timeOfDay := Day
	(&timeOfDay).MarshalGQL(os.Stdout)
	// Output: "day"
}
