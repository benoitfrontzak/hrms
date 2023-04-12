package handlers

// getTableNameDB returns the name of the config table as label in the DB
func getTableNameDB(ct string) string {
	var dbTable string
	switch ct {
	case "Category":
		dbTable = "CONFIG_LEAVE_CATEGORY"
	case "Status":
		dbTable = "CONFIG_LEAVE_STATUS"
	default:
		dbTable = ""
	}

	return dbTable
}
