package pg

import (
	"context"
	"log"
	"time"
)

// generate all entitled leaves to LEAVE_EMPLOYEE for employee_id
func (l *Leave) CreateEntitled(eid, uid int) error {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch all available leave definition
	// since employee is new we only care about the lowest seniority
	stmt := `SELECT ldd.leave_definition_id  ,
					ldd.entitled ,
					ld.calculation_method_id 
			 FROM "LEAVE_DEFINITION_DETAILS" ldd, "LEAVE_DEFINITION" ld 
			 WHERE ldd.seniority < 2 
			 AND ldd.soft_delete = 0
			 AND ldd.leave_definition_id = ld.id `

	rows, err := db.QueryContext(ctx, stmt)
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
	// if calculation is 1,2,3 (not earned) entitled is straight
	// if calculation is 4,5,6 (earned) entitled 0
	for _, entry := range all {
		var entitled int
		switch entry.CalculationID {
		case 1, 2, 3:
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

// create new leave
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
		log.Println("entry:", entry)
		_, err := insertLeaveDetails(entry.RequestedDate, entry.IsHalf, newID, l.EmployeeID)
		if err != nil {
			return 0, err
		}
	}
	// return error
	return newID, nil
}

// insert leave details (requested dates)
func insertLeaveDetails(rd string, half, id, user int) (int, error) {
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

// update claim by id (TODO!)
func (l *Leave) Update(id int) error {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which update an employee (soft delete)
	stmt := `UPDATE public."CLAIM" SET soft_delete=1 WHERE id=$1;`

	// executes SQL query
	_, err := db.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	// return error
	return nil
}

// approve claim by id
func (l *Leave) Approve(rowID, userID int) error {
	now := time.Now().Format("2006-01-02")

	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which update the claim (approve it)
	stmt := `UPDATE public."LEAVE_APPLICATION" SET status_id=4, approved_at=$1, approved_by=$2 WHERE id=$3;`

	// executes SQL query
	_, err := db.ExecContext(ctx, stmt, now, userID, rowID)
	if err != nil {
		return err
	}

	// return error
	return nil
}

// reject claim by id
func (l *Leave) Reject(rowID, userID int, reason string) error {
	now := time.Now().Format("2006-01-02")

	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which update the claim (approve it)
	stmt := `UPDATE public."LEAVE_APPLICATION" SET status_id=3, approved_at=$1, approved_by=$2, rejected_reason=$3 WHERE id=$4;`

	// executes SQL query
	_, err := db.ExecContext(ctx, stmt, now, userID, reason, rowID)
	if err != nil {
		return err
	}

	// return error
	return nil
}

// soft delete claim by id
func (l *Leave) Delete(id int) error {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which update an employee (soft delete)
	stmt := `UPDATE public."LEAVE_APPLICATION" SET soft_delete=1 WHERE id=$1;`

	// executes SQL query
	_, err := db.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	// return error
	return nil
}

// get all my leaves
func (l *Leave) GetAllMyEntitledLeave(eid int) ([]*EntitledLeave, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch employee summary
	query := `SELECT le.id, 
					 le.entitled, 
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

// get all my leaves
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
		details, err := getAllMyLeaveDetails(myLeaves.ID)
		if err != nil {
			return nil, err
		}
		myLeaves.Details = details

		all = append(all, &myLeaves)
	}

	// return all employee summary rows
	return all, nil
}

// get all leaves by date
func (l *Leave) GetAllLeaveToday() (*AllCurentLeave, error) {
	// set today's & tomorrow's date
	n := time.Now()
	t := n.AddDate(0, 0, 1)

	// fetch today's leaves
	today, err := getLeaveByDate(n.Format("2006-01-02"))
	if err != nil {
		return nil, err
	}

	// fetch tomorrow's leave
	tomorrow, err := getLeaveByDate(t.Format("2006-01-02"))
	if err != nil {
		return nil, err
	}

	// store all information
	all := AllCurentLeave{
		Today:    today,
		Tomorrow: tomorrow,
	}

	return &all, nil
}

// get all my leaves
func getAllMyLeaveDetails(id int) ([]*LeaveDetails, error) {
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

// get all leave by status
func (c *Leave) GetAllLeave() (*AllLeave, error) {

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

// helpers
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
		details, err := getAllMyLeaveDetails(myLeaves.ID)
		if err != nil {
			return nil, err
		}
		myLeaves.Details = details

		all = append(all, &myLeaves)
	}

	// return all employee summary rows
	return all, nil
}

func getLeaveByDate(myDate string) ([]*ApprovedLeave, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch employee summary
	query := `SELECT ld.code as leave_code,
					 ld.description as leave_definition,
   					 la.description,
   					 la.employee_id ,
   					 lad.requested_date, 
   					 lad.is_half 
			  FROM public."LEAVE_DEFINITION" ld, public."LEAVE_APPLICATION" la,  public."LEAVE_APPLICATION_DETAILS" lad 
			  WHERE lad.requested_date = $1
			  AND lad.leave_application_id = la.id
			  AND la.leave_definition_id = ld.id 
			  AND la.status_id = 4`

	// executes SQL query
	rows, err := db.QueryContext(ctx, query, myDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// populate returned rows to employee summary struct
	var all []*ApprovedLeave
	for rows.Next() {
		var myLeaves ApprovedLeave

		err = rows.Scan(
			&myLeaves.Code,
			&myLeaves.Description,
			&myLeaves.Reason,
			&myLeaves.EmployeeID,
			&myLeaves.RequestedDate,
			&myLeaves.IsHalf,
		)
		if err != nil {
			return nil, err
		}

		all = append(all, &myLeaves)
	}

	// return all employee summary rows
	return all, nil
}
