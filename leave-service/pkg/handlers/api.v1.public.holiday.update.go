package handlers

import (
	"net/http"
	"time"
)

// CreatePH is the handler which receives a PH payload from the broker
// convert it and inserts it to the DB
func (rep *Repository) UpdatePH(w http.ResponseWriter, r *http.Request) {
	// extract payload from request
	err := rep.readJSON(w, r, &rep.App.Models.PublicHoliday)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// insert payload to employee's config table DB
	err = rep.App.Models.PublicHoliday.Update()
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// response to be sent
	answer := jsonResponse{
		Error:     false,
		Message:   "Public Holiday successfully updated",
		Data:      err,
		CreatedAt: time.Now().Format("02-Jan-2006 15:04:05"),
		CreatedBy: "p.ConnectedUser",
	}

	rep.writeJSON(w, http.StatusAccepted, answer)

}
