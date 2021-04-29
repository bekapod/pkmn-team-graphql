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

	types := []model.Type{
		dragon,
		fairy,
		poison,
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
		Types: []model.Type{},
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
		Types: []model.Type{},
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
		Types: []model.Type{},
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
		Types: []model.Type{},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestType_GetTypeById(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM types WHERE types.id.*").
		WithArgs("a82aa044-d8fd-43b3-9dd6-0ce0bfb29fb1").
		WillReturnRows(mockRowsForGetTypeById(false))

	exp := dragon
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

	mock.ExpectQuery("SELECT .* FROM types WHERE types.id.*").
		WithArgs("a82aa044-d8fd-43b3-9dd6-0ce0bfb29fb1").
		WillReturnError(errors.New("I am Error."))

	_, err := NewType(db).GetTypeById(context.Background(), "a82aa044-d8fd-43b3-9dd6-0ce0bfb29fb1")
	if err == nil {
		t.Error("expected an error but got nil")
	}
}

func TestType_GetTypeById_WithNoRows(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM types WHERE types.id.*").
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

func TestType_TypesByIdDataLoader(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM types WHERE types.id.*").
		WithArgs(dragon.ID, "", fairy.ID).
		WillReturnRows(mockRowsForTypesByIdDataLoader(false, false, false, []string{dragon.ID, "", fairy.ID}))

	got, err := NewType(db).TypesByIdDataLoader(context.Background())([]string{dragon.ID, "", fairy.ID})
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	exp := []*model.Type{
		&dragon,
		nil,
		&fairy,
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestType_TypesByIdDataLoader_WithQueryError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM types WHERE types.id.*").
		WithArgs(dragon.ID, "", fairy.ID).
		WillReturnRows(mockRowsForTypesByIdDataLoader(false, false, false, []string{dragon.ID, "", fairy.ID})).
		WillReturnError(errors.New("I am Error."))

	got, err := NewType(db).TypesByIdDataLoader(context.Background())([]string{dragon.ID, "", fairy.ID})
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := []*model.Type{
		nil,
		nil,
		nil,
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestType_TypesByIdDataLoader_WithScanError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM types WHERE types.id.*").
		WithArgs(dragon.ID, "", fairy.ID).
		WillReturnRows(mockRowsForTypesByIdDataLoader(false, false, true, []string{dragon.ID, "", fairy.ID}))

	got, err := NewType(db).TypesByIdDataLoader(context.Background())([]string{dragon.ID, "", fairy.ID})
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := []*model.Type{
		nil,
		nil,
		nil,
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestType_TypesByIdDataLoader_WithRowError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM types WHERE types.id.*").
		WithArgs(dragon.ID, "", fairy.ID).
		WillReturnRows(mockRowsForTypesByIdDataLoader(false, true, false, []string{dragon.ID, "", fairy.ID}))

	got, err := NewType(db).TypesByIdDataLoader(context.Background())([]string{dragon.ID, "", fairy.ID})
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := []*model.Type{
		&dragon,
		nil,
		nil,
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

var dragon = model.Type{
	ID:   Dragon.ID,
	Name: Dragon.Name,
	Slug: Dragon.Slug,
	NoDamageTo: model.TypeList{
		Total: 1,
		Types: []model.Type{
			Fairy,
		},
	},
	HalfDamageTo: model.TypeList{
		Total: 1,
		Types: []model.Type{
			Steel,
		},
	},
	DoubleDamageTo: model.TypeList{
		Total: 1,
		Types: []model.Type{
			Dragon,
		},
	},
	NoDamageFrom: model.TypeList{
		Total: 0,
		Types: nil,
	},
	HalfDamageFrom: model.TypeList{
		Total: 4,
		Types: []model.Type{
			Grass,
			Electric,
			Fire,
			Water,
		},
	},
	DoubleDamageFrom: model.TypeList{
		Total: 3,
		Types: []model.Type{
			Ice,
			Fairy,
			Dragon,
		},
	},
}

var fairy = model.Type{
	ID:   Fairy.ID,
	Name: Fairy.Name,
	Slug: Fairy.Slug,
	NoDamageTo: model.TypeList{
		Total: 0,
		Types: nil,
	},
	HalfDamageTo: model.TypeList{
		Total: 3,
		Types: []model.Type{
			Steel,
			Poison,
			Fire,
		},
	},
	DoubleDamageTo: model.TypeList{
		Total: 3,
		Types: []model.Type{
			Fighting,
			Dark,
			Dragon,
		},
	},
	NoDamageFrom: model.TypeList{
		Total: 1,
		Types: []model.Type{
			Dragon,
		},
	},
	HalfDamageFrom: model.TypeList{
		Total: 3,
		Types: []model.Type{
			Bug,
			Fighting,
			Dark,
		},
	},
	DoubleDamageFrom: model.TypeList{
		Total: 2,
		Types: []model.Type{
			Steel,
			Poison,
		},
	},
}

var poison = model.Type{
	ID:   Poison.ID,
	Name: Poison.Name,
	Slug: Poison.Slug,
	NoDamageTo: model.TypeList{
		Total: 1,
		Types: []model.Type{
			Steel,
		},
	},
	HalfDamageTo: model.TypeList{
		Total: 4,
		Types: []model.Type{
			Ghost,
			Ground,
			Poison,
			Rock,
		},
	},
	DoubleDamageTo: model.TypeList{
		Total: 2,
		Types: []model.Type{
			Grass,
			Fairy,
		},
	},
	NoDamageFrom: model.TypeList{
		Total: 0,
		Types: nil,
	},
	HalfDamageFrom: model.TypeList{
		Total: 5,
		Types: []model.Type{
			Grass,
			Poison,
			Bug,
			Fighting,
			Fairy,
		},
	},
	DoubleDamageFrom: model.TypeList{
		Total: 2,
		Types: []model.Type{
			Ground,
			Psychic,
		},
	},
}

func mockRowsForGetTypes(empty bool, hasRowError bool, hasScanError bool) *sqlmock.Rows {
	if hasScanError {
		rows := sqlmock.NewRows([]string{"id"})
		rows.AddRow("9f61694f-34f0-4531-b5e4-aff9a3d9edde")
		return rows
	}
	rows := sqlmock.NewRows([]string{"id", "name", "slug", "no_damage_to", "half_damage_to", "double_damage_to", "no_damage_from", "half_damage_from", "double_damage_from"})
	if !empty {
		rows.AddRow("a82aa044-d8fd-43b3-9dd6-0ce0bfb29fb1", "Dragon", "dragon", `{"{\"id\": \"a248c127-8e9c-4f87-8513-c5dbc3385011\", \"name\": \"Fairy\", \"slug\": \"fairy\"}"}`, `{"{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}"}`, `{"{\"id\": \"a82aa044-d8fd-43b3-9dd6-0ce0bfb29fb1\", \"name\": \"Dragon\", \"slug\": \"dragon\"}"}`, nil, `{"{\"id\": \"07b9eb0f-e676-4649-bf2e-0e5ef2c2c2e3\", \"name\": \"Grass\", \"slug\": \"grass\"}","{\"id\": \"b3468930-5d60-418f-aaaf-f16cbc93f08d\", \"name\": \"Electric\", \"slug\": \"electric\"}","{\"id\": \"d43f57ab-e5b3-4912-a667-9f237d21d391\", \"name\": \"Fire\", \"slug\": \"fire\"}","{\"id\": \"de384e1c-89aa-44de-88fe-5e914e468f2b\", \"name\": \"Water\", \"slug\": \"water\"}"}`, `{"{\"id\": \"1dcc9d3c-55d4-4d33-809a-d1580c6e6542\", \"name\": \"Ice\", \"slug\": \"ice\"}","{\"id\": \"a248c127-8e9c-4f87-8513-c5dbc3385011\", \"name\": \"Fairy\", \"slug\": \"fairy\"}","{\"id\": \"a82aa044-d8fd-43b3-9dd6-0ce0bfb29fb1\", \"name\": \"Dragon\", \"slug\": \"dragon\"}"}`).
			AddRow("a248c127-8e9c-4f87-8513-c5dbc3385011", "Fairy", "fairy", nil, `{"{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}","{\"id\": \"42b31825-de68-4c1c-bea1-b32a290f1fef\", \"name\": \"Poison\", \"slug\": \"poison\"}","{\"id\": \"d43f57ab-e5b3-4912-a667-9f237d21d391\", \"name\": \"Fire\", \"slug\": \"fire\"}"}`, `{"{\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}","{\"id\": \"9ca47516-fff8-4f5e-8eb5-582c1f7c05af\", \"name\": \"Dark\", \"slug\": \"dark\"}","{\"id\": \"a82aa044-d8fd-43b3-9dd6-0ce0bfb29fb1\", \"name\": \"Dragon\", \"slug\": \"dragon\"}"}`, `{"{\"id\": \"a82aa044-d8fd-43b3-9dd6-0ce0bfb29fb1\", \"name\": \"Dragon\", \"slug\": \"dragon\"}"}`, `{"{\"id\": \"56dddb9a-3623-43c5-8228-ea24d598afe7\", \"name\": \"Bug\", \"slug\": \"bug\"}","{\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}","{\"id\": \"9ca47516-fff8-4f5e-8eb5-582c1f7c05af\", \"name\": \"Dark\", \"slug\": \"dark\"}"}`, `{"{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}","{\"id\": \"42b31825-de68-4c1c-bea1-b32a290f1fef\", \"name\": \"Poison\", \"slug\": \"poison\"}"}`).
			AddRow("42b31825-de68-4c1c-bea1-b32a290f1fef", "Poison", "poison", `{"{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}"}`, `{"{\"id\": \"027f1455-8e6f-4891-8c62-d75bb6c49dae\", \"name\": \"Ghost\", \"slug\": \"ghost\"}","{\"id\": \"1b7d7950-305a-48fa-a771-01f7bc4dad8d\", \"name\": \"Ground\", \"slug\": \"ground\"}","{\"id\": \"42b31825-de68-4c1c-bea1-b32a290f1fef\", \"name\": \"Poison\", \"slug\": \"poison\"}","{\"id\": \"5179f383-b765-4cc7-b9f9-8b1a3ba93019\", \"name\": \"Rock\", \"slug\": \"rock\"}"}`, `{"{\"id\": \"07b9eb0f-e676-4649-bf2e-0e5ef2c2c2e3\", \"name\": \"Grass\", \"slug\": \"grass\"}","{\"id\": \"a248c127-8e9c-4f87-8513-c5dbc3385011\", \"name\": \"Fairy\", \"slug\": \"fairy\"}"}`, nil, `{"{\"id\": \"07b9eb0f-e676-4649-bf2e-0e5ef2c2c2e3\", \"name\": \"Grass\", \"slug\": \"grass\"}","{\"id\": \"42b31825-de68-4c1c-bea1-b32a290f1fef\", \"name\": \"Poison\", \"slug\": \"poison\"}","{\"id\": \"56dddb9a-3623-43c5-8228-ea24d598afe7\", \"name\": \"Bug\", \"slug\": \"bug\"}","{\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}","{\"id\": \"a248c127-8e9c-4f87-8513-c5dbc3385011\", \"name\": \"Fairy\", \"slug\": \"fairy\"}"}`, `{"{\"id\": \"1b7d7950-305a-48fa-a771-01f7bc4dad8d\", \"name\": \"Ground\", \"slug\": \"ground\"}","{\"id\": \"2222c839-3c6e-4727-b6b5-a946bb8af5fa\", \"name\": \"Psychic\", \"slug\": \"psychic\"}"}`)
	}
	if hasRowError {
		rows.RowError(0, errors.New("scan error"))
	}
	return rows
}

func mockRowsForGetTypeById(empty bool) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"id", "name", "slug", "no_damage_to", "half_damage_to", "double_damage_to", "no_damage_from", "half_damage_from", "double_damage_from"})
	if !empty {
		rows.AddRow("a82aa044-d8fd-43b3-9dd6-0ce0bfb29fb1", "Dragon", "dragon", `{"{\"id\": \"a248c127-8e9c-4f87-8513-c5dbc3385011\", \"name\": \"Fairy\", \"slug\": \"fairy\"}"}`, `{"{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}"}`, `{"{\"id\": \"a82aa044-d8fd-43b3-9dd6-0ce0bfb29fb1\", \"name\": \"Dragon\", \"slug\": \"dragon\"}"}`, nil, `{"{\"id\": \"07b9eb0f-e676-4649-bf2e-0e5ef2c2c2e3\", \"name\": \"Grass\", \"slug\": \"grass\"}","{\"id\": \"b3468930-5d60-418f-aaaf-f16cbc93f08d\", \"name\": \"Electric\", \"slug\": \"electric\"}","{\"id\": \"d43f57ab-e5b3-4912-a667-9f237d21d391\", \"name\": \"Fire\", \"slug\": \"fire\"}","{\"id\": \"de384e1c-89aa-44de-88fe-5e914e468f2b\", \"name\": \"Water\", \"slug\": \"water\"}"}`, `{"{\"id\": \"1dcc9d3c-55d4-4d33-809a-d1580c6e6542\", \"name\": \"Ice\", \"slug\": \"ice\"}","{\"id\": \"a248c127-8e9c-4f87-8513-c5dbc3385011\", \"name\": \"Fairy\", \"slug\": \"fairy\"}","{\"id\": \"a82aa044-d8fd-43b3-9dd6-0ce0bfb29fb1\", \"name\": \"Dragon\", \"slug\": \"dragon\"}"}`)
	}
	return rows
}

func mockRowsForTypesByIdDataLoader(empty bool, hasRowError bool, hasScanError bool, ids []string) *sqlmock.Rows {
	if hasScanError {
		rows := sqlmock.NewRows([]string{"id"})
		rows.AddRow("9f61694f-34f0-4531-b5e4-aff9a3d9edde")
		return rows
	}
	rows := sqlmock.NewRows([]string{"id", "name", "slug", "no_damage_to", "half_damage_to", "double_damage_to", "no_damage_from", "half_damage_from", "double_damage_from"})
	if !empty {
		rows.AddRow(ids[0], "Dragon", "dragon", `{"{\"id\": \"a248c127-8e9c-4f87-8513-c5dbc3385011\", \"name\": \"Fairy\", \"slug\": \"fairy\"}"}`, `{"{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}"}`, `{"{\"id\": \"a82aa044-d8fd-43b3-9dd6-0ce0bfb29fb1\", \"name\": \"Dragon\", \"slug\": \"dragon\"}"}`, nil, `{"{\"id\": \"07b9eb0f-e676-4649-bf2e-0e5ef2c2c2e3\", \"name\": \"Grass\", \"slug\": \"grass\"}","{\"id\": \"b3468930-5d60-418f-aaaf-f16cbc93f08d\", \"name\": \"Electric\", \"slug\": \"electric\"}","{\"id\": \"d43f57ab-e5b3-4912-a667-9f237d21d391\", \"name\": \"Fire\", \"slug\": \"fire\"}","{\"id\": \"de384e1c-89aa-44de-88fe-5e914e468f2b\", \"name\": \"Water\", \"slug\": \"water\"}"}`, `{"{\"id\": \"1dcc9d3c-55d4-4d33-809a-d1580c6e6542\", \"name\": \"Ice\", \"slug\": \"ice\"}","{\"id\": \"a248c127-8e9c-4f87-8513-c5dbc3385011\", \"name\": \"Fairy\", \"slug\": \"fairy\"}","{\"id\": \"a82aa044-d8fd-43b3-9dd6-0ce0bfb29fb1\", \"name\": \"Dragon\", \"slug\": \"dragon\"}"}`).
			AddRow(ids[2], "Fairy", "fairy", nil, `{"{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}","{\"id\": \"42b31825-de68-4c1c-bea1-b32a290f1fef\", \"name\": \"Poison\", \"slug\": \"poison\"}","{\"id\": \"d43f57ab-e5b3-4912-a667-9f237d21d391\", \"name\": \"Fire\", \"slug\": \"fire\"}"}`, `{"{\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}","{\"id\": \"9ca47516-fff8-4f5e-8eb5-582c1f7c05af\", \"name\": \"Dark\", \"slug\": \"dark\"}","{\"id\": \"a82aa044-d8fd-43b3-9dd6-0ce0bfb29fb1\", \"name\": \"Dragon\", \"slug\": \"dragon\"}"}`, `{"{\"id\": \"a82aa044-d8fd-43b3-9dd6-0ce0bfb29fb1\", \"name\": \"Dragon\", \"slug\": \"dragon\"}"}`, `{"{\"id\": \"56dddb9a-3623-43c5-8228-ea24d598afe7\", \"name\": \"Bug\", \"slug\": \"bug\"}","{\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}","{\"id\": \"9ca47516-fff8-4f5e-8eb5-582c1f7c05af\", \"name\": \"Dark\", \"slug\": \"dark\"}"}`, `{"{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}","{\"id\": \"42b31825-de68-4c1c-bea1-b32a290f1fef\", \"name\": \"Poison\", \"slug\": \"poison\"}"}`).
			AddRow(ids[2], "Fairy", "fairy", nil, `{"{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}","{\"id\": \"42b31825-de68-4c1c-bea1-b32a290f1fef\", \"name\": \"Poison\", \"slug\": \"poison\"}","{\"id\": \"d43f57ab-e5b3-4912-a667-9f237d21d391\", \"name\": \"Fire\", \"slug\": \"fire\"}"}`, `{"{\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}","{\"id\": \"9ca47516-fff8-4f5e-8eb5-582c1f7c05af\", \"name\": \"Dark\", \"slug\": \"dark\"}","{\"id\": \"a82aa044-d8fd-43b3-9dd6-0ce0bfb29fb1\", \"name\": \"Dragon\", \"slug\": \"dragon\"}"}`, `{"{\"id\": \"a82aa044-d8fd-43b3-9dd6-0ce0bfb29fb1\", \"name\": \"Dragon\", \"slug\": \"dragon\"}"}`, `{"{\"id\": \"56dddb9a-3623-43c5-8228-ea24d598afe7\", \"name\": \"Bug\", \"slug\": \"bug\"}","{\"id\": \"9093f701-0f10-4e59-aff7-05748b23f953\", \"name\": \"Fighting\", \"slug\": \"fighting\"}","{\"id\": \"9ca47516-fff8-4f5e-8eb5-582c1f7c05af\", \"name\": \"Dark\", \"slug\": \"dark\"}"}`, `{"{\"id\": \"05cd51bd-23ca-4736-b8ec-aa93aca68a8b\", \"name\": \"Steel\", \"slug\": \"steel\"}","{\"id\": \"42b31825-de68-4c1c-bea1-b32a290f1fef\", \"name\": \"Poison\", \"slug\": \"poison\"}"}`)
	}
	if hasRowError {
		rows.RowError(1, errors.New("row error"))
	}
	return rows
}
