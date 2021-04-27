package repository

import "database/sql"

type Evolution struct {
	db *sql.DB
}

func NewEvolution(db *sql.DB) Evolution {
	return Evolution{
		db: db,
	}
}
