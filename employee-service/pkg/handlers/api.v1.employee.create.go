package handlers

import (
	"bytes"
	"employee/pkg/pg"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// CreateEmployee is the handler which receives
// an employee payload from the broker
// convert it and inserts it to the DB
func (rep *Repository) Create(w http.ResponseWriter, r *http.Request) {

	// extract payload from request
	var p pg.Employee
	err := rep.readJSON(w, r, &p)
	if err != nil {
		log.Println("err decoding:", err)
		rep.errorJSON(w, err)
		return
	}

	// insert payload to employee DB
	rep.App.Models.Employee = p
	rowID, err := rep.App.Models.Employee.Insert()
	if err != nil {
		log.Println("err db:", err)
		rep.errorJSON(w, err)
		return
	}

	// convert payload to employee (employee-service) & user (authentication-service)
	// newEmployee := convertToEmployee(p)
	newUser := convertToUser(p, rowID)

	// insert all the child table
	// (emergency contact, spouse, employment, statutory, addition&deduction, bank, attachment )
	err = rep.App.Models.Employee.InsertChilds(rowID, p.CreatedBy)
	if err != nil {
		log.Println("err db childs:", err)
		rep.errorJSON(w, err)
		return
	}

	// send user payload to be inserted by authentication-service
	err = sendUserToAuthentication(newUser)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// store new employee info
	u := User{
		ID:    rowID,
		Email: p.PrimaryEmail,
	}

	// response to be sent
	answer := jsonResponse{
		Error:     false,
		Message:   "employee successfully created",
		Data:      u,
		CreatedAt: time.Now().Format("02-Jan-2006 15:04:05"),
		CreatedBy: p.ConnectedUser,
	}

	rep.writeJSON(w, http.StatusAccepted, answer)

}

// convert payloadNewEmployee to authPayload
func convertToUser(p pg.Employee, eid int) authPayload {
	return authPayload{
		Email:      p.PrimaryEmail,
		Password:   "Thundersoft@123",
		Fullname:   p.Fullname,
		Nickname:   p.Nickname,
		Active:     p.IsActive,
		Role:       p.Role,
		EmployeeID: eid,
	}
}

// send user payload to be inserted by authentication-service
func sendUserToAuthentication(p authPayload) error {
	// Encode payload to []byte
	encoded, err := json.Marshal(p)
	if err != nil {
		return err
	}

	// send http POST request to authentication-service
	url := authenticationService + "api/v1/authentication/createUser"
	_, err = http.Post(url, "application/json", bytes.NewBuffer(encoded))
	if err != nil {
		return err
	}

	return nil
}
