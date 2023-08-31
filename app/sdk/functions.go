package sdk

import "github.com/lib/pq"

func IsDublicateTableErr(err error) bool {
	if pqErr, ok := err.(*pq.Error); ok && pqErr.Code.Name() == "duplicate_table" {
		return true
	}
	return false
}

func IsUniqueViolationErr(err error) bool {
	if pqErr, ok := err.(*pq.Error); ok && pqErr.Code.Name() == "unique_violation" {
		return true
	}

	return false
}
