package db

import (
	"log"

	_ "github.com/lib/pq"
	_ "github.com/mattes/migrate/source/file"
)

type Place struct {
	ID      int32
	Name    string
	Zipcode string
	Country string
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
