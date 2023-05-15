package pg

import (
	"context"
	"log"
	"strconv"
	"time"
)

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

// insert a new claim definition & all claim definition details relative
func (cd *ClaimDefinition) Insert() (int, error) {
	// insert new claim definition and fetch rowID inserted
	rowID, err := cd.insertClaimDefinition()
	if err != nil {
		return 0, err
	}
	// insert all claim definition details
	for _, entry := range cd.Details {
		_, err = insertClaimDefinitionDetails(entry.Seniority, entry.Limitation, cd.CreatedBy, rowID)
		if err != nil {
			return 0, err
		}
	}

	return rowID, nil
}

// update claim definition by id
func (cd *ClaimDefinition) Update() error {
	// update claim definition
	err := cd.updateClaimDefinition()
	if err != nil {
		log.Println("error at updateCDef", err)
		return err
	}

	// update all claim definition details
	for _, entry := range cd.Details {
		_, err = insertClaimDefinitionDetails(entry.Seniority, entry.Limitation, cd.CreatedBy, cd.ID)
		if err != nil {
			log.Println("error at insCDef", err)
			return err
		}
	}

	// update all claim definition details
	for _, entry := range cd.DetailsUpdate {
		err = updateClaimDefinitionDetails(entry.Seniority, entry.Limitation, cd.CreatedBy, entry.ID)
		if err != nil {
			log.Println("error at upCDef", err)
			return err
		}
	}

	// delete all claim definition details
	if len(cd.DetailsDeleted) > 0 {
		for _, entry := range cd.DetailsDeleted {
			rowID, err := strconv.Atoi(entry)
			log.Println("rowID", rowID)
			if err != nil {
				return err
			}
			err = deleteClaimDefinitionDetails(rowID)
			if err != nil {
				log.Println("error at delCDef", err)
				return err
			}
		}
	}

	return nil
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
					 cd.limitation, 
					 cd.doc_required, 
					 cd.created_at, 
					 cd.created_by, 
					 cd.updated_at, 
					 cd.updated_by
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
			&def.Limitation,
			&def.DocRequired,
			&def.CreatedAt,
			&def.CreatedBy,
			&def.UpdatedAt,
			&def.UpdatedBy,
		)
		if err != nil {
			return nil, err
		}

		// fetch all leave definition details
		details, err := GetClaimDefinitionDetails(def.ID)
		if err != nil {
			return nil, err
		}
		def.Details = details

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
					 cd.limitation, 
					 cd.doc_required,
					 cd.created_at, 
					 cd.created_by, 
					 cd.updated_at, 
					 cd.updated_by
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
			&def.Limitation,
			&def.DocRequired,
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

	// return all employee summary rows
	return all, nil
}

// insert a new claim definition
func (cd *ClaimDefinition) insertClaimDefinition() (int, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which insert a new employee
	stmt := `INSERT INTO public."CLAIM_DEFINITION" (active, "name", description,
			category_id, limitation, doc_required, created_by, updated_by) 
			VALUES($1, $2, $3, $4, $5, $6, $7, $8) returning id;`

	// executes SQL query (set SQL parameters and cacth rowID)
	var newID int

	err := db.QueryRowContext(ctx, stmt,
		cd.Active,
		cd.Name,
		cd.Description,
		cd.Category,
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
func (cd *ClaimDefinition) updateClaimDefinition() error {
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
				limitation=$5, 
				doc_required=$6, 
				soft_delete=0
			 WHERE id=$7;`

	// executes SQL query
	_, err := db.ExecContext(ctx, stmt,
		cd.Active,
		cd.Name,
		cd.Description,
		cd.Category,
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

//  ---------------- Definition Detail Section -----------------

// Get all claim definition details
func GetClaimDefinitionDetails(id int) ([]*ClaimDefinitionDetails, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch claim definition details
	query := `SELECT id, 
					 seniority,  
					 limitation,
					 claim_definition_id, 
					 soft_delete, 
					 created_at, 
					 created_by, 
					 updated_at, 
					 updated_by 
			  FROM public."CLAIM_DEFINITION_DETAILS" 
			  WHERE soft_delete = 0
			  AND claim_definition_id = $1
			  ORDER BY seniority DESC;`
	//   ORDER BY id;` //To be validated

	// executes SQL query
	rows, err := db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// populate returned rows to claim definition details struct
	var all []*ClaimDefinitionDetails
	for rows.Next() {
		var def ClaimDefinitionDetails
		err := rows.Scan(
			&def.ID,
			&def.Seniority,
			&def.Limitation,
			&def.ClaimDefinitionID,
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

	// return all claim definition details rows
	return all, nil
}

// insert claim definition details
func insertClaimDefinitionDetails(seniority int, limitation float32, user int, rowID int) (int, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which insert a new claim definition
	stmt := `INSERT INTO public."CLAIM_DEFINITION_DETAILS" (seniority, limitation,
			claim_definition_id, created_by, updated_by)
			VALUES($1, $2, $3, $4, $5) returning id;`

	log.Println("Output in insertClaimDefinitionDetails: ", seniority, limitation, user, rowID)

	// executes SQL query (set SQL parameters and cacth rowID)
	var newID int

	err := db.QueryRowContext(ctx, stmt,
		seniority,
		limitation,
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

// Delete claim definition detail by id
func deleteClaimDefinitionDetails(id int) error {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which update an employee (soft delete)
	stmt := `UPDATE public."CLAIM_DEFINITION_DETAILS" SET soft_delete=1 WHERE id=$1;`

	// executes SQL query
	_, err := db.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	// return error
	return nil
}

// Update the claim definition detail
func updateClaimDefinitionDetails(seniority int, limitation float32, user int, rowID int) error {
	now := time.Now().Format("2006-01-02")
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	log.Println("arg", seniority, limitation, user, rowID)

	// SQL statement which update an employee (soft delete)
	stmt := `UPDATE public."CLAIM_DEFINITION_DETAILS" 
			 SET seniority=$1, 
			 	limitation=$2, 
				updated_at=$3, 
				updated_by=$4
			 WHERE id=$5;`

	// executes SQL query
	_, err := db.ExecContext(ctx, stmt,
		seniority,
		limitation,
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
