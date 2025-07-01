package utils

import "database/sql"

func ToNullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: s != ""}
}

func ToNullBool(s bool) sql.NullBool {
	return sql.NullBool{Bool: s, Valid: true}
}
