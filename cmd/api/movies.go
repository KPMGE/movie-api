package main

import (
	"fmt"
	"net/http"
	"time"

	"movie.api.kpmge/internal/data"
	"movie.api.kpmge/internal/validator"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string       `json:"title"`
		Year    int32        `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()

	v.Check(input.Title != "", "title", "must be provided")
	v.Check(len(input.Title) < 500, "title", "must be less than 500 bytes")
	v.Check(input.Year != 0, "year", "year must be provided")
	v.Check(input.Year <= int32(time.Now().Year()), "year", "year must not be in the future")
	v.Check(input.Runtime != 0, "runtime", "runtime must be provided")
	v.Check(input.Runtime > 0, "runtime", "runtime must be a positive integer")
	v.Check(input.Genres != nil, "genres", "must be provided")
	v.Check(len(input.Genres) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(input.Genres) <= 10, "genres", "must not contain more than 10 genres")
	v.Check(validator.Unique(input.Genres), "genres", "must not contain duplicate genres")

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	fmt.Fprintf(w, "%+v", input)
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "the party is over",
		Year:      2003,
		Runtime:   200,
		Genres:    []string{"comedy", "drama", "fiction", "romance"},
		Version:   2302,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
