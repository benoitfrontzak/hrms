package pg

// Models is the type for this package. Note that any model that is included as a member
// in this type is available to us throughout the application, anywhere that the
// app variable is used, provided that the model is also added in the New function.
type Models struct {
	Employee     Employee
	ConfigTables ConfigTables
}

// Bank is the structure which holds one employee's bank details
type Bank struct {
	ID            int    `json:"id,string,omitempty"`
	BankName      string `json:"bankName"`
	AccountNumber string `json:"accountNumber"`
	AccountName   string `json:"accountName"`
	CreatedAt     string `json:"createdAt,omitempty"`
	CreatedBy     int    `json:"createdBy,string,omitempty"`
	UpdatedAt     string `json:"updatedAt,omitempty"`
	UpdatedBy     int    `json:"updatedBy,string,omitempty"`
}

// EmergencyContact is the structure which holds one employee's emergency contact.
type EmergencyContact struct {
	ID             int    `json:"id,string,omitempty"`
	FullnameEC     string `json:"fullnameEC"`
	MobileEC       string `json:"mobileEC"`
	RelationshipEC int    `json:"relationshipEC,string"`
	CreatedAt      string `json:"createdAt,omitempty"`
	CreatedBy      int    `json:"createdBy,string,omitempty"`
	UpdatedAt      string `json:"updatedAt,omitempty"`
	UpdatedBy      int    `json:"updatedBy,string,omitempty"`
}

// EmergencyContact is the structure which holds one employee's emergency contact.
type OtherInformation struct {
	ID         int    `json:"id,string,omitempty"`
	Education  string `json:"education"`
	Experience string `json:"experience"`
	Notes      string `json:"notes"`
	CreatedAt  string `json:"createdAt,omitempty"`
	CreatedBy  int    `json:"createdBy,string,omitempty"`
	UpdatedAt  string `json:"updatedAt,omitempty"`
	UpdatedBy  int    `json:"updatedBy,string,omitempty"`
}

// Employment is the structure which holds one employee's employment details
type Employment struct {
	ID                  int    `json:"id,string,omitempty"`
	JobTitle            string `json:"jobTitle"`
	Department          int    `json:"department,string"`
	Superior            int    `json:"superior,string"`
	Supervisor          int    `json:"supervisor,string"`
	EmployeeType        int    `json:"employeeType,string"`
	WagesType           int    `json:"wagesType,string"`
	BasicRate           string `json:"basicRate"`
	PayFrequency        int    `json:"payFrequency,string"`
	PaymentBy           int    `json:"paymentBy,string"`
	BankPayout          int    `json:"bankPayout,string"`
	Group               int    `json:"group,string"`
	Branch              int    `json:"branch,string"`
	Project             int    `json:"project,string"`
	Overtime            int    `json:"overtime,string"`
	WorkingPermitExpiry string `json:"workingPermitExpiry"`
	JoinDate            string `json:"joinDate"`
	ConfirmDate         string `json:"confirmDate"`
	ResignDate          string `json:"resignDate"`
	CreatedAt           string `json:"createdAt,omitempty"`
	CreatedBy           int    `json:"createdBy,string,omitempty"`
	UpdatedAt           string `json:"updatedAt,omitempty"`
	UpdatedBy           int    `json:"updatedBy,string,omitempty"`
}

// EmploymentArchive is the structure which holds one employee's employment archive details
type EmploymentArchive struct {
	ID                  int
	JobTitle            string
	Department          string
	Superior            int
	Supervisor          int
	EmployeeType        string
	WagesType           string
	BasicRate           string
	PayFrequency        string
	PaymentBy           string
	BankPayout          string
	Group               string
	Branch              string
	Project             string
	Overtime            string
	WorkingPermitExpiry string
	JoinDate            string
	ConfirmDate         string
	ResignDate          string
	CreatedAt           string
	CreatedBy           int
}

