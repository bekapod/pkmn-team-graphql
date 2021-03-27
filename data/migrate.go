package data

import (
	"bekapod/pkmn-team-graphql/config"
	"bekapod/pkmn-team-graphql/log"
	"database/sql"
	"embed"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/johejo/golang-migrate-extra/source/iofs"
)

//go:embed migrations/*.sql
var migrationFiles embed.FS

func MigrateDatabase(db *sql.DB, cfg config.Config) {
	log.Logger.Info("checking database migrations")

	databaseDriver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Logger.Fatalf("unable to create migration instance from database: %s", err)
	}

	d, err := iofs.New(migrationFiles, "migrations")
	if err != nil {
		log.Logger.Fatalf("unable to load migration files from embedded filesystem: %s", err)
	}

	am, err := migrate.NewWithInstance("iofs", d, cfg.DatabaseName, databaseDriver)
	if err != nil {
		log.Logger.Fatalf("failed to load migration files from source driver: %s", err)
	}

	if err := am.Up(); err != nil && err != migrate.ErrNoChange {
		log.Logger.Fatalf("failed to migrate database: %s", err)
	}

	log.Logger.Info("database is up-to-date, all migrations applied")
}
