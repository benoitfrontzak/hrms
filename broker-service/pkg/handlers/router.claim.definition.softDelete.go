package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// forward the post request from the front-end to employee-service to create a new employee
func SoftDeleteClaimDefinition(w http.ResponseWriter, r *http.Request) {
	// send post request to employee-service and collect the response
	url := claimService + "api/v1/claim/definition/softDelete"
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
	// log to employee-service collection
	l := rpcPayload{
		Collection: "claim",
		Name:       "definition delete",
		Data:       fmt.Sprintf("entries %s successfully deleted", answer.Data),
		CreatedAt:  answer.CreatedAt,
		CreatedBy:  answer.CreatedBy,
	}
	LogItemViaRPC(l)

	// send response to front-end
	writeJSON(w, http.StatusAccepted, answer)
}
