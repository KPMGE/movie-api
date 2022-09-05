package main

import (
	"fmt"
	"net/http"
	"time"

	"movie.api.kpmge/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	// write messa into the response
	fmt.Fprintln(w, "create a new movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	movie := data.Movie{
		ID:        100,
		CreatedAt: time.Now(),
		Title:     "the party is over",
		Year:      2003,
		Runtime:   200,
		Genres:    []string{"comedy", "drama", "fiction", "romance"},
		Version:   2302,
	}

	app.writeJSON(w, http.StatusOK, movie, nil)

	// write message into the response
	fmt.Fprintf(w, "show details of the movie: %d", id)
}
