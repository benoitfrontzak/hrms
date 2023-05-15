package pg

import (
	"context"
	"log"
	"time"
)

// helper for create new claim (set default value)
func createMyDefaultValue() claimDefaultValue {

	return claimDefaultValue{
		StatusID:       2,
		ApprovedAt:     "0001-01-01",
		ApprovedBy:     0,
		ApprovedAmount: 0,
		RejectedReason: "",
		SoftDelete:     0,
	}
}

// create new claim
func (c *Claim) Insert() (int, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which update an employee (soft delete)
	stmt := `INSERT INTO public."CLAIM_APPLICATION" (claim_definition_id, 
										 description, 
										 amount, 
										 status_id, 
										 approved_at, 
										 approved_by, 
										 approved_amount, 
										 rejected_reason, 
										 employee_id, 
										 soft_delete,
										 created_by,
										 updated_by) 
			 VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, 0, $10, $11) returning id;`

	// store new row id
	var newID int

	// get claim default value
	cdv := createMyDefaultValue()

	// executes SQL query (set SQL parameters and cacth rowID)
	err := db.QueryRowContext(ctx, stmt,
		c.ClaimDefinition,
		c.Description,
		c.Amount,
		cdv.StatusID,
		cdv.ApprovedAt,
		cdv.ApprovedBy,
		cdv.ApprovedAmount,
		cdv.RejectedReason,
		c.EmployeeID,
		c.CreatedBy,
		c.UpdatedBy,
	).Scan(&newID)
	if err != nil {
		return 0, err
	}

	// return error
	return newID, nil
}

// update claim by id (TODO!)
func (c *Claim) Update(id int) error {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which update an employee (soft delete)
	stmt := `UPDATE public."CLAIM_APPLICATION" SET soft_delete=1 WHERE id=$1;`

	// executes SQL query
	_, err := db.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	// return error
	return nil
}

// approve claim by id
func (c *Claim) Approve(rowID, userID, amount int) error {
	now := time.Now().Format("2006-01-02")

	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which update the claim (approve it)
	stmt := `UPDATE public."CLAIM_APPLICATION" SET status_id=4, approved_at=$1, approved_by=$2, approved_amount=$3, updated_by=$2, updated_at=$1 WHERE id=$4;`

	// executes SQL query
	_, err := db.ExecContext(ctx, stmt, now, userID, amount, rowID)
	if err != nil {
		return err
	}

	// return error
	return nil
}

// reject claim by id
func (c *Claim) Reject(rowID, userID int, reason string) error {
	now := time.Now().Format("2006-01-02")

	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which update the claim (approve it)
	stmt := `UPDATE public."CLAIM_APPLICATION" SET status_id=3, approved_at=$1, approved_by=$2, approved_amount=0 , rejected_reason=$3, updated_by=$2, updated_at=$1 WHERE id=$4;`

	// executes SQL query
	_, err := db.ExecContext(ctx, stmt, now, userID, reason, rowID)
	if err != nil {
		return err
	}

	// return error
	return nil
}

// soft delete claim by id
func (c *Claim) Delete(id int) error {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which update an employee (soft delete)
	stmt := `UPDATE public."CLAIM_APPLICATION" SET soft_delete=1 WHERE id=$1;`

	// executes SQL query
	_, err := db.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	// return error
	return nil
}

