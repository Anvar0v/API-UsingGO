package dtos

import "time"

type CreateMovieDto struct{
	Name        string    `json:"name"`
	ReleaseDate time.Time `json:"release_date"`
	AuthorId    int32     `json:"author_id"`
}