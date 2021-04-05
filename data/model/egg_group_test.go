package model

import (
	"testing"

	"github.com/go-test/deep"
)

func TestNewEggGroupList(t *testing.T) {
	eggGroups := []EggGroup{
		{
			ID: "123-456",
		},
		{
			ID: "456-789",
		},
	}

	exp := EggGroupList{
		Total:     2,
		EggGroups: eggGroups,
	}

	got := NewEggGroupList(eggGroups)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestNewEmptyEggGroupList(t *testing.T) {
	exp := EggGroupList{
		Total:     0,
		EggGroups: []EggGroup{},
	}

	got := NewEmptyEggGroupList()
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestEggGroupList_AddEggGroup(t *testing.T) {
	eggGroups := EggGroupList{}
	eggGroup1 := EggGroup{}
	eggGroup2 := EggGroup{}
	eggGroups.AddEggGroup(eggGroup1)
	eggGroups.AddEggGroup(eggGroup2)

	if eggGroups.Total != 2 {
		t.Errorf("expected Total of 2, but got %d instead", eggGroups.Total)
	}
}

func TestEggGroup_Scan(t *testing.T) {
	exp := EggGroup{
		ID:   "1f0958a0-48ca-4160-9f18-7e5f06d96d27",
		Name: "Fairy",
		Slug: "fairy",
	}
	got := EggGroup{}
	err := (&got).Scan([]uint8{123, 34, 105, 100, 34, 58, 32, 34, 49, 102, 48, 57, 53, 56, 97, 48, 45, 52, 56, 99, 97, 45, 52, 49, 54, 48, 45, 57, 102, 49, 56, 45, 55, 101, 53, 102, 48, 54, 100, 57, 54, 100, 50, 55, 34, 44, 32, 34, 110, 97, 109, 101, 34, 58, 32, 34, 70, 97, 105, 114, 121, 34, 44, 32, 34, 115, 108, 117, 103, 34, 58, 32, 34, 102, 97, 105, 114, 121, 34, 125})

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestEggGroupList_Scan_Error(t *testing.T) {
	got := EggGroup{}
	err := (&got).Scan([]uint8{123, 34, 115, 108, 111, 116, 34, 58, 32, 49, 44, 32, 34, 116, 121, 112, 101, 34, 58, 32, 123, 34, 105, 100, 34, 58, 32, 34, 48, 55, 98, 57, 101, 98, 125})

	if err == nil {
		t.Error("expected an error but got nil")
	}
}

func TestEggGroupList_Scan_EggGroupError(t *testing.T) {
	got := EggGroup{}
	err := (&got).Scan(5)

	if err == nil {
		t.Error("expected an error but got nil")
	}
}
