package handlers

import (
	"bytes"
	"employee/pkg/pg"
	"encoding/json"
	"net/http"
	"time"
)

// UpdateEmployee is the handler which receives a POST request
// from the broker to update emplioyee information to the DB
func (rep *Repository) Update(w http.ResponseWriter, r *http.Request) {

	// extract payload from request
	var p pg.EmployeeFull
	err := rep.readJSON(w, r, &p)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// update employee's information to DB
	err = rep.App.Models.Employee.UpdateAllEmployeeInformation(p)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// store updated employee info
	u := User{
		ID:    p.Employee.ID,
		Email: p.Employee.PrimaryEmail,
	}

	// update user information
	user := authPayload{
		Email:      p.Employee.PrimaryEmail,
		Fullname:   p.Employee.Fullname,
		Nickname:   p.Employee.Nickname,
		Active:     p.Employee.IsActive,
		Role:       p.Employee.Role,
		EmployeeID: p.Employee.ID,
	}
	err = updateUser(user)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// response to be sent
	answer := jsonResponse{
		Error:     false,
		Message:   "employee successfully updated",
		Data:      u,
		CreatedAt: time.Now().Format("02-Jan-2006 15:04:05"),
		CreatedBy: p.Employee.ConnectedUser,
	}

	rep.writeJSON(w, http.StatusAccepted, answer)

}

// send user payload to be inserted by authentication-service
func updateUser(p authPayload) error {
	// Encode payload to []byte
	encoded, err := json.Marshal(p)
	if err != nil {
		return err
	}

	// send http POST request to authentication-service
	url := authenticationService + "api/v1/authentication/updateUser"
	_, err = http.Post(url, "application/json", bytes.NewBuffer(encoded))
	if err != nil {
		return err
	}

	return nil
}
