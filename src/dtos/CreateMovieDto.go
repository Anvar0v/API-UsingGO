package dtos

import "time"

type CreateMovieDto struct {
	Name        string    `json:"name"`
	ReleaseDate time.Time `json:"release_date"`
	AuthorName      string    `json:"author_name"`
}
