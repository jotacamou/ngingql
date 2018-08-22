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

func AllPlaces(wrapper func(place *Place)) {
	res, err := db.Query("SELECT id, name, zipcode, country FROM places")
	if err != nil {
		log.Println("error fetching all places:", err)
	}
	//defer db.Close()

	for res.Next() {
		place := Place{}
		err := res.Scan(&place.ID, &place.Name, &place.Zipcode, &place.Country)
		if err != nil {
			log.Fatal(err)
		}
		wrapper(&place)
	}
}

type Place struct {
	ID      string
	Name    string
	Zipcode string
	Country string
}

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
}

func FindUserByID(id string) *User {
	user := &User{}
	err := db.QueryRow("SELECT id, name, email, password FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		log.Println("error fetching user by id", err)
	}
	return user
}
