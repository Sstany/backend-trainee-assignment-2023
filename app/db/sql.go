package db

import (
	"errors"
	"time"
)

const (
	queryFetchUser    = `SELECT * FROM users WHERE userId = $1`
	queryFetchSegment = `SELECT * FROM segments WHERE name = $1`

	queryInsertUser          = `INSERT INTO users(username) VALUES ($1) RETURNING userId`
	queryCreateUsersTable    = `CREATE TABLE users (userId serial PRIMARY KEY,username text NOT NULL)`
	queryCreateSegmentsTable = `CREATE TABLE segments (segmentId serial PRIMARY KEY,name text NOT NULL UNIQUE)`
	queryInsertSegment       = `INSERT INTO segments(name) VALUES ($1) RETURNING segmentId`
	queryFetchAllSegments    = `SELECT * FROM segments`
	queryDeleteSegment       = `DELETE FROM segments WHERE name = $1 RETURNING *`
	queryDeleteSegmentUser   = `DELETE FROM segments_users su WHERE (su.userId = $1 and su.segmentId = $2)`
	queryInsertSegmentUser   = `INSERT INTO segments_users (userId, segmentId) VALUES ($1,$2)`
)

const (
	queryCreateSegmentsUsersTable = `CREATE TABLE segments_users (
		segmentId   int REFERENCES segments (segmentId) ON UPDATE CASCADE ON DELETE CASCADE,
		userId int REFERENCES users (userId) ON UPDATE CASCADE,
		CONSTRAINT segment_user_pkey PRIMARY KEY (segmentId, userId))`
	querySelectAllUserSegments = `SELECT s.segmentId, name
		FROM users u
		JOIN segments_users su on su.userId = u.userId
		JOIN segments s on su.segmentId = s.segmentId
		WHERE u.userId = $1`
)

const (
	defaultTimeout  = 5 * time.Second
	extendedTimeout = 30 * time.Second
)

var (
	errNothingChanged = errors.New("Nothing changed")
)