// get all my claims (all years)
func (c *Claim) GetAllMyClaim(eid int) ([]*Claim, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch employee summary
	query := `SELECT c.id, 
					 c.claim_definition_id,
					 cd.name as claim_definition,
					 c.description, 
					 c.amount,
					 ccc.name AS category,
					 c.status_id, 
					 ccs.name as status, 
					 c.approved_at, 
					 c.approved_by, 
					 c.approved_amount, 
					 c.rejected_reason,
					 c.created_at,
					 c.created_by,
					 c.updated_at,
					 c.updated_by
			  FROM public."CLAIM_APPLICATION" c, public."CLAIM_DEFINITION" cd, public."CONFIG_STATUS" ccs, public."CONFIG_CATEGORY" ccc
			  WHERE c.soft_delete = 0
			  AND c.employee_id = $1
			  AND c.claim_definition_id = cd.id
			  AND cd.category_id = ccc.id
			  and c.status_id = ccs.id
			  ORDER BY c.id;`

	// executes SQL query
	rows, err := db.QueryContext(ctx, query, eid)
	if err != nil {
		log.Println("the error in query", err)
		return nil, err
	}
	defer rows.Close()

	// populate returned rows to employee summary struct
	var all []*Claim
	for rows.Next() {
		var myClaims Claim
		err := rows.Scan(
			&myClaims.ID,
			&myClaims.ClaimDefinitionID,
			&myClaims.ClaimDefinition,
			&myClaims.Description,
			&myClaims.Amount,
			&myClaims.Category,
			&myClaims.StatusID,
			&myClaims.Status,
			&myClaims.ApprovedAt,
			&myClaims.ApprovedBy,
			&myClaims.ApprovedAmount,
			&myClaims.RejectedReason,
			&myClaims.CreatedAt,
			&myClaims.CreatedBy,
			&myClaims.UpdatedAt,
			&myClaims.UpdatedBy,
		)
		if err != nil {
			log.Println("the error in scan", err)
			return nil, err
		}

		all = append(all, &myClaims)
	}

	// return all employee summary rows
	return all, nil
}

// get all my claims by status (actual year total)
func (c *Claim) GetAllMyYearlyClaim(eid int) (*MyClaims, error) {

	// status id
	// 0	not defined
	// 1	draft
	// 2	pending
	// 3	rejected
	// 4	approved

	var all MyClaims

	approved, err := getMyYearlyClaimByStatus(4, eid)
	if err != nil {
		return nil, err
	}
	pending, err := getMyYearlyClaimByStatus(2, eid)
	if err != nil {
		return nil, err
	}
	rejected, err := getMyYearlyClaimByStatus(3, eid)
	if err != nil {
		return nil, err
	}

	all.Approved = approved
	all.Pending = pending
	all.Rejected = rejected

	return &all, nil
}

// get all my claim by status (actual year details)
func (c *Claim) GetAllMyYearlyClaimDetails(eid int) (*AllClaim, error) {

	// status id
	// 0	not defined
	// 1	draft
	// 2	pending
	// 3	rejected
	// 4	approved

	var all AllClaim

	approved, err := getAllMyClaimByStatus(4, eid)
	if err != nil {
		return nil, err
	}
	pending, err := getAllMyClaimByStatus(2, eid)
	if err != nil {
		return nil, err
	}
	rejected, err := getAllMyClaimByStatus(3, eid)
	if err != nil {
		return nil, err
	}

	all.Approved = approved
	all.Pending = pending
	all.Rejected = rejected

	return &all, nil
}

// get all claim by status
func (c *Claim) GetAllClaim() (*AllClaim, error) {

	// status id
	// 0	not defined
	// 1	draft
	// 2	pending
	// 3	rejected
	// 4	approved

	var all AllClaim

	approved, err := getAllClaimByStatus(4)
	if err != nil {
		return nil, err
	}
	pending, err := getAllClaimByStatus(2)
	if err != nil {
		return nil, err
	}
	rejected, err := getAllClaimByStatus(3)
	if err != nil {
		return nil, err
	}

	all.Approved = approved
	all.Pending = pending
	all.Rejected = rejected

	return &all, nil
}

