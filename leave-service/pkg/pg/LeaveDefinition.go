package pg

import (
	"context"
	"strconv"
	"time"
)

// get all leave definitions
func (ld *LeaveDefinition) GetLeaveDefinition() ([]*LeaveDefinition, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch leave definition
	query := `SELECT ld.id, 
					 ld.code,  
					 ld.description,
					 ld.is_unpaid,
					 ld.is_encashable,
					 ld.replacement_required,
					 ld.doc_required, 
					 ld.expiry_month, 
					 ld.gender_id, 
					 cg.name as gender, 
					 ld.limitation_id, 
					 cl.name as limitation, 
					 ld.calculation_method_id, 
					 ccm.name as calculation_method, 
					 ld.soft_delete, 
					 ld.created_at, 
					 ld.created_by, 
					 ld.updated_at, 
					 ld.updated_by 
			  FROM public."LEAVE_DEFINITION" ld, public."CONFIG_GENDER" cg, public."CONFIG_LIMITATION" cl, public."CONFIG_CALCULATION_METHOD" ccm
			  WHERE ld.soft_delete = 0
			  AND ld.gender_id = cg.id
			  AND ld.limitation_id = cl.id
			  AND ld.calculation_method_id = ccm.id
			  ORDER BY ld.id;`

	// executes SQL query
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// populate returned rows to leave definition struct
	var all []*LeaveDefinition
	for rows.Next() {
		var def LeaveDefinition
		err := rows.Scan(
			&def.ID,
			&def.Code,
			&def.Description,
			&def.Unpaid,
			&def.Encashable,
			&def.ReplacementRequired,
			&def.DocRequired,
			&def.Expiry,
			&def.GenderID,
			&def.Gender,
			&def.LimitationID,
			&def.Limitation,
			&def.CalculationID,
			&def.Calculation,
			&def.SoftDelete,
			&def.CreatedAt,
			&def.CreatedBy,
			&def.UpdatedAt,
			&def.UpdatedBy,
		)
		if err != nil {
			return nil, err
		}

		// fetch all leave definition details
		details, err := GetLeaveDefinitionDetails(def.ID)
		if err != nil {
			return nil, err
		}
		def.Details = details

		all = append(all, &def)
	}

	// return all leave definition rows
	return all, nil
}

// get one leave definition by id
func (ld *LeaveDefinition) GetLeaveDefinitionID(lid int) (*LeaveDefinition, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch leave definition
	query := `SELECT ld.id, 
					 ld.code,  
					 ld.description,
					 ld.is_unpaid,
					 ld.is_encashable,
					 ld.replacement_required,
					 ld.doc_required, 
					 ld.expiry_month, 
					 ld.gender_id, 
					 cg.name as gender, 
					 ld.limitation_id, 
					 cl.name as limitation, 
					 ld.calculation_method_id, 
					 ccm.name as calculation_method, 
					 ld.soft_delete, 
					 ld.created_at, 
					 ld.created_by, 
					 ld.updated_at, 
					 ld.updated_by 
			  FROM public."LEAVE_DEFINITION" ld, public."CONFIG_GENDER" cg, public."CONFIG_LIMITATION" cl, public."CONFIG_CALCULATION_METHOD" ccm
			  WHERE ld.id = $1 
			  AND ld.soft_delete = 0
			  AND ld.gender_id = cg.id
			  AND ld.limitation_id = cl.id
			  AND ld.calculation_method_id = ccm.id
			  ORDER BY ld.id;`

	// executes SQL query
	row := db.QueryRowContext(ctx, query, lid)

	// populate returned rows to leave definition struct
	var def LeaveDefinition

	err := row.Scan(
		&def.ID,
		&def.Code,
		&def.Description,
		&def.Unpaid,
		&def.Encashable,
		&def.ReplacementRequired,
		&def.DocRequired,
		&def.Expiry,
		&def.GenderID,
		&def.Gender,
		&def.LimitationID,
		&def.Limitation,
		&def.CalculationID,
		&def.Calculation,
		&def.SoftDelete,
		&def.CreatedAt,
		&def.CreatedBy,
		&def.UpdatedAt,
		&def.UpdatedBy,
	)
	if err != nil {
		return nil, err
	}

	// fetch all leave definition details
	details, err := GetLeaveDefinitionDetails(lid)
	if err != nil {
		return nil, err
	}
	def.Details = details

	// return all leave definition rows
	return &def, nil
}

