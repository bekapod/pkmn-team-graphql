package model

import (
	"os"
	"testing"

	"github.com/go-test/deep"
)

type TestDataGender struct {
	input Gender
	exp   interface{}
}

type TestDataGenderUnmarshal struct {
	input    Gender
	hasError bool
}

func TestGender_IsValid(t *testing.T) {
	genders := []TestDataGender{
		{Male, true},
		{Female, true},
		{Unknown, true},
		{"Something else", false},
	}

	for _, item := range genders {
		if got := item.input.IsValid(); item.exp != got {
			t.Errorf("expected '%t' but got '%t' instead", item.exp, got)
		}
	}
}

func TestGender_String(t *testing.T) {
	genders := []TestDataGender{
		{Female, "female"},
		{Male, "male"},
		{Unknown, "unknown"},
	}

	for _, item := range genders {
		if got := item.input.String(); item.exp != got {
			t.Errorf("expected '%s' but got '%s' instead", item.exp, got)
		}
	}
}

func TestGender_UnmarshalGQL(t *testing.T) {
	genders := []TestDataGenderUnmarshal{
		{Female, false},
		{Male, false},
		{Unknown, false},
		{"Something else", true},
	}

	for _, item := range genders {
		if got := (&item.input).UnmarshalGQL(item.input.String()); item.hasError && got == nil {
			t.Error("expected an error but didn't get one")
		}
	}
}

func TestGender_UnmarshalGQL_Error(t *testing.T) {
	gender := Female
	got := (&gender).UnmarshalGQL(5)
	if got == nil {
		t.Error("expected an error but didn't get one")
	}
}

func TestGender_Scan(t *testing.T) {
	exp := Male
	got := Male
	err := (&got).Scan("male")

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestGender_Scan_Nil(t *testing.T) {
	var got Gender
	err := (&got).Scan(nil)

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestGender_Scan_Invalid(t *testing.T) {
	got := Female
	err := (&got).Scan("Sky")

	if err == nil {
		t.Error("expected an error but got nil")
	}
}

func TestGender_Scan_TypeError(t *testing.T) {
	got := Female
	err := (&got).Scan(5)

	if err == nil {
		t.Error("expected an error but got nil")
	}
}

func ExampleGender_MarshalGQL() {
	gender := Female
	(&gender).MarshalGQL(os.Stdout)
	// Output: "female"
}
