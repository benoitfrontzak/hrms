package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type payloadNewEmployee struct {
	Firstname         string `json:"firstname"`
	Middlename        string `json:"middlename"`
	Familyname        string `json:"familyname"`
	Givenname         string `json:"givenname"`
	EmployeeCode      string `json:"employeeCode"`
	Gender            string `json:"gender"`
	Maritalstatus     string `json:"maritalstatus"`
	IcNumber          string `json:"icNumber"`
	PassportNumber    string `json:"passportNumber"`
	PassportExpiryAt  string `json:"passportExpiryAt"`
	Nationality       string `json:"nationality"`
	Residence         string `json:"residence"`
	ImmigrationNumber string `json:"immigrationNumber"`
	Birthdate         string `json:"birthdate"`
	Race              string `json:"race"`
	Religion          string `json:"religion"`
	IsDisabled        string `json:"isDisabled"`
	IsActive          string `json:"isActive"`
	IsForeigner       string `json:"isForeigner"`
	Eclaim            string `json:"eclaim"`
	Eleave            string `json:"eleave"`
	Streetaddr1       string `json:"streetaddr1"`
	Streetaddr2       string `json:"streetaddr2"`
	Zip               string `json:"zip"`
	City              string `json:"city"`
	State             string `json:"state"`
	Country           string `json:"country"`
	PrimaryPhone      string `json:"primaryPhone"`
	SecondaryPhone    string `json:"secondaryPhone"`
	PrimaryEmail      string `json:"primaryEmail"`
	SecondaryEmail    string `json:"secondaryEmail"`
	FullnameEC        string `json:"fullnameEC"`
	MobileEC          string `json:"mobileEC"`
	RelationshipEC    string `json:"relationshipEC"`
}

// CreateEmployee receive a post request from the front-end
// it send payloadNewEmployee to employee-service
func (rep *Repository) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	// Extract payload from request
	var p payloadNewEmployee
	err := rep.readJSON(w, r, &p)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// Send payload to employee-service
	err = insertEmployee(p)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	rep.writeJSON(w, http.StatusAccepted, "successfully sent")
}

func insertEmployee(p payloadNewEmployee) error {
	// Encode payload
	encoded, err := json.Marshal(p)
	if err != nil {
		log.Println("encode error:", err)
	}

	employeeServiceURL := "http://employee-service:8882/createEmployee"

	response, err := http.Post(employeeServiceURL, "application/json", bytes.NewBuffer(encoded))
	if err != nil {
		return err
	}
	defer response.Body.Close()
	log.Println("response form employee-service: ", response)
	return nil
}
