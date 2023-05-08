package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

// forward the post request from the front-end to leave-service to create a new leave definition
// when created, update LEAVE_EMPLOYEE with the new definition
func CreateLeaveDefinition(w http.ResponseWriter, r *http.Request) {
	// send post request to leave-service and collect the response
	url := leaveService + "api/v1/leave/definition/create"
	resp, err := http.Post(url, "application/json", r.Body)
	if err != nil {
		errorJSON(w, err)
		return
	}
	defer resp.Body.Close()

	// decode response
	var answer jsonResponse
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&answer)
	if err != nil {
		errorJSON(w, err)
		return
	}
	if answer.Error {
		errorJSON(w, errors.New(answer.Message))
		return
	}

	// decode leave definition information from answer
	leaveID, calculationMethodID := getLeaveDefinitionFromResponse(answer.Data)

	// update LEAVE_EMPLOYEE DB with new leave definition
	createNewLeaveEmployeeEntitled(leaveID, calculationMethodID)

	// log to leave-service collection
	l := rpcPayload{
		Collection: "leave",
		Name:       "create definition",
		Data:       fmt.Sprintf("new leave definition successfully created for table with id %d", leaveID),
		CreatedAt:  answer.CreatedAt,
		CreatedBy:  answer.CreatedBy,
	}
	LogItemViaRPC(l)

	// send response to front-end
	writeJSON(w, http.StatusAccepted, answer)
}

func createNewLeaveEmployeeEntitled(leaveID, calculationMethodID int) {
	// get all active employees with seniority
	allEmployees, err := getActiveEmployeeSeniority()
	if err != nil {
		log.Println(err)
	}

	// select all leave definition details (max entitled per seniority per year)
	// for earned calculation method (calculation_method_id = 2)
	lDef, err := getDefinitionsByID(leaveID)
	if err != nil {
		log.Println(err)
	}

	genderEmployees := []*employeeList{}
	// when gender is not for both (gender_id = 2)
	if lDef.GenderID != 2 {
		for _, e := range allEmployees {
			if lDef.GenderID == e.GenderID {
				log.Println(e)
				genderEmployees = append(genderEmployees, e)
			}
		}
	} else {
		genderEmployees = allEmployees
	}

	// case calculation method is yearly (id=1)
	if calculationMethodID == 1 {
		// create new definition for each employee
		for _, e := range genderEmployees {
		out1:
			// for each seniority
			for _, d := range lDef.Details {
				if e.Seniority >= d.Seniority {
					_ = createEntitledByDefinition(lDef.ID, e.ID, float64(d.Entitled))
					continue out1
				}
			}

		}
	}

	// case calculation method is earn (id=2)
	if calculationMethodID == 2 {
		// calculate entitlement of one month for each seniority
		for _, d := range lDef.Details {
			d.ToAdd = float64(d.Entitled) / 12
		}

		// create new definition for each employee
		for _, e := range genderEmployees {
		out2:
			// for each seniority
			for _, d := range lDef.Details {
				if e.Seniority >= d.Seniority {
					_ = createEntitledByDefinition(lDef.ID, e.ID, d.ToAdd)
					continue out2
				}
			}

		}
	}
}
