package pg

import (
	"context"
	"log"
	"time"
)

// create new public holiday
func (ph *PublicHoliday) Insert() (int, error) {
	now := time.Now().Format("2006-01-02")

	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which update an employee (soft delete)
	stmt := `INSERT INTO public."PUBLIC_HOLIDAYS" (ph_date, "name", description, soft_delete, created_at, created_by, updated_at, updated_by) 
			 VALUES($1, $2, $3, 0, $4, $5, $6, $7)
			 returning id;`

	// store new row id
	var newID int

	// executes SQL query (set SQL parameters and cacth rowID)
	err := db.QueryRowContext(ctx, stmt,
		ph.Date,
		ph.Name,
		ph.Description,
		now,
		ph.CreatedBy,
		now,
		ph.CreatedBy,
	).Scan(&newID)
	if err != nil {
		return 0, err
	}

	// return error
	return newID, nil
}

// update public holiday by rowID
func (ph *PublicHoliday) Update(uid int) error {
	now := time.Now().Format("2006-01-02")

	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which update an employee
	stmt := `UPDATE public."PUBLIC_HOLIDAYS" 
			 SET ph_date=$1, 
			 	 "name"=$2, 
				 description=$3, 
				 updated_at=$4, 
				 updated_by=$5 
			 WHERE id=$6`

	// executes SQL query
	_, err := db.ExecContext(ctx, stmt, ph.Date, ph.Name, ph.Description, now, uid)
	if err != nil {
		return err
	}

	return nil
}

// soft delete public holiday by rowID
func (ph *PublicHoliday) Delete(id, uid int) error {
	now := time.Now().Format("2006-01-02")

	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which update an employee (soft delete)
	stmt := `UPDATE public."PUBLIC_HOLIDAYS" SET soft_delete=1, updated_at=$1, updated_by=$2 WHERE id=$3;`

	// executes SQL query
	_, err := db.ExecContext(ctx, stmt, now, uid, id)
	if err != nil {
		return err
	}

	// return error
	return nil
}

// get all public holidays
func (ph *PublicHoliday) GetAllPublicHolidays() ([]*PublicHoliday, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch employee summary
	query := `SELECT id, 
					 ph_date, 
					 "name", 
					 description,
					 created_at, 
					 created_by, 
					 updated_at, 
					 updated_by 
			  FROM public."PUBLIC_HOLIDAYS"
			  WHERE soft_delete=0
			  ORDER BY id;`

	// executes SQL query
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		log.Println("err in SQL", err)
		return nil, err
	}
	defer rows.Close()

	// populate returned rows to employee summary struct
	var all []*PublicHoliday
	for rows.Next() {
		var myPH PublicHoliday

		err = rows.Scan(
			&myPH.ID,
			&myPH.Date,
			&myPH.Name,
			&myPH.Description,
			&myPH.CreatedAt,
			&myPH.CreatedBy,
			&myPH.UpdatedAt,
			&myPH.UpdatedBy,
		)
		if err != nil {
			return nil, err
		}
		all = append(all, &myPH)
	}

	// return all public holidays rows
	return all, nil
}
