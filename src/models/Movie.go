package models

import "time"

type Movie struct {
	Id          int32     `json:"id"`
	Name        string    `json:"name"`
	ReleaseDate time.Time `json:"release_date"`
	AuthorName  string    `json:"author_name"`
}
