package handlers

import (
	"employee/pkg/pg"
	"log"
	"net/http"
	"strconv"
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

// AllUsers returns all users from the DB
func (rep *Repository) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	// Extract payload from request
	var p payloadNewEmployee
	err := rep.readJSON(w, r, &p)
	if err != nil {
		log.Println("error while decoding: ", err)
		rep.errorJSON(w, err)
		return
	}
	newEmployee := convertPayloadNewEmployee(p)
	rowID, err := rep.App.Models.Employee.Insert(newEmployee)
	if err != nil {
		log.Println("error while inserted: ", err)
	}

	log.Println("employee successfully inserted: ", rowID)

	rep.writeJSON(w, http.StatusAccepted, "employee successfully inserted")

}

func convertPayloadNewEmployee(p payloadNewEmployee) pg.Employee {
	gender, _ := strconv.Atoi(p.Gender)
	marital, _ := strconv.Atoi(p.Maritalstatus)
	nationality, _ := strconv.Atoi(p.Nationality)
	residence, _ := strconv.Atoi(p.Residence)

	var race, religion, isDisabled, isActive, isForeigner, eclaim, eleave, relationship int
	var err error
	race, err = strconv.Atoi(p.Race)
	if err != nil {
		race = 0
	}
	religion, err = strconv.Atoi(p.Religion)
	if err != nil {
		religion = 0
	}
	isDisabled, err = strconv.Atoi(p.IsDisabled)
	if err != nil {
		isDisabled = 0
	}
	isActive, err = strconv.Atoi(p.IsActive)
	if err != nil {
		isActive = 0
	}
	isForeigner, err = strconv.Atoi(p.IsForeigner)
	if err != nil {
		isForeigner = 0
	}
	eclaim, err = strconv.Atoi(p.Eclaim)
	if err != nil {
		eclaim = 0
	}
	eleave, err = strconv.Atoi(p.Eleave)
	if err != nil {
		eleave = 0
	}
	relationship, err = strconv.Atoi(p.RelationshipEC)
	if err != nil {
		relationship = 0
	}

	return pg.Employee{
		Firstname:         p.Firstname,
		Middlename:        p.Middlename,
		Familyname:        p.Familyname,
		Givenname:         p.Givenname,
		EmployeeCode:      p.EmployeeCode,
		Gender:            gender,
		Maritalstatus:     marital,
		IcNumber:          p.IcNumber,
		PassportNumber:    p.PassportNumber,
		PassportExpiryAt:  p.PassportExpiryAt,
		Nationality:       nationality,
		Residence:         residence,
		ImmigrationNumber: p.ImmigrationNumber,
		Birthdate:         p.Birthdate,
		Race:              race,
		Religion:          religion,
		IsDisabled:        isDisabled,
		IsActive:          isActive,
		IsForeigner:       isForeigner,
		Eclaim:            eclaim,
		Eleave:            eleave,
		Streetaddr1:       p.Streetaddr1,
		Streetaddr2:       p.Streetaddr2,
		Zip:               p.Zip,
		City:              p.City,
		State:             p.State,
		Country:           p.Country,
		PrimaryPhone:      p.PrimaryPhone,
		SecondaryPhone:    p.SecondaryPhone,
		PrimaryEmail:      p.PrimaryEmail,
		SecondaryEmail:    p.SecondaryEmail,
		FullnameEC:        p.FullnameEC,
		MobileEC:          p.MobileEC,
		RelationshipEC:    relationship,
	}
}