// helpers returning all claims by status (all employees, all years)
func getAllClaimByStatus(status int) ([]*Claim, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch employee summary
	query := `SELECT c.id, 
					 c.claim_definition_id,
					 cd.name as claim_definition,
					 cd.doc_required,
					 c.description, 
					 c.amount, 
					 ccc.name as category,
					 c.status_id, 
					 ccs.name as status, 
					 c.approved_at, 
					 c.approved_by, 
					 c.approved_amount, 
					 c.rejected_reason, 
					 c.employee_id,
					 c.created_at,
					 c.created_by,
					 c.updated_at,
					 c.updated_by
			  FROM public."CLAIM_APPLICATION" c, public."CLAIM_DEFINITION" cd, public."CONFIG_CATEGORY" ccc, public."CONFIG_STATUS" ccs
			  WHERE c.status_id = $1
			  AND c.soft_delete = 0
			  AND cd.category_id = ccc.id
			  AND c.claim_definition_id = cd.id
			  and c.status_id = ccs.id
			  ORDER BY c.id;`

	// executes SQL query
	rows, err := db.QueryContext(ctx, query, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// populate returned rows to employee summary struct
	var all []*Claim
	for rows.Next() {
		var myClaims Claim
		err := rows.Scan(
			&myClaims.ID,
			&myClaims.ClaimDefinitionID,
			&myClaims.ClaimDefinition,
			&myClaims.ClaimDefinitionDocRequired,
			&myClaims.Description,
			&myClaims.Amount,
			&myClaims.Category,
			&myClaims.StatusID,
			&myClaims.Status,
			&myClaims.ApprovedAt,
			&myClaims.ApprovedBy,
			&myClaims.ApprovedAmount,
			&myClaims.RejectedReason,
			&myClaims.EmployeeID,
			&myClaims.CreatedAt,
			&myClaims.CreatedBy,
			&myClaims.UpdatedAt,
			&myClaims.UpdatedBy,
		)
		if err != nil {
			return nil, err
		}

		all = append(all, &myClaims)
	}

	// return all employee summary rows
	return all, nil
}

// helpers returning one employee's claims by status (current year)
func getAllMyClaimByStatus(status, eid int) ([]*Claim, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch employee summary
	query := `SELECT c.id, 
					 c.claim_definition_id,
					 cd.name as claim_definition,
					 c.description, 
					 c.amount, 
					 ccc.name as category, 
					 c.status_id, 
					 ccs.name as status, 
					 c.approved_at, 
					 c.approved_by, 
					 c.approved_amount, 
					 c.rejected_reason, 
					 c.employee_id,
					 c.created_at,
					 c.created_by,
					 c.updated_at,
					 c.updated_by
			  FROM public."CLAIM_APPLICATION" c, public."CLAIM_DEFINITION" cd, public."CONFIG_CATEGORY" ccc, public."CONFIG_STATUS" ccs
			  WHERE c.status_id = $1
			  AND c.employee_id  = $2
			  AND c.claim_definition_id = cd.id
			  AND c.status_id = ccs.id
			  AND date_part('year', c.created_at) = date_part('year', CURRENT_DATE)
			  ORDER BY c.id;`

	// executes SQL query
	rows, err := db.QueryContext(ctx, query, status, eid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// populate returned rows to employee summary struct
	var all []*Claim
	for rows.Next() {
		var myClaims Claim
		err := rows.Scan(
			&myClaims.ID,
			&myClaims.ClaimDefinitionID,
			&myClaims.ClaimDefinition,
			&myClaims.Description,
			&myClaims.Amount,
			&myClaims.Category,
			&myClaims.StatusID,
			&myClaims.Status,
			&myClaims.ApprovedAt,
			&myClaims.ApprovedBy,
			&myClaims.ApprovedAmount,
			&myClaims.RejectedReason,
			&myClaims.EmployeeID,
			&myClaims.CreatedAt,
			&myClaims.CreatedBy,
			&myClaims.UpdatedAt,
			&myClaims.UpdatedBy,
		)
		if err != nil {
			return nil, err
		}

		all = append(all, &myClaims)
	}

	// return all employee summary rows
	return all, nil
}

// helpers returning one employee's claims by status (total current year)
func getMyYearlyClaimByStatus(status, eid int) (float32, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch employee summary
	query := `SELECT COALESCE(SUM(amount),0) as total
			  FROM public."CLAIM_APPLICATION" c
			  WHERE c.status_id = $1
			  AND c.employee_id  = $2
			  AND date_part('year', c.created_at) = date_part('year', CURRENT_DATE)`

	// executes SQL query
	row := db.QueryRowContext(ctx, query, status, eid)

	var total float32

	// populate returned row to total
	err := row.Scan(
		&total,
	)

	if err != nil {
		return 0, err
	}

	// return all employee summary rows
	return total, nil
}
