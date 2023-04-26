package handlers

import (
	"log"
	"net/http"
)

// receives a get request from broker to fetch all employee
func (rep *Repository) GetAll(w http.ResponseWriter, r *http.Request) {
	// get all employee
	all, err := rep.App.Models.Employee.GetAll()
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// response to be sent
	answer := jsonResponse{
		Error:   false,
		Message: "all employees successfully fetched",
		Data:    all,
	}

	log.Println(answer)

	rep.writeJSON(w, http.StatusAccepted, answer)

}
