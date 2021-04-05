package model

import (
	"testing"

	"github.com/go-test/deep"
)

func TestNewTypeList(t *testing.T) {
	types := []Type{
		{
			ID: "123-456",
		},
		{
			ID: "456-789",
		},
	}

	exp := TypeList{
		Total: 2,
		Types: types,
	}

	got := NewTypeList(types)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestNewEmptyTypeList(t *testing.T) {
	exp := TypeList{
		Total: 0,
		Types: []Type{},
	}

	got := NewEmptyTypeList()
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestTypeList_AddType(t *testing.T) {
	types := TypeList{}
	type1 := &Type{}
	type2 := &Type{}
	types.AddType(type1)
	types.AddType(type2)

	if types.Total != 2 {
		t.Errorf("expected Total of 2, but got %d instead", types.Total)
	}
}

func TestType_Scan(t *testing.T) {
	exp := Type{
		ID:   "1b7d7950-305a-48fa-a771-01f7bc4dad8d",
		Name: "Ground",
		Slug: "ground",
	}
	got := Type{}
	err := (&got).Scan([]uint8{123, 34, 105, 100, 34, 58, 32, 34, 49, 98, 55, 100, 55, 57, 53, 48, 45, 51, 48, 53, 97, 45, 52, 56, 102, 97, 45, 97, 55, 55, 49, 45, 48, 49, 102, 55, 98, 99, 52, 100, 97, 100, 56, 100, 34, 44, 32, 34, 110, 97, 109, 101, 34, 58, 32, 34, 71, 114, 111, 117, 110, 100, 34, 44, 32, 34, 115, 108, 117, 103, 34, 58, 32, 34, 103, 114, 111, 117, 110, 100, 34, 125})

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestTypeList_Scan_Error(t *testing.T) {
	got := Type{}
	err := (&got).Scan([]uint8{123, 34, 115, 108, 111, 116, 34, 58, 32, 49, 44, 32, 34, 116, 121, 112, 101, 34, 58, 32, 123, 34, 105, 100, 34, 58, 32, 34, 48, 55, 98, 57, 101, 98, 125})

	if err == nil {
		t.Error("expected an error but got nil")
	}
}

func TestTypeList_Scan_AbilityError(t *testing.T) {
	got := Type{}
	err := (&got).Scan(5)

	if err == nil {
		t.Error("expected an error but got nil")
	}
}
