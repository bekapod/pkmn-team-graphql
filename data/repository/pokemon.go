package repository

import "database/sql"

type Pokemon struct {
	db *sql.DB
}

func NewPokemon(db *sql.DB) Pokemon {
	return Pokemon{
		db: db,
	}
}
