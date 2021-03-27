package repository

import (
	"bekapod/pkmn-team-graphql/data/model"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

var (
	ErrNoPokemon = errors.New("no pokemon found")
)

func (r Pokemon) GetAllPokemon(ctx context.Context) (*model.PokemonList, error) {
	pokemon := model.NewEmptyPokemonList()

	rows, err := r.db.QueryContext(
		ctx,
		"SELECT id FROM pokemon",
	)
	if err != nil {
		return &pokemon, fmt.Errorf("error fetching all pokemon: %w", err)
	}

	defer rows.Close()
	for rows.Next() {
		var pkmn model.Pokemon
		err := rows.Scan(&pkmn.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				return &pokemon, ErrNoPokemon
			}
			return &pokemon, fmt.Errorf("error scanning result in GetAllPokemon: %w", err)
		}
		pokemon.AddPokemon(&pkmn)
	}

	return &pokemon, nil
}
