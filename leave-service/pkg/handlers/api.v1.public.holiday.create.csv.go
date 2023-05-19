package handlers

import (
	"leave/pkg/pg"
	"net/http"
	"time"
)

// CreatePH is the handler which receives a PH payload from the broker
// convert it and inserts it to the DB
func (rep *Repository) CreateCSV(w http.ResponseWriter, r *http.Request) {
	// extract payload from request
	var p pg.CSV
	err := rep.readJSON(w, r, &p)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	for _, ph := range p.PublicHolidays {
		rep.App.Models.PublicHoliday.Date = ph.Date
		rep.App.Models.PublicHoliday.Name = ph.Name
		rep.App.Models.PublicHoliday.Description = ph.Description
		rep.App.Models.PublicHoliday.CreatedBy = p.CreatedBy
		// insert payload to employee's config table DB
		_, err := rep.App.Models.PublicHoliday.Insert()
		if err != nil {
			rep.errorJSON(w, err)
			return
		}
	}

	// response to be sent
	answer := jsonResponse{
		Error:     false,
		Message:   "new Public Holiday successfully created",
		Data:      "rowID",
		CreatedAt: time.Now().Format("02-Jan-2006 15:04:05"),
		CreatedBy: "p.ConnectedUser",
	}

	rep.writeJSON(w, http.StatusAccepted, answer)

}
