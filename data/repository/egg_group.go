package repository

import "database/sql"

type EggGroup struct {
	db *sql.DB
}

func NewEggGroup(db *sql.DB) EggGroup {
	return EggGroup{
		db: db,
	}
}
