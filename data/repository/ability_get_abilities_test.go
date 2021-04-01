package repository

import (
	"bekapod/pkmn-team-graphql/data/model"
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-test/deep"
)

func TestAbility_GetAbilities(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM abilities ORDER BY slug ASC").
		WillReturnRows(mockRowsForGetAbilities(false, false, false))

	abilities := []*model.Ability{
		{
			ID:     "2064de60-5a0f-46df-859d-7ba0acdc9aed",
			Slug:   "sap-sipper",
			Name:   "Sap Sipper",
			Effect: "Absorbs grass moves, raising Attack one stage.",
		},
		{
			ID:     "557336fa-392b-4429-812f-cca5fb92eceb",
			Slug:   "magic-bounce",
			Name:   "Magic Bounce",
			Effect: "Reflects most non-damaging moves back at their user.",
		},
	}

	exp := model.NewAbilityList(abilities)
	got, err := NewAbility(db).GetAbilities(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if diff := deep.Equal(&exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestAbility_GetAbilities_WithQueryError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM abilities ORDER BY slug ASC").
		WillReturnError(errors.New("I am Error."))

	got, err := NewAbility(db).GetAbilities(context.Background())
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := &model.AbilityList{
		Total:     0,
		Abilities: []*model.Ability{},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestAbility_GetAbilities_WithScanError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM abilities ORDER BY slug ASC").
		WillReturnRows(mockRowsForGetAbilities(false, false, true))

	got, err := NewAbility(db).GetAbilities(context.Background())
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := &model.AbilityList{
		Total:     0,
		Abilities: []*model.Ability{},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestAbility_GetAbilities_WithRowError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM abilities ORDER BY slug ASC").
		WillReturnRows(mockRowsForGetAbilities(false, true, false))

	got, err := NewAbility(db).GetAbilities(context.Background())
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := &model.AbilityList{
		Total:     0,
		Abilities: []*model.Ability{},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestAbility_GetAbilities_WithNoRows(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM abilities ORDER BY slug ASC").
		WillReturnRows(mockRowsForGetAbilities(true, false, false))

	got, err := NewAbility(db).GetAbilities(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	exp := &model.AbilityList{
		Total:     0,
		Abilities: []*model.Ability{},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestAbility_GetAbilityById(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM abilities WHERE id.*").
		WithArgs("9f61694f-34f0-4531-b5e4-aff9a3d9edde").
		WillReturnRows(mockRowsForGetAbilityById(false))

	exp := model.Ability{
		ID:     "2064de60-5a0f-46df-859d-7ba0acdc9aed",
		Slug:   "sap-sipper",
		Name:   "Sap Sipper",
		Effect: "Absorbs grass moves, raising Attack one stage.",
	}
	got, err := NewAbility(db).GetAbilityById(context.Background(), "9f61694f-34f0-4531-b5e4-aff9a3d9edde")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if diff := deep.Equal(&exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestAbility_GetAbilityById_WithQueryError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM abilities WHERE id.*").
		WithArgs("9f61694f-34f0-4531-b5e4-aff9a3d9edde").
		WillReturnError(errors.New("I am Error."))

	_, err := NewAbility(db).GetAbilityById(context.Background(), "9f61694f-34f0-4531-b5e4-aff9a3d9edde")
	if err == nil {
		t.Error("expected an error but got nil")
	}
}

func TestAbility_GetAbilityById_WithNoRows(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM abilities WHERE id.*").
		WithArgs("9f61694f-34f0-4531-b5e4-aff9a3d9edde").
		WillReturnRows(mockRowsForGetAbilityById(true))

	got, err := NewAbility(db).GetAbilityById(context.Background(), "9f61694f-34f0-4531-b5e4-aff9a3d9edde")
	if err == nil {
		t.Error("expected an error but got nil")
	}

	var exp *model.Ability = nil

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestAbility_AbilitiesByPokemonIdDataLoader(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM abilities LEFT JOIN pokemon_ability ON abilities.id = pokemon_ability.ability_id WHERE pokemon_ability.pokemon_id IN (.*)").
		WithArgs("49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a").
		WillReturnRows(mockRowsForAbilitiesByPokemonIdDataLoader(false, false, false, []string{"49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a"}))

	got, err := NewAbility(db).AbilitiesByPokemonIdDataLoader(context.Background())([]string{"49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a"})
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	exp := []*model.AbilityList{
		{
			Total: 1,
			Abilities: []*model.Ability{
				{
					ID:     "2064de60-5a0f-46df-859d-7ba0acdc9aed",
					Slug:   "sap-sipper",
					Name:   "Sap Sipper",
					Effect: "Absorbs grass moves, raising Attack one stage.",
				},
			},
		},
		nil,
		{
			Total: 2,
			Abilities: []*model.Ability{
				{
					ID:     "2064de60-5a0f-46df-859d-7ba0acdc9aed",
					Slug:   "sap-sipper",
					Name:   "Sap Sipper",
					Effect: "Absorbs grass moves, raising Attack one stage.",
				},
				{
					ID:     "557336fa-392b-4429-812f-cca5fb92eceb",
					Slug:   "magic-bounce",
					Name:   "Magic Bounce",
					Effect: "Reflects most non-damaging moves back at their user.",
				},
			},
		},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestAbility_AbilitiesByPokemonIdDataLoader_WithQueryError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM abilities LEFT JOIN pokemon_ability ON abilities.id = pokemon_ability.ability_id WHERE pokemon_ability.pokemon_id IN (.*)").
		WithArgs("49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a").
		WillReturnRows(mockRowsForAbilitiesByPokemonIdDataLoader(false, false, false, []string{"49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a"})).
		WillReturnError(errors.New("I am Error."))

	got, err := NewAbility(db).AbilitiesByPokemonIdDataLoader(context.Background())([]string{"49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a"})
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := []*model.AbilityList{
		{
			Total:     0,
			Abilities: []*model.Ability{},
		},
		{
			Total:     0,
			Abilities: []*model.Ability{},
		},
		{
			Total:     0,
			Abilities: []*model.Ability{},
		},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestAbility_AbilitiesByPokemonIdDataLoader_WithScanError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM abilities LEFT JOIN pokemon_ability ON abilities.id = pokemon_ability.ability_id WHERE pokemon_ability.pokemon_id IN (.*)").
		WithArgs("49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a").
		WillReturnRows(mockRowsForAbilitiesByPokemonIdDataLoader(false, false, true, []string{"49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a"}))

	got, err := NewAbility(db).AbilitiesByPokemonIdDataLoader(context.Background())([]string{"49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a"})
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := []*model.AbilityList{
		{
			Total:     0,
			Abilities: []*model.Ability{},
		},
		{
			Total:     0,
			Abilities: []*model.Ability{},
		},
		{
			Total:     0,
			Abilities: []*model.Ability{},
		},
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestAbility_AbilitiesByPokemonIdDataLoader_WithRowError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("SELECT .* FROM abilities LEFT JOIN pokemon_ability ON abilities.id = pokemon_ability.ability_id WHERE pokemon_ability.pokemon_id IN (.*)").
		WithArgs("49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a").
		WillReturnRows(mockRowsForAbilitiesByPokemonIdDataLoader(false, true, false, []string{"49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a"}))

	got, err := NewAbility(db).AbilitiesByPokemonIdDataLoader(context.Background())([]string{"49653637-1d35-4138-98eb-14305a2741a0", "742eb94e-829a-4bd2-a409-428167a389da", "49de1627-e7b3-4a54-8d42-0ed7c795f28a"})
	if err == nil {
		t.Error("expected an error but got nil")
	}

	exp := []*model.AbilityList{
		{
			Total: 1,
			Abilities: []*model.Ability{
				{
					ID:     "2064de60-5a0f-46df-859d-7ba0acdc9aed",
					Slug:   "sap-sipper",
					Name:   "Sap Sipper",
					Effect: "Absorbs grass moves, raising Attack one stage.",
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

func mockRowsForGetAbilities(empty bool, hasRowError bool, hasScanError bool) *sqlmock.Rows {
	if hasScanError {
		rows := sqlmock.NewRows([]string{"id"})
		rows.AddRow("2064de60-5a0f-46df-859d-7ba0acdc9aed")
		return rows
	}
	rows := sqlmock.NewRows([]string{"id", "name", "slug", "effect"})
	if !empty {
		rows.AddRow("2064de60-5a0f-46df-859d-7ba0acdc9aed", "Sap Sipper", "sap-sipper", "Absorbs grass moves, raising Attack one stage.").
			AddRow("557336fa-392b-4429-812f-cca5fb92eceb", "Magic Bounce", "magic-bounce", "Reflects most non-damaging moves back at their user.")
	}
	if hasRowError {
		rows.RowError(0, errors.New("scan error"))
	}
	return rows
}

func mockRowsForGetAbilityById(empty bool) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"id", "name", "slug", "effect"})
	if !empty {
		rows.AddRow("2064de60-5a0f-46df-859d-7ba0acdc9aed", "Sap Sipper", "sap-sipper", "Absorbs grass moves, raising Attack one stage.")
	}
	return rows
}

func mockRowsForAbilitiesByPokemonIdDataLoader(empty bool, hasRowError bool, hasScanError bool, ids []string) *sqlmock.Rows {
	if hasScanError {
		rows := sqlmock.NewRows([]string{"id"})
		rows.AddRow("2064de60-5a0f-46df-859d-7ba0acdc9aed")
		return rows
	}
	rows := sqlmock.NewRows([]string{"id", "name", "slug", "effect", "pokemon_ability.pokemon_id"})
	if !empty {
		rows.AddRow("2064de60-5a0f-46df-859d-7ba0acdc9aed", "Sap Sipper", "sap-sipper", "Absorbs grass moves, raising Attack one stage.", ids[0]).
			AddRow("2064de60-5a0f-46df-859d-7ba0acdc9aed", "Sap Sipper", "sap-sipper", "Absorbs grass moves, raising Attack one stage.", ids[2]).
			AddRow("557336fa-392b-4429-812f-cca5fb92eceb", "Magic Bounce", "magic-bounce", "Reflects most non-damaging moves back at their user.", ids[2])
	}
	if hasRowError {
		rows.RowError(1, errors.New("scan error"))
	}
	return rows
}
