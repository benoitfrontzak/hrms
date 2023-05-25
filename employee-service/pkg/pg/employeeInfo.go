package pg

import "context"

// fetch employee seniority by id
func (e *Employee) GetSeniorityByID(id int) (int, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch employee summary
	query := `SELECT date_part('year', age(e2.join_date))as joined
			  FROM public."EMPLOYEE" e, "EMPLOYMENT" e2 
			  WHERE e.id = $1
			  AND e2.employee_id = e.id;`

	// executes SQL query
	row := db.QueryRowContext(ctx, query, id)

	var seniority int

	// populate returned row to total
	err := row.Scan(
		&seniority,
	)
	if err != nil {
		return 0, err
	}

	// return all employee summary rows
	return seniority, nil

}

// fetch employee seniority by id
func (e *Employee) FindEmployeeEmail(id int) (string, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch employee summary
	query := `SELECT primary_email
			  FROM public."EMPLOYEE" e
			  WHERE e.id = $1`

	// executes SQL query
	row := db.QueryRowContext(ctx, query, id)

	var email string

	// populate returned row to total
	err := row.Scan(
		&email,
	)
	if err != nil {
		return "", err
	}

	// return all employee summary rows
	return email, nil

}

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
	employmentArchive, err := getEmployeeEmploymentArchiveByID(employee.ID)
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
	statutoryArchive, err := getEmployeeStatutoryArchiveByID(employee.ID)
	if err != nil {
		return nil, err
	}
	payrollItems, err := getEmployeePayrollItemsByID(employee.ID)
	if err != nil {
		return nil, err
	}
	myEmp := EmployeeFull{
		Employee:          employee,
		EmergencyContact:  ec,
		OtherInformation:  oi,
		Spouse:            spouse,
		Employment:        employment,
		EmploymentArchive: employmentArchive,
		Statutory:         statutory,
		StatutoryArchive:  statutoryArchive,
		Bank:              bank,
		PayrollItems:      payrollItems,
	}
	return &myEmp, nil
}

// fetch all employee info by email
func (e *Employee) GetEmployeeInfoByEmail(email string) (*EmployeeFull, error) {
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
					 e.fullname, e.nickname,
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
		&el.Fullname, &el.Nickname,
		&el.JoinDate, &el.ConfirmDate,
	)
	if err != nil {
		return nil, err
	}
	// return employee full row
	return &el, nil
}

