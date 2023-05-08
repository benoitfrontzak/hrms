package handlers

import (
	"net/http"
)

// receives a get request from broker to fetch all employee info
func (rep *Repository) GetSeniorityByID(w http.ResponseWriter, r *http.Request) {
	// extract employee id from request
	var u User
	err := rep.readJSON(w, r, &u)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	seniority, err := rep.App.Models.Employee.GetSeniorityByID(u.ID)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// response to be sent
	answer := jsonResponse{
		Error:   false,
		Message: "employee summary successfully fetched",
		Data:    seniority,
	}

	rep.writeJSON(w, http.StatusAccepted, answer)

}
