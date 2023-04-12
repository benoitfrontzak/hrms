package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

// employeeList holds one employee information
type employeeList struct {
	ID        int
	Seniority int // number of years from join date
	GenderID  int
}

// Structure holding one leave Definition
type LeaveDefinition struct {
	ID       int                       `json:"rowid,string,omitempty"`
	GenderID int                       `json:"genderid,string"`
	Details  []*LeaveDefinitionDetails `json:"details"`
}

// Structure holding one leave Definition calculation
type LeaveDefinitionDetails struct {
	Entitled  int     `json:",string"`
	Seniority int     `json:",string"`
	ToAdd     float64 `json:",string"`
}

func getLeaveDefinitionFromResponse(data any) (leaveID, calculationMethodID int) {
	lID := data.(map[string]interface{})["rowid"]
	cID := data.(map[string]interface{})["calculationid"]

	leaveID, _ = strconv.Atoi(lID.(string))
	calculationMethodID, _ = strconv.Atoi(cID.(string))

	return
}

func getActiveEmployeeSeniority() ([]*employeeList, error) {
	// get all active employees with seniority
	url := employeeService + "api/v1/employee/get/allActive"

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// store response
	type myjsonResponse struct {
		Error   bool            `json:"error"`
		Message string          `json:"message"`
		Data    []*employeeList `json:"data,omitempty"`
	}
	var answer myjsonResponse

	// decode response
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&answer)
	if err != nil {
		return nil, err
	}
	if answer.Error {
		return nil, errors.New(answer.Message)
	}

	return answer.Data, nil
}

func getDefinitionsByID(lid int) (*LeaveDefinition, error) {
	// leave service URL
	url := leaveService + "api/v1/leave/definition/get/id"

	// structure holding the payload to send
	type payload struct {
		ID int
	}

	// encode payload to send
	p := payload{ID: lid}
	json_data, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	// get leave definition information
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// store response
	type myjsonResponse struct {
		Error   bool             `json:"error"`
		Message string           `json:"message"`
		Data    *LeaveDefinition `json:"data,omitempty"`
	}
	var answer myjsonResponse

	// decode response
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&answer)
	if err != nil {
		return nil, err
	}
	if answer.Error {
		return nil, errors.New(answer.Message)
	}

	return answer.Data, nil
}

func createEntitledByDefinition(lid, eid int, entitled float64) error {
	// leave service URL
	url := leaveService + "api/v1/leave/definition/create/employee/entitled"

	// structure holding the payload to send
	type payload struct {
		LeaveID    int
		EmployeeID int
		Entitled   float64
	}

	// encode payload to send
	p := payload{
		LeaveID:    lid,
		EmployeeID: eid,
		Entitled:   entitled,
	}
	json_data, err := json.Marshal(p)
	if err != nil {
		return err
	}

	// get leave definition information
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var answer jsonResponse
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&answer)
	if err != nil {
		return err
	}

	return nil
}