// get all leave definition details
func GetLeaveDefinitionDetails(id int) ([]*LeaveDefinitionDetails, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch leave definition details
	query := `SELECT ldc.id, 
					 ldc.seniority,  
					 ldc.entitled,
					 ldc.leave_definition_id, 
					 ldc.soft_delete, 
					 ldc.created_at, 
					 ldc.created_by, 
					 ldc.updated_at, 
					 ldc.updated_by 
			  FROM public."LEAVE_DEFINITION_DETAILS" ldc
			  WHERE ldc.soft_delete = 0
			  AND ldc.leave_definition_id = $1
			  ORDER BY ldc.id;`

	// executes SQL query
	rows, err := db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// populate returned rows to leave definition details struct
	var all []*LeaveDefinitionDetails
	for rows.Next() {
		var def LeaveDefinitionDetails
		err := rows.Scan(
			&def.ID,
			&def.Seniority,
			&def.Entitled,
			&def.LeaveDefinitionID,
			&def.SoftDelete,
			&def.CreatedAt,
			&def.CreatedBy,
			&def.UpdatedAt,
			&def.UpdatedBy,
		)
		if err != nil {
			return nil, err
		}

		all = append(all, &def)
	}

	// return all leave definition details rows
	return all, nil
}

// insert a new leave definition & all leave definition details relative
// update new leave definition to leave employee
func (ld *LeaveDefinition) Insert() (int, error) {
	// insert new leave definition and fetch rowID inserted
	rowID, err := ld.insertLeaveDefinition()
	if err != nil {
		return 0, err
	}

	// insert all leave definition details
	for _, entry := range ld.Details {
		_, err = insertLeaveDefinitionDetails(entry.Seniority, entry.Entitled, ld.CreatedBy, rowID)
		if err != nil {
			return 0, err
		}
	}

	return rowID, nil
}

// insert leave definition
func (ld *LeaveDefinition) insertLeaveDefinition() (int, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which insert a new leave definition
	stmt := `INSERT INTO public."LEAVE_DEFINITION" (code, description, is_unpaid, is_encashable, replacement_required, doc_required, expiry_month, gender_id, limitation_id, calculation_method_id, created_by, updated_by) 
			 VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) returning id;`

	// executes SQL query (set SQL parameters and cacth rowID)
	var newID int

	err := db.QueryRowContext(ctx, stmt,
		ld.Code,
		ld.Description,
		ld.Unpaid,
		ld.Encashable,
		ld.ReplacementRequired,
		ld.DocRequired,
		ld.Expiry,
		ld.GenderID,
		ld.LimitationID,
		ld.CalculationID,
		ld.CreatedBy,
		ld.UpdatedBy,
	).Scan(&newID)
	if err != nil {
		return 0, err
	}

	// return rowID of created new employee
	return newID, nil
}

// insert leave definition
func insertLeaveDefinitionDetails(seniority, entitled, user, rowID int) (int, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which insert a new leave definition
	stmt := `INSERT INTO public."LEAVE_DEFINITION_DETAILS" (seniority, entitled, leave_definition_id, created_by, updated_by)
			 VALUES($1, $2, $3, $4, $5) returning id;`

	// executes SQL query (set SQL parameters and cacth rowID)
	var newID int

	err := db.QueryRowContext(ctx, stmt,
		seniority,
		entitled,
		rowID,
		user,
		user,
	).Scan(&newID)
	if err != nil {
		return 0, err
	}

	// return rowID of created new employee
	return newID, nil
}

