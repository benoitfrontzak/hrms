package pg

import (
	"context"
	"time"
)

// create new leave application
func (l *Leave) Insert() (int, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which update an employee (soft delete)
	stmt := `INSERT INTO public."LEAVE_APPLICATION" (leave_definition_id, 
													 description, 
													 status_id, 
													 rejected_reason, 
													 approved_at, 
													 approved_by, 
													 employee_id,
													 created_by, 
													 updated_by) 
			 VALUES($1, $2, 2, '', '0001-01-01', 0, $3, $4, $5) returning id;`

	// store new row id
	var newID int

	// executes SQL query (set SQL parameters and cacth rowID)
	err := db.QueryRowContext(ctx, stmt,
		l.LeaveDefinitionID,
		l.Description,
		l.EmployeeID,
		l.EmployeeID,
		l.EmployeeID,
	).Scan(&newID)
	if err != nil {
		return 0, err
	}

	// insert all requested dates to leave details
	for _, entry := range l.Details {
		_, err := insertRequestedDates(entry.RequestedDate, entry.IsHalf, newID, l.EmployeeID)
		if err != nil {
			return 0, err
		}
	}
	// return error
	return newID, nil
}

// insert requested dates to leave details
func insertRequestedDates(rd string, half, id, user int) (int, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which update an employee (soft delete)
	stmt := `INSERT INTO public."LEAVE_APPLICATION_DETAILS" (requested_date, 
															 is_half, 
															 leave_application_id, 
													 		 created_by, 
													 		 updated_by) 
			 VALUES($1, $2, $3, $4, $5) returning id;`

	// store new row id
	var newID int

	// executes SQL query (set SQL parameters and cacth rowID)
	err := db.QueryRowContext(ctx, stmt,
		rd,
		half,
		id,
		user,
		user,
	).Scan(&newID)
	if err != nil {
		return 0, err
	}

	return newID, nil
}

// soft delete leave by id
func (l *Leave) Delete(id, uid int) error {
	now := time.Now().Format("2006-01-02")

	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which update an employee (soft delete)
	stmt := `UPDATE public."LEAVE_APPLICATION" SET soft_delete=1, updated_at=$1, updated_by=$2 WHERE id=$3;`

	// executes SQL query
	_, err := db.ExecContext(ctx, stmt, now, uid, id)
	if err != nil {
		return err
	}

	// return error
	return nil
}

