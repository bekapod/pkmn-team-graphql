package repository

import (
	"bekapod/pkmn-team-graphql/data/model"
	"bekapod/pkmn-team-graphql/log"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

var (
	ErrNoMove = errors.New("no move found")
)

func (r Move) GetMoves(ctx context.Context) (*model.MoveList, error) {
	moves := model.NewEmptyMoveList()

	rows, err := r.db.QueryContext(
		ctx,
		`SELECT moves.*, types.name, types.slug
		FROM moves
			LEFT JOIN types ON moves.type_id = types.id
		ORDER BY moves.slug ASC`,
	)
	if err != nil {
		return &moves, fmt.Errorf("error fetching all moves in GetAllMoves: %w", err)
	}

	defer rows.Close()
	for rows.Next() {
		var m model.Move
		err := rows.Scan(&m.ID, &m.Name, &m.Slug, &m.Accuracy, &m.PP, &m.Power, &m.DamageClass, &m.Effect, &m.EffectChance, &m.Target, &m.Type.ID, &m.Type.Name, &m.Type.Slug)
		if err != nil {
			return &moves, fmt.Errorf("error scanning result in GetAllMoves: %w", err)
		}
		moves.AddMove(m)
	}
	err = rows.Err()
	if err != nil {
		return &moves, fmt.Errorf("error after fetching all moves in GetAllMoves: %w", err)
	}

	return &moves, nil
}

func (r Move) GetMoveById(ctx context.Context, id string) (*model.Move, error) {
	m := model.Move{}

	err := r.db.QueryRowContext(
		ctx,
		`SELECT moves.*, types.name, types.slug
		FROM moves
			LEFT JOIN types ON moves.type_id = types.id
		WHERE moves.id = $1`,
		id,
	).Scan(&m.ID, &m.Name, &m.Slug, &m.Accuracy, &m.PP, &m.Power, &m.DamageClass, &m.Effect, &m.EffectChance, &m.Target, &m.Type.ID, &m.Type.Name, &m.Type.Slug)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNoMove
		}
		return nil, fmt.Errorf("error scanning result in GetMoveById %s: %w", id, err)
	}

	return &m, nil
}

func (r Move) MovesByPokemonIdDataLoader(ctx context.Context) func(pokemonIds []string) ([]*model.MoveList, []error) {
	return func(pokemonIds []string) ([]*model.MoveList, []error) {
		movesByPokemonId := map[string]*model.MoveList{}
		placeholders := make([]string, len(pokemonIds))
		args := make([]interface{}, len(pokemonIds))
		for i := 0; i < len(pokemonIds); i++ {
			placeholders[i] = fmt.Sprintf("$%d", i+1)
			args[i] = pokemonIds[i]
		}

		query := `SELECT moves.*, types.name, types.slug, pokemon_move.pokemon_id
		FROM moves
			LEFT JOIN types ON moves.type_id = types.id
			LEFT JOIN pokemon_move ON moves.id = pokemon_move.move_id
		WHERE pokemon_move.pokemon_id IN (` + strings.Join(placeholders, ",") + `)`

		rows, err := r.db.QueryContext(ctx,
			query,
			args...,
		)
		if err != nil {
			moveList := make([]*model.MoveList, len(pokemonIds))
			emptyMoveList := model.NewEmptyMoveList()
			errors := make([]error, len(pokemonIds))
			for i := range pokemonIds {
				moveList[i] = &emptyMoveList
				errors[i] = fmt.Errorf("error fetching moves for pokemon in MovesByPokemonIdDataLoader: %w", err)
			}
			return moveList, errors
		}

		defer rows.Close()
		for rows.Next() {
			var m model.Move
			var pokemonId string
			err := rows.Scan(&m.ID, &m.Name, &m.Slug, &m.Accuracy, &m.PP, &m.Power, &m.DamageClass, &m.Effect, &m.EffectChance, &m.Target, &m.Type.ID, &m.Type.Name, &m.Type.Slug, &pokemonId)
			if err != nil {
				moveList := make([]*model.MoveList, len(pokemonIds))
				emptyMoveList := model.NewEmptyMoveList()
				errors := make([]error, len(pokemonIds))
				for i := range pokemonIds {
					moveList[i] = &emptyMoveList
					errors[i] = fmt.Errorf("error scanning result moves for pokemon in MovesByPokemonIdDataLoader: %w", err)
				}
				return moveList, errors
			}

			_, ok := movesByPokemonId[pokemonId]
			if !ok {
				ml := model.NewEmptyMoveList()
				movesByPokemonId[pokemonId] = &ml
			}

			movesByPokemonId[pokemonId].AddMove(m)
		}

		moveList := make([]*model.MoveList, len(pokemonIds))
		for i, id := range pokemonIds {
			moveList[i] = movesByPokemonId[id]
			i++
		}

		err = rows.Err()
		if err != nil {
			errors := make([]error, len(pokemonIds))
			for i := range pokemonIds {
				errors[i] = fmt.Errorf("error after fetching moves for pokemon in MovesByPokemonIdDataLoader: %w", err)
			}
			return moveList, errors
		}

		return moveList, nil
	}
}

