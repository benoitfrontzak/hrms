package pg

import "context"

// inserts a new payroll item for a specific employee into the database, and returns the ID of the newly inserted row
func (e *Employee) InsertNewPI(p PayrollItem) (int, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which insert a new employee
	stmt := `INSERT INTO public."CONFIG_PAYROLL_ITEM" (payroll_type_id, code, description, start_period, end_period, amount, pay_epf, pay_socso_eif, pay_hrdf, pay_tax, is_fixed, soft_delete) 
			 VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, 0) returning id;`

	// executes SQL query (set SQL parameters and cacth rowID)
	var newID int

	err := db.QueryRowContext(ctx, stmt,
		p.Type,
		p.Code,
		p.Description,
		p.Start,
		p.End,
		p.Amount,
		p.PayEPF,
		p.PaySOCSO,
		p.PayHRDF,
		p.PayTax,
		p.IsFixed,
	).Scan(&newID)
	if err != nil {
		return 0, err
	}

	// return rowID of created new employee
	return newID, nil
}

// inserts a new payroll item for a specific employee into the database, and returns the ID of the newly inserted row
func (e *Employee) UpdatePI(p PayrollItem) error {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which insert a new employee
	stmt := `UPDATE public."CONFIG_PAYROLL_ITEM" 
			 SET code=$1, 
				 description=$2, 
				 start_period=$3, 
				 end_period=$4, 
				 amount=$5
			 WHERE id=$6;`

	_, err := db.ExecContext(ctx, stmt,
		p.Code,
		p.Description,
		p.Start,
		p.End,
		p.Amount,
		p.ID,
	)
	if err != nil {
		return err
	}

	// return rowID of created new employee
	return nil
}
