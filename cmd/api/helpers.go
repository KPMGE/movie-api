package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// envelope type for wrapping json response
type envelope map[string]any

func (app *application) readIDParam(r *http.Request) (int64, error) {
	// return slice with parameter names and values
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)

	if err != nil || id < 1 {
		return 0, errors.New("Invalid id parameter!")
	}

	return id, nil
}

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	// decode json body into the target destination
	err := json.NewDecoder(r.Body).Decode(dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		switch {
		// if the error has a syntaxError
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly formed JSON (at character %d)", syntaxError.Offset)

		// if the error is a ErrUnexpectedEOF
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly formed JSON")

		// if the error has a unmarshalTypeError
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)

		// if the body is empty
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")

		case errors.As(err, &invalidUnmarshalError):
			// it's fine to panic here, cuz if that error occured, probably it's an error from the developer
			// so, it makes sense to catch it and stop the application as soon as possible
			panic(err)

		default:
			return err
		}
	}
	return nil
}

// converts entry in json and write it to the response
func (app *application) writeJSON(w http.ResponseWriter, status int, data any, headers http.Header) error {
	// this works exactly as the Marshal function, but it indents the json,
	// which is good when using curl.
	js, err := json.MarshalIndent(data, "", "  ")

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
