package handlers

import (
	"net/http"
)

// forward the post request from the front-end to leave-service to get all my leaves
// fetch all others needed informations related to myLeave page
func AllMyLeave(w http.ResponseWriter, r *http.Request) {
	// need to extract payload since r.Body is a ReadCloser
	var p User
	err := readJSON(w, r, &p)
	if err != nil {
		errorJSON(w, err)
		return
	}

	type allNeeded struct {
		MyEntitled          interface{}
		MyLeaves            interface{}
		AllLeaveDefinitions interface{}
		AllEmployees        interface{}
	}
	var all allNeeded

	// create all post requests
	postRequests := []Request{
		{URL: leaveService + "api/v1/myleave/get/today", Data: p},
		{URL: leaveService + "api/v1/myleave/get/all", Data: p},
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
		if p.URL == leaveService+"api/v1/myleave/get/today" {
			all.MyEntitled = p.Data
		}
		if p.URL == leaveService+"api/v1/myleave/get/all" {
			all.MyLeaves = p.Data
		}
	}
	// create all get requests
	var noData any
	getRequests := []Request{
		{URL: employeeService + "api/v1/employee/get/all", Data: noData},
		{URL: leaveService + "api/v1/leave/definition/get/all", Data: noData},
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
		if g.URL == leaveService+"api/v1/leave/definition/get/all" {
			all.AllLeaveDefinitions = g.Data
		}
	}

	// send response to front-end
	writeJSON(w, http.StatusAccepted, all)
}
