package services

import (
	"database/sql"
	"log"
	"movies_api/config"
	"movies_api/models"
)

var db sql.DB

func init() {
	config.Connect()
	db = *config.GetDbObject()
}

func GetMovies() []models.Movie {
	var movies []models.Movie

	row, err := db.Query("SELECT * FROM movies")

	if err != nil {
		log.Fatal("There is problem with retrieving data from DB", row)
	}

	return movies
}
