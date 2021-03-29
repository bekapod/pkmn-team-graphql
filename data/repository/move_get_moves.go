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
	ErrNoMoves = errors.New("no moves found")
	ErrNoMove  = errors.New("no move found")
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

func (r Move) GetMoveById(ctx context.Context, id string) (*model.Move, error) {
	m := model.Move{}

	err := r.db.QueryRowContext(
		ctx,
		"SELECT id, name, slug, accuracy, pp, power, damage_class_enum, effect, effect_chance, target, type_id FROM moves WHERE id = $1",
		id,
	).Scan(&m.ID, &m.Name, &m.Slug, &m.Accuracy, &m.PP, &m.Power, &m.DamageClass, &m.Effect, &m.EffectChance, &m.Target, &m.TypeId)
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

		query := "SELECT id, name, slug, accuracy, pp, power, damage_class_enum, effect, effect_chance, target, type_id, pokemon_move.pokemon_id FROM moves LEFT JOIN pokemon_move ON moves.id = pokemon_move.move_id WHERE pokemon_move.pokemon_id IN (" + strings.Join(placeholders, ",") + ")"

		log.Logger.WithField("args", args).Debug(query)
		rows, err := r.db.QueryContext(ctx,
			query,
			args...,
		)
		if err != nil {
			panic(fmt.Errorf("error fetching moves for pokemon: %w", err))
		}

		defer rows.Close()
		for rows.Next() {
			var m model.Move
			var pokemonId string
			err := rows.Scan(&m.ID, &m.Name, &m.Slug, &m.Accuracy, &m.PP, &m.Power, &m.DamageClass, &m.Effect, &m.EffectChance, &m.Target, &m.TypeId, &pokemonId)
			if err != nil {
				panic(fmt.Errorf("error scanning result in MovesByPokemonId: %w", err))
			}

			_, ok := movesByPokemonId[pokemonId]
			if !ok {
				ml := model.NewEmptyMoveList()
				movesByPokemonId[pokemonId] = &ml
			}

			movesByPokemonId[pokemonId].AddMove(&m)
		}

		moveList := make([]*model.MoveList, len(pokemonIds))
		for i, id := range pokemonIds {
			moveList[i] = movesByPokemonId[id]
			i++
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

		query := "SELECT id, name, slug, accuracy, pp, power, damage_class_enum, effect, effect_chance, target, type_id FROM moves WHERE type_id IN (" + strings.Join(placeholders, ",") + ")"

		log.Logger.WithField("args", args).Debug(query)
		rows, err := r.db.QueryContext(ctx,
			query,
			args...,
		)
		if err != nil {
			panic(fmt.Errorf("error fetching moves for type: %w", err))
		}

		defer rows.Close()
		for rows.Next() {
			var m model.Move
			err := rows.Scan(&m.ID, &m.Name, &m.Slug, &m.Accuracy, &m.PP, &m.Power, &m.DamageClass, &m.Effect, &m.EffectChance, &m.Target, &m.TypeId)
			if err != nil {
				panic(fmt.Errorf("error scanning result in MovesByTypeId: %w", err))
			}

			_, ok := movesByTypeId[m.TypeId]
			if !ok {
				ml := model.NewEmptyMoveList()
				movesByTypeId[m.TypeId] = &ml
			}

			movesByTypeId[m.TypeId].AddMove(&m)
		}

		moveList := make([]*model.MoveList, len(typeIds))
		for i, id := range typeIds {
			moveList[i] = movesByTypeId[id]
			i++
		}

		return moveList, nil
	}
}
