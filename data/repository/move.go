package repository

import "database/sql"

type Move struct {
	db *sql.DB
}

func NewMove(db *sql.DB) Move {
	return Move{
		db: db,
	}
}
