package repository

import (
	"bekapod/pkmn-team-graphql/data/model"
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-test/deep"
)

func TestType_GetTypes(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM types ORDER BY slug ASC").
		WillReturnRows(mockRowsForGetTypes(false, false, false))

	types := []*model.Type{
		{
			ID:   "a82aa044-d8fd-43b3-9dd6-0ce0bfb29fb1",
			Slug: "dragon",
			Name: "Dragon",
		},
		{
			ID:   "a248c127-8e9c-4f87-8513-c5dbc3385011",
			Slug: "fairy",
			Name: "Fairy",
		},
		{
			ID:   "42b31825-de68-4c1c-bea1-b32a290f1fef",
			Slug: "poison",
			Name: "Poison",
		},
	}

	exp := model.NewTypeList(types)
	got, err := NewType(db).GetTypes(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if diff := deep.Equal(&exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestType_GetTypes_WithQueryError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM types ORDER BY slug ASC").
		WillReturnError(errors.New("I am Error."))

	got, err := NewType(db).GetTypes(context.Background())
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := &model.TypeList{
		Total: 0,
		Types: []*model.Type{},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestType_GetTypes_WithScanError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM types ORDER BY slug ASC").
		WillReturnRows(mockRowsForGetTypes(false, false, true))

	got, err := NewType(db).GetTypes(context.Background())
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := &model.TypeList{
		Total: 0,
		Types: []*model.Type{},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestType_GetTypes_WithRowError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM types ORDER BY slug ASC").
		WillReturnRows(mockRowsForGetTypes(false, true, false))

	got, err := NewType(db).GetTypes(context.Background())
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := &model.TypeList{
		Total: 0,
		Types: []*model.Type{},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestType_GetTypes_WithNoRows(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM types ORDER BY slug ASC").
		WillReturnRows(mockRowsForGetTypes(true, false, false))

	got, err := NewType(db).GetTypes(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	exp := &model.TypeList{
		Total: 0,
		Types: []*model.Type{},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestType_GetTypeById(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM types WHERE id.*").
		WithArgs("a82aa044-d8fd-43b3-9dd6-0ce0bfb29fb1").
		WillReturnRows(mockRowsForGetTypeById(false))

	exp := model.Type{
		ID:   "a82aa044-d8fd-43b3-9dd6-0ce0bfb29fb1",
		Slug: "dragon",
		Name: "Dragon",
	}
	got, err := NewType(db).GetTypeById(context.Background(), "a82aa044-d8fd-43b3-9dd6-0ce0bfb29fb1")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if diff := deep.Equal(&exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestType_GetTypeById_WithQueryError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM types WHERE id.*").
		WithArgs("a82aa044-d8fd-43b3-9dd6-0ce0bfb29fb1").
		WillReturnError(errors.New("I am Error."))

	_, err := NewType(db).GetTypeById(context.Background(), "a82aa044-d8fd-43b3-9dd6-0ce0bfb29fb1")
	if err == nil {
		t.Error("expected an error but got nil")
	}
}

func TestType_GetTypeById_WithNoRows(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM types WHERE id.*").
		WithArgs("a82aa044-d8fd-43b3-9dd6-0ce0bfb29fb1").
		WillReturnRows(mockRowsForGetTypeById(true))

	got, err := NewType(db).GetTypeById(context.Background(), "a82aa044-d8fd-43b3-9dd6-0ce0bfb29fb1")
	if err == nil {
		t.Error("expected an error but got nil")
	}

	var exp *model.Type = nil

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func mockRowsForGetTypes(empty bool, hasRowError bool, hasScanError bool) *sqlmock.Rows {
	if hasScanError {
		rows := sqlmock.NewRows([]string{"id"})
		rows.AddRow("9f61694f-34f0-4531-b5e4-aff9a3d9edde")
		return rows
	}
	rows := sqlmock.NewRows([]string{"id", "name", "slug"})
	if !empty {
		rows.AddRow("a82aa044-d8fd-43b3-9dd6-0ce0bfb29fb1", "Dragon", "dragon").
			AddRow("a248c127-8e9c-4f87-8513-c5dbc3385011", "Fairy", "fairy").
			AddRow("42b31825-de68-4c1c-bea1-b32a290f1fef", "Poison", "poison")
	}
	if hasRowError {
		rows.RowError(0, errors.New("scan error"))
	}
	return rows
}

func mockRowsForGetTypeById(empty bool) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"id", "name", "slug"})
	if !empty {
		rows.AddRow("a82aa044-d8fd-43b3-9dd6-0ce0bfb29fb1", "Dragon", "dragon")
	}
	return rows
}
