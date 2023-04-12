package pg

// Models is the type for this package. Note that any model that is included as a member
// in this type is available to us throughout the application, anywhere that the
// app variable is used, provided that the model is also added in the New function.
type Models struct {
	User User
}

// User is the structure which holds one user from the database.
type User struct {
	ID         int    `json:"id,string"`
	Email      string `json:"email"`
	Fullname   string `json:"fullname,omitempty"`
	Nickname   string `json:"nickname,omitempty"`
	Password   string `json:"password,omitempty"`
	Active     int    `json:"active"`
	Role       int    `json:"role"`
	EmployeeID int    `json:"employeeID"`
	CreatedAt  string `json:"createdAt,omitempty"`
	CreatedBy  int    `json:"createdBy,string,omitempty"`
	UpdatedAt  string `json:"updatedAt,omitempty"`
	UpdatedBy  int    `json:"updatedBy,string,omitempty"`
}

type ChangePassword struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
	CreatedBy   int    `json:"createdBy,string"`
	UpdatedBy   int    `json:"updatedBy,string"`
}

// {"oldPassword":"riqporip","newPassword":"Benoit@76","confirmPassword":"Benoit@76","createdBy":"125","updatedBy":"125"}
