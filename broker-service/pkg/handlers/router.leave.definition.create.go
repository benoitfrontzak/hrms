package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// forward the post request from the front-end to claim-service to create a new claim definition
func CreateLeaveDefinition(w http.ResponseWriter, r *http.Request) {
	// send post request to employee-service and collect the response
	url := leaveService + "api/v1/leave/definition/create"
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
	// decode answer.Data
	// ct, err := extractResponseCT(answer.Data)
	// if err != nil {
	// 	errorJSON(w, err)
	// 	return
	// }

	// log to employee-service collection
	l := rpcPayload{
		Collection: "leave",
		Name:       "create definition",
		// Data:       fmt.Sprintf("new entry successfully created for table %s with id %d", ct.Table, ct.RowID),
		Data:      fmt.Sprintf("new entry successfully created for table with id"),
		CreatedAt: answer.CreatedAt,
		CreatedBy: answer.CreatedBy,
	}
	LogItemViaRPC(l)

	// send response to front-end
	writeJSON(w, http.StatusAccepted, answer)
}
