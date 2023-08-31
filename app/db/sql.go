package db

import "time"

const (
	queryFetchUser = `SELECT * FROM users WHERE userId = $1`
)

const (
	defaultTimeout = 5 * time.Second
)
