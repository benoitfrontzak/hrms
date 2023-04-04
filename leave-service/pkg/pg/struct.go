package pg

// Models is the type for this package. Note that any model that is included as a member
// in this type is available to us throughout the application, anywhere that the
// app variable is used, provided that the model is also added in the New function.
type Models struct {
	Leave           Leave
	LeaveDefinition LeaveDefinition
	ConfigTables    ConfigTables
}

// Structure holding config tables
type ConfigTables struct {
	Limitation  []configT
	Calculation []configT
}

// Structure holding the generic config table columns
type configT struct {
	ID   int
	Name string
}

// Structure holding all status leaves
type AllLeave struct {
	Approved []*Leave
	Pending  []*Leave
	Rejected []*Leave
}

// Structure holding all approved leave summary of today & tomorrow
type AllCurentLeave struct {
	Today    []*ApprovedLeave
	Tomorrow []*ApprovedLeave
}

// Structure holding one approved leave summary
type ApprovedLeave struct {
	Code          string
	Description   string
	Reason        string
	EmployeeID    int
	RequestedDate string
	IsHalf        int
}

// Structure holding one entitled leave (got one per leave definition)
type EntitledLeave struct {
	ID                  int     `json:"rowid,string,omitempty"`
	Entitled            float32 `json:"entitled,string"`
	Taken               float32 `json:"taken,string"`
	LeaveDefinitionID   int     `json:"leaveDefinition,string"`
	LeaveDefinitionCode string  `json:"leaveDefinitionCode"`
	LeaveDefinitionName string  `json:"leaveDefinitionName"`
	EmployeeID          int     `json:"employeeid,string"`
	SoftDelete          int     `json:"softDelete,string"`
	CreatedAt           string  `json:"createdAt"`
	CreatedBy           int     `json:"createdBy"`
	UpdatedAt           string  `json:"updatedAt"`
	UpdatedBy           int     `json:"updatedBy"`
}

// Structure holding one leave (application)
type Leave struct {
	ID                  int             `json:"rowid,string,omitempty"`
	LeaveDefinitionID   int             `json:"leaveDefinition,string"`
	LeaveDefinitionCode string          `json:"leaveDefinitionCode"`
	LeaveDefinitionName string          `json:"leaveDefinitionName"`
	Description         string          `json:"description"`
	StatusID            int             `json:"statusid,string"`
	Status              string          `json:"status"`
	ApprovedAt          string          `json:"approvedAt"`
	ApprovedBy          int             `json:"approvedBy,string"`
	RejectedReason      string          `json:"rejectedReason"`
	Details             []*LeaveDetails `json:"details"`
	EmployeeID          int             `json:"employeeid,string"`
	SoftDelete          int             `json:"softDelete,string"`
	CreatedAt           string          `json:"createdAt"`
	CreatedBy           int             `json:"createdBy"`
	UpdatedAt           string          `json:"updatedAt"`
	UpdatedBy           int             `json:"updatedBy"`
}

// Structure holding one leave details (application dates)
type LeaveDetails struct {
	ID            int    `json:"rowid,string,omitempty"`
	LeaveID       int    `json:"leaveID,string"`
	RequestedDate string `json:"requestedDate"`
	IsHalf        int    `json:"isHalf"`
	SoftDelete    int    `json:"softDelete,string"`
	CreatedAt     string `json:"createdAt"`
	CreatedBy     int    `json:"createdBy"`
	UpdatedAt     string `json:"updatedAt"`
	UpdatedBy     int    `json:"updatedBy"`
}

// Structure holding one leave Definition
type LeaveDefinition struct {
	ID                  int                       `json:"rowid,string,omitempty"`
	Code                string                    `json:"code"`
	Description         string                    `json:"description"`
	Unpaid              int                       `json:"unpaid,string"`
	Encashable          int                       `json:"encashable,string"`
	ReplacementRequired int                       `json:"replacementRequired,string"`
	DocRequired         int                       `json:"docRequired,string"`
	Expiry              string                    `json:"expiry"`
	GenderID            int                       `json:"genderid,string"`
	Gender              string                    `json:"gender"`
	LimitationID        int                       `json:"limitationid,string"`
	Limitation          string                    `json:"limitation"`
	CalculationID       int                       `json:"calculationid,string"`
	Calculation         string                    `json:"calculation"`
	Details             []*LeaveDefinitionDetails `json:"details"`
	DetailsDeleted      []string                  `json:"detailsDeleted"`
	SoftDelete          int                       `json:"softDelete,string"`
	CreatedAt           string                    `json:"createdAt"`
	CreatedBy           int                       `json:"createdBy"`
	UpdatedAt           string                    `json:"updatedAt"`
	UpdatedBy           int                       `json:"updatedBy"`
}

// Structure holding one leave Definition calculation
type LeaveDefinitionDetails struct {
	ID                int    `json:"rowid,string,omitempty"`
	Seniority         int    `json:"seniority,string"`
	Entitled          int    `json:"entitled,string"`
	LeaveDefinitionID int    `json:"leaveDefinitionid,string"`
	SoftDelete        int    `json:"softDelete,string"`
	CreatedAt         string `json:"createdAt"`
	CreatedBy         int    `json:"createdBy,string"`
	UpdatedAt         string `json:"updatedAt"`
	UpdatedBy         int    `json:"updatedBy,string"`
}

// Structure holding one leave Definition calculation & entitled
type LeaveDefinitionEntitled struct {
	ID            int
	Entitled      int
	CalculationID int
}
