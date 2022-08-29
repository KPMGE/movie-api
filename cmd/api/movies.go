package main

import (
	"fmt"
	"net/http"
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

	// write message into the response
	fmt.Fprintf(w, "show details of the movie: %d", id)
}
