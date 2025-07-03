package utils

import (
	"database/sql"
	"strings"
)

const (
	PathSeparator = "/"
)

type Integer interface {
	~int | ~int32 | ~int64
}

func ToNullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: s != ""}
}

func NullStringToString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}

func ToNullBool(s bool) sql.NullBool {
	return sql.NullBool{Bool: s, Valid: true}
}

func NullBoolToBool(nb sql.NullBool) bool {
	if nb.Valid {
		return nb.Bool
	}
	return false
}

func IsEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}

func AppendWithSeparator(base, part, sep string) string {
	if strings.TrimSpace(part) == "" {
		return base
	}
	if base == "" {
		return part
	}
	base = strings.TrimRight(base, sep)
	part = strings.TrimLeft(part, sep)
	return base + sep + part
}

func CountParts[T Integer](s, sep string) T {
	parts := strings.Split(s, sep)
	var result []string
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	var count T
	count = T(len(result))
	return count
}
