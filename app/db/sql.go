package db

import "time"

const (
	queryFetchUser           = `SELECT * FROM users WHERE userId = $1`
	queryInsertUser          = `INSERT INTO users(username) VALUES ($1) RETURNING userId`
	queryCreateUsersTable    = `CREATE TABLE users (userId serial PRIMARY KEY,username text NOT NULL)`
	queryCreateSegmentsTable = `CREATE TABLE segments (segmentId serial PRIMARY KEY,name text NOT NULL UNIQUE)`
)

const (
	defaultTimeout  = 5 * time.Second
	extendedTimeout = 30 * time.Second
)
