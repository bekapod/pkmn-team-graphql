package repository

import "database/sql"

type Type struct {
	db *sql.DB
}

func NewType(db *sql.DB) Type {
	return Type{
		db: db,
	}
}
