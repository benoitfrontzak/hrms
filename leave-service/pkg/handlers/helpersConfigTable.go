package handlers

// getTableNameDB returns the name of the config table as label in the DB
func getTableNameDB(ct string) string {
	var dbTable string
	switch ct {
	case "Category":
		dbTable = "CONFIG_CLAIM_CATEGORY"
	case "Status":
		dbTable = "CONFIG_CLAIM_STATUS"
	default:
		dbTable = ""
	}

	return dbTable
}
