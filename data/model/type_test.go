package model

import (
	"bekapod/pkmn-team-graphql/data/db"
	"reflect"
	"testing"

	"github.com/go-test/deep"
)

func TestNewTypeFromDb(t *testing.T) {
	typ := db.TypeModel{
		InnerType: db.InnerType{
			ID:   "123",
			Slug: "some-type",
			Name: "Some Type",
		},
	}
	exp := Type{
		ID:   typ.ID,
		Slug: typ.Slug,
		Name: typ.Name,
	}

	got := NewTypeFromDb(typ)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestNewTypeEdgeFromDb(t *testing.T) {
	typ := db.TypeModel{
		InnerType: db.InnerType{
			ID:   "123",
			Slug: "some-type",
			Name: "Some Type",
		},
	}
	exp := TypeEdge{
		Cursor: typ.ID,
		Node: &Type{
			ID:   typ.ID,
			Slug: typ.Slug,
			Name: typ.Name,
		},
	}

	got := NewTypeEdgeFromDb(typ)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestNewEmptyTypeConnection(t *testing.T) {
	exp := TypeConnection{
		PageInfo: &PageInfo{
			HasNextPage:     false,
			HasPreviousPage: false,
		},
		Edges: []*TypeEdge{},
	}

	got := NewEmptyTypeConnection()
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestTypeConnection_AddType(t *testing.T) {
	types := NewEmptyTypeConnection()
	type1 := &TypeEdge{}
	type2 := &TypeEdge{}
	types.AddEdge(type1)
	types.AddEdge(type2)

	if *types.PageInfo.StartCursor != type1.Cursor {
		t.Errorf("expected start cursor to be %s but got %s", type1.Cursor, *types.PageInfo.StartCursor)
	}

	if *types.PageInfo.EndCursor != type2.Cursor {
		t.Errorf("expected start cursor to be %s but got %s", type2.Cursor, *types.PageInfo.StartCursor)
	}

	if !reflect.DeepEqual([]*TypeEdge{type1, type2}, types.Edges) {
		t.Errorf("the types added do not match")
	}
}
