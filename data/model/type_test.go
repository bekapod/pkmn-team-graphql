package model

import (
	"reflect"
	"testing"

	"github.com/go-test/deep"
)

func TestNewTypeList(t *testing.T) {
	types := []*Type{
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
		Types: []*Type{},
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

	if !reflect.DeepEqual([]*Type{type1, type2}, types.Types) {
		t.Errorf("the types added do not match")
	}
}
