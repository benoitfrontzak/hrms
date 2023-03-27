package handlers

import (
	"net/http"
)

// receives a get request from broker to fetch all employee info
func (rep *Repository) GetAllActive(w http.ResponseWriter, r *http.Request) {

	all, err := rep.App.Models.Employee.GetAllActive()
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// response to be sent
	answer := jsonResponse{
		Error:   false,
		Message: "active employee list successfully fetched",
		Data:    all,
	}

	rep.writeJSON(w, http.StatusAccepted, answer)

}
