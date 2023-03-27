package pg

import (
	"context"
)

// fetch all rows of CONFIG_* (t) and returns it as []configT
func getAllRowsConfigTable(t string) ([]configT, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch config_t
	query := `SELECT id, name FROM "` + t + `" WHERE soft_delete = 0 ORDER BY id`

	// executes SQL query
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// populate each row to configT
	var myCT []configT
	for rows.Next() {
		var ct configT
		err := rows.Scan(
			&ct.ID,
			&ct.Name,
		)
		if err != nil {
			return nil, err
		}
		// populate []configT
		myCT = append(myCT, ct)
	}

	// return all config_t rows
	return myCT, nil
}

// fetch all config tables
func (ct *ConfigTables) GetAllConfigTables() (*ConfigTables, error) {
	var err error

	ct.Limitation, err = getAllRowsConfigTable("CONFIG_LIMITATION")
	if err != nil {
		return nil, err
	}
	ct.Calculation, err = getAllRowsConfigTable("CONFIG_CALCULATION_METHOD")
	if err != nil {
		return nil, err
	}

	return ct, nil
}

// inserts a new ct entry into the database, and returns the ID of the newly inserted row
func (ct *ConfigTables) Insert(table, value string) (int, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which insert a new entry to the config table
	stmt := `INSERT INTO public."` + table + `" ("name") VALUES($1) returning id;`

	// executes SQL query (set SQL parameters and cacth rowID)
	var newID int

	err := db.QueryRowContext(ctx, stmt, value).Scan(&newID)
	if err != nil {
		return 0, err
	}

	// return rowID of created new employee
	return newID, nil
}

// inserts a new ct entry into the database, and returns the ID of the newly inserted row
func (ct *ConfigTables) Update(table, value string, rowID int) (int, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which insert a new entry to the config table
	stmt := `UPDATE public."` + table + `" SET name = $1 WHERE id=$2;`

	// executes SQL query (set SQL parameters and cacth rowID)
	var newID int

	_, err := db.ExecContext(ctx, stmt, value, rowID)
	if err != nil {
		return 0, err
	}

	// return rowID of created new employee
	return newID, nil
}

// inserts a new ct entry into the database, and returns the ID of the newly inserted row
func (ct *ConfigTables) Delete(table, rowID string) error {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which insert a new entry to the config table
	stmt := `UPDATE public."` + table + `" SET soft_delete = 1 WHERE id=$1;`

	// executes SQL query (set SQL parameters)
	_, err := db.ExecContext(ctx, stmt, rowID)
	if err != nil {
		return err
	}

	return nil
}
