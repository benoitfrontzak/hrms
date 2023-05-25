package handlers

import (
	"encoding/json"
	"net/http"
)

// forward the get request from the front-end to leave-service to fetch all leaves
func AllLeaveUser(w http.ResponseWriter, r *http.Request) {

	// send get request to leave-service and collect the response
	url := leaveService + "api/v1/leave/get/user"
	resp, err := http.Post(url, "application/json", r.Body)
	if err != nil {
		errorJSON(w, err)
		return
	}
	defer resp.Body.Close()

	// decode leave-service response
	var answer jsonResponse
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&answer)
	if err != nil {
		errorJSON(w, err)
		return
	}

	// send response to front-end
	writeJSON(w, http.StatusAccepted, answer)
}
