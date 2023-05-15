package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// forward the post request from the front-end to leave-service to create a new leave
func LeaveConfigTableSoftDelete(w http.ResponseWriter, r *http.Request) {
	// send post request to leave-service and collect the response
	url := leaveService + "api/v1/leave/configTable/softDelete"
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
	if answer.Error {
		errorJSON(w, errors.New(answer.Message))
		return
	}
	// decode answer.Data
	ct, err := extractResponseDeleteCT(answer.Data)
	if err != nil {
		errorJSON(w, err)
		return
	}

	// log to employee-service collection
	l := rpcPayload{
		Collection: "leave",
		Name:       "config table",
		Data:       fmt.Sprintf("entries %s successfully deleted from table %s", ct.List, ct.Table),
		CreatedAt:  answer.CreatedAt,
		CreatedBy:  answer.CreatedBy,
	}
	LogItemViaRPC(l)

	// send response to front-end
	writeJSON(w, http.StatusAccepted, answer)
}
