package utils

import "database/sql"

func NullIntToInt64(n sql.NullInt64, defaultVal int64) int64 {
	if n.Valid {
		return n.Int64
	}
	return defaultVal
}

func ToNullInt64(val int64) sql.NullInt64 {
	return sql.NullInt64{
		Int64: val,
		Valid: true,
	}
}
