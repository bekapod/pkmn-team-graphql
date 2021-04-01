package repository

import "database/sql"

type Ability struct {
	db *sql.DB
}

func NewAbility(db *sql.DB) Ability {
	return Ability{
		db: db,
	}
}
