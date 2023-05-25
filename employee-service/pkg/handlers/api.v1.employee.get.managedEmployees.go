package handlers

import (
	"net/http"
)

// receives a get request from broker to fetch all employee info
func (rep *Repository) GetManagedEmployees(w http.ResponseWriter, r *http.Request) {
	// extract employee id from request
	var u User
	err := rep.readJSON(w, r, &u)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	all, err := rep.App.Models.Employee.GetManagedEmployees(u.ID)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// response to be sent
	rep.writeJSON(w, http.StatusAccepted, all)

}