// Spouse is the structure which holds one employee's spouse details
type Spouse struct {
	ID                    int    `json:"id,string,omitempty"`
	Fullname              string `json:"fullname"`
	ICNumber              string `json:"icNumber"`
	PassporNumber         string `json:"passporNumber"`
	IsDisabled            int    `json:"isDisabled,string"`
	IsWorking             int    `json:"isWorking,string"`
	TaxNumber             string `json:"taxNumber"`
	DeductibleChildNumber string `json:"deductibleChildNumber"`
	DeductibleChildAmount string `json:"deductibleChildAmount"`
	PrimaryPhone          string `json:"primaryPhone"`
	SecondaryPhone        string `json:"secondaryPhone"`
	Streetaddr1           string `json:"streetaddr1"`
	Streetaddr2           string `json:"streetaddr2"`
	City                  string `json:"city"`
	State                 string `json:"state"`
	Zip                   string `json:"zip"`
	Country               int    `json:"country,string"`
	CreatedAt             string `json:"createdAt,omitempty"`
	CreatedBy             int    `json:"createdBy,string,omitempty"`
	UpdatedAt             string `json:"updatedAt,omitempty"`
	UpdatedBy             int    `json:"updatedBy,string,omitempty"`
}

// Statutory is the structure which holds one employee's statutory details
type Statutory struct {
	ID                int     `json:"id,string,omitempty"`
	EPFTable          int     `json:"epfTable,string"`
	EPFNumber         string  `json:"epfNumber"`
	EPFInitial        string  `json:"epfInitial"`
	NK                string  `json:"nk"`
	EPFBorne          int     `json:"epfBorne,string"`
	SOCSOCategory     int     `json:"socsoCategory,string"`
	SOCSONumber       string  `json:"socsoNumber"`
	SOCSOStatus       int     `json:"socsoStatus,string"`
	SOCSOBorne        int     `json:"socsoBorne,string"`
	ContributeEIS     int     `json:"contributeEIS,string"`
	EISBorne          int     `json:"eisBorne,string"`
	TaxStatus         int     `json:"taxStatus,string"`
	TaxNumber         string  `json:"taxNumber"`
	TaxBranch         int     `json:"taxBranch,string"`
	EASerial          string  `json:"eaSerial"`
	TaxBorne          int     `json:"taxBorne,string"`
	ForeignWorkerLevy int     `json:"foreignWorkerLevy,string"`
	ZakatNumber       string  `json:"zakatNumber"`
	ZakatAmount       float32 `json:"zakatAmount,string"`
	TabungHajiNumber  string  `json:"tabungHajiNumber"`
	TabungHajiAmount  float32 `json:"tabungHajiAmount,string"`
	ASNNumber         string  `json:"asnNumber"`
	ASNAmount         float32 `json:"asnAmount,string"`
	ContributeHRDF    int     `json:"contributeHRDF,string"`
	CreatedAt         string  `json:"createdAt,omitempty"`
	CreatedBy         int     `json:"createdBy,string,omitempty"`
	UpdatedAt         string  `json:"updatedAt,omitempty"`
	UpdatedBy         int     `json:"updatedBy,string,omitempty"`
}

// StatutoryArchive is the structure which holds one employee's statutory archive details
type StatutoryArchive struct {
	ID                int
	EPFTable          string
	EPFNumber         string
	EPFInitial        string
	NK                string
	EPFBorne          int
	SOCSOCategory     string
	SOCSONumber       string
	SOCSOStatus       string
	SOCSOBorne        int
	ContributeEIS     int
	EISBorne          int
	TaxStatus         string
	TaxNumber         string
	TaxBranch         string
	EASerial          string
	TaxBorne          int
	ForeignWorkerLevy string
	ZakatNumber       string
	ZakatAmount       float32
	TabungHajiNumber  string
	TabungHajiAmount  float32
	ASNNumber         string
	ASNAmount         float32
	ContributeHRDF    int
	CreatedAt         string
	CreatedBy         int
	UpdatedAt         string
	UpdatedBy         int
}

