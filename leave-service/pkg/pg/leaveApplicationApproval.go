package pg

import (
	"context"
	"time"
)

// approve leave by id
func (l *Leave) Approve(rowID, userID int) error {
	now := time.Now().Format("2006-01-02")

	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which update the leave (approve it)
	stmt := `UPDATE public."LEAVE_APPLICATION" SET status_id=4, approved_at=$1, approved_by=$2 WHERE id=$3;`

	// executes SQL query
	_, err := db.ExecContext(ctx, stmt, now, userID, rowID)
	if err != nil {
		return err
	}

	// return error
	return nil
}

// reject leave by id
func (l *Leave) Reject(rowID, userID int, reason string) error {
	now := time.Now().Format("2006-01-02")

	// canceling this context releases resources associated with it
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// SQL statement which update the leave (approve it)
	stmt := `UPDATE public."LEAVE_APPLICATION" SET status_id=3, approved_at=$1, approved_by=$2, rejected_reason=$3 WHERE id=$4;`

	// executes SQL query
	_, err := db.ExecContext(ctx, stmt, now, userID, reason, rowID)
	if err != nil {
		return err
	}

	// return error
	return nil
}
