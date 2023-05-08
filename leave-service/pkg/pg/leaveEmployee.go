package pg

import (
	"context"
	"time"
)

// generate all entitled leaves to LEAVE_EMPLOYEE for new employee_id
func (l *Leave) CreateEntitled(eid, uid, genderid int) error {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// check which gender to skip: if male then skip female...
	skipGender := 1 // skip male by default
	if genderid == 1 {
		skipGender = 0 // if employee is male, then skip female
	}

	// SQL statement which fetch all available leave definition
	// since employee is new we only care about the lowest seniority
	stmt := `SELECT ldd.leave_definition_id  ,
					ldd.entitled ,
					ld.calculation_method_id 
			 FROM "LEAVE_DEFINITION_DETAILS" ldd, "LEAVE_DEFINITION" ld 
			 WHERE ldd.seniority < 2 
			 AND ldd.soft_delete = 0
			 AND ld.gender_id <> $1
			 AND ldd.leave_definition_id = ld.id `

	rows, err := db.QueryContext(ctx, stmt, skipGender)
	if err != nil {
		return err
	}
	defer rows.Close()

	// populate returned rows
	var all []*LeaveDefinitionEntitled
	for rows.Next() {
		var entitled LeaveDefinitionEntitled

		err = rows.Scan(
			&entitled.ID,
			&entitled.Entitled,
			&entitled.CalculationID,
		)
		if err != nil {
			return err
		}

		all = append(all, &entitled)
	}

	// insert all entitled leaves to EMPLOYEE_LEAVE
	// if calculation is 1 (not earned) entitled is straight
	// if calculation is 2 (earned) entitled 0
	for _, entry := range all {
		var entitled int
		switch entry.CalculationID {
		case 1:
			entitled = entry.Entitled
		default:
			entitled = 0
		}
		_, err := InsertEntitled(entry.ID, entitled, eid, uid)
		if err != nil {
			return err
		}
	}

	// return error
	return nil
}

// create new entitled
func InsertEntitled(definitionID, entitled, eid, uid int) (int, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which update an employee (soft delete)
	stmt := `INSERT INTO public."LEAVE_EMPLOYEE" (entitled, taken, credits, leave_definition_id, employee_id, soft_delete, created_at, created_by, updated_at, updated_by) 
			 VALUES($1, 0, 0, $2, $3, 0, now(), $4, now(), $5) returning id;`

	// store new row id
	var newID int

	// executes SQL query (set SQL parameters and cacth rowID)
	err := db.QueryRowContext(ctx, stmt,
		entitled,
		definitionID,
		eid,
		uid,
		uid,
	).Scan(&newID)
	if err != nil {
		return 0, err
	}

	// return error
	return newID, nil
}

// get all my entitled leaves
func (l *Leave) GetAllMyEntitledLeave(eid int) ([]*EntitledLeave, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch employee summary
	query := `SELECT le.id, 
					 le.entitled, 
					 le.credits, 
					 le.taken, 
					 le.leave_definition_id, 
					 ld.code as leave_code,
					 ld.description as leave_definition,
					 le.created_at, 
					 le.created_by, 
					 le.updated_at, 
					 le.updated_by
			  FROM public."LEAVE_EMPLOYEE" le, public."LEAVE_DEFINITION" ld
			  WHERE le.soft_delete = 0
			  AND le.employee_id = $1
			  AND le.leave_definition_id = ld.id
			  ORDER BY le.id;`

	// executes SQL query
	rows, err := db.QueryContext(ctx, query, eid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// populate returned rows to employee summary struct
	var all []*EntitledLeave
	for rows.Next() {
		var myLeaves EntitledLeave

		err = rows.Scan(
			&myLeaves.ID,
			&myLeaves.Entitled,
			&myLeaves.Credits,
			&myLeaves.Taken,
			&myLeaves.LeaveDefinitionID,
			&myLeaves.LeaveDefinitionCode,
			&myLeaves.LeaveDefinitionName,
			&myLeaves.CreatedAt,
			&myLeaves.CreatedBy,
			&myLeaves.UpdatedAt,
			&myLeaves.UpdatedBy,
		)
		if err != nil {
			return nil, err
		}

		all = append(all, &myLeaves)
	}

	// return all employee summary rows
	return all, nil
}

// increment taken number days of leave id for employee id (when application is requested)
func (l *Leave) IncrementTakenLeave(lid, eid int, total float32) error {
	now := time.Now().Format("2006-01-02")

	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which update an employee
	stmt := `UPDATE "LEAVE_EMPLOYEE" 
			 SET taken = taken + $1, updated_at = $2, updated_by = $3
			 WHERE leave_definition_id = $4
			 AND employee_id = $5`

	// executes SQL query
	_, err := db.ExecContext(ctx, stmt, total, now, eid, lid, eid)
	if err != nil {
		return err
	}
	return nil
}

// decrement taken number days of leave id for employee id (when application rejected)
func (l *Leave) DecrementTakenLeave(lid, eid int, total float32) error {
	now := time.Now().Format("2006-01-02")

	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which update an employee
	stmt := `UPDATE "LEAVE_EMPLOYEE" 
			 SET taken = taken - $1, updated_at = $2, updated_by = $3
			 WHERE leave_definition_id = $4
			 AND employee_id = $5`

	// executes SQL query
	_, err := db.ExecContext(ctx, stmt, total, now, eid, lid, eid)
	if err != nil {
		return err
	}
	return nil
}

// update credits of one specific employee leave by rowID
func (l *Leave) UpdateCredits(rowID, eid int, credits float32) error {
	now := time.Now().Format("2006-01-02")

	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which update an employee
	stmt := `UPDATE "LEAVE_EMPLOYEE" 
			 SET credits = $1, updated_at = $2, updated_by = $3
			 WHERE id = $4`

	// executes SQL query
	_, err := db.ExecContext(ctx, stmt, credits, now, eid, rowID)
	if err != nil {
		return err
	}
	return nil
}
