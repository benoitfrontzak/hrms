package pg

import "context"

// fetch all employee info by id
func (e *Employee) GetAllEmployeeInfo(id int) (*EmployeeFull, error) {
	employee, err := getEmployeeFullByID(id)
	if err != nil {
		return nil, err
	}
	bank, err := getEmployeeBankByID(employee.ID)
	if err != nil {
		return nil, err
	}
	ec, err := getEmployeeEmergencyContactByID(employee.ID)
	if err != nil {
		return nil, err
	}
	oi, err := getEmployeeOtherInformationByID(employee.ID)
	if err != nil {
		return nil, err
	}
	employment, err := getEmployeeEmmploymentByID(employee.ID)
	if err != nil {
		return nil, err
	}
	spouse, err := getEmployeeSpouseByID(employee.ID)
	if err != nil {
		return nil, err
	}
	statutory, err := getEmployeeStatutoryByID(employee.ID)
	if err != nil {
		return nil, err
	}

	myEmp := EmployeeFull{
		Employee:         employee,
		EmergencyContact: ec,
		OtherInformation: oi,
		Spouse:           spouse,
		Employment:       employment,
		Statutory:        statutory,
		Bank:             bank,
	}
	return &myEmp, nil
}

// fetch all employee info by email
func (e *Employee) GetAllEmployeeInfoByEmail(email string) (*EmployeeFull, error) {
	employee, err := getEmployeeFullByEmail(email)
	if err != nil {
		return nil, err
	}
	bank, err := getEmployeeBankByID(employee.ID)
	if err != nil {
		return nil, err
	}
	ec, err := getEmployeeEmergencyContactByID(employee.ID)
	if err != nil {
		return nil, err
	}
	oi, err := getEmployeeOtherInformationByID(employee.ID)
	if err != nil {
		return nil, err
	}
	employment, err := getEmployeeEmmploymentByID(employee.ID)
	if err != nil {
		return nil, err
	}
	spouse, err := getEmployeeSpouseByID(employee.ID)
	if err != nil {
		return nil, err
	}
	statutory, err := getEmployeeStatutoryByID(employee.ID)
	if err != nil {
		return nil, err
	}

	myEmp := EmployeeFull{
		Employee:         employee,
		Bank:             bank,
		EmergencyContact: ec,
		OtherInformation: oi,
		Employment:       employment,
		Spouse:           spouse,
		Statutory:        statutory,
	}
	return &myEmp, nil
}

// fetch all employee info by id for claim application
func (e *Employee) GetEmployeeLeaveInfo(id int) (*employeeLeave, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch employee summary
	query := `SELECT e.id,
					 e.employee_code,
					 e.firstname, e.middlename, e.lastname, e.givenname,
				 	 em.join_date, em.confirm_date
			FROM public."EMPLOYEE" e, public."EMPLOYMENT" em	
			WHERE e.id  = $1
			AND em.employee_id = e.id`

	// executes SQL query
	row := db.QueryRowContext(ctx, query, id)

	// populate returned row to employee full struct
	var el employeeLeave
	err := row.Scan(
		&el.ID,
		&el.Code,
		&el.Firstname, &el.Middlename, &el.Lastname, &el.Nickname,
		&el.JoinDate, &el.ConfirmDate,
	)
	if err != nil {
		return nil, err
	}
	// return employee full row
	return &el, nil
}

