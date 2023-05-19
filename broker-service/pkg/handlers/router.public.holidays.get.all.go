package handlers

import (
	"encoding/json"
	"net/http"
)

// forward the get request from the front-end to leave-service to fetch all public holidays
func AllPublicHolidays(w http.ResponseWriter, r *http.Request) {

	// send get request to employee-service and collect the response
	url := leaveService + "api/v1/publicHoliday/get/all"

	resp, err := http.Get(url)
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

	// send response to front-end
	writeJSON(w, http.StatusAccepted, answer)
}