// get leave definition id of leave application id
func (l *Leave) GetLeaveDefinitionID(applicationID int) (int, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch leave definition
	// numberOfDays = total day requested - (sum(HalfDays) / 2)
	// if half day is checked we store 1...
	query := `SELECT leave_definition_id
			  FROM "LEAVE_APPLICATION" la
			  WHERE la.id = $1`

	// executes SQL query
	row := db.QueryRowContext(ctx, query, applicationID)

	// populate returned rows to leave definition struct
	var id int

	err := row.Scan(
		&id,
	)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// get all leaves of employee id
func (l *Leave) GetAllMyLeave(eid int) ([]*Leave, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch employee summary
	query := `SELECT la.id, 
					 la.leave_definition_id, 
					 ld.code as leave_code,
					 ld.description as leave_definition, 
					 la.description, 
					 la.status_id, 
					 cs.name as status_id, 
					 la.rejected_reason, 
					 la.approved_at, 
					 la.approved_by,
					 la.created_at, 
					 la.created_by, 
					 la.updated_at, 
					 la.updated_by
			  FROM public."LEAVE_APPLICATION" la, public."LEAVE_DEFINITION" ld, public."CONFIG_STATUS" cs
			  WHERE la.soft_delete = 0
			  AND la.employee_id = $1
			  AND la.leave_definition_id = ld.id
			  AND la.status_id = cs.id
			  ORDER BY la.id;`

	// executes SQL query
	rows, err := db.QueryContext(ctx, query, eid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// populate returned rows to employee summary struct
	var all []*Leave
	for rows.Next() {
		var myLeaves Leave

		err = rows.Scan(
			&myLeaves.ID,
			&myLeaves.LeaveDefinitionID,
			&myLeaves.LeaveDefinitionCode,
			&myLeaves.LeaveDefinitionName,
			&myLeaves.Description,
			&myLeaves.StatusID,
			&myLeaves.Status,
			&myLeaves.RejectedReason,
			&myLeaves.ApprovedAt,
			&myLeaves.ApprovedBy,
			&myLeaves.CreatedAt,
			&myLeaves.CreatedBy,
			&myLeaves.UpdatedAt,
			&myLeaves.UpdatedBy,
		)
		if err != nil {
			return nil, err
		}

		// fetch leave details
		details, err := getAllMyRequestedDates(myLeaves.ID)
		if err != nil {
			return nil, err
		}
		myLeaves.Details = details

		all = append(all, &myLeaves)
	}

	// return all employee summary rows
	return all, nil
}

// get all leaves by status (pending | approved | rejected)
func (l *Leave) GetAllLeaveManager(list string) (*AllLeave, error) {

	// status id
	// 0	not defined
	// 1	draft
	// 2	pending
	// 3	rejected
	// 4	approved

	var all AllLeave

	approved, err := getAllLeaveManagerByStatus(4, list)
	if err != nil {
		return nil, err
	}
	pending, err := getAllLeaveManagerByStatus(2, list)
	if err != nil {
		return nil, err
	}
	rejected, err := getAllLeaveManagerByStatus(3, list)
	if err != nil {
		return nil, err
	}

	all.Approved = approved
	all.Pending = pending
	all.Rejected = rejected

	return &all, nil
}

// get all leaves by status (pending | approved | rejected)
func (l *Leave) GetAllLeaveUser(eid int) (*AllLeave, error) {

	// status id
	// 0	not defined
	// 1	draft
	// 2	pending
	// 3	rejected
	// 4	approved

	var all AllLeave

	approved, err := getAllLeaveUserByStatus(4, eid)
	if err != nil {
		return nil, err
	}
	pending, err := getAllLeaveUserByStatus(2, eid)
	if err != nil {
		return nil, err
	}
	rejected, err := getAllLeaveUserByStatus(3, eid)
	if err != nil {
		return nil, err
	}

	all.Approved = approved
	all.Pending = pending
	all.Rejected = rejected

	return &all, nil
}

// get all leaves by status (pending | approved | rejected)
func (l *Leave) GetAllLeave() (*AllLeave, error) {

	// status id
	// 0	not defined
	// 1	draft
	// 2	pending
	// 3	rejected
	// 4	approved

	var all AllLeave

	approved, err := getAllLeaveByStatus(4)
	if err != nil {
		return nil, err
	}
	pending, err := getAllLeaveByStatus(2)
	if err != nil {
		return nil, err
	}
	rejected, err := getAllLeaveByStatus(3)
	if err != nil {
		return nil, err
	}

	all.Approved = approved
	all.Pending = pending
	all.Rejected = rejected

	return &all, nil
}

// calculate number of requested dates of a leave id (minus half days)
func (l *Leave) GetRequestedDatesNumber(lid int) (float32, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch leave definition
	// numberOfDays = total day requested - (sum(HalfDays) / 2)
	// if half day is checked we store 1...
	query := `SELECT count(id) - (cast(sum(is_half) as decimal)/2) as numberOfDays
			  FROM "LEAVE_APPLICATION_DETAILS" lad 
			  WHERE lad.leave_application_id = $1`

	// executes SQL query
	row := db.QueryRowContext(ctx, query, lid)

	// populate returned rows to leave definition struct
	var numb float32

	err := row.Scan(
		&numb,
	)
	if err != nil {
		return 0, err
	}

	return numb, nil
}

/*
helpers
*/
// get all leaves by status
func getAllLeaveManagerByStatus(status int, list string) ([]*Leave, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch employee summary
	query := `SELECT la.id, 
					 la.leave_definition_id, 
					 ld.code as leave_code,
					 ld.description as leave_definition, 
					 la.description, 
					 la.status_id, 
					 cs.name as status_id, 
					 la.rejected_reason, 
					 la.approved_at, 
					 la.approved_by,
					 la.employee_id,
					 la.created_at, 
					 la.created_by, 
					 la.updated_at, 
					 la.updated_by
			 FROM public."LEAVE_APPLICATION" la, public."LEAVE_DEFINITION" ld, public."CONFIG_STATUS" cs
			 WHERE la.soft_delete = 0
			 AND la.status_id = $1
			 AND la.leave_definition_id = ld.id
			 AND la.status_id = cs.id
			 AND la.employee_id IN (` + list + `)
			 ORDER BY la.id;`

	// executes SQL query
	rows, err := db.QueryContext(ctx, query, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// populate returned rows to employee summary struct
	var all []*Leave
	for rows.Next() {
		var myLeaves Leave
		err := rows.Scan(
			&myLeaves.ID,
			&myLeaves.LeaveDefinitionID,
			&myLeaves.LeaveDefinitionCode,
			&myLeaves.LeaveDefinitionName,
			&myLeaves.Description,
			&myLeaves.StatusID,
			&myLeaves.Status,
			&myLeaves.RejectedReason,
			&myLeaves.ApprovedAt,
			&myLeaves.ApprovedBy,
			&myLeaves.EmployeeID,
			&myLeaves.CreatedAt,
			&myLeaves.CreatedBy,
			&myLeaves.UpdatedAt,
			&myLeaves.UpdatedBy,
		)
		if err != nil {
			return nil, err
		}

		// fetch leave details
		details, err := getAllMyRequestedDates(myLeaves.ID)
		if err != nil {
			return nil, err
		}
		myLeaves.Details = details

		all = append(all, &myLeaves)
	}

	// return all employee summary rows
	return all, nil
}

// get all leaves by status
func getAllLeaveUserByStatus(status, eid int) ([]*Leave, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch employee summary
	query := `SELECT la.id, 
					 la.leave_definition_id, 
					 ld.code as leave_code,
					 ld.description as leave_definition, 
					 la.description, 
					 la.status_id, 
					 cs.name as status_id, 
					 la.rejected_reason, 
					 la.approved_at, 
					 la.approved_by,
					 la.employee_id,
					 la.created_at, 
					 la.created_by, 
					 la.updated_at, 
					 la.updated_by
			 FROM public."LEAVE_APPLICATION" la, public."LEAVE_DEFINITION" ld, public."CONFIG_STATUS" cs
			 WHERE la.soft_delete = 0
			 AND la.status_id = $1
			 AND la.leave_definition_id = ld.id
			 AND la.status_id = cs.id
			 AND la.employee_id = $2
			 ORDER BY la.id;`

	// executes SQL query
	rows, err := db.QueryContext(ctx, query, status, eid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// populate returned rows to employee summary struct
	var all []*Leave
	for rows.Next() {
		var myLeaves Leave
		err := rows.Scan(
			&myLeaves.ID,
			&myLeaves.LeaveDefinitionID,
			&myLeaves.LeaveDefinitionCode,
			&myLeaves.LeaveDefinitionName,
			&myLeaves.Description,
			&myLeaves.StatusID,
			&myLeaves.Status,
			&myLeaves.RejectedReason,
			&myLeaves.ApprovedAt,
			&myLeaves.ApprovedBy,
			&myLeaves.EmployeeID,
			&myLeaves.CreatedAt,
			&myLeaves.CreatedBy,
			&myLeaves.UpdatedAt,
			&myLeaves.UpdatedBy,
		)
		if err != nil {
			return nil, err
		}

		// fetch leave details
		details, err := getAllMyRequestedDates(myLeaves.ID)
		if err != nil {
			return nil, err
		}
		myLeaves.Details = details

		all = append(all, &myLeaves)
	}

	// return all employee summary rows
	return all, nil
}

// get all leaves by status
func getAllLeaveByStatus(status int) ([]*Leave, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch employee summary
	query := `SELECT la.id, 
					 la.leave_definition_id, 
					 ld.code as leave_code,
					 ld.description as leave_definition, 
					 la.description, 
					 la.status_id, 
					 cs.name as status_id, 
					 la.rejected_reason, 
					 la.approved_at, 
					 la.approved_by,
					 la.employee_id,
					 la.created_at, 
					 la.created_by, 
					 la.updated_at, 
					 la.updated_by
			 FROM public."LEAVE_APPLICATION" la, public."LEAVE_DEFINITION" ld, public."CONFIG_STATUS" cs
			 WHERE la.soft_delete = 0
			 AND la.status_id = $1
			 AND la.leave_definition_id = ld.id
			 AND la.status_id = cs.id
			 ORDER BY la.id;`

	// executes SQL query
	rows, err := db.QueryContext(ctx, query, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// populate returned rows to employee summary struct
	var all []*Leave
	for rows.Next() {
		var myLeaves Leave
		err := rows.Scan(
			&myLeaves.ID,
			&myLeaves.LeaveDefinitionID,
			&myLeaves.LeaveDefinitionCode,
			&myLeaves.LeaveDefinitionName,
			&myLeaves.Description,
			&myLeaves.StatusID,
			&myLeaves.Status,
			&myLeaves.RejectedReason,
			&myLeaves.ApprovedAt,
			&myLeaves.ApprovedBy,
			&myLeaves.EmployeeID,
			&myLeaves.CreatedAt,
			&myLeaves.CreatedBy,
			&myLeaves.UpdatedAt,
			&myLeaves.UpdatedBy,
		)
		if err != nil {
			return nil, err
		}

		// fetch leave details
		details, err := getAllMyRequestedDates(myLeaves.ID)
		if err != nil {
			return nil, err
		}
		myLeaves.Details = details

		all = append(all, &myLeaves)
	}

	// return all employee summary rows
	return all, nil
}

// get all requested dates of leave id from leaves details
func getAllMyRequestedDates(id int) ([]*LeaveDetails, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch employee summary
	query := `SELECT lad.id, 
					 lad.requested_date, 
					 lad.is_half, 
					 lad.leave_application_id, 
					 lad.created_at, 
					 lad.created_by, 
					 lad.updated_at, 
					 lad.updated_by
			  FROM public."LEAVE_APPLICATION_DETAILS" lad
			  WHERE lad.soft_delete = 0
			  AND lad.leave_application_id = $1
			  ORDER BY lad.id;`

	// executes SQL query
	rows, err := db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// populate returned rows to employee summary struct
	var all []*LeaveDetails
	for rows.Next() {
		var myLD LeaveDetails
		err := rows.Scan(
			&myLD.ID,
			&myLD.RequestedDate,
			&myLD.IsHalf,
			&myLD.LeaveID,
			&myLD.CreatedAt,
			&myLD.CreatedBy,
			&myLD.UpdatedAt,
			&myLD.UpdatedBy,
		)
		if err != nil {
			return nil, err
		}

		all = append(all, &myLD)
	}

	// return all employee summary rows
	return all, nil
}
