package repository

import (
	"bekapod/pkmn-team-graphql/data/model"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"
)

var (
	ErrNoTypes = errors.New("no types found")
	ErrNoType  = errors.New("no type found")
)

var TypeColumns = "types.id, types.name, types.slug"
var TemplatedTypeRelation = `(
		SELECT
	 		array_agg(DISTINCT jsonb_build_object('name', t.name, 'slug', t.slug, 'id', t.id))
		FROM type_damage_relations
			LEFT JOIN types as t on type_damage_relations.related_type_id = t.id
		WHERE type_damage_relations.type_id = types.id AND type_damage_relations.damage_relation_enum = '%s'
	)`
var AllTypeRelations = fmt.Sprintf(TemplatedTypeRelation, "no-damage-to") + ` as no_damage_to,
	` + fmt.Sprintf(TemplatedTypeRelation, "half-damage-to") + ` as half_damage_to,
	` + fmt.Sprintf(TemplatedTypeRelation, "double-damage-to") + ` as double_damage_to,
	` + fmt.Sprintf(TemplatedTypeRelation, "no-damage-from") + ` as no_damage_from,
	` + fmt.Sprintf(TemplatedTypeRelation, "half-damage-from") + ` as half_damage_from,
	` + fmt.Sprintf(TemplatedTypeRelation, "double-damage-from") + ` as double_damage_from`

func (r Type) GetTypes(ctx context.Context) (*model.TypeList, error) {
	types := model.NewEmptyTypeList()

	rows, err := r.db.QueryContext(
		ctx,
		`SELECT `+TypeColumns+`, `+AllTypeRelations+`
		FROM types
		ORDER BY slug ASC`,
	)
	if err != nil {
		return &types, fmt.Errorf("error fetching all types: %w", err)
	}

	defer rows.Close()
	for rows.Next() {
		var t model.Type
		err := rows.Scan(&t.ID, &t.Name, &t.Slug, pq.Array(&t.NoDamageTo.Types), pq.Array(&t.HalfDamageTo.Types), pq.Array(&t.DoubleDamageTo.Types), pq.Array(&t.NoDamageFrom.Types), pq.Array(&t.HalfDamageFrom.Types), pq.Array(&t.DoubleDamageFrom.Types))
		if err != nil {
			return &types, fmt.Errorf("error scanning result in GetAllTypes: %w", err)
		}
		t.NoDamageTo.Total = len(t.NoDamageTo.Types)
		t.HalfDamageTo.Total = len(t.HalfDamageTo.Types)
		t.DoubleDamageTo.Total = len(t.DoubleDamageTo.Types)
		t.NoDamageFrom.Total = len(t.NoDamageFrom.Types)
		t.HalfDamageFrom.Total = len(t.HalfDamageFrom.Types)
		t.DoubleDamageFrom.Total = len(t.DoubleDamageFrom.Types)
		types.AddType(&t)
	}
	err = rows.Err()
	if err != nil {
		return &types, fmt.Errorf("error after fetching all types in GetAllTypes: %w", err)
	}

	return &types, nil
}

func (r Type) GetTypeById(ctx context.Context, id string) (*model.Type, error) {
	t := model.Type{}

	err := r.db.QueryRowContext(
		ctx,
		`SELECT `+TypeColumns+`, `+AllTypeRelations+`
		FROM types
		WHERE types.id = $1
		ORDER BY slug ASC`,
		id,
	).Scan(&t.ID, &t.Name, &t.Slug, pq.Array(&t.NoDamageTo.Types), pq.Array(&t.HalfDamageTo.Types), pq.Array(&t.DoubleDamageTo.Types), pq.Array(&t.NoDamageFrom.Types), pq.Array(&t.HalfDamageFrom.Types), pq.Array(&t.DoubleDamageFrom.Types))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNoType
		}
		return nil, fmt.Errorf("error scanning result in GetTypeById %s: %w", id, err)
	}

	t.NoDamageTo.Total = len(t.NoDamageTo.Types)
	t.HalfDamageTo.Total = len(t.HalfDamageTo.Types)
	t.DoubleDamageTo.Total = len(t.DoubleDamageTo.Types)
	t.NoDamageFrom.Total = len(t.NoDamageFrom.Types)
	t.HalfDamageFrom.Total = len(t.HalfDamageFrom.Types)
	t.DoubleDamageFrom.Total = len(t.DoubleDamageFrom.Types)
	return &t, nil
}
