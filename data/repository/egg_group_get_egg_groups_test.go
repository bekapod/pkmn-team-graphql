package repository

import (
	"bekapod/pkmn-team-graphql/data/model"
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-test/deep"
)

func TestEggGroup_EggGroupsByPokemonIdDataLoader(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM egg_groups LEFT JOIN pokemon_egg_group ON egg_groups.id = pokemon_egg_group.egg_group_id WHERE pokemon_egg_group.pokemon_id IN (.*)").
		WithArgs("49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a").
		WillReturnRows(mockRowsForEggGroupsByPokemonIdDataLoader(false, false, false, []string{"49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a"}))

	got, err := NewEggGroup(db).EggGroupsByPokemonIdDataLoader(context.Background())([]string{"49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a"})
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	exp := []*model.EggGroupList{
		{
			Total: 1,
			EggGroups: []model.EggGroup{
				{
					ID:   "9f61694f-34f0-4531-b5e4-aff9a3d9edde",
					Slug: "plant",
					Name: "Plant",
				},
			},
		},
		nil,
		{
			Total: 2,
			EggGroups: []model.EggGroup{
				{
					ID:   "9f61694f-34f0-4531-b5e4-aff9a3d9edde",
					Slug: "plant",
					Name: "Plant",
				},
				{
					ID:   "fd22f390-1745-4af2-adbe-9d9eca2db086",
					Slug: "dragon",
					Name: "Dragon",
				},
			},
		},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestEggGroup_EggGroupsByPokemonIdDataLoader_WithQueryError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM egg_groups LEFT JOIN pokemon_egg_group ON egg_groups.id = pokemon_egg_group.egg_group_id WHERE pokemon_egg_group.pokemon_id IN (.*)").
		WithArgs("49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a").
		WillReturnRows(mockRowsForEggGroupsByPokemonIdDataLoader(false, false, false, []string{"49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a"})).
		WillReturnError(errors.New("I am Error."))

	got, err := NewEggGroup(db).EggGroupsByPokemonIdDataLoader(context.Background())([]string{"49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a"})
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := []*model.EggGroupList{
		{
			Total:     0,
			EggGroups: []model.EggGroup{},
		},
		{
			Total:     0,
			EggGroups: []model.EggGroup{},
		},
		{
			Total:     0,
			EggGroups: []model.EggGroup{},
		},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestEggGroup_EggGroupsByPokemonIdDataLoader_WithScanError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM egg_groups LEFT JOIN pokemon_egg_group ON egg_groups.id = pokemon_egg_group.egg_group_id WHERE pokemon_egg_group.pokemon_id IN (.*)").
		WithArgs("49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a").
		WillReturnRows(mockRowsForEggGroupsByPokemonIdDataLoader(false, false, true, []string{"49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a"}))

	got, err := NewEggGroup(db).EggGroupsByPokemonIdDataLoader(context.Background())([]string{"49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a"})
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := []*model.EggGroupList{
		{
			Total:     0,
			EggGroups: []model.EggGroup{},
		},
		{
			Total:     0,
			EggGroups: []model.EggGroup{},
		},
		{
			Total:     0,
			EggGroups: []model.EggGroup{},
		},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestEggGroup_EggGroupsByPokemonIdDataLoader_WithRowError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM egg_groups LEFT JOIN pokemon_egg_group ON egg_groups.id = pokemon_egg_group.egg_group_id WHERE pokemon_egg_group.pokemon_id IN (.*)").
		WithArgs("49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a").
		WillReturnRows(mockRowsForEggGroupsByPokemonIdDataLoader(false, true, false, []string{"49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a"}))

	got, err := NewEggGroup(db).EggGroupsByPokemonIdDataLoader(context.Background())([]string{"49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a"})
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := []*model.EggGroupList{
		{
			Total: 1,
			EggGroups: []model.EggGroup{
				{
					ID:   "9f61694f-34f0-4531-b5e4-aff9a3d9edde",
					Slug: "plant",
					Name: "Plant",
				},
			},
		},
		nil,
		nil,
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func mockRowsForEggGroupsByPokemonIdDataLoader(empty bool, hasRowError bool, hasScanError bool, ids []string) *sqlmock.Rows {
	if hasScanError {
		rows := sqlmock.NewRows([]string{"id"})
		rows.AddRow("9f61694f-34f0-4531-b5e4-aff9a3d9edde")
		return rows
	}
	rows := sqlmock.NewRows([]string{"id", "name", "slug", "pokemon_eggGroup.pokemon_id"})
	if !empty {
		rows.AddRow("9f61694f-34f0-4531-b5e4-aff9a3d9edde", "Plant", "plant", ids[0]).
			AddRow("9f61694f-34f0-4531-b5e4-aff9a3d9edde", "Plant", "plant", ids[2]).
			AddRow("fd22f390-1745-4af2-adbe-9d9eca2db086", "Dragon", "dragon", ids[2])
	}
	if hasRowError {
		rows.RowError(1, errors.New("scan error"))
	}
	return rows
}
