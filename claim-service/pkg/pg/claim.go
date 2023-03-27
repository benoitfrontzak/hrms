package pg

import (
	"context"
	"strconv"
	"time"
)

// helper for create new claim (set default value)
func createMyDefaultValue(category string) claimDefaultValue {
	categoryID, _ := strconv.Atoi(category)

	return claimDefaultValue{
		CategoryID:     categoryID,
		StatusID:       2,
		ApprovedAt:     "0001-01-01",
		ApprovedBy:     0,
		ApprovedAmount: 0,
		ApprovedReason: "",
		SoftDelete:     0,
	}
}

// create new claim
func (c *Claim) Insert() (int, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which update an employee (soft delete)
	stmt := `INSERT INTO public."CLAIM" (claim_definition_id, 
										 "name", 
										 description, 
										 amount, 
										 category_id, 
										 status_id, 
										 approved_at, 
										 approved_by, 
										 approved_amount, 
										 approved_reason, 
										 employee_id, 
										 soft_delete,
										 created_by,
										 updated_by) 
			 VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, 0, $12, $13) returning id;`

	// store new row id
	var newID int

	// get claim default value
	cdv := createMyDefaultValue(c.Category)

	// executes SQL query (set SQL parameters and cacth rowID)
	err := db.QueryRowContext(ctx, stmt,
		c.ClaimDefinition,
		c.Name,
		c.Description,
		c.Amount,
		cdv.CategoryID,
		cdv.StatusID,
		cdv.ApprovedAt,
		cdv.ApprovedBy,
		cdv.ApprovedAmount,
		cdv.ApprovedReason,
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
func (c *Claim) Approve(rowID, userID, amount int) error {
	now := time.Now().Format("2006-01-02")

	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which update the claim (approve it)
	stmt := `UPDATE public."CLAIM" SET status_id=4, approved_at=$1, approved_by=$2, approved_amount=$3, updated_by=$2, updated_at=$1 WHERE id=$4;`

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
	stmt := `UPDATE public."CLAIM" SET status_id=3, approved_at=$1, approved_by=$2, approved_amount=0 , approved_reason=$3, updated_by=$2, updated_at=$1 WHERE id=$4;`

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
	stmt := `UPDATE public."CLAIM" SET soft_delete=1 WHERE id=$1;`

	// executes SQL query
	_, err := db.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	// return error
	return nil
}

// get all claim (active, inactive, deleted)
func (c *Claim) GetAllMyClaim(eid int) ([]*Claim, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch employee summary
	query := `SELECT c.id, 
					 c.claim_definition_id,
					 cd.name as claim_definition,
					 c."name", 
					 c.description, 
					 c.amount, 
					 c.category_id, 
					 ccc.name as category, 
					 c.status_id, 
					 ccs.name as status, 
					 c.approved_at, 
					 c.approved_by, 
					 c.approved_amount, 
					 c.approved_reason,
					 c.created_at,
					 c.created_by,
					 c.updated_at,
					 c.updated_by
			  FROM public."CLAIM" c, public."CLAIM_DEFINITION" cd, public."CONFIG_CATEGORY" ccc, public."CONFIG_STATUS" ccs
			  WHERE c.soft_delete = 0
			  AND c.employee_id = $1
			  AND c.claim_definition_id = cd.id
			  and c.category_id = ccc.id
			  and c.status_id = ccs.id
			  ORDER BY c.id;`

	// executes SQL query
	rows, err := db.QueryContext(ctx, query, eid)
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
			&myClaims.Name,
			&myClaims.Description,
			&myClaims.Amount,
			&myClaims.CategoryID,
			&myClaims.Category,
			&myClaims.StatusID,
			&myClaims.Status,
			&myClaims.ApprovedAt,
			&myClaims.ApprovedBy,
			&myClaims.ApprovedAmount,
			&myClaims.ApprovedReason,
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

// helpers
func getAllClaimByStatus(status int) ([]*Claim, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch employee summary
	query := `SELECT c.id, 
					 c.claim_definition_id,
					 cd.name as claim_definition,
					 c."name", 
					 c.description, 
					 c.amount, 
					 c.category_id, 
					 ccc.name as category, 
					 c.status_id, 
					 ccs.name as status, 
					 c.approved_at, 
					 c.approved_by, 
					 c.approved_amount, 
					 c.approved_reason, 
					 c.employee_id,
					 c.created_at,
					 c.created_by,
					 c.updated_at,
					 c.updated_by
			  FROM public."CLAIM" c, public."CLAIM_DEFINITION" cd, public."CONFIG_CATEGORY" ccc, public."CONFIG_STATUS" ccs
			  WHERE c.status_id = $1
			  AND c.claim_definition_id = cd.id
			  and c.category_id = ccc.id
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
			&myClaims.Name,
			&myClaims.Description,
			&myClaims.Amount,
			&myClaims.CategoryID,
			&myClaims.Category,
			&myClaims.StatusID,
			&myClaims.Status,
			&myClaims.ApprovedAt,
			&myClaims.ApprovedBy,
			&myClaims.ApprovedAmount,
			&myClaims.ApprovedReason,
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
