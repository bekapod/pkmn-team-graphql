package data

import (
	"bekapod/pkmn-team-graphql/log"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

const retryAttempts = 5

func NewDB(dsn string) *sql.DB {
	db, err := sql.Open("postgres", dsn)
	log.Logger.Info("connecting to database")

	if err != nil {
		log.Logger.Fatalf("unable to connect to the database: %s", err)
	}

	tries := retryAttempts
	for {
		err := db.Ping()
		if err == nil {
			break
		}

		time.Sleep(time.Second * 1)
		tries--
		log.Logger.Infof("database is not available (err: %s), retrying %d more time(s)", err, tries)

		if tries == 0 {
			log.Logger.Fatalf("database %s did not become available within %d connection attempts", dsn, retryAttempts)
		}
	}

	return db
}
