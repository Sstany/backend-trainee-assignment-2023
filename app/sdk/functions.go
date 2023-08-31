package sdk

import "github.com/lib/pq"

func IsDublicateTableErr(err error) bool {
	if pqErr, ok := err.(*pq.Error); ok && pqErr.Code.Name() == "duplicate table" {
		return true
	}
	return false
}
