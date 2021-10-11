package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/oglinuk/restful-go/internal/pkg/config"
)

// Open opens an sql database based on the configuration and execute the
// given schema. If a seed file exists, call seed.
func Open(schema string) *sql.DB {
	cfg := config.Get()

	db, err := sql.Open(cfg.Database.Driver, cfg.Database.File)
	if err != nil {
		log.Fatalf("database::Open::sql.Open: %s\n", err.Error())
	}

	_, err = db.Exec(schema)
	if err != nil {
		log.Fatalf("database::Open::db.Exec: %s\n", err.Error())
	}

	return db
}
