package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// forward the get request from the front-end to employee-service to fetch all employee summary
func AllEmployees(w http.ResponseWriter, r *http.Request) {
	// send get request to employee-service and collect the response
	url := employeeService + "api/v1/employee/get/all"
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
		log.Println("err with broker is", err)
		errorJSON(w, err)
		return
	}

	// send response to front-end
	writeJSON(w, http.StatusAccepted, answer)
}
