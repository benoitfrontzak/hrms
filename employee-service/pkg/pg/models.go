package pg

import (
	"database/sql"
	"time"
)

// db timeout
const dbTimeout = time.Second * 3

// db connection
var db *sql.DB

// New is the function used to create an instance of the data package. It returns the type
// Model, which embeds all the types we want to be available to our application.
func New(dbPool *sql.DB) Models {
	db = dbPool

	return Models{
		Employee:     Employee{},
		ConfigTables: ConfigTables{},
	}
}
