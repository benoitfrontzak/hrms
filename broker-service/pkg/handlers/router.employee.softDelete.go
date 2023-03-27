package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// forward the get request from the front-end to employee-service to fetch all employee summary
func SoftDeleteEmployee(w http.ResponseWriter, r *http.Request) {
	// send post request to employee-service and collect the response
	url := employeeService + "api/v1/employee/softDelete"
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
	if err != nil || answer.Error {
		errorJSON(w, err)
		return
	}

	// log to employee-service collection
	l := rpcPayload{
		Collection: "employee",
		Name:       "deleted",
		Data:       fmt.Sprintf("%s", answer.Data),
		CreatedAt:  answer.CreatedAt,
		CreatedBy:  answer.CreatedBy,
	}
	LogItemViaRPC(l)

	// send response to front-end
	writeJSON(w, http.StatusAccepted, answer)
}
