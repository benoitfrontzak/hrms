package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// forward the post request from the front-end to employee-service to create a new employee
func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	// send post request to employee-service and collect the response
	url := employeeService + "api/v1/employee/create"
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
	// new employee email
	nee := answer.Data.(map[string]interface{})["Email"]

	// log to employee-service collection
	l := rpcPayload{
		Collection: "employee",
		Name:       "created",
		Data:       fmt.Sprintf("employee %s successfully created", nee),
		CreatedAt:  answer.CreatedAt,
		CreatedBy:  answer.CreatedBy,
	}
	LogItemViaRPC(l)

	// send response to front-end
	writeJSON(w, http.StatusAccepted, answer)
}