// fetch all employees managed by manager uid
func (e *Employee) GetManagedEmployees(id int) ([]int, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch employee summary
	query := `SELECT e.employee_id  
			  FROM "EMPLOYMENT" e
			  WHERE e.employee_id = $1
			  OR e.superior_id = $2
			  OR e.supervisor_id = $3`

	// executes SQL query
	rows, err := db.QueryContext(ctx, query, id, id, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// populate returned row to employee full struct
	var list []int
	for rows.Next() {
		var eid int
		err := rows.Scan(
			&eid,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, eid)
	}
	if err != nil {
		return nil, err
	}
	// return employee full row
	return list, nil
}

// helpers
func getEmployeeFullByID(id int) (*Employee, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch employee summary
	query := `SELECT e.id,
					 e.employee_code,
					 e.fullname, e.nickname,
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
		&e.Fullname, &e.Nickname,
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
					 e.fullname, e.nickname,
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
		&e.Fullname, &e.Nickname,
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
func getEmployeeEmploymentArchiveByID(id int) ([]*EmploymentArchive, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch employee summary
	query := `SELECT e.id,
					 e.job_title,
					 ced."name" as department,
					 e.superior_id,	e.supervisor_id,
					 cet."name" as employee_type,
					 cew."name" as wages_type,
					 e.basic_rate,
					 cepf."name"  as pay_frequency, cepb."name"  as payment_by, ceb."name"  as bank_payout,
					 ceg."name"  as group, ceb2."name"  as branch, cep."name" as project,
					 ceo."name"  as overtime,
					 e.working_permit_expiry,
					 e.join_date, e.confirm_date, e.resign_date,
					 e.created_at, e.created_by
			  FROM "EMPLOYMENT_ARCHIVE" e, "CONFIG_EMPLOYMENT_DEPARTMENT" ced, "CONFIG_EMPLOYMENT_TYPE" cet, 
					"CONFIG_EMPLOYMENT_WAGES" cew, "CONFIG_EMPLOYMENT_PAY_FREQUENCY" cepf, "CONFIG_EMPLOYMENT_PAYMENT_BY" cepb, 
					"CONFIG_EMPLOYMENT_BANK" ceb, "CONFIG_EMPLOYMENT_GROUP" ceg , "CONFIG_EMPLOYMENT_BRANCH" ceb2, "CONFIG_EMPLOYMENT_PROJECT" cep, "CONFIG_EMPLOYMENT_OVERTIME" ceo 
			  WHERE e.employee_id  = $1
			  AND e.department_id = ced.id 
			  AND e.employee_type_id = cet.id 
			  AND e.wages_type_id = cew.id 
			  AND e.pay_frequency_id = cepf.id 
			  AND e.payment_by_id = cepb.id 
			  AND e.bank_payout_id = ceb.id 
			  AND e.group_id = ceg.id 
			  AND e.branch_id = ceb2.id 
			  AND e.project_id = cep.id 
			  AND e.overtime_id = ceo.id`

	// executes SQL query
	rows, err := db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// populate returned row to employee full struct
	var all []*EmploymentArchive
	for rows.Next() {
		var e EmploymentArchive
		err := rows.Scan(
			&e.ID,
			&e.JobTitle,
			&e.Department,
			&e.Superior, &e.Supervisor,
			&e.EmployeeType,
			&e.WagesType,
			&e.BasicRate,
			&e.PayFrequency, &e.PaymentBy, &e.BankPayout,
			&e.Group, &e.Branch, &e.Project,
			&e.Overtime,
			&e.WorkingPermitExpiry,
			&e.JoinDate, &e.ConfirmDate, &e.ResignDate,
			&e.CreatedAt, &e.CreatedBy,
		)
		if err != nil {
			return nil, err
		}
		all = append(all, &e)
	}
	// return employee full row
	return all, nil
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
func getEmployeeStatutoryArchiveByID(id int) ([]*StatutoryArchive, error) {
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
					 s.contribute_hrdf,
					 s.created_at, s.created_by
			FROM "STATUTORY_ARCHIVE" s WHERE s.employee_id  = $1`

	// executes SQL query
	rows, err := db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// populate returned rows to statutory struct
	var all []*StatutoryArchive
	for rows.Next() {
		var s StatutoryArchive
		err := rows.Scan(
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
			&s.CreatedAt, &s.CreatedBy,
		)
		if err != nil {
			return nil, err
		}

		all = append(all, &s)
	}
	// return employee full row
	return all, nil
}
func getEmployeePayrollItemsByID(id int) ([]*PayrollItem, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch employee summary
	query := `select 	ad.id,
						cpt."name" as type,
						cpi.code, 
						cpi.description,
						cpi.pay_epf,
						cpi.pay_socso_eif,
						cpi.pay_hrdf,
						cpi.pay_tax,					
						ad.start_period,
						ad.end_period,
						ad.amount,
						ad.created_at,
						ad.created_by,
						ad.updated_at,
						ad.updated_by 
				FROM "ADDITION_DEDUCTION" ad, "CONFIG_PAYROLL_ITEM" cpi, "CONFIG_PAYROLL_TYPE" cpt 
				WHERE ad.employee_id = $1 
				AND ad.payroll_item_id = cpi.id 
				AND cpi.payroll_type_id = cpt.id 
				AND ad.soft_delete=0`

	// executes SQL query
	rows, err := db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// populate returned rows to statutory struct
	var all []*PayrollItem
	for rows.Next() {
		var pi PayrollItem
		err := rows.Scan(
			&pi.ID,
			&pi.Type,
			&pi.Code,
			&pi.Description,
			&pi.PayEPF,
			&pi.PaySOCSO,
			&pi.PayHRDF,
			&pi.PayTax,
			&pi.Start,
			&pi.End,
			&pi.Amount,
			&pi.CreatedAt, &pi.CreatedBy,
			&pi.UpdatedAt, &pi.UpdatedBy,
		)
		if err != nil {
			return nil, err
		}

		all = append(all, &pi)
	}
	// return employee full row
	return all, nil
}
