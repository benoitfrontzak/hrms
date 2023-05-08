package pg

import (
	"context"
	"time"
)

// get all approved leaves by date (today & tomorrow)
func (l *Leave) GetAllLeaveToday() (*AllCurentLeave, error) {
	// set today's & tomorrow's date
	n := time.Now()
	t := n.AddDate(0, 0, 1)

	// fetch today's leaves
	today, err := getLeaveByDate(n.Format("2006-01-02"))
	if err != nil {
		return nil, err
	}

	// fetch tomorrow's leave
	tomorrow, err := getLeaveByDate(t.Format("2006-01-02"))
	if err != nil {
		return nil, err
	}

	// store all information
	all := AllCurentLeave{
		Today:    today,
		Tomorrow: tomorrow,
	}

	return &all, nil
}

func getLeaveByDate(myDate string) ([]*ApprovedLeave, error) {
	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which fetch employee summary
	query := `SELECT ld.code as leave_code,
					 ld.description as leave_definition,
   					 la.description,
   					 la.employee_id ,
   					 lad.requested_date, 
   					 lad.is_half 
			  FROM public."LEAVE_DEFINITION" ld, public."LEAVE_APPLICATION" la,  public."LEAVE_APPLICATION_DETAILS" lad 
			  WHERE lad.requested_date = $1
			  AND lad.leave_application_id = la.id
			  AND la.leave_definition_id = ld.id 
			  AND la.status_id = 4`

	// executes SQL query
	rows, err := db.QueryContext(ctx, query, myDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// populate returned rows to employee summary struct
	var all []*ApprovedLeave
	for rows.Next() {
		var myLeaves ApprovedLeave

		err = rows.Scan(
			&myLeaves.Code,
			&myLeaves.Description,
			&myLeaves.Reason,
			&myLeaves.EmployeeID,
			&myLeaves.RequestedDate,
			&myLeaves.IsHalf,
		)
		if err != nil {
			return nil, err
		}

		all = append(all, &myLeaves)
	}

	// return all employee summary rows
	return all, nil
}
