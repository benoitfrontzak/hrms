package pg

import (
	"context"
	"database/sql"
	"log"
	"time"
)

const dbTimeout = time.Second * 3

var db *sql.DB

// New is the function used to create an instance of the data package. It returns the type
// Model, which embeds all the types we want to be available to our application.
func New(dbPool *sql.DB) Models {
	db = dbPool

	return Models{
		Employee: Employee{},
	}
}

// Models is the type for this package. Note that any model that is included as a member
// in this type is available to us throughout the application, anywhere that the
// app variable is used, provided that the model is also added in the New function.
type Models struct {
	Employee Employee
}

// Employee is the structure which holds one employee to be inserted to the database.
type Employee struct {
	Firstname         string `json:"firstname"`
	Middlename        string `json:"middlename"`
	Familyname        string `json:"familyname"`
	Givenname         string `json:"givenname"`
	EmployeeCode      string `json:"employeeCode"`
	Gender            int    `json:"gender"`
	Maritalstatus     int    `json:"maritalstatus"`
	IcNumber          string `json:"icNumber"`
	PassportNumber    string `json:"passportNumber"`
	PassportExpiryAt  string `json:"passportExpiryAt"`
	Nationality       int    `json:"nationality"`
	Residence         int    `json:"residence"`
	ImmigrationNumber string `json:"immigrationNumber"`
	Birthdate         string `json:"birthdate"`
	Race              int    `json:"race"`
	Religion          int    `json:"religion"`
	IsDisabled        int    `json:"isDisabled"`
	IsActive          int    `json:"isActive"`
	IsForeigner       int    `json:"isForeigner"`
	Eclaim            int    `json:"eclaim"`
	Eleave            int    `json:"eleave"`
	Streetaddr1       string `json:"streetaddr1"`
	Streetaddr2       string `json:"streetaddr2"`
	Zip               string `json:"zip"`
	City              string `json:"city"`
	State             string `json:"state"`
	Country           string `json:"country"`
	PrimaryPhone      string `json:"primaryPhone"`
	SecondaryPhone    string `json:"secondaryPhone"`
	PrimaryEmail      string `json:"primaryEmail"`
	SecondaryEmail    string `json:"secondaryEmail"`
	FullnameEC        string `json:"fullnameEC"`
	MobileEC          string `json:"mobileEC"`
	RelationshipEC    int    `json:"relationshipEC"`
}

// Insert inserts a new user into the database, and returns the ID of the newly inserted row
func (e *Employee) Insert(emp Employee) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var newID int
	stmt := `INSERT INTO public."EMPLOYEE"
	(created_at, created_by, updated_at, updated_by, employee_code, firstname, middlename, lastname, givenname, ic_number, passport_number, passport_expiry_at, birthdate, nationality_id, residence_country_id, marital_status_id, gender_id, address_street1, address_street2, address_city, address_state, address_zip, address_country, primary_phone, secondary_phone, primary_email, secondary_email, is_foreigner, immigration_number, is_disabled, is_active, has_eleave, has_eclaim, race_id, religion_id)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35) returning id;`

	err := db.QueryRowContext(ctx, stmt,
		time.Now().String(),
		1,
		time.Now().String(),
		1,
		emp.EmployeeCode,
		emp.Firstname,
		emp.Middlename,
		emp.Familyname,
		emp.Givenname,
		emp.IcNumber,
		emp.PassportNumber,
		emp.PassportExpiryAt,
		emp.Birthdate,
		emp.Nationality,
		emp.Residence,
		emp.Maritalstatus,
		emp.Gender,
		emp.Streetaddr1,
		emp.Streetaddr2,
		emp.City,
		emp.State,
		emp.Zip,
		emp.Country,
		emp.PrimaryPhone,
		emp.SecondaryPhone,
		emp.PrimaryEmail,
		emp.SecondaryEmail,
		emp.IsForeigner,
		emp.ImmigrationNumber,
		emp.IsDisabled,
		emp.IsActive,
		emp.Eleave,
		emp.Eclaim,
		emp.Race,
		emp.Religion,
	).Scan(&newID)

	if err != nil {
		log.Println("SQL error: ", err)
		return 0, err
	}

	return newID, nil
}

// GetAll returns a slice of all users, sorted by last name
// func (u *Employee) GetAll() ([]*Employee, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
// 	defer cancel()

// 	query := `select id, email, firstname, lastname, nickname, password, active, role, created_at, updated_at
// 	from users order by lastname`

// 	rows, err := db.QueryContext(ctx, query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var users []*Employee

// 	for rows.Next() {
// 		var user Employee
// 		err := rows.Scan(
// 			&user.ID,
// 			&user.Email,
// 			&user.Firstname,
// 			&user.Lastname,
// 			&user.Nickname,
// 			&user.Password,
// 			&user.Active,
// 			&user.Role,
// 			&user.CreatedAt,
// 			&user.UpdatedAt,
// 			// &user.TerminatedAt,
// 		)
// 		if err != nil {
// 			log.Println("Error scanning", err)
// 			return nil, err
// 		}

// 		users = append(users, &user)
// 	}

// 	return users, nil
// }

// GetByEmail returns one user by email
// func (u *Employee) GetByEmail(email string) (*Employee, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
// 	defer cancel()

// 	query := `select id, email, firstname, lastname, nickname, password, active, role, created_at, updated_at  from users where email = $1`

// 	var user Employee
// 	row := db.QueryRowContext(ctx, query, email)

// 	err := row.Scan(
// 		&user.ID,
// 		&user.Email,
// 		&user.Firstname,
// 		&user.Lastname,
// 		&user.Nickname,
// 		&user.Password,
// 		&user.Active,
// 		&user.Role,
// 		&user.CreatedAt,
// 		&user.UpdatedAt,
// 	)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &user, nil
// }

// GetOne returns one user by id
// func (u *Employee) GetOne(id int) (*Employee, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
// 	defer cancel()

// 	query := `select id, email, firstname, lastname, password, active, created_at, updated_at from users where id = $1`

// 	var user Employee
// 	row := db.QueryRowContext(ctx, query, id)

// 	err := row.Scan(
// 		&user.ID,
// 		&user.Email,
// 		&user.Firstname,
// 		&user.Lastname,
// 		&user.Nickname,
// 		&user.Password,
// 		&user.Active,
// 		&user.Role,
// 		&user.CreatedAt,
// 		&user.UpdatedAt,
// 		// &user.TerminatedAt,
// 	)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &user, nil
// }

// Update updates one user in the database, using the information
// stored in the receiver u
// func (u *Employee) Update() error {
// 	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
// 	defer cancel()

// 	stmt := `update users set
// 		email = $1,
// 		firstname = $2,
// 		lastname = $3,
// 		nickname = $4,
// 		active = $5,
// 		role = $6,
// 		updated_at = $7
// 		terminated_at = $8
// 		where id = $9
// 	`

// 	_, err := db.ExecContext(ctx, stmt,
// 		u.Email,
// 		u.Firstname,
// 		u.Lastname,
// 		u.Nickname,
// 		u.Active,
// 		u.Role,
// 		time.Now(),
// 		// u.TerminatedAt,
// 		u.ID,
// 	)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// DeleteByID deletes one user from the database, by ID
// func (u *Employee) DeleteByID(id int) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
// 	defer cancel()

// 	stmt := `delete from users where id = $1`

// 	_, err := db.ExecContext(ctx, stmt, id)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
