package repository

import (
	"bekapod/pkmn-team-graphql/data/db"
	"bekapod/pkmn-team-graphql/data/model"
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

func (r PokemonMove) PokemonMoveByIdDataLoader(ctx context.Context) func(ids []string) ([]*model.PokemonMove, []error) {
	return func(ids []string) ([]*model.PokemonMove, []error) {
		pokemonMovesById := map[string]*model.PokemonMove{}
		results, err := r.client.PokemonMove.FindMany(db.PokemonMove.ID.In(ids)).Exec(ctx)
		pokemonMoves := make([]*model.PokemonMove, len(ids))

		if err != nil {
			errors := make([]error, len(ids))
			for i, id := range ids {
				errors[i] = fmt.Errorf("error loading pokemon move by id %s in dataloader %w", id, err)
			}

			return pokemonMoves, errors
		}

		for _, result := range results {
			pm := model.NewPokemonMoveFromDb(result)
			pokemonMovesById[result.ID] = &pm
		}

		for i, id := range ids {
			pokemonMoves[i] = pokemonMovesById[id]
		}

		return pokemonMoves, nil
	}
}

func (r PokemonMove) PokemonMoveByPokemonIdDataLoader(ctx context.Context) func(ids []string) ([]*model.PokemonMoveList, []error) {
	return func(ids []string) ([]*model.PokemonMoveList, []error) {
		pokemonMoveListsById := map[string]*model.PokemonMoveList{}
		results, err := r.client.PokemonMove.
			FindMany(db.PokemonMove.PokemonID.In(ids)).
			Exec(ctx)
		pokemonMoveLists := make([]*model.PokemonMoveList, len(ids))

		if err != nil {
			errors := make([]error, len(ids))
			for i, id := range ids {
				ml := model.NewEmptyPokemonMoveList()
				pokemonMoveLists[i] = &ml
				errors[i] = fmt.Errorf("error loading pokemon move by pokemon id %s in dataloader %w", id, err)
			}

			return pokemonMoveLists, errors
		}

		if len(results) == 0 {
			for i := range ids {
				ml := model.NewEmptyPokemonMoveList()
				pokemonMoveLists[i] = &ml
			}

			return pokemonMoveLists, nil
		}

		for _, result := range results {
			ml := pokemonMoveListsById[result.PokemonID]
			if ml == nil {
				empty := model.NewEmptyPokemonMoveList()
				pokemonMoveListsById[result.PokemonID] = &empty
			}
			m := model.NewPokemonMoveFromDb(result)
			pokemonMoveListsById[result.PokemonID].AddPokemonMove(&m)
		}

		for i, id := range ids {
			pokemonMoveList := pokemonMoveListsById[id]

			if pokemonMoveList == nil {
				empty := model.NewEmptyPokemonMoveList()
				pokemonMoveList = &empty
			}

			pokemonMoveLists[i] = pokemonMoveList
		}

		return pokemonMoveLists, nil
	}
}

func (r PokemonMove) PokemonMoveByMoveIdDataLoader(ctx context.Context) func(ids []string) ([]*model.PokemonMoveList, []error) {
	return func(ids []string) ([]*model.PokemonMoveList, []error) {
		pokemonMoveListsById := map[string]*model.PokemonMoveList{}
		results, err := r.client.PokemonMove.
			FindMany(db.PokemonMove.MoveID.In(ids)).
			Exec(ctx)
		pokemonMoveLists := make([]*model.PokemonMoveList, len(ids))

		if err != nil {
			errors := make([]error, len(ids))
			for i, id := range ids {
				ml := model.NewEmptyPokemonMoveList()
				pokemonMoveLists[i] = &ml
				errors[i] = fmt.Errorf("error loading pokemon move by move id %s in dataloader %w", id, err)
			}

			return pokemonMoveLists, errors
		}

		if len(results) == 0 {
			for i := range ids {
				ml := model.NewEmptyPokemonMoveList()
				pokemonMoveLists[i] = &ml
			}

			return pokemonMoveLists, nil
		}

		for _, result := range results {
			ml := pokemonMoveListsById[result.MoveID]
			if ml == nil {
				empty := model.NewEmptyPokemonMoveList()
				pokemonMoveListsById[result.MoveID] = &empty
			}
			m := model.NewPokemonMoveFromDb(result)
			pokemonMoveListsById[result.MoveID].AddPokemonMove(&m)
		}

		for i, id := range ids {
			pokemonMoveList := pokemonMoveListsById[id]

			if pokemonMoveList == nil {
				empty := model.NewEmptyPokemonMoveList()
				pokemonMoveList = &empty
			}

			pokemonMoveLists[i] = pokemonMoveList
		}

		return pokemonMoveLists, nil
	}
}
