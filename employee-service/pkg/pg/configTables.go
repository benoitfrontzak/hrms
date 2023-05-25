package pg

import (
	"context"
	"log"
)

// fetch all rows of CONFIG_* (t) and returns it as []configT
func getAllRowsConfigTable(t string) ([]configT, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch config_t
	query := `SELECT id, name FROM "` + t + `" WHERE soft_delete = 0 ORDER BY id`

	// executes SQL query
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// populate each row to configT
	var myCT []configT
	for rows.Next() {
		var ct configT
		err := rows.Scan(
			&ct.ID,
			&ct.Name,
		)
		if err != nil {
			return nil, err
		}
		// populate []configT
		myCT = append(myCT, ct)
	}

	// return all config_t rows
	return myCT, nil
}

// fetch all rows of CONFIG_COMMON_COUNTRY and returns it as []configC
func getAllRowsCountryTable(t string) ([]configC, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch config_country_common
	query := `SELECT id, name, code2, code3, nationality FROM "` + t + `" order by id`

	// executes SQL query
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		log.Println("reached rows ", rows)
		return nil, err
	}
	defer rows.Close()

	// populate each row to configC
	var myCC []configC
	for rows.Next() {
		var cc configC
		err := rows.Scan(
			&cc.ID,
			&cc.Name,
			&cc.Code2,
			&cc.Code3,
			&cc.Nationality,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		myCC = append(myCC, cc)
	}

	// return all country rows
	return myCC, nil
}

// fetch all rows of CONFIG_COMMON_COUNTRY and returns it as []configC
func getAllRowsPayrollItem() ([]PayrollItem, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch config_country_common
	query := `SELECT cpi.id,
				  	 cpt.name as payroll_type,
				  	 cpi.code,
				  	 cpi.description,
				  	 cpi.start_period,
				  	 cpi.end_period,
				  	 cpi.amount,
				  	 cpi.pay_epf,
				  	 cpi.pay_socso_eif,
				  	 cpi.pay_hrdf,
				  	 cpi.pay_tax,
				  	 cpi.is_fixed 
			  FROM "CONFIG_PAYROLL_ITEM" cpi, "CONFIG_PAYROLL_TYPE" cpt 
			  WHERE cpi.payroll_type_id = cpt.id
			  AND cpi.soft_delete = 0
			  ORDER by id`

	// executes SQL query
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		log.Println("reached rows ", rows)
		return nil, err
	}
	defer rows.Close()

	// populate each row to configC
	var myPI []PayrollItem
	for rows.Next() {
		var pi PayrollItem
		err := rows.Scan(
			&pi.ID,
			&pi.Type,
			&pi.Code,
			&pi.Description,
			&pi.Start,
			&pi.End,
			&pi.Amount,
			&pi.PayEPF,
			&pi.PaySOCSO,
			&pi.PayHRDF,
			&pi.PayTax,
			&pi.IsFixed,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		myPI = append(myPI, pi)
	}

	// return all country rows
	return myPI, nil
}

