package handlers

import (
	"net/http"
	"strconv"
	"strings"
)

// GetAllLeave is the handler which returns all leave definition
func (rep *Repository) GetAllLeaveManager(w http.ResponseWriter, r *http.Request) {
	// extract payload from request
	var p []int

	err := rep.readJSON(w, r, &p)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// Convert each element to a string
	strSlice := make([]string, len(p))
	for i, num := range p {
		strSlice[i] = strconv.Itoa(num)
	}

	// Join the string elements with a separator
	result := strings.Join(strSlice, ", ")

	// get leaves (approved, rejected, pending)
	all, err := rep.App.Models.Leave.GetAllLeaveManager(result)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// response to be sent
	answer := jsonResponse{
		Error:   false,
		Message: "leave successfully fetched",
		Data:    all,
	}

	rep.writeJSON(w, http.StatusAccepted, answer)

}