// update leave definition by id
func (ld *LeaveDefinition) Update() error {
	// update leave definition
	err := ld.updateLeaveDefinition()
	if err != nil {
		return err
	}

	// update all leave definition details
	for _, entry := range ld.Details {
		err = updateLeaveDefinitionDetails(entry.Seniority, entry.Entitled, ld.CreatedBy, entry.ID)
		if err != nil {
			return err
		}
	}

	// delete all leave definition details
	if len(ld.DetailsDeleted) > 0 {
		for _, entry := range ld.DetailsDeleted {
			rowID, err := strconv.Atoi(entry)
			if err != nil {
				return err
			}
			err = deleteLeaveDefinitionDetails(rowID)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (ld *LeaveDefinition) updateLeaveDefinition() error {
	now := time.Now().Format("2006-01-02")
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which update an employee (soft delete)
	stmt := `UPDATE public."LEAVE_DEFINITION" 
			 SET code=$1, 
			 	 description=$2, 
				 is_unpaid=$3, 
				 is_encashable=$4, 
				 replacement_required=$5, 
				 doc_required=$6, 
				 expiry_month=$7, 
				 gender_id=$8, 
				 limitation_id=$9, 
				 calculation_method_id=$10, 
				 updated_at=$11, 
				 updated_by=$12
			 WHERE id=$13;`

	// executes SQL query
	_, err := db.ExecContext(ctx, stmt,
		ld.Code,
		ld.Description,
		ld.Unpaid,
		ld.Encashable,
		ld.ReplacementRequired,
		ld.DocRequired,
		ld.Expiry,
		ld.GenderID,
		ld.LimitationID,
		ld.CalculationID,
		now,
		ld.UpdatedBy,
		ld.ID,
	)
	if err != nil {
		return err
	}

	// return error
	return nil
}
func updateLeaveDefinitionDetails(seniority, entitled, user, rowID int) error {
	now := time.Now().Format("2006-01-02")
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which update an employee (soft delete)
	stmt := `UPDATE public."LEAVE_DEFINITION_DETAILS" 
			 SET seniority=$1, 
			 	 entitled=$2, 
				 updated_at=$3, 
				 updated_by=$4
			 WHERE id=$5;`

	// executes SQL query
	_, err := db.ExecContext(ctx, stmt,
		seniority,
		entitled,
		now,
		user,
		rowID,
	)
	if err != nil {
		return err
	}

	// return error
	return nil
}
func deleteLeaveDefinitionDetails(id int) error {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which update an employee (soft delete)
	stmt := `UPDATE public."LEAVE_DEFINITION_DETAILS" SET soft_delete=1 WHERE id=$1;`

	// executes SQL query
	_, err := db.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	// return error
	return nil
}

// soft delete claim definition by id
func (ld *LeaveDefinition) Delete(id int) error {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which update an employee (soft delete)
	stmt := `UPDATE public."LEAVE_DEFINITION" SET soft_delete=1 WHERE id=$1;`

	// executes SQL query
	_, err := db.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	// return error
	return nil
}

// create the entitlement for a specific leave definition id of an employee id
func (ld *LeaveDefinition) CreateEntitledByDefinition(lid, eid int, entitled float64) error {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which update an employee (soft delete)
	stmt := `INSERT INTO public."LEAVE_EMPLOYEE" (entitled, taken, credits, leave_definition_id, employee_id, soft_delete, created_at, created_by, updated_at, updated_by) 
			 VALUES($1, 0, 0, $2, $3, 0, now(), 0, now(), 0)returning id;`

	// executes SQL query
	_, err := db.ExecContext(ctx, stmt, entitled, lid, eid)
	if err != nil {
		return err
	}
	return nil
}
