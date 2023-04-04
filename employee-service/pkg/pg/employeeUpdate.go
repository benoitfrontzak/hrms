package pg

import (
	"context"
	"log"
	"time"
)

// update employee's information and all child tables
func (e *Employee) UpdateAllEmployeeInformation(emp EmployeeFull) error {
	// employee id
	eid := emp.Employee.ID

	// user id
	uid := emp.Employee.UpdatedBy

	// update all employee's tables
	uEp := emp.Employee.updateEmployee()
	if uEp != nil {
		log.Println("uEp err: ", uEp)
		return uEp
	}

	uEC := emp.EmergencyContact.updateEmergencyContact(eid, uid)
	if uEC != nil {
		log.Println("uEC err: ", uEC)
		return uEC
	}

	uOI := emp.OtherInformation.updateOtherInformation(eid, uid)
	if uOI != nil {
		log.Println("uOI err: ", uOI)
		return uOI
	}

	uSp := emp.Spouse.updateSpouse(eid, uid)
	if uSp != nil {
		log.Println("uSp err: ", uSp)
		return uSp
	}

	uEm := emp.Employment.updateEmployment(eid, uid)
	log.Println("eid, uid")
	log.Println(eid, uid)
	if uEm != nil {
		log.Println("uEm err: ", uEm)
		return uEm
	}

	uSt := emp.Statutory.updateStatutory(eid, uid)
	if uSt != nil {
		log.Println("uSt err: ", uSt)
		return uSt
	}

	uBk := emp.Bank.updateBank(eid, uid)
	if uBk != nil {
		log.Println("uBk err: ", uBk)
		return uBk
	}

	return nil
}

// Update one employee's information
func (e *Employee) updateEmployee() error {
	now := time.Now().Format("2006-01-02")

	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which update an employee
	stmt := `UPDATE public."EMPLOYEE"
			 SET employee_code=$1, 
			 	 fullname=$2, nickname=$3, 
				 ic_number=$4, passport_number=$5, passport_expiry_at=$6, 
				 birthdate=$7, 
				 nationality_id=$8, residence_country_id=$9, 
				 marital_status_id=$10, gender_id=$11, race_id=$12, religion_id=$13, 
				 address_street1=$14, address_street2=$15, address_city=$16, address_state=$17, address_zip=$18, address_country_id=$19, 
				 primary_phone=$20, secondary_phone=$21, 
				 primary_email=$22, secondary_email=$23, 
				 is_foreigner=$24, immigration_number=$25, 
				 is_disabled=$26, is_active=$27, 
				 user_role_id=$28,
				 updated_at=$29,
				 updated_by=$30
			 WHERE id=$31;`

	// executes SQL query
	_, err := db.ExecContext(ctx, stmt,
		e.EmployeeCode,
		e.Fullname, e.Nickname,
		e.IcNumber, e.PassportNumber, e.PassportExpiryAt,
		e.Birthdate,
		e.Nationality, e.Residence,
		e.Maritalstatus, e.Gender, e.Race, e.Religion,
		e.Streetaddr1, e.Streetaddr2, e.City, e.State, e.Zip, e.Country,
		e.PrimaryPhone, e.SecondaryPhone,
		e.PrimaryEmail, e.SecondaryEmail,
		e.IsForeigner, e.ImmigrationNumber,
		e.IsDisabled, e.IsActive,
		e.Role,
		now,
		e.UpdatedBy,
		e.ID,
	)
	if err != nil {
		return err
	}

	// return error
	return nil
}

// Update one employee's emergency contact information
func (ec *EmergencyContact) updateEmergencyContact(eid, uid int) error {
	now := time.Now().Format("2006-01-02")

	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which update an employee
	stmt := `UPDATE public."EMERGENCY_CONTACT"
			 SET fullname=$1, 
			 phone=$2, 
			 relationship_id=$3,
			 updated_at=$4,
			 updated_by=$5
			 WHERE employee_id=$6;`

	// executes SQL query
	_, err := db.ExecContext(ctx, stmt,
		ec.FullnameEC,
		ec.MobileEC,
		ec.RelationshipEC,
		now,
		uid,
		eid,
	)
	if err != nil {
		return err
	}

	// return error
	return nil
}

// Update one employee's emergency contact information
func (oi *OtherInformation) updateOtherInformation(eid, uid int) error {
	now := time.Now().Format("2006-01-02")

	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which update an employee
	stmt := `UPDATE public."OTHER_INFORMATION"
			 SET education=$1, 
			 	 experience=$2, 
			 	 notes=$3,
				 updated_at=$4,
				 updated_by=$5
			 WHERE employee_id=$6;`

	// executes SQL query
	_, err := db.ExecContext(ctx, stmt,
		oi.Education,
		oi.Experience,
		oi.Notes,
		now,
		uid,
		eid,
	)
	if err != nil {
		return err
	}

	// return error
	return nil
}

