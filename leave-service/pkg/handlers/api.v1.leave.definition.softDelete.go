package handlers

import (
	"net/http"
	"strconv"
	"time"
)

// receives a get request from broker to fetch all employee
func (rep *Repository) SoftDeleteLeaveDefinition(w http.ResponseWriter, r *http.Request) {
	// extract payload from request
	var p listID

	err := rep.readJSON(w, r, &p)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// soft delete leave definition one by one
	// remove leave definition from employees leave
	for _, eid := range p.List {
		id, _ := strconv.Atoi(eid)
		// soft delete definition
		err = rep.App.Models.LeaveDefinition.Delete(id)
		if err != nil {
			rep.errorJSON(w, err)
			return
		}

		// remove entitled leave definition from employees
		err = rep.App.Models.LeaveDefinition.RemoveEntitledByDefinition(id)
		if err != nil {
			rep.errorJSON(w, err)
			return
		}
	}

	// response to be sent
	answer := jsonResponse{
		Error:     false,
		Message:   "leave definition successfully deleted",
		Data:      p.List,
		CreatedAt: time.Now().Format("02-Jan-2006 15:04:05"),
		CreatedBy: p.ConnectedUser,
	}

	rep.writeJSON(w, http.StatusAccepted, answer)

}
