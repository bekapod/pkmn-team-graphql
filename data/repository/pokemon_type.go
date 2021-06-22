package repository

import (
	"bekapod/pkmn-team-graphql/data/db"
	"bekapod/pkmn-team-graphql/data/model"
	"context"
	"fmt"
)

type PokemonType struct {
	client *db.PrismaClient
}

func NewPokemonType(client *db.PrismaClient) PokemonType {
	return PokemonType{
		client: client,
	}
}

func (r PokemonType) PokemonTypeByPokemonIdDataLoader(ctx context.Context) func(ids []string) ([]*model.PokemonTypeConnection, []error) {
	return func(ids []string) ([]*model.PokemonTypeConnection, []error) {
		pokemonTypeConnectionsById := map[string]*model.PokemonTypeConnection{}
		results, err := r.client.PokemonType.
			FindMany(db.PokemonType.PokemonID.In(ids)).
			Exec(ctx)
		pokemonTypeConnections := make([]*model.PokemonTypeConnection, len(ids))

		if err != nil {
			errors := make([]error, len(ids))
			for i, id := range ids {
				tl := model.NewEmptyPokemonTypeConnection()
				pokemonTypeConnections[i] = &tl
				errors[i] = fmt.Errorf("error loading pokemon type by pokemon id %s in dataloader %w", id, err)
			}

			return pokemonTypeConnections, errors
		}

		if len(results) == 0 {
			for i := range ids {
				tl := model.NewEmptyPokemonTypeConnection()
				pokemonTypeConnections[i] = &tl
			}

			return pokemonTypeConnections, nil
		}

		for _, result := range results {
			tl := pokemonTypeConnectionsById[result.PokemonID]
			if tl == nil {
				empty := model.NewEmptyPokemonTypeConnection()
				pokemonTypeConnectionsById[result.PokemonID] = &empty
			}
			t := model.NewPokemonTypeFromDb(result)
			pokemonTypeConnectionsById[result.PokemonID].AddEdge(&t)
		}

		for i, id := range ids {
			pokemonTypeConnection := pokemonTypeConnectionsById[id]

			if pokemonTypeConnection == nil {
				empty := model.NewEmptyPokemonTypeConnection()
				pokemonTypeConnection = &empty
			}

			pokemonTypeConnections[i] = pokemonTypeConnection
		}

		return pokemonTypeConnections, nil
	}
}

func (r PokemonType) PokemonTypeByTypeIdDataLoader(ctx context.Context) func(ids []string) ([]*model.PokemonWithTypeConnection, []error) {
	return func(ids []string) ([]*model.PokemonWithTypeConnection, []error) {
		pokemonTypeConnectionsById := map[string]*model.PokemonWithTypeConnection{}
		results, err := r.client.PokemonType.
			FindMany(db.PokemonType.TypeID.In(ids)).
			Exec(ctx)
		pokemonTypeConnections := make([]*model.PokemonWithTypeConnection, len(ids))

		if err != nil {
			errors := make([]error, len(ids))
			for i, id := range ids {
				tl := model.NewEmptyPokemonWithTypeConnection()
				pokemonTypeConnections[i] = &tl
				errors[i] = fmt.Errorf("error loading pokemon type by type id %s in dataloader %w", id, err)
			}

			return pokemonTypeConnections, errors
		}

		if len(results) == 0 {
			for i := range ids {
				tl := model.NewEmptyPokemonWithTypeConnection()
				pokemonTypeConnections[i] = &tl
			}

			return pokemonTypeConnections, nil
		}

		for _, result := range results {
			tl := pokemonTypeConnectionsById[result.TypeID]
			if tl == nil {
				empty := model.NewEmptyPokemonWithTypeConnection()
				pokemonTypeConnectionsById[result.TypeID] = &empty
			}
			t := model.NewPokemonWithTypeFromDb(result)
			pokemonTypeConnectionsById[result.TypeID].AddEdge(&t)
		}

		for i, id := range ids {
			pokemonTypeConnection := pokemonTypeConnectionsById[id]

			if pokemonTypeConnection == nil {
				empty := model.NewEmptyPokemonWithTypeConnection()
				pokemonTypeConnection = &empty
			}

			pokemonTypeConnections[i] = pokemonTypeConnection
		}

		return pokemonTypeConnections, nil
	}
}
