package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
)

// forward the post request from the front-end to claim-service to update a claim definition
func UpdateClaimDefinition(w http.ResponseWriter, r *http.Request) {
	// send post request to employee-service and collect the response
	url := claimService + "api/v1/claim/definition/update"
	resp, err := http.Post(url, "application/json", r.Body)
	if err != nil {
		errorJSON(w, err)
		return
	}
	defer resp.Body.Close()

	// decode employee-service response
	var answer jsonResponse
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&answer)
	if err != nil {
		errorJSON(w, err)
		return
	}
	if answer.Error {
		errorJSON(w, errors.New(answer.Message))
		return
	}

	// send response to front-end
	writeJSON(w, http.StatusAccepted, answer)
}
