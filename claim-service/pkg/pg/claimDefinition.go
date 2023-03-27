package pg

import "context"

// get all claim definition (active, inactive, deleted)
func (cd *ClaimDefinition) GetClaimDefinition() (*AllClaimDefinition, error) {
	var all AllClaimDefinition

	active, err := getClaimDefinitionByStatus(1)
	if err != nil {
		return nil, err
	}
	inactive, err := getClaimDefinitionByStatus(0)
	if err != nil {
		return nil, err
	}
	deleted, err := getClaimDefinitionDeleted()
	if err != nil {
		return nil, err
	}

	all.Active = active
	all.Inactive = inactive
	all.Deleted = deleted

	return &all, nil
}

// get all claim definition by status: active | inactive
func getClaimDefinitionByStatus(active int) ([]*ClaimDefinition, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch employee summary
	query := `SELECT cd.id, 
					 cd.active, 
					 cd."name", 
					 cd.description,
					 ccc.name as category,
					 cd.category_id, 
					 cd.confirmation_required, 
					 cd.seniority_required, 
					 cd.limitation, 
					 cd.doc_required 
			  FROM public."CLAIM_DEFINITION" cd, public."CONFIG_CATEGORY" ccc
			  WHERE cd.soft_delete = 0
			  AND cd.active = $1
			  AND cd.category_id = ccc.id
			  ORDER BY cd.id;`

	// executes SQL query
	rows, err := db.QueryContext(ctx, query, active)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// populate returned rows to employee summary struct
	var all []*ClaimDefinition
	for rows.Next() {
		var def ClaimDefinition
		err := rows.Scan(
			&def.ID,
			&def.Active,
			&def.Name,
			&def.Description,
			&def.Category,
			&def.CategoryID,
			&def.ConfirmationRequired,
			&def.SeniorityRequired,
			&def.Limitation,
			&def.DocRequired,
		)
		if err != nil {
			return nil, err
		}

		all = append(all, &def)
	}

	// return all employee summary rows
	return all, nil
}

// get all claim definition deleted
func getClaimDefinitionDeleted() ([]*ClaimDefinition, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch employee summary
	query := `SELECT cd.id, 
					 cd.active, 
					 cd."name", 
					 cd.description,
					 ccc.name as category,
					 cd.category_id, 
					 cd.confirmation_required, 
					 cd.seniority_required, 
					 cd.limitation, 
					 cd.doc_required 
			  FROM public."CLAIM_DEFINITION" cd, public."CONFIG_CATEGORY" ccc
			  WHERE cd.soft_delete = 1
			  AND cd.category_id = ccc.id
			  ORDER BY cd.id;`

	// executes SQL query
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// populate returned rows to employee summary struct
	var all []*ClaimDefinition
	for rows.Next() {
		var def ClaimDefinition
		err := rows.Scan(
			&def.ID,
			&def.Active,
			&def.Name,
			&def.Description,
			&def.Category,
			&def.CategoryID,
			&def.ConfirmationRequired,
			&def.SeniorityRequired,
			&def.Limitation,
			&def.DocRequired,
		)
		if err != nil {
			return nil, err
		}

		all = append(all, &def)
	}

	// return all employee summary rows
	return all, nil
}

// insert a new claim definition
func (cd *ClaimDefinition) Insert() (int, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which insert a new employee
	stmt := `INSERT INTO public."CLAIM_DEFINITION" (active, "name", description, category_id, confirmation_required, seniority_required, limitation, doc_required, created_by, updated_by) 
			 VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) returning id;`

	// executes SQL query (set SQL parameters and cacth rowID)
	var newID int

	err := db.QueryRowContext(ctx, stmt,
		cd.Active,
		cd.Name,
		cd.Description,
		cd.Category,
		cd.ConfirmationRequired,
		cd.SeniorityRequired,
		cd.Limitation,
		cd.DocRequired,
		cd.CreatedBy,
		cd.UpdatedBy,
	).Scan(&newID)
	if err != nil {
		return 0, err
	}

	// return rowID of created new employee
	return newID, nil
}

// update claim definition by id
func (cd *ClaimDefinition) Update() error {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which update an employee (soft delete)
	stmt := `UPDATE public."CLAIM_DEFINITION" 
			 SET 
				active=$1, 
				"name"=$2, 
				description=$3, 
				category_id=$4, 
				confirmation_required=$5, 
				seniority_required=$6, 
				limitation=$7, 
				doc_required=$8, 
				soft_delete=0
			 WHERE id=$9;`

	// executes SQL query
	_, err := db.ExecContext(ctx, stmt,
		cd.Active,
		cd.Name,
		cd.Description,
		cd.Category,
		cd.ConfirmationRequired,
		cd.SeniorityRequired,
		cd.Limitation,
		cd.DocRequired,
		cd.ID,
	)
	if err != nil {
		return err
	}

	// return error
	return nil
}

// soft delete claim definition by id
func (cd *ClaimDefinition) Delete(id int) error {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which update an employee (soft delete)
	stmt := `UPDATE public."CLAIM_DEFINITION" SET soft_delete=1 WHERE id=$1;`

	// executes SQL query
	_, err := db.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	// return error
	return nil
}
