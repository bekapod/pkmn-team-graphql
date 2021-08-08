package repository

import (
	"bekapod/pkmn-team-graphql/data/db"
	"bekapod/pkmn-team-graphql/data/model"
	"bekapod/pkmn-team-graphql/log"
	"context"
	"fmt"
)

type PokemonMove struct {
	client *db.PrismaClient
}

func NewPokemonMove(client *db.PrismaClient) PokemonMove {
	return PokemonMove{
		client: client,
	}
}

func (r PokemonMove) PokemonMoveByPokemonIdDataLoader(ctx context.Context) func(ids []string) ([]*model.PokemonMoveConnection, []error) {
	return func(ids []string) ([]*model.PokemonMoveConnection, []error) {
		pokemonMoveConnectionsById := map[string]*model.PokemonMoveConnection{}
		results, err := r.client.PokemonMove.
			FindMany(db.PokemonMove.PokemonID.In(ids)).
			Exec(ctx)
		pokemonMoveConnections := make([]*model.PokemonMoveConnection, len(ids))

		if err != nil {
			errors := make([]error, len(ids))
			for i, id := range ids {
				ml := model.NewEmptyPokemonMoveConnection()
				pokemonMoveConnections[i] = &ml
				log.Logger.WithField("id", id).WithError(err).WithContext(ctx).Error("error loading pokemon move by pokemon id")
				errors[i] = fmt.Errorf("error loading pokemon move by pokemon id %s", id)
			}

			return pokemonMoveConnections, errors
		}

		if len(results) == 0 {
			for i := range ids {
				ml := model.NewEmptyPokemonMoveConnection()
				pokemonMoveConnections[i] = &ml
			}

			return pokemonMoveConnections, nil
		}

		for _, result := range results {
			ml := pokemonMoveConnectionsById[result.PokemonID]
			if ml == nil {
				empty := model.NewEmptyPokemonMoveConnection()
				pokemonMoveConnectionsById[result.PokemonID] = &empty
			}
			m := model.NewPokemonMoveEdgeFromDb(result)
			pokemonMoveConnectionsById[result.PokemonID].AddEdge(&m)
		}

		for i, id := range ids {
			pokemonMoveConnection := pokemonMoveConnectionsById[id]

			if pokemonMoveConnection == nil {
				empty := model.NewEmptyPokemonMoveConnection()
				pokemonMoveConnection = &empty
			}

			pokemonMoveConnections[i] = pokemonMoveConnection
		}

		return pokemonMoveConnections, nil
	}
}

func (r PokemonMove) PokemonMoveByMoveIdDataLoader(ctx context.Context) func(ids []string) ([]*model.PokemonWithMoveConnection, []error) {
	return func(ids []string) ([]*model.PokemonWithMoveConnection, []error) {
		pokemonMoveConnectionsById := map[string]*model.PokemonWithMoveConnection{}
		results, err := r.client.PokemonMove.
			FindMany(db.PokemonMove.MoveID.In(ids)).
			Exec(ctx)
		pokemonMoveConnections := make([]*model.PokemonWithMoveConnection, len(ids))

		if err != nil {
			errors := make([]error, len(ids))
			for i, id := range ids {
				ml := model.NewEmptyPokemonWithMoveConnection()
				pokemonMoveConnections[i] = &ml
				log.Logger.WithField("id", id).WithError(err).WithContext(ctx).Error("error loading pokemon move by move id")
				errors[i] = fmt.Errorf("error loading pokemon move by move id %s", id)
			}

			return pokemonMoveConnections, errors
		}

		if len(results) == 0 {
			for i := range ids {
				ml := model.NewEmptyPokemonWithMoveConnection()
				pokemonMoveConnections[i] = &ml
			}

			return pokemonMoveConnections, nil
		}

		for _, result := range results {
			ml := pokemonMoveConnectionsById[result.MoveID]
			if ml == nil {
				empty := model.NewEmptyPokemonWithMoveConnection()
				pokemonMoveConnectionsById[result.MoveID] = &empty
			}
			m := model.NewPokemonWithMoveEdgeFromDb(result)
			pokemonMoveConnectionsById[result.MoveID].AddEdge(&m)
		}

		for i, id := range ids {
			pokemonMoveConnection := pokemonMoveConnectionsById[id]

			if pokemonMoveConnection == nil {
				empty := model.NewEmptyPokemonWithMoveConnection()
				pokemonMoveConnection = &empty
			}

			pokemonMoveConnections[i] = pokemonMoveConnection
		}

		return pokemonMoveConnections, nil
	}
}
