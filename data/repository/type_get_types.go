package repository

import (
	"bekapod/pkmn-team-graphql/data/model"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

var (
	ErrNoTypes = errors.New("no types found")
	ErrNoType  = errors.New("no type found")
)

func (r Type) GetTypes(ctx context.Context) (*model.TypeList, error) {
	types := model.NewEmptyTypeList()

	rows, err := r.db.QueryContext(
		ctx,
		"SELECT id, name, slug FROM types ORDER BY slug ASC",
	)
	if err != nil {
		return &types, fmt.Errorf("error fetching all types: %w", err)
	}

	defer rows.Close()
	for rows.Next() {
		var t model.Type
		err := rows.Scan(&t.ID, &t.Name, &t.Slug)
		if err != nil {
			return &types, fmt.Errorf("error scanning result in GetAllTypes: %w", err)
		}
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
		"SELECT id, name, slug FROM types WHERE id = $1",
		id,
	).Scan(&t.ID, &t.Name, &t.Slug)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNoType
		}
		return nil, fmt.Errorf("error scanning result in GetTypeById %s: %w", id, err)
	}

	return &t, nil
}
