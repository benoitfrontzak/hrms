package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// forward the post request from the front-end to authentication-service to update the user's password
func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	// send post request to employee-service and collect the response
	url := authService + "api/v1/authentication/update/password"
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
	// decode answer.Data
	// ct, err := extractResponseCT(answer.Data)
	// if err != nil {
	// 	errorJSON(w, err)
	// 	return
	// }

	// log to authentication service log on changing password
	l := rpcPayload{
		Collection: "authentication",
		Name:       "change password",
		Data:       fmt.Sprintf("new password successfully updated for for user...TODO"),
		CreatedAt:  answer.CreatedAt,
		CreatedBy:  answer.CreatedBy,
	}
	LogItemViaRPC(l)

	// send response to front-end
	writeJSON(w, http.StatusAccepted, answer)
}
