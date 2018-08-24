package db

import (
	"log"

	_ "github.com/lib/pq"
	_ "github.com/mattes/migrate/source/file"
)

type Activity struct {
	ID               int32
	Title            string
	Description      string
	Owner            string
	Price            float32
	Fee              float32
	MainPicture      []string
	SlideshowPicture []string
	Duration         string
	PlaceID          int32
	Hour             []string
	Tag              []string
}

func AllActivities(wrapper func(activity *Activity)) {
	res, err := db.Query("SELECT id ,title ,description ,owner ,price ,fee ,main_picture ,slideshow_picture ,duration ,place_id ,hour ,tag FROM activities")
	if err != nil {
		log.Println("error fetching all activities:", err)
	}
	//defer db.Close()

	for res.Next() {
		activity := Activity{}
		err := res.Scan(&activity.ID, &activity.Title, &activity.Description, &activity.Owner, &activity.Price, &activity.Fee, &activity.MainPicture, &activity.SlideshowPicture, &activity.Duration, &activity.PlaceID, &activity.Hour, &activity.Tag)
		if err != nil {
			log.Fatal(err)
		}
		wrapper(&activity)
	}
}
