package repository

import (
	"bekapod/pkmn-team-graphql/data/db"
	"bekapod/pkmn-team-graphql/data/model"
	"bekapod/pkmn-team-graphql/log"
	"context"
	"fmt"
)

type PokemonEvolution struct {
	client *db.PrismaClient
}

func NewPokemonEvolution(client *db.PrismaClient) PokemonEvolution {
	return PokemonEvolution{
		client: client,
	}
}

func (r PokemonEvolution) PokemonEvolutionByFromPokemonIdDataLoader(ctx context.Context) func(ids []string) ([]*model.PokemonEvolutionConnection, []error) {
	return func(ids []string) ([]*model.PokemonEvolutionConnection, []error) {
		pokemonEvolutionConnectionsById := map[string]*model.PokemonEvolutionConnection{}
		results, err := r.client.PokemonEvolution.
			FindMany(db.PokemonEvolution.FromPokemonID.In(ids)).
			Exec(ctx)
		pokemonEvolutionConnections := make([]*model.PokemonEvolutionConnection, len(ids))

		if err != nil {
			errors := make([]error, len(ids))
			for i, id := range ids {
				el := model.NewEmptyPokemonEvolutionConnection()
				pokemonEvolutionConnections[i] = &el
				log.Logger.WithField("id", id).WithError(err).WithContext(ctx).Error("error loading pokemon evolution by FROM pokemon id")
				errors[i] = fmt.Errorf("error loading pokemon evolution by from pokemon id %s", id)
			}

			return pokemonEvolutionConnections, errors
		}

		if len(results) == 0 {
			for i := range ids {
				el := model.NewEmptyPokemonEvolutionConnection()
				pokemonEvolutionConnections[i] = &el
			}

			return pokemonEvolutionConnections, nil
		}

		for _, result := range results {
			el := pokemonEvolutionConnectionsById[result.FromPokemonID]
			if el == nil {
				empty := model.NewEmptyPokemonEvolutionConnection()
				pokemonEvolutionConnectionsById[result.FromPokemonID] = &empty
			}
			e := model.NewPokemonEvolutionEdgeFromDb(result)
			e.Node.PokemonID = e.Node.ToPokemonID
			pokemonEvolutionConnectionsById[result.FromPokemonID].AddEdge(&e)
		}

		for i, id := range ids {
			pokemonEvolutionConnection := pokemonEvolutionConnectionsById[id]

			if pokemonEvolutionConnection == nil {
				empty := model.NewEmptyPokemonEvolutionConnection()
				pokemonEvolutionConnection = &empty
			}

			pokemonEvolutionConnections[i] = pokemonEvolutionConnection
		}

		return pokemonEvolutionConnections, nil
	}
}

func (r PokemonEvolution) PokemonEvolutionByToPokemonIdDataLoader(ctx context.Context) func(ids []string) ([]*model.PokemonEvolutionConnection, []error) {
	return func(ids []string) ([]*model.PokemonEvolutionConnection, []error) {
		pokemonEvolutionConnectionsById := map[string]*model.PokemonEvolutionConnection{}
		results, err := r.client.PokemonEvolution.
			FindMany(db.PokemonEvolution.ToPokemonID.In(ids)).
			Exec(ctx)
		pokemonEvolutionConnections := make([]*model.PokemonEvolutionConnection, len(ids))

		if err != nil {
			errors := make([]error, len(ids))
			for i, id := range ids {
				el := model.NewEmptyPokemonEvolutionConnection()
				pokemonEvolutionConnections[i] = &el
				log.Logger.WithField("id", id).WithError(err).WithContext(ctx).Error("error loading pokemon evolution by TO pokemon id")
				errors[i] = fmt.Errorf("error loading pokemon evolution by to pokemon id %s", id)
			}

			return pokemonEvolutionConnections, errors
		}

		if len(results) == 0 {
			for i := range ids {
				el := model.NewEmptyPokemonEvolutionConnection()
				pokemonEvolutionConnections[i] = &el
			}

			return pokemonEvolutionConnections, nil
		}

		for _, result := range results {
			el := pokemonEvolutionConnectionsById[result.ToPokemonID]
			if el == nil {
				empty := model.NewEmptyPokemonEvolutionConnection()
				pokemonEvolutionConnectionsById[result.ToPokemonID] = &empty
			}
			e := model.NewPokemonEvolutionEdgeFromDb(result)
			e.Node.PokemonID = e.Node.FromPokemonID
			pokemonEvolutionConnectionsById[result.ToPokemonID].AddEdge(&e)
		}

		for i, id := range ids {
			pokemonEvolutionConnection := pokemonEvolutionConnectionsById[id]

			if pokemonEvolutionConnection == nil {
				empty := model.NewEmptyPokemonEvolutionConnection()
				pokemonEvolutionConnection = &empty
			}

			pokemonEvolutionConnections[i] = pokemonEvolutionConnection
		}

		return pokemonEvolutionConnections, nil
	}
}
