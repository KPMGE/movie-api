package data

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
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

type MovieModel struct {
	DB *sql.DB
}

func (m MovieModel) Insert(movie *Movie) error {
	query := `
    INSERT INTO movies (title, year, runtime, genres) 
    VALUES($1, $2, $3, $4)
    RETURNING id, created_at, version`

	// slice containing the values for the placeholder movie struct.
	// NOTE: when working with slices, we must use pq.Array
	args := []any{movie.Title, movie.Year, movie.Runtime, pq.Array(movie.Genres)}

	// executing query against database and copying id, created_at and version to the movie struct
	return m.DB.QueryRow(query, args...).Scan(&movie.ID, &movie.CreatedAt, &movie.Version)
}

func (m MovieModel) Get(id int64) (*Movie, error) {
	return nil, nil
}

func (m MovieModel) Update(movie *Movie) error {
	return nil
}

func (m MovieModel) Delete(id int64) error {
	return nil
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
