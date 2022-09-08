package data

import (
	"time"

	"movie.api.kpmge/internal/validator"
)

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Runtime   Runtime   `json:"runtime"`
	Genres    []string  `json:"genres"`
	Version   int32     `json:"version"`
}

func ValidateMovie(v *validator.Validator, m *Movie) {
	v.Check(m.Title != "", "title", "must be provided")
	v.Check(len(m.Title) < 500, "title", "must be less than 500 bytes")
	v.Check(m.Year != 0, "year", "year must be provided")
	v.Check(m.Year <= int32(time.Now().Year()), "year", "year must not be in the future")
	v.Check(m.Runtime != 0, "runtime", "runtime must be provided")
	v.Check(m.Runtime > 0, "runtime", "runtime must be a positive integer")
	v.Check(m.Genres != nil, "genres", "must be provided")
	v.Check(len(m.Genres) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(m.Genres) <= 10, "genres", "must not contain more than 10 genres")
	v.Check(validator.Unique(m.Genres), "genres", "must not contain duplicate genres")
}