// Update one employee's spouse information
func (s *Spouse) updateSpouse(eid, uid int) error {
	now := time.Now().Format("2006-01-02")

	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which update an employee
	stmt := `UPDATE public."SPOUSE"
			 SET fullname=$1, 
				 ic_number=$2, passport_number=$3, 
				 is_disabled=$4, is_working=$5, 
				 tax_number=$6, deductible_child_number=$7, deductible_child_amount=$8, 
				 primary_phone=$9, secondary_phone=$10, 
				 address_street1=$11, address_street2=$12, address_city=$13, address_state=$14, address_zip=$15, address_country_id=$16,
				 updated_at=$17,
			 	 updated_by=$18
			 WHERE employee_id=$19;`

	// executes SQL query
	_, err := db.ExecContext(ctx, stmt,
		s.Fullname,
		s.ICNumber, s.PassporNumber,
		s.IsDisabled, s.IsWorking,
		s.TaxNumber, s.DeductibleChildNumber, s.DeductibleChildAmount,
		s.PrimaryPhone, s.SecondaryPhone,
		s.Streetaddr1, s.Streetaddr2, s.City, s.State, s.Zip, s.Country,
		now,
		uid,
		eid,
	)
	if err != nil {
		return err
	}

	// return error
	return nil
}

// Update one employee's employment information
func (e *Employment) updateEmployment(eid, uid int) error {
	now := time.Now().Format("2006-01-02")

	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which update an employee
	stmt := `UPDATE public."EMPLOYMENT"
			 SET job_title=$1, 
			 	 department_id=$2, 
				 superior_id=$3, supervisor_id=$4, 
				 employee_type_id=$5, wages_type_id=$6, basic_rate=$7, 
				 pay_frequency_id=$8, payment_by_id=$9, bank_payout_id=$10, 
				 default_rule_id=$11, 
				 group_id=$12, branch_id=$13, project_id=$14, 
				 overtime_id=$15, 
				 working_permit_expiry=$16, 
				 join_date=$17, confirm_date=$18, resign_date=$19,
				 updated_at=$20,
				 updated_by=$21				 
			 WHERE employee_id=$22;`

	// executes SQL query
	_, err := db.ExecContext(ctx, stmt,
		e.JobTitle,
		e.Department,
		e.Superior, e.Supervisor,
		e.EmployeeType, e.WagesType, e.BasicRate,
		e.PayFrequency, e.PaymentBy, e.BankPayout,
		e.DefaultRule,
		e.Group, e.Branch, e.Project,
		e.Overtime,
		e.WorkingPermitExpiry,
		e.JoinDate, e.ConfirmDate, e.ResignDate,
		now,
		uid,
		eid,
	)
	if err != nil {
		return err
	}

	// return error
	return nil
}

// Update one employee's staturory information
func (s *Statutory) updateStatutory(eid, uid int) error {
	now := time.Now().Format("2006-01-02")

	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which update an employee
	stmt := `UPDATE public."STATUTORY"
			 SET epf_table_id=$1, epf_number=$2, epf_initial=$3, 
				 nk=$4, epf_is_borne=$5, 
				 socso_category_id=$6, socso_number=$7, socso_status_id=$8, socso_is_borne=$9, 
				 contribute_eis=$10, eis_is_borne=$11, 
				 tax_status_id=$12, tax_number=$13, tax_branch_id=$14, 
				 ea_serial_number=$15, tax_is_borne=$16, 
				 foreign_worker_levy_id=$17, 
				 zakat_number=$18, zakat_amount=$19, 
				 tabung_haji_number=$20, tabung_haji_amount=$21, 
				 asn_number=$22, asn_amount=$23, 
				 contribute_hrdf=$24,
				 updated_at=$25,
				 updated_by=$26
			 WHERE employee_id=$27;`

	// executes SQL query
	_, err := db.ExecContext(ctx, stmt,
		s.EPFTable, s.EPFNumber, s.EPFInitial,
		s.NK, s.EPFBorne,
		s.SOCSOCategory, s.SOCSONumber, s.SOCSOStatus, s.SOCSOBorne,
		s.ContributeEIS, s.EISBorne,
		s.TaxStatus, s.TaxNumber, s.TaxBranch,
		s.EASerial, s.TaxBorne,
		s.ForeignWorkerLevy,
		s.ZakatNumber, s.ZakatAmount,
		s.TabungHajiNumber, s.TabungHajiAmount,
		s.ASNNumber, s.ASNAmount,
		s.ContributeHRDF,
		now,
		uid,
		eid,
	)
	if err != nil {
		return err
	}

	// return error
	return nil
}

// Update one employee's bank information
func (b *Bank) updateBank(eid, uid int) error {
	now := time.Now().Format("2006-01-02")

	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which update an employee
	stmt := `UPDATE public."BANK"
			 SET bank_name=$1, account_number=$2, account_name=$3, updated_at=$4, updated_by=$5
			 WHERE employee_id=$6;`

	log.Println("now: ", now)
	log.Println("stmt: ", stmt)
	log.Println("$1:", b.BankName, "$2:", b.AccountNumber, "$3:", b.AccountName, "$4:", now, "$5:", uid, "$6:", eid)
	log.Println("$4: now is ", now)
	// executes SQL query
	_, err := db.ExecContext(ctx, stmt,
		b.BankName, b.AccountNumber, b.AccountName,
		now,
		uid,
		eid,
	)
	if err != nil {
		return err
	}

	// return error
	return nil
}
