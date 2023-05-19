package handlers

import (
	"net/http"
	"strconv"
	"time"
)

// receives a get request from broker to fetch all employee
func (rep *Repository) SoftDeletePublicHoliday(w http.ResponseWriter, r *http.Request) {
	// extract payload from request
	var p listID

	err := rep.readJSON(w, r, &p)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// soft delete Public Holiday one by one
	for _, eid := range p.List {
		id, _ := strconv.Atoi(eid)
		// soft delete definition
		err = rep.App.Models.PublicHoliday.Delete(id, p.UserID)
		if err != nil {
			rep.errorJSON(w, err)
			return
		}
	}

	// response to be sent
	answer := jsonResponse{
		Error:     false,
		Message:   "Public Holiday successfully deleted",
		Data:      p.List,
		CreatedAt: time.Now().Format("02-Jan-2006 15:04:05"),
		CreatedBy: p.ConnectedUser,
	}

	rep.writeJSON(w, http.StatusAccepted, answer)

}
