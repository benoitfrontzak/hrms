package pg

import "context"

// inserts a new employee into the database, and returns the ID of the newly inserted row
func (e *Employee) Insert() (int, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which insert a new employee
	stmt := `INSERT INTO public."EMPLOYEE"
	(employee_code, fullname, nickname, ic_number, passport_number, passport_expiry_at, birthdate, nationality_id, 
		residence_country_id, marital_status_id, gender_id, address_street1, address_street2, address_city, address_state, address_zip, address_country_id, primary_phone, 
		secondary_phone, primary_email, secondary_email, is_foreigner, immigration_number, is_disabled, is_active, race_id, religion_id, user_role_id, created_by, updated_by)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30) returning id;`

	// executes SQL query (set SQL parameters and cacth rowID)
	var newID int

	err := db.QueryRowContext(ctx, stmt,
		e.EmployeeCode,
		e.Fullname,
		e.Nickname,
		e.IcNumber,
		e.PassportNumber,
		e.PassportExpiryAt,
		e.Birthdate,
		e.Nationality,
		e.Residence,
		e.Maritalstatus,
		e.Gender,
		e.Streetaddr1,
		e.Streetaddr2,
		e.City,
		e.State,
		e.Zip,
		e.Country,
		e.PrimaryPhone,
		e.SecondaryPhone,
		e.PrimaryEmail,
		e.SecondaryEmail,
		e.IsForeigner,
		e.ImmigrationNumber,
		e.IsDisabled,
		e.IsActive,
		e.Race,
		e.Religion,
		e.Role,
		e.CreatedBy,
		e.UpdatedBy,
	).Scan(&newID)
	if err != nil {
		return 0, err
	}

	// return rowID of created new employee
	return newID, nil
}

// inserts new employee id to all child tables (empty values)
func (e *Employee) InsertChilds(eid, user int) error {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// insert EMERGENCY_CONTACT
	stmt := `INSERT INTO public."EMERGENCY_CONTACT"	(fullname, phone, relationship_id, employee_id, created_by, updated_by)	VALUES('', '', 0, $1, $2, $3);`
	_, err := db.ExecContext(ctx, stmt, eid, user, user)
	if err != nil {
		return err
	}

	// insert BANK
	stmt = `INSERT INTO public."BANK" (bank_name, account_number, account_name, employee_id, created_by, updated_by) VALUES('', '', '', $1, $2, $3);`
	_, err = db.ExecContext(ctx, stmt, eid, user, user)
	if err != nil {
		return err
	}

	// insert OTHER_INFORMATION
	stmt = `INSERT INTO public."OTHER_INFORMATION" (education, experience, notes, employee_id, created_by, updated_by) VALUES('', '', '', $1, $2, $3);`
	_, err = db.ExecContext(ctx, stmt, eid, user, user)
	if err != nil {
		return err
	}

	// insert SPOUSE
	stmt = `INSERT INTO public."SPOUSE"
			(fullname, ic_number, passport_number, is_disabled, is_working, tax_number, deductible_child_number, deductible_child_amount, primary_phone, secondary_phone, address_street1, address_street2, address_city, address_state, address_zip, address_country_id, employee_id, created_by, updated_by)
			VALUES('', '', '', 0, 1, '', 0, 0, '', '', '', '', '', '', '', 0, $1, $2, $3);`
	_, err = db.ExecContext(ctx, stmt, eid, user, user)
	if err != nil {
		return err
	}

	// insert EMPLOYMENT
	stmt = `INSERT INTO public."EMPLOYMENT"
			(job_title, department_id, superior_id, supervisor_id, employee_type_id, wages_type_id, basic_rate, pay_frequency_id, payment_by_id, bank_payout_id, group_id, branch_id, project_id, overtime_id, working_permit_expiry, join_date, confirm_date, resign_date, employee_id, created_by, updated_by)
			VALUES('', 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, '0001-01-01', '0001-01-01', '0001-01-01', '0001-01-01', $1, $2, $3);`
	_, err = db.ExecContext(ctx, stmt, eid, user, user)
	if err != nil {
		return err
	}

	// insert STATUTORY
	stmt = `INSERT INTO public."STATUTORY"
			(epf_table_id, epf_number, epf_initial, nk, epf_is_borne, socso_category_id, socso_number, socso_status_id, socso_is_borne, contribute_eis, eis_is_borne, tax_status_id, tax_number, tax_branch_id, ea_serial_number, tax_is_borne, foreign_worker_levy_id, zakat_number, zakat_amount, tabung_haji_number, tabung_haji_amount, asn_number, asn_amount, contribute_hrdf, employee_id, created_by, updated_by)
			VALUES(0, '', '', '', 0, 0, '', 0, 0, 0, 0, 0, '', 0, '', 0, 0, '', 0, '', 0, '', 0, 0, $1, $2, $3);`
	_, err = db.ExecContext(ctx, stmt, eid, user, user)
	if err != nil {
		return err
	}

	return nil
}
