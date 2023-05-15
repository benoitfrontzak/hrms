package pg

// Models is the type for this package. Note that any model that is included as a member
// in this type is available to us throughout the application, anywhere that the
// app variable is used, provided that the model is also added in the New function.
type Models struct {
	Claim           Claim
	ClaimDefinition ClaimDefinition
	ConfigTables    ConfigTables
}

// Structure holding config tables
type ConfigTables struct {
	Category []configT
	Status   []configT
}

// Structure holding the generic config table columns
type configT struct {
	ID   int
	Name string
}

// Structure holding all my yearly claims total
type MyClaims struct {
	Approved float32
	Pending  float32
	Rejected float32
}

// Structure holding all status claims
type AllClaim struct {
	Approved []*Claim
	Pending  []*Claim
	Rejected []*Claim
}

// Structure holding one Claim (application)
type Claim struct {
	ID                         int     `json:"rowid,string,omitempty"`
	ClaimDefinitionID          int     `json:"claimDefinitionID,string,omitempty"`
	ClaimDefinition            string  `json:"claimDefinition"`
	ClaimDefinitionDocRequired int     `json:"claimDefinitionDocRequired"`
	Description                string  `json:"description"`
	Amount                     float32 `json:"amount,string,omitempty"`
	CategoryID                 int     `json:"categoryid,string,omitempty"`
	Category                   string  `json:"category"`
	StatusID                   int     `json:"statusid,string,omitempty"`
	Status                     string  `json:"status"`
	ApprovedAt                 string  `json:"approvedAt"`
	ApprovedBy                 int     `json:"approvedBy,string,omitempty"`
	ApprovedAmount             float32 `json:"approvedAmount,string,omitempty"`
	RejectedReason             string  `json:"rejectedReason"`
	EmployeeID                 int     `json:"employeeid,string,omitempty"`
	SoftDelete                 int     `json:"softDelete,string,omitempty"`
	CreatedAt                  string  `json:"createdAt,omitempty"`
	CreatedBy                  int     `json:"createdBy,string,omitempty"`
	UpdatedAt                  string  `json:"updatedAt,omitempty"`
	UpdatedBy                  int     `json:"updatedBy,string,omitempty"`
}

// Structure holding default value when creating a new claim
type claimDefaultValue struct {
	CategoryID     int
	StatusID       int
	ApprovedAt     string
	ApprovedBy     int
	ApprovedAmount float32
	RejectedReason string
	SoftDelete     int
}

// Structure holding all Claim Definition (active, inactive, deleted)
type AllClaimDefinition struct {
	Active   []*ClaimDefinition
	Inactive []*ClaimDefinition
	Deleted  []*ClaimDefinition
}

// Structure holding one Claim Definition
type ClaimDefinition struct {
	ID             int                       `json:"rowid,string,omitempty"`
	Active         int                       `json:"active,string"`
	Name           string                    `json:"name"`
	Description    string                    `json:"description"`
	Category       string                    `json:"category"`
	CategoryID     int                       `json:"categoryid,string,omitempty"`
	Limitation     float32                   `json:"limitation,string,omitempty"`
	DocRequired    int                       `json:"docRequired,string,omitempty"`
	Details        []*ClaimDefinitionDetails `json:"details"`
	DetailsUpdate  []*ClaimDefinitionDetails `json:"detailsUpdate"`
	DetailsDeleted []string                  `json:"detailsDeleted"`
	SoftDelete     int                       `json:"softDelete,string,omitempty"`
	CreatedAt      string                    `json:"createdAt,omitempty"`
	CreatedBy      int                       `json:"createdBy,string,omitempty"`
	UpdatedAt      string                    `json:"updatedAt,omitempty"`
	UpdatedBy      int                       `json:"updatedBy,string,omitempty"`
}

// Structure holding one Claim Definition detail
type ClaimDefinitionDetails struct {
	ID                int     `json:"rowid,string,omitempty"`
	Seniority         int     `json:"seniority,string"`
	Limitation        float32 `json:"limitation,string,omitempty"`
	ClaimDefinitionID int     `json:"claimDefinitionID,string,omitempty"`
	SoftDelete        int     `json:"softDelete,string,omitempty"`
	CreatedAt         string  `json:"createdAt,omitempty"`
	CreatedBy         int     `json:"createdBy,string,omitempty"`
	UpdatedAt         string  `json:"updatedAt,omitempty"`
	UpdatedBy         int     `json:"updatedBy,string,omitempty"`
}
