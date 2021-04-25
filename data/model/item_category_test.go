package model

import (
	"os"
	"testing"

	"github.com/go-test/deep"
)

type TestDataItemCategory struct {
	input ItemCategory
	exp   interface{}
}

type TestDataItemCategoryUnmarshal struct {
	input    ItemCategory
	hasError bool
}

func TestItemCategory_IsValid(t *testing.T) {
	itemCategorys := []TestDataItemCategory{
		{OtherCategory, true},
		{ApricornBalls, true},
		{HeldItems, true},
		{"Something else", false},
	}

	for _, item := range itemCategorys {
		if got := item.input.IsValid(); item.exp != got {
			t.Errorf("expected '%t' but got '%t' instead", item.exp, got)
		}
	}
}

func TestItemCategory_String(t *testing.T) {
	itemCategorys := []TestDataItemCategory{
		{OtherCategory, "other"},
		{ApricornBalls, "apricorn-balls"},
		{HeldItems, "held-items"},
	}

	for _, item := range itemCategorys {
		if got := item.input.String(); item.exp != got {
			t.Errorf("expected '%s' but got '%s' instead", item.exp, got)
		}
	}
}

func TestItemCategory_UnmarshalGQL(t *testing.T) {
	itemCategorys := []TestDataItemCategoryUnmarshal{
		{OtherCategory, false},
		{HeldItems, false},
		{BadHeldItems, false},
		{"Something else", true},
	}

	for _, item := range itemCategorys {
		if got := (&item.input).UnmarshalGQL(item.input.String()); item.hasError && got == nil {
			t.Error("expected an error but didn't get one")
		}
	}
}

func TestItemCategory_UnmarshalGQL_Error(t *testing.T) {
	itemCategory := OtherCategory
	got := (&itemCategory).UnmarshalGQL(5)
	if got == nil {
		t.Error("expected an error but didn't get one")
	}
}

func TestItemCategory_Scan(t *testing.T) {
	exp := OtherCategory
	got := OtherCategory
	err := (&got).Scan("other")

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestItemCategory_Scan_Nil(t *testing.T) {
	var got ItemCategory
	err := (&got).Scan(nil)

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestItemCategory_Scan_Invalid(t *testing.T) {
	got := OtherCategory
	err := (&got).Scan("Sky")

	if err == nil {
		t.Error("expected an error but got nil")
	}
}

func TestItemCategory_Scan_TypeError(t *testing.T) {
	got := OtherCategory
	err := (&got).Scan(5)

	if err == nil {
		t.Error("expected an error but got nil")
	}
}

func ExampleItemCategory_MarshalGQL() {
	itemCategory := OtherCategory
	(&itemCategory).MarshalGQL(os.Stdout)
	// Output: "other"
}
