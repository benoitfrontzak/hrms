package pg

import "context"

// soft delete employee by id
func (e *Employee) SoftDeleteEmployeeByID(id int) error {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which update an employee (soft delete)
	stmt := `UPDATE public."EMPLOYEE" SET soft_delete=1 WHERE id=$1;`

	// executes SQL query
	_, err := db.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	// return error
	return nil
}
