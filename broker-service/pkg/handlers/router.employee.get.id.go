package handlers

import (
	"log"
	"net/http"
)

// forward the get request from the front-end to employee-service to fetch all employee info
func EmployeeByID(w http.ResponseWriter, r *http.Request) {
	// need to extract payload since r.Body is a ReadCloser
	var p User
	err := readJSON(w, r, &p)
	if err != nil {
		log.Println(err)
		errorJSON(w, err)
		return
	}

	type allNeeded struct {
		AllEmployees interface{}
		EmployeeCT   interface{}
		EmployeeInfo interface{}
	}
	var all allNeeded

	// create all post requests
	postRequests := []Request{
		{URL: employeeService + "api/v1/employee/get/id", Data: p},
	}

	// create a channel to receive responses and errors
	postResponseChan := make(chan *MyResponse)

	// get all post responses
	allPostResp, err := makeConcurrentPostRequests(postRequests, postResponseChan)
	if err != nil {
		errorJSON(w, err)
		return
	}

	// assign appropriate answer to allNeeded
	for _, p := range allPostResp {
		if p.URL == employeeService+"api/v1/employee/get/id" {
			all.EmployeeInfo = p.Data
		}
	}

	// create all get requests
	var noData any
	getRequests := []Request{
		{URL: employeeService + "api/v1/employee/get/all", Data: noData},
		{URL: employeeService + "api/v1/employee/configTables/get", Data: noData},
	}

	// create a channel to receive responses and errors
	getResponseChan := make(chan *MyResponse)

	// get all get responses
	allGetResp, err := makeConcurrentGetRequests(getRequests, getResponseChan)
	if err != nil {
		errorJSON(w, err)
		return
	}

	for _, g := range allGetResp {
		if g.URL == employeeService+"api/v1/employee/get/all" {
			all.AllEmployees = g.Data
		}
		if g.URL == employeeService+"api/v1/employee/configTables/get" {
			all.EmployeeCT = g.Data
		}
	}

	// send response to front-end
	writeJSON(w, http.StatusAccepted, all)
}
