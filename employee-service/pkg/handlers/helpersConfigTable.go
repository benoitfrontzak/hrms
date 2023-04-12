package handlers

// getTableNameDB returns the name of the config table as label in the DB
func getTableNameDB(ct string) string {
	var dbTable string
	switch ct {
	case "Race":
		dbTable = "CONFIG_COMMON_RACE"
	case "Relationship":
		dbTable = "CONFIG_COMMON_RELATIONSHIP"
	case "Religion":
		dbTable = "CONFIG_COMMON_RELIGION"
	case "EmploymentBank":
		dbTable = "CONFIG_EMPLOYMENT_BANK"
	case "EmploymentBranch":
		dbTable = "CONFIG_EMPLOYMENT_BRANCH"
	case "EmploymentDepartment":
		dbTable = "CONFIG_EMPLOYMENT_DEPARTMENT"
	case "EmploymentGroup":
		dbTable = "CONFIG_EMPLOYMENT_GROUP"
	case "EmploymentOT":
		dbTable = "CONFIG_EMPLOYMENT_OVERTIME"
	case "EmploymentPaymentBy":
		dbTable = "CONFIG_EMPLOYMENT_PAYMENT_BY"
	case "EmploymentPayFrequency":
		dbTable = "CONFIG_EMPLOYMENT_PAY_FREQUENCY"
	case "EmploymentProject":
		dbTable = "CONFIG_EMPLOYMENT_PROJECT"
	case "EmploymentType":
		dbTable = "CONFIG_EMPLOYMENT_TYPE"
	case "EmploymentWages":
		dbTable = "CONFIG_EMPLOYMENT_WAGES"
	case "StatutoryEPF":
		dbTable = "CONFIG_STATUTORY_EPF"
	case "StatutoryForeignLevy":
		dbTable = "CONFIG_STATUTORY_FOREIGN_LEVY"
	case "StatutorySOCSOCategory":
		dbTable = "CONFIG_STATUTORY_SOCSO_CATEGORY"
	case "StatutorySOCSOStatus":
		dbTable = "CONFIG_STATUTORY_SOCSO_STATUS"
	case "StatutoryTaxBranch":
		dbTable = "CONFIG_STATUTORY_TAX_BRANCH"
	case "StatutoryTaxStatus":
		dbTable = "CONFIG_STATUTORY_TAX_STATUS"
	default:
		dbTable = ""
	}

	return dbTable
}
