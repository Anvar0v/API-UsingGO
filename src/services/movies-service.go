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

	row,err := db.Query("SELECT * FROM movies LEFT JOIN authors where movies.id = authors.Id")

	if err != nil{
		log.Fatal("There is problem in retrieving data from DB", row)
	}

	// for row.Next() {
			
	// }

	// movies = append(movies, newMovie)

	return movies
}
