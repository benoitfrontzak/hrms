package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// forward the post request from the front-end to employee-service to create a new employee
func UpdateEmployeeCT(w http.ResponseWriter, r *http.Request) {
	// send post request to employee-service and collect the response
	url := employeeService + "api/v1/employee/configTable/update"
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
	ct, err := extractResponseCT(answer.Data)
	if err != nil {
		errorJSON(w, err)
		return
	}

	// log to employee-service collection
	l := rpcPayload{
		Collection: "employee",
		Name:       "config table",
		Data:       fmt.Sprintf("new entry successfully created for table %s with id %d", ct.Table, ct.RowID),
		CreatedAt:  answer.CreatedAt,
		CreatedBy:  answer.CreatedBy,
	}
	LogItemViaRPC(l)

	// send response to front-end
	writeJSON(w, http.StatusAccepted, answer)
}
