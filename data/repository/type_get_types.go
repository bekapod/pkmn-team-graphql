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
			if err == sql.ErrNoRows {
				return &types, ErrNoTypes
			}
			return &types, fmt.Errorf("error scanning result in GetAllTypes: %w", err)
		}
		types.AddType(&t)
	}

	return &types, nil
}
