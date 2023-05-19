package handlers

import (
	"net/http"
)

// forward the post request from the front-end to claim-service to get all my claims
func AllMyClaim(w http.ResponseWriter, r *http.Request) {
	// need to extract payload since r.Body is a ReadCloser
	var p User
	err := readJSON(w, r, &p)
	if err != nil {
		errorJSON(w, err)
		return
	}

	type allNeeded struct {
		MyClaims             interface{}
		MySeniority          interface{}
		AllClaimsDefinitions interface{}
		AllEmployees         interface{}
	}
	var all allNeeded

	// create all post requests
	postRequests := []Request{
		{URL: claimService + "api/v1/myclaim/get/all", Data: p},
		{URL: employeeService + "api/v1/employee/get/seniority", Data: p},
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
		if p.URL == claimService+"api/v1/myclaim/get/all" {
			all.MyClaims = p.Data
		}
		if p.URL == employeeService+"api/v1/employee/get/seniority" {
			all.MySeniority = p.Data
		}
	}
	// create all get requests
	var noData any
	getRequests := []Request{
		{URL: employeeService + "api/v1/employee/get/all", Data: noData},
		{URL: claimService + "api/v1/claim/definition/get/all", Data: noData},
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
		if g.URL == claimService+"api/v1/claim/definition/get/all" {
			all.AllClaimsDefinitions = g.Data
		}
	}

	// send response to front-end
	writeJSON(w, http.StatusAccepted, all)
}