// fetch all config tables
func (ct *ConfigTables) GetAllConfigTables() (*ConfigTables, error) {
	var err error
	ct.Country, err = getAllRowsCountryTable("CONFIG_COMMON_COUNTRY")
	if err != nil {
		return nil, err
	}
	ct.EmploymentBank, err = getAllRowsConfigTable("CONFIG_EMPLOYMENT_BANK")
	if err != nil {
		return nil, err
	}
	ct.EmploymentBranch, err = getAllRowsConfigTable("CONFIG_EMPLOYMENT_BRANCH")
	if err != nil {
		return nil, err
	}
	ct.EmploymentDepartment, err = getAllRowsConfigTable("CONFIG_EMPLOYMENT_DEPARTMENT")
	if err != nil {
		return nil, err
	}
	ct.EmploymentGroup, err = getAllRowsConfigTable("CONFIG_EMPLOYMENT_GROUP")
	if err != nil {
		return nil, err
	}
	ct.EmploymentOT, err = getAllRowsConfigTable("CONFIG_EMPLOYMENT_OVERTIME")
	if err != nil {
		return nil, err
	}
	ct.EmploymentPaymentBy, err = getAllRowsConfigTable("CONFIG_EMPLOYMENT_PAYMENT_BY")
	if err != nil {
		return nil, err
	}
	ct.EmploymentPayFrequency, err = getAllRowsConfigTable("CONFIG_EMPLOYMENT_PAY_FREQUENCY")
	if err != nil {
		return nil, err
	}
	ct.EmploymentProject, err = getAllRowsConfigTable("CONFIG_EMPLOYMENT_PROJECT")
	if err != nil {
		return nil, err
	}
	ct.EmploymentType, err = getAllRowsConfigTable("CONFIG_EMPLOYMENT_TYPE")
	if err != nil {
		return nil, err
	}
	ct.EmploymentWages, err = getAllRowsConfigTable("CONFIG_EMPLOYMENT_WAGES")
	if err != nil {
		return nil, err
	}
	ct.Race, err = getAllRowsConfigTable("CONFIG_COMMON_RACE")
	if err != nil {
		return nil, err
	}
	ct.Relationship, err = getAllRowsConfigTable("CONFIG_COMMON_RELATIONSHIP")
	if err != nil {
		return nil, err
	}
	ct.Religion, err = getAllRowsConfigTable("CONFIG_COMMON_RELIGION")
	if err != nil {
		return nil, err
	}
	ct.StatutoryEPF, err = getAllRowsConfigTable("CONFIG_STATUTORY_EPF")
	if err != nil {
		return nil, err
	}
	ct.StatutoryForeignLevy, err = getAllRowsConfigTable("CONFIG_STATUTORY_FOREIGN_LEVY")
	if err != nil {
		return nil, err
	}
	ct.StatutorySOCSOCategory, err = getAllRowsConfigTable("CONFIG_STATUTORY_SOCSO_CATEGORY")
	if err != nil {
		return nil, err
	}
	ct.StatutorySOCSOStatus, err = getAllRowsConfigTable("CONFIG_STATUTORY_SOCSO_STATUS")
	if err != nil {
		return nil, err
	}
	ct.StatutoryTaxBranch, err = getAllRowsConfigTable("CONFIG_STATUTORY_TAX_STATUS")
	if err != nil {
		return nil, err
	}
	ct.StatutoryTaxStatus, err = getAllRowsConfigTable("CONFIG_STATUTORY_TAX_BRANCH")
	if err != nil {
		return nil, err
	}
	ct.PayrollType, err = getAllRowsConfigTable("CONFIG_PAYROLL_TYPE")
	if err != nil {
		return nil, err
	}
	ct.PayrollItem, err = getAllRowsPayrollItem()
	if err != nil {
		return nil, err
	}

	return ct, nil
}

// inserts a new ct entry into the database, and returns the ID of the newly inserted row
func (ct *ConfigTables) Insert(table, value string) (int, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which insert a new entry to the config table
	stmt := `INSERT INTO public."` + table + `" ("name") VALUES($1) returning id;`

	// executes SQL query (set SQL parameters and cacth rowID)
	var newID int

	err := db.QueryRowContext(ctx, stmt, value).Scan(&newID)
	if err != nil {
		return 0, err
	}

	// return rowID of created new employee
	return newID, nil
}

// inserts a new ct entry into the database, and returns the ID of the newly inserted row
func (ct *ConfigTables) Update(table, value string, rowID int) (int, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which insert a new entry to the config table
	stmt := `UPDATE public."` + table + `" SET name = $1 WHERE id=$2;`

	// executes SQL query (set SQL parameters and cacth rowID)
	var newID int

	_, err := db.ExecContext(ctx, stmt, value, rowID)
	if err != nil {
		return 0, err
	}

	// return rowID of created new employee
	return newID, nil
}

// inserts a new ct entry into the database, and returns the ID of the newly inserted row
func (ct *ConfigTables) Delete(table, rowID string) error {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which insert a new entry to the config table
	stmt := `UPDATE public."` + table + `" SET soft_delete = 1 WHERE id=$1;`

	// executes SQL query (set SQL parameters)
	_, err := db.ExecContext(ctx, stmt, rowID)
	if err != nil {
		return err
	}

	return nil
}