// helpers
func getEmployeeFullByID(id int) (*Employee, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch employee summary
	query := `SELECT e.id,
					 e.employee_code,
					 e.firstname, e.middlename, e.lastname, e.givenname,
				 	 e.ic_number,  e.passport_number, e.passport_expiry_at,
					 e.birthdate,
					 e.nationality_id, e.residence_country_id,
					 e.marital_status_id,		
					 e.gender_id,
					 e.race_id, e.religion_id,
					 e.address_street1, e.address_street2, e.address_city, e.address_state, e.address_zip, e.address_country_id,
					 e.primary_phone, e.secondary_phone,		
					 e.primary_email, e.secondary_email,
					 e.is_foreigner, e.immigration_number,
					 e.is_disabled, e.is_active,
					 e.user_role_id
			FROM "EMPLOYEE" e	
			WHERE e.id  = $1`

	// executes SQL query
	row := db.QueryRowContext(ctx, query, id)

	// populate returned row to employee full struct
	var e Employee
	err := row.Scan(
		&e.ID,
		&e.EmployeeCode,
		&e.Firstname, &e.Middlename, &e.Familyname, &e.Givenname,
		&e.IcNumber, &e.PassportNumber, &e.PassportExpiryAt,
		&e.Birthdate,
		&e.Nationality, &e.Residence,
		&e.Maritalstatus,
		&e.Gender,
		&e.Race, &e.Religion,
		&e.Streetaddr1, &e.Streetaddr2, &e.City, &e.State, &e.Zip, &e.Country,
		&e.PrimaryPhone, &e.SecondaryPhone,
		&e.PrimaryEmail, &e.SecondaryEmail,
		&e.IsForeigner, &e.ImmigrationNumber,
		&e.IsDisabled, &e.IsActive,
		&e.Role,
	)
	if err != nil {
		return nil, err
	}
	// return employee full row
	return &e, nil
}
func getEmployeeFullByEmail(email string) (*Employee, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch employee summary
	query := `SELECT e.id,
					 e.employee_code,
					 e.firstname, e.middlename, e.lastname, e.givenname,
				 	 e.ic_number,  e.passport_number, e.passport_expiry_at,
					 e.birthdate,
					 e.nationality_id, e.residence_country_id,
					 e.marital_status_id,		
					 e.gender_id,
					 e.race_id, e.religion_id,
					 e.address_street1, e.address_street2, e.address_city, e.address_state, e.address_zip, e.address_country_id,
					 e.primary_phone, e.secondary_phone,		
					 e.primary_email, e.secondary_email,
					 e.is_foreigner, e.immigration_number,
					 e.is_disabled, e.is_active,
					 e.user_role_id
			FROM "EMPLOYEE" e	
			WHERE e.primary_email  = $1`

	// executes SQL query
	row := db.QueryRowContext(ctx, query, email)

	// populate returned row to employee full struct
	var e Employee
	err := row.Scan(
		&e.ID,
		&e.EmployeeCode,
		&e.Firstname, &e.Middlename, &e.Familyname, &e.Givenname,
		&e.IcNumber, &e.PassportNumber, &e.PassportExpiryAt,
		&e.Birthdate,
		&e.Nationality, &e.Residence,
		&e.Maritalstatus,
		&e.Gender,
		&e.Race, &e.Religion,
		&e.Streetaddr1, &e.Streetaddr2, &e.City, &e.State, &e.Zip, &e.Country,
		&e.PrimaryPhone, &e.SecondaryPhone,
		&e.PrimaryEmail, &e.SecondaryEmail,
		&e.IsForeigner, &e.ImmigrationNumber,
		&e.IsDisabled, &e.IsActive,
		&e.Role,
	)
	if err != nil {
		return nil, err
	}
	// return employee full row
	return &e, nil
}
func getEmployeeBankByID(id int) (*Bank, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch employee summary
	query := `SELECT b.id, b.bank_name, b.account_number, b.account_name FROM "BANK" b WHERE b.employee_id  = $1`

	// executes SQL query
	row := db.QueryRowContext(ctx, query, id)

	// populate returned row to employee full struct
	var bank Bank
	err := row.Scan(
		&bank.ID,
		&bank.BankName,
		&bank.AccountNumber,
		&bank.AccountName,
	)
	if err != nil {
		return nil, err
	}
	// return employee full row
	return &bank, nil
}
func getEmployeeEmergencyContactByID(id int) (*EmergencyContact, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch employee summary
	query := `SELECT ec.id, ec.fullname, ec.phone , ec.relationship_id FROM "EMERGENCY_CONTACT" ec WHERE ec.employee_id  = $1`

	// executes SQL query
	row := db.QueryRowContext(ctx, query, id)

	// populate returned row to employee full struct
	var ec EmergencyContact
	err := row.Scan(
		&ec.ID,
		&ec.FullnameEC,
		&ec.MobileEC,
		&ec.RelationshipEC,
	)
	if err != nil {
		return nil, err
	}
	// return employee full row
	return &ec, nil
}
func getEmployeeOtherInformationByID(id int) (*OtherInformation, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch employee summary
	query := `SELECT oi.id, oi.education, oi.experience , oi.notes FROM "OTHER_INFORMATION" oi WHERE oi.employee_id  = $1`

	// executes SQL query
	row := db.QueryRowContext(ctx, query, id)

	// populate returned row to employee full struct
	var oi OtherInformation
	err := row.Scan(
		&oi.ID,
		&oi.Education,
		&oi.Experience,
		&oi.Notes,
	)
	if err != nil {
		return nil, err
	}
	// return employee full row
	return &oi, nil
}
func getEmployeeEmmploymentByID(id int) (*Employment, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch employee summary
	query := `SELECT e.id,
					 e.job_title,
					 e.department_id,
					 e.superior_id,	e.supervisor_id,
					 e.employee_type_id,
					 e.wages_type_id,
					 e.basic_rate,
					 e.pay_frequency_id, e.payment_by_id, e.bank_payout_id,
					 e.default_rule_id,
					 e.group_id, e.branch_id, e.project_id,
					 e.overtime_id,
					 e.working_permit_expiry,
					 e.join_date, e.confirm_date, e.resign_date
			FROM "EMPLOYMENT" e WHERE e.employee_id  = $1`

	// executes SQL query
	row := db.QueryRowContext(ctx, query, id)

	// populate returned row to employee full struct
	var e Employment
	err := row.Scan(
		&e.ID,
		&e.JobTitle,
		&e.Department,
		&e.Superior, &e.Supervisor,
		&e.EmployeeType,
		&e.WagesType,
		&e.BasicRate,
		&e.PayFrequency, &e.PaymentBy, &e.BankPayout,
		&e.DefaultRule,
		&e.Group, &e.Branch, &e.Project,
		&e.Overtime,
		&e.WorkingPermitExpiry,
		&e.JoinDate, &e.ConfirmDate, &e.ResignDate,
	)
	if err != nil {
		return nil, err
	}
	// return employee full row
	return &e, nil
}
func getEmployeeSpouseByID(id int) (*Spouse, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch employee summary
	query := `SELECT s.id,
					 s.fullname,
					 s.ic_number, s.passport_number,
					 s.is_disabled, s.is_working,
					 s.tax_number,
					 s.deductible_child_number, s.deductible_child_amount,
					 s.primary_phone, s.secondary_phone,
					 s.address_street1, s.address_street2,
					 s.address_city, s.address_state,
					 s.address_zip, s.address_country_id
			FROM "SPOUSE" s WHERE s.employee_id  = $1`

	// executes SQL query
	row := db.QueryRowContext(ctx, query, id)

	// populate returned row to employee full struct
	var s Spouse
	err := row.Scan(
		&s.ID,
		&s.Fullname,
		&s.ICNumber, &s.PassporNumber,
		&s.IsDisabled, &s.IsWorking,
		&s.TaxNumber,
		&s.DeductibleChildNumber, &s.DeductibleChildAmount,
		&s.PrimaryPhone, &s.SecondaryPhone,
		&s.Streetaddr1, &s.Streetaddr2,
		&s.City, &s.State,
		&s.Zip, &s.Country,
	)
	if err != nil {
		return nil, err
	}
	// return employee full row
	return &s, nil
}
func getEmployeeStatutoryByID(id int) (*Statutory, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch employee summary
	query := `SELECT s.id,
					 s.epf_table_id, s.epf_number, s.epf_initial,
					 s.nk,
					 s.epf_is_borne,
					 s.socso_category_id, s.socso_number, s.socso_status_id, s.socso_is_borne,
					 s.contribute_eis, s.eis_is_borne,
					 s.tax_status_id, s.tax_number, s.tax_branch_id,
					 s.ea_serial_number,
					 s.tax_is_borne,
					 s.foreign_worker_levy_id,
					 s.zakat_number, s.zakat_amount,
					 s.tabung_haji_number, s.tabung_haji_amount,
					 s.asn_number, s.asn_amount,
					 s.contribute_hrdf
			FROM "STATUTORY" s WHERE s.employee_id  = $1`

	// executes SQL query
	row := db.QueryRowContext(ctx, query, id)

	// populate returned row to employee full struct
	var s Statutory
	err := row.Scan(
		&s.ID,
		&s.EPFTable, &s.EPFNumber, &s.EPFInitial,
		&s.NK,
		&s.EPFBorne,
		&s.SOCSOCategory, &s.SOCSONumber, &s.SOCSOStatus, &s.SOCSOBorne,
		&s.ContributeEIS, &s.EISBorne,
		&s.TaxStatus, &s.TaxNumber, &s.TaxBranch,
		&s.EASerial,
		&s.TaxBorne,
		&s.ForeignWorkerLevy,
		&s.ZakatNumber, &s.ZakatAmount,
		&s.TabungHajiNumber, &s.TabungHajiAmount,
		&s.ASNNumber, &s.ASNAmount,
		&s.ContributeHRDF,
	)
	if err != nil {
		return nil, err
	}
	// return employee full row
	return &s, nil
}
