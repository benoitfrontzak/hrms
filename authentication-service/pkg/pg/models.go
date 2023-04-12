package pg

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const dbTimeout = time.Second * 3

var db *sql.DB

// New is the function used to create an instance of the data package. It returns the type
// Model, which embeds all the types we want to be available to our application.
func New(dbPool *sql.DB) Models {
	db = dbPool

	return Models{
		User: User{},
	}
}

// GetAll returns a slice of all users, sorted by id
func (u *User) GetAll() ([]*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, email, fullname, nickname, password, active, role, employee_id from "USERS" order by id`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User

	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.Fullname,
			&user.Nickname,
			&user.Password,
			&user.Active,
			&user.Role,
			&user.EmployeeID,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

// GetByEmail returns one user by email
func (u *User) GetByEmail(email string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, email, fullname, nickname, password, active, role, employee_id from "USERS" where email = $1 AND active = 1`

	var user User
	row := db.QueryRowContext(ctx, query, email)

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Fullname,
		&user.Nickname,
		&user.Password,
		&user.Active,
		&user.Role,
		&user.EmployeeID,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetOne returns one user by id
func (u *User) GetOne(id int) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, email, fullname, nickname, password, active, role, employee_id from "USERS" where id = $1`

	var user User
	row := db.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Fullname,
		&user.Nickname,
		&user.Password,
		&user.Active,
		&user.Role,
		&user.EmployeeID,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Update updates one user in the database, using the information
// stored in the receiver u
func (u *User) Update(user User) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `UPDATE "USERS" 
			 SET
				email = $1,
				fullname = $2,
				nickname = $3,
				active = $4,
				role = $5
			 WHERE employee_id = $6`

	_, err := db.ExecContext(ctx, stmt,
		user.Email,
		user.Fullname,
		user.Nickname,
		user.Active,
		user.Role,
		user.EmployeeID,
	)

	if err != nil {
		return err
	}

	return nil
}

// DeleteByID deletes one user from the database, by ID
func (u *User) DeleteByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `delete from "USERS" where id = $1`

	_, err := db.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	return nil
}

// Insert inserts a new user into the database, and returns the ID of the newly inserted row
func (u *User) Insert(user User) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return 0, err
	}

	var newID int
	stmt := `insert into "USERS" (email, fullname, nickname, password, active, role, employee_id)
		values ($1, $2, $3, $4, $5, $6, $7) returning id`

	err = db.QueryRowContext(ctx, stmt,
		user.Email,
		user.Fullname,
		user.Nickname,
		hashedPassword,
		user.Active,
		user.Role,
		user.EmployeeID,
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}

// ResetPassword is the method we will use to change a user's password.
func (u *User) UpdatePassword() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	log.Println("password to be hashed is ", u.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
	if err != nil {
		return err
	}

	stmt := `update "USERS" set password = $1 where employee_id = $2`
	_, err = db.ExecContext(ctx, stmt, hashedPassword, u.EmployeeID)
	if err != nil {
		return err
	}

	return nil
}

// PasswordMatches uses Go's bcrypt package to compare a user supplied password
// with the hash we have stored for a given user in the database. If the password
// and hash match, we return true; otherwise, we return false.
func (u *User) PasswordMatches(plainText string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainText))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			// invalid password
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}

// Get user password for compare
func (u *User) GetPassword(id int) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	log.Println("employeeID is ", id)

	stmt := `SELECT password FROM "USERS" WHERE employee_id = $1`
	row := db.QueryRowContext(ctx, stmt, id)

	var pwd string
	err := row.Scan(&pwd)
	if err != nil {
		return "", err
	}

	return pwd, nil
}
