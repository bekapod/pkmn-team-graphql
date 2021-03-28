package repository

import (
	"bekapod/pkmn-team-graphql/data/model"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

var (
	ErrNoMoves = errors.New("no moves found")
)

func (r Move) GetMoves(ctx context.Context) (*model.MoveList, error) {
	moves := model.NewEmptyMoveList()

	rows, err := r.db.QueryContext(
		ctx,
		"SELECT id, name, slug, accuracy, pp, power, damage_class_enum, effect, effect_chance, target, type_id FROM moves ORDER BY slug ASC",
	)
	if err != nil {
		return &moves, fmt.Errorf("error fetching all moves: %w", err)
	}

	defer rows.Close()
	for rows.Next() {
		var m model.Move
		err := rows.Scan(&m.ID, &m.Name, &m.Slug, &m.Accuracy, &m.PP, &m.Power, &m.DamageClass, &m.Effect, &m.EffectChance, &m.Target, &m.TypeId)
		if err != nil {
			if err == sql.ErrNoRows {
				return &moves, ErrNoMoves
			}
			return &moves, fmt.Errorf("error scanning result in GetAllMoves: %w", err)
		}
		moves.AddMove(&m)
	}

	return &moves, nil
}
