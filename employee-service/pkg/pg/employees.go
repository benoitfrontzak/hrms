package pg

import "context"

// fetch all active employees and returns []*employeeList
func (e *Employee) GetAllActive() ([]*employeeList, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch employee summary
	query := `SELECT 	e.ID,
						e.employee_code AS code,
						LOWER(e.firstname),
						LOWER(e.middlename),
						UPPER(e.lastname),
						LOWER(e.givenname)
			  	FROM "EMPLOYEE" e
				WHERE e.is_active = 1
				AND e.ID <> 0
				AND e.soft_delete = 0`

	// executes SQL query
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// populate returned rows to employee summary struct
	var all []*employeeList
	for rows.Next() {
		var emp employeeList
		err := rows.Scan(
			&emp.ID,
			&emp.Code,
			&emp.Firstname,
			&emp.Middlename,
			&emp.Lastname,
			&emp.Nickname,
		)
		if err != nil {
			return nil, err
		}

		all = append(all, &emp)
	}

	// return all employee summary rows
	return all, nil
}

// fetch all employee summary and returns it as []*EmployeeSummary,
// filter rows by active status and not soft deleted
func (e *Employee) GetAll() (*AllEmployees, error) {
	active, err := getAllByStatus(1)
	if err != nil {
		return nil, err
	}

	inactive, err := getAllByStatus(0)
	if err != nil {
		return nil, err
	}

	deleted, err := getAllDeleted()
	if err != nil {
		return nil, err
	}

	all := AllEmployees{
		Active:   active,
		Inactive: inactive,
		Deleted:  deleted,
	}
	return &all, nil
}

// helpers
func getAllByStatus(active int) ([]*EmployeeSummary, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch employee summary
	query := `SELECT 	e.ID,
						e.employee_code AS code,
						CONCAT(UPPER(e.lastname), ' ', LOWER(e.firstname), ' ' , LOWER(e.middlename)) AS fullname,
						e.primary_email AS email,
						e.primary_phone AS mobile,
						e.birthdate,
						ccr."name" AS race,
						e.gender_id AS gender
			  	FROM "EMPLOYEE" e, "CONFIG_COMMON_RACE" ccr
				WHERE e.is_active = $1
				AND e.ID <> 0
				AND e.soft_delete = 0
				AND e.race_id = ccr.id`

	// executes SQL query
	rows, err := db.QueryContext(ctx, query, active)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// populate returned rows to employee summary struct
	var all []*EmployeeSummary
	for rows.Next() {
		var emp EmployeeSummary
		err := rows.Scan(
			&emp.ID,
			&emp.Code,
			&emp.Fullname,
			&emp.Email,
			&emp.Mobile,
			&emp.Birthdate,
			&emp.Race,
			&emp.Gender,
		)
		if err != nil {
			return nil, err
		}

		all = append(all, &emp)
	}

	// return all employee summary rows
	return all, nil
}
func getAllDeleted() ([]*EmployeeSummary, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch employee summary
	query := `SELECT 	e.ID,
						e.employee_code AS code,
						CONCAT(UPPER(e.lastname), ' ', LOWER(e.firstname), ' ' , LOWER(e.middlename)) AS fullname,
						e.primary_email AS email,
						e.primary_phone AS mobile,
						e.birthdate,
						ccr."name" AS race,
						e.gender_id AS gender
			  	FROM "EMPLOYEE" e, "CONFIG_COMMON_RACE" ccr
				WHERE e.soft_delete = 1
				AND e.ID <> 0
				AND e.race_id = ccr.id`

	// executes SQL query
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// populate returned rows to employee summary struct
	var all []*EmployeeSummary
	for rows.Next() {
		var emp EmployeeSummary
		err := rows.Scan(
			&emp.ID,
			&emp.Code,
			&emp.Fullname,
			&emp.Email,
			&emp.Mobile,
			&emp.Birthdate,
			&emp.Race,
			&emp.Gender,
		)
		if err != nil {
			return nil, err
		}

		all = append(all, &emp)
	}

	// return all employee summary rows
	return all, nil
}
