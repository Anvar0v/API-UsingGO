package models

import "time"

type Movie struct {
	Id          int32     `json:"id"`
	Name        string    `json:"name"`
	ReleaseDate time.Time `json:"release_date"`
	AuthorId    int32     `json:"author_id"`
}