// Employee is the structure which holds one employee full details
type Employee struct {
	ID                int    `json:"id,string,omitempty"`
	EmployeeCode      string `json:"employeeCode"`
	Fullname          string `json:"fullname"`
	Nickname          string `json:"nickname"`
	IcNumber          string `json:"icNumber"`
	PassportNumber    string `json:"passportNumber"`
	PassportExpiryAt  string `json:"passportExpiryAt"`
	Birthdate         string `json:"birthdate"`
	Nationality       int    `json:"nationality,string"`
	Residence         int    `json:"residence,string"`
	Maritalstatus     int    `json:"maritalstatus,string"`
	Gender            int    `json:"gender,string"`
	Race              int    `json:"race,string"`
	Religion          int    `json:"religion,string"`
	Streetaddr1       string `json:"streetaddr1"`
	Streetaddr2       string `json:"streetaddr2"`
	City              string `json:"city"`
	Zip               string `json:"zip"`
	State             string `json:"state"`
	Country           int    `json:"country,string"`
	PrimaryPhone      string `json:"primaryPhone"`
	SecondaryPhone    string `json:"secondaryPhone"`
	PrimaryEmail      string `json:"primaryEmail"`
	SecondaryEmail    string `json:"secondaryEmail"`
	IsForeigner       int    `json:"isForeigner,string"`
	ImmigrationNumber string `json:"immigrationNumber"`
	IsDisabled        int    `json:"isDisabled,string"`
	IsActive          int    `json:"isActive,string"`
	Role              int    `json:"role,string"`
	ConnectedUser     string `json:"connectedUser,omitempty"`
	CreatedAt         string `json:"createdAt,omitempty"`
	CreatedBy         int    `json:"createdBy,omitempty"`
	UpdatedAt         string `json:"updatedAt,omitempty"`
	UpdatedBy         int    `json:"updatedBy,omitempty"`
}

// Employee is the structure which holds one employee full details
type EmployeeFull struct {
	Employee          *Employee            `json:"Employee"`
	EmergencyContact  *EmergencyContact    `json:"EmergencyContact"`
	OtherInformation  *OtherInformation    `json:"OtherInformation"`
	Spouse            *Spouse              `json:"Spouse"`
	Employment        *Employment          `json:"Employment"`
	EmploymentArchive []*EmploymentArchive `json:"EmploymentArchive"`
	Statutory         *Statutory           `json:"Statutory"`
	StatutoryArchive  []*StatutoryArchive  `json:"StatutoryArchive"`
	Bank              *Bank                `json:"Bank"`
}

// EmployeeSummary is the structure which holds one employee summary
type EmployeeSummary struct {
	ID        int
	Code      string
	Fullname  string
	Nickname  string
	Email     string
	Mobile    string
	Birthdate string
	Race      string
	Gender    int
}

// AllEmployeeSummary is the structure which holds all employees summary (active, inactive & deleted)
type AllEmployees struct {
	Active   []*EmployeeSummary
	Inactive []*EmployeeSummary
	Deleted  []*EmployeeSummary
}

// ConfigTables is the structure which holds all employee's config tables
type ConfigTables struct {
	Country                []configC
	Race                   []configT
	Relationship           []configT
	Religion               []configT
	EmploymentBank         []configT
	EmploymentBranch       []configT
	EmploymentDepartment   []configT
	EmploymentGroup        []configT
	EmploymentOT           []configT
	EmploymentPaymentBy    []configT
	EmploymentPayFrequency []configT
	EmploymentProject      []configT
	EmploymentRule         []configT
	EmploymentType         []configT
	EmploymentWages        []configT
	StatutoryEPF           []configT
	StatutoryForeignLevy   []configT
	StatutorySOCSOCategory []configT
	StatutorySOCSOStatus   []configT
	StatutoryTaxBranch     []configT
	StatutoryTaxStatus     []configT
}

// configT holds one config table struct
type configT struct {
	ID   string
	Name string
}

// configC holds country config table struct
type configC struct {
	ID          string
	Name        string
	Code2       string
	Code3       string
	Nationality string
}

// employeeList holds the name of one employee
type employeeList struct {
	ID       int
	Code     string
	Fullname string
	Nickname string
}

// employeeLeave holds the information required for leave application
type employeeLeave struct {
	ID          int
	Code        string
	Fullname    string
	Nickname    string
	JoinDate    string
	ConfirmDate string
}
