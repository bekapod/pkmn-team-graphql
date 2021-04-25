package model

import (
	"os"
	"testing"

	"github.com/go-test/deep"
)

type TestDataItemAttribute struct {
	input ItemAttribute
	exp   interface{}
}

type TestDataItemAttributeUnmarshal struct {
	input    ItemAttribute
	hasError bool
}

func TestItemAttribute_IsValid(t *testing.T) {
	itemAttributes := []TestDataItemAttribute{
		{Countable, true},
		{UsableInBattle, true},
		{Consumable, true},
		{"Something else", false},
	}

	for _, item := range itemAttributes {
		if got := item.input.IsValid(); item.exp != got {
			t.Errorf("expected '%t' but got '%t' instead", item.exp, got)
		}
	}
}

func TestItemAttribute_String(t *testing.T) {
	itemAttributes := []TestDataItemAttribute{
		{Countable, "countable"},
		{UsableInBattle, "usable-in-battle"},
		{Consumable, "consumable"},
	}

	for _, item := range itemAttributes {
		if got := item.input.String(); item.exp != got {
			t.Errorf("expected '%s' but got '%s' instead", item.exp, got)
		}
	}
}

func TestItemAttribute_UnmarshalGQL(t *testing.T) {
	itemAttributes := []TestDataItemAttributeUnmarshal{
		{Countable, false},
		{UsableInBattle, false},
		{Consumable, false},
		{"Something else", true},
	}

	for _, item := range itemAttributes {
		if got := (&item.input).UnmarshalGQL(item.input.String()); item.hasError && got == nil {
			t.Error("expected an error but didn't get one")
		}
	}
}

func TestItemAttribute_UnmarshalGQL_Error(t *testing.T) {
	itemAttribute := Consumable
	got := (&itemAttribute).UnmarshalGQL(5)
	if got == nil {
		t.Error("expected an error but didn't get one")
	}
}

func TestItemAttribute_Scan(t *testing.T) {
	exp := UsableInBattle
	got := UsableInBattle
	err := (&got).Scan("usable-in-battle")

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestItemAttribute_Scan_Nil(t *testing.T) {
	var got ItemAttribute
	err := (&got).Scan(nil)

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestItemAttribute_Scan_Invalid(t *testing.T) {
	got := Countable
	err := (&got).Scan("Sky")

	if err == nil {
		t.Error("expected an error but got nil")
	}
}

func TestItemAttribute_Scan_TypeError(t *testing.T) {
	got := Consumable
	err := (&got).Scan(5)

	if err == nil {
		t.Error("expected an error but got nil")
	}
}

func ExampleItemAttribute_MarshalGQL() {
	itemAttribute := UsableInBattle
	(&itemAttribute).MarshalGQL(os.Stdout)
	// Output: "usable-in-battle"
}
