package repository

import (
	"bekapod/pkmn-team-graphql/data/db"
	"bekapod/pkmn-team-graphql/data/model"
	"bekapod/pkmn-team-graphql/log"
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
		log.Logger.WithError(err).WithContext(ctx).Error("error getting moves")
		return &moves, fmt.Errorf("error getting moves")
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
		log.Logger.WithField("id", id).WithContext(ctx).Info("couldn't find move by id")
		return nil, fmt.Errorf("couldn't find move by id: %s", id)
	}

	if err != nil {
		log.Logger.WithField("id", id).WithError(err).WithContext(ctx).Error("error getting move by id")
		return nil, fmt.Errorf("error getting move by id: %s", id)
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
				log.Logger.WithField("id", id).WithError(err).WithContext(ctx).Error("error loading move by id")
				errors[i] = fmt.Errorf("error loading move by id %s", id)
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
				log.Logger.WithField("id", id).WithError(err).WithContext(ctx).Error("error loading moves by type id")
				errors[i] = fmt.Errorf("error loading moves by type id %s", id)
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