func (r Move) MovesByTypeIdDataLoader(ctx context.Context) func(typeIds []string) ([]*model.MoveList, []error) {
	return func(typeIds []string) ([]*model.MoveList, []error) {
		movesByTypeId := map[string]*model.MoveList{}
		placeholders := make([]string, len(typeIds))
		args := make([]interface{}, len(typeIds))
		for i := 0; i < len(typeIds); i++ {
			placeholders[i] = fmt.Sprintf("$%d", i+1)
			args[i] = typeIds[i]
		}

		query := `SELECT moves.*, types.name, types.slug
		FROM moves
			LEFT JOIN types ON moves.type_id = types.id
		WHERE type_id IN (` + strings.Join(placeholders, ",") + `)`

		log.Logger.WithField("args", args).Debug(query)
		rows, err := r.db.QueryContext(ctx,
			query,
			args...,
		)
		if err != nil {
			moveList := make([]*model.MoveList, len(typeIds))
			emptyMoveList := model.NewEmptyMoveList()
			errors := make([]error, len(typeIds))
			for i := range typeIds {
				moveList[i] = &emptyMoveList
				errors[i] = fmt.Errorf("error fetching moves for type in MovesByTypeIdDataLoader: %w", err)
			}
			return moveList, errors
		}

		defer rows.Close()
		for rows.Next() {
			var m model.Move
			err := rows.Scan(&m.ID, &m.Name, &m.Slug, &m.Accuracy, &m.PP, &m.Power, &m.DamageClass, &m.Effect, &m.EffectChance, &m.Target, &m.Type.ID, &m.Type.Name, &m.Type.Slug)
			if err != nil {
				moveList := make([]*model.MoveList, len(typeIds))
				emptyMoveList := model.NewEmptyMoveList()
				errors := make([]error, len(typeIds))
				for i := range typeIds {
					moveList[i] = &emptyMoveList
					errors[i] = fmt.Errorf("error scanning result moves for type in MovesByTypeIdDataLoader: %w", err)
				}
				return moveList, errors
			}

			_, ok := movesByTypeId[m.Type.ID]
			if !ok {
				ml := model.NewEmptyMoveList()
				movesByTypeId[m.Type.ID] = &ml
			}

			movesByTypeId[m.Type.ID].AddMove(m)
		}

		moveList := make([]*model.MoveList, len(typeIds))
		for i, id := range typeIds {
			moveList[i] = movesByTypeId[id]
			i++
		}

		err = rows.Err()
		if err != nil {
			errors := make([]error, len(typeIds))
			for i := range typeIds {
				errors[i] = fmt.Errorf("error after fetching moves for type in MovesByTypeIdDataLoader: %w", err)
			}
			return moveList, errors
		}

		return moveList, nil
	}
}
