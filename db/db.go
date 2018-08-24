package db

import (
	"database/sql"

	"log"

	_ "github.com/lib/pq"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
)

const dbURL = "postgres://postgres@localhost:5432/tour1?sslmode=disable"

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"tour1", driver)
	if err != nil {
		log.Println("Error occurred during migration", err)
	}
	m.Steps(4)
}
