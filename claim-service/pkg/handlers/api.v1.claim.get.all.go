package handlers

import (
	"net/http"
)

// GetAllClaim is the handler which returns all claim definition
func (rep *Repository) GetAllClaim(w http.ResponseWriter, r *http.Request) {
	// get claim definition (active, inactive, deleted)
	all, err := rep.App.Models.Claim.GetAllClaim()
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// response to be sent
	answer := jsonResponse{
		Error:   false,
		Message: "claim definition successfully fetched",
		Data:    all,
	}

	rep.writeJSON(w, http.StatusAccepted, answer)

}
