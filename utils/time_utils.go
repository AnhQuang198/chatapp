package utils

import (
	"database/sql"
	"time"
)

func NullTimeToTime(nt sql.NullTime) *time.Time {
	if nt.Valid {
		return &nt.Time
	}
	return nil
}
