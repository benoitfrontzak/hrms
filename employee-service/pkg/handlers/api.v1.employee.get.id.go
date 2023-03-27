package handlers

import (
	"net/http"
	"strconv"
)

// receives a get request from broker to fetch all employee info
func (rep *Repository) GetByID(w http.ResponseWriter, r *http.Request) {
	// extract employee id from request
	var id string
	err := rep.readJSON(w, r, &id)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// get all employee info
	eid, err := strconv.Atoi(id)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}
	all, err := rep.App.Models.Employee.GetAllEmployeeInfo(eid)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// response to be sent
	answer := jsonResponse{
		Error:   false,
		Message: "employee summary successfully fetched",
		Data:    all,
	}

	rep.writeJSON(w, http.StatusAccepted, answer)

}
