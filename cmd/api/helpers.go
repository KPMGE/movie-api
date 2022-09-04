package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) readIDParam(r *http.Request) (int64, error) {
	// return slice with parameter names and values
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)

	if err != nil || id < 1 {
		return 0, errors.New("Invalid id parameter!")
	}

	return id, nil
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data any, headers http.Header) error {
	js, err := json.Marshal(data)

	if err != nil {
		return err
	}

	// append new line, so that it's easier to see on terminal applications
	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	// set the content type header to json
	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
