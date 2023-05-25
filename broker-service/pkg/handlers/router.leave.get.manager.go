package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// forward the get request from the front-end to leave-service to fetch all leaves
func AllLeaveManager(w http.ResponseWriter, r *http.Request) {
	// extract payload from request (manager uid)
	var p User

	err := readJSON(w, r, &p)
	if err != nil {
		log.Println("json err", err)
		errorJSON(w, err)
		return
	}

	// Convert data to JSON
	jsonData, err := json.Marshal(p)
	if err != nil {
		errorJSON(w, err)
		return
	}

	// get employees managed by manager uid
	jsonList, err := getManagerEmployees(jsonData)
	if err != nil {
		errorJSON(w, err)
		return
	}

	// // send post request to leave-service and collect the response
	url := leaveService + "api/v1/leave/get/manager"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonList))
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

// fetch manage employees of manager uid
func getManagerEmployees(payload []byte) ([]byte, error) {
	// fetch list of all employee_id managed by manager uid
	url := employeeService + "api/v1/employee/get/managedEmployees"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// decode employee-service response
	var list []int
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&list)
	if err != nil {
		return nil, err
	}

	// Convert list to JSON
	jsonList, err := json.Marshal(list)
	if err != nil {
		return nil, err
	}

	return jsonList, err
}
