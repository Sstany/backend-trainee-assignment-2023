package db

import "time"

const (
	queryFetchUser           = `SELECT * FROM users WHERE userId = $1`
	queryInsertUser          = `INSERT INTO users(username) VALUES ($1) RETURNING userId`
	queryCreateUsersTable    = `CREATE TABLE users (userId serial PRIMARY KEY,username text NOT NULL)`
	queryCreateSegmentsTable = `CREATE TABLE segments (segmentId serial PRIMARY KEY,name text NOT NULL UNIQUE)`
	queryInsertSegment       = `INSERT INTO segments(name) VALUES ($1) RETURNING segmentId`
	queryFetchAllSegments    = `SELECT * FROM segments`
)

const (
	defaultTimeout  = 5 * time.Second
	extendedTimeout = 30 * time.Second
)
