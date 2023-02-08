package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// readJSON tries to read the body of a request and converts it into JSON
func (rep *Repository) readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1048576 // one megabyte

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must have only a single JSON value")
	}

	return nil
}

// writeJSON takes a response status code and arbitrary data and writes a json response to the client
func (rep *Repository) writeJSON(w http.ResponseWriter, status int, data any, token string) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Set header
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Authorization", "Bearer: "+token)
	// Write header
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

// errorJSON takes an error, and optionally a response status code, and generates and sends
// a json error response
func (rep *Repository) errorJSON(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload jsonResponse
	payload.Error = true
	payload.Message = err.Error()

	return rep.writeJSON(w, statusCode, payload, "")
}
