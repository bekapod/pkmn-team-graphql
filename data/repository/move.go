package repository

import (
	"bekapod/pkmn-team-graphql/data/db"
	"bekapod/pkmn-team-graphql/data/model"
	"context"
	"errors"
	"fmt"
)

type Move struct {
	client *db.PrismaClient
}

func NewMove(client *db.PrismaClient) Move {
	return Move{
		client: client,
	}
}

func (r Move) GetMoves(ctx context.Context) (*model.MoveConnection, error) {
	moves := model.NewEmptyMoveConnection()

	results, err := r.client.Move.FindMany().Exec(ctx)

	if errors.Is(err, db.ErrNotFound) {
		return &moves, nil
	}

	if err != nil {
		return &moves, fmt.Errorf("error getting moves: %s", err)
	}

	for _, result := range results {
		m := model.NewMoveEdgeFromDb(result)
		moves.AddEdge(&m)
	}

	return &moves, nil
}

func (r Move) GetMoveById(ctx context.Context, id string) (*model.Move, error) {
	result, err := r.client.Move.FindUnique(db.Move.ID.Equals(id)).Exec(ctx)

	if errors.Is(err, db.ErrNotFound) {
		return nil, fmt.Errorf("couldn't find move by id: %s", id)
	}

	if err != nil {
		return nil, fmt.Errorf("error getting move by id: %s, error: %s", id, err)
	}

	m := model.NewMoveFromDb(*result)
	return &m, nil
}

func (r Move) MoveByIdDataLoader(ctx context.Context) func(ids []string) ([]*model.Move, []error) {
	return func(ids []string) ([]*model.Move, []error) {
		movesById := map[string]*model.Move{}
		results, err := r.client.Move.FindMany(db.Move.ID.In(ids)).Exec(ctx)
		moves := make([]*model.Move, len(ids))

		if err != nil {
			errors := make([]error, len(ids))
			for i, id := range ids {
				errors[i] = fmt.Errorf("error loading move by id %s in dataloader %w", id, err)
			}

			return moves, errors
		}

		for _, result := range results {
			m := model.NewMoveFromDb(result)
			movesById[m.ID] = &m
		}

		for i, id := range ids {
			moves[i] = movesById[id]
		}

		return moves, nil
	}
}

func (r Move) MovesByTypeIdDataLoader(ctx context.Context) func(ids []string) ([]*model.MoveConnection, []error) {
	return func(ids []string) ([]*model.MoveConnection, []error) {
		moveConnectionsById := map[string]*model.MoveConnection{}
		results, err := r.client.Move.
			FindMany(db.Move.TypeID.In(ids)).
			Exec(ctx)
		moveConnections := make([]*model.MoveConnection, len(ids))

		if err != nil {
			errors := make([]error, len(ids))
			for i, id := range ids {
				ml := model.NewEmptyMoveConnection()
				moveConnections[i] = &ml
				errors[i] = fmt.Errorf("error loading moves by type id %s in dataloader %w", id, err)
			}

			return moveConnections, errors
		}

		if len(results) == 0 {
			for i := range ids {
				ml := model.NewEmptyMoveConnection()
				moveConnections[i] = &ml
			}

			return moveConnections, nil
		}

		for _, result := range results {
			ml := moveConnectionsById[result.TypeID]
			if ml == nil {
				empty := model.NewEmptyMoveConnection()
				moveConnectionsById[result.TypeID] = &empty
			}
			m := model.NewMoveEdgeFromDb(result)
			moveConnectionsById[result.TypeID].AddEdge(&m)
		}

		for i, id := range ids {
			moveConnection := moveConnectionsById[id]

			if moveConnection == nil {
				empty := model.NewEmptyMoveConnection()
				moveConnection = &empty
			}

			moveConnections[i] = moveConnection
		}

		return moveConnections, nil
	}
}
