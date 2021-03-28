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
		"SELECT id, name, slug, pokedex_id, sprite, hp, attack, defense, special_attack, special_defense, speed, is_baby, is_legendary, is_mythical, description FROM pokemon ORDER BY pokedex_id, slug ASC",
	)
	if err != nil {
		return &pokemon, fmt.Errorf("error fetching all pokemon: %w", err)
	}

	defer rows.Close()
	for rows.Next() {
		var pkmn model.Pokemon
		err := rows.Scan(&pkmn.ID, &pkmn.Name, &pkmn.Slug, &pkmn.PokedexId, &pkmn.Sprite, &pkmn.HP, &pkmn.Attack, &pkmn.Defense, &pkmn.SpecialAttack, &pkmn.SpecialDefense, &pkmn.Speed, &pkmn.IsBaby, &pkmn.IsLegendary, &pkmn.IsMythical, &pkmn.Description)
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
