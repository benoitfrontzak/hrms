package handlers

import (
	"net/http"
)

// GetAllPublicHoliday is the handler which returns all public holidays
func (rep *Repository) GetAllPublicHoliday(w http.ResponseWriter, r *http.Request) {
	// get all public holidays
	all, err := rep.App.Models.PublicHoliday.GetAllPublicHolidays()
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// response to be sent
	answer := jsonResponse{
		Error:   false,
		Message: "public holidays successfully fetched",
		Data:    all,
	}

	rep.writeJSON(w, http.StatusAccepted, answer)

}
