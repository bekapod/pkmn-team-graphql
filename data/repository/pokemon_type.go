package repository

import "database/sql"

type PokemonType struct {
	db *sql.DB
}

func NewPokemonType(db *sql.DB) PokemonType {
	return PokemonType{
		db: db,
	}
}
