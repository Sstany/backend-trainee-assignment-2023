package db

import (
	"errors"
	"time"
)

const (
	queryCreateSegmentsTable = `CREATE TABLE segments (segmentId serial PRIMARY KEY,name text NOT NULL UNIQUE)`
	queryFetchSegment        = `SELECT * FROM segments WHERE name = $1`
	queryFetchAllSegments    = `SELECT * FROM segments`
	queryInsertSegment       = `INSERT INTO segments(name) VALUES ($1) RETURNING segmentId`
	queryDeleteSegment       = `DELETE FROM segments WHERE name = $1 RETURNING *`
)

//
// Users
//

const (
	queryCreateUsersTable = `CREATE TABLE users (userId serial PRIMARY KEY,username text NOT NULL)`
	queryFetchUser        = `SELECT * FROM users WHERE userId = $1`
	queryFetchAllUsers    = `SELECT * FROM users`
	queryInsertUser       = `INSERT INTO users(username) VALUES ($1) RETURNING userId`
	queryDeleteUser       = `DELETE FROM users WHERE userId = $1 RETURNING *`
)

//
// History
//

const (
	queryCreateHistoryTable = `CREATE TABLE history (
		historyId serial PRIMARY KEY,
		userId int NOT NULL,
		time TIMESTAMP NOT NULL,
		type text NOT NULL,
		segment text NOT NULL)`

	queryFetchUserHistory = `SELECT * FROM history WHERE userId = $1 AND time > $2 AND time < $3`
	queryInsertHistory    = `INSERT INTO history (userId, time, type, segment)
VALUES ($1,$2, $3, $4);`
)

//
// Segments and Users
//

const (
	queryCreateSegmentsUsersTable = `CREATE TABLE segments_users (
		segmentId   int REFERENCES segments (segmentId) ON UPDATE CASCADE ON DELETE CASCADE,
		userId int REFERENCES users (userId) ON UPDATE CASCADE,
		CONSTRAINT segment_user_pkey PRIMARY KEY (segmentId, userId))`
	queryFetchAllUserSegments = `SELECT s.segmentId, name
		FROM users u
		JOIN segments_users su on su.userId = u.userId
		JOIN segments s on su.segmentId = s.segmentId
		WHERE u.userId = $1`

	queryDeleteSegmentUser = `DELETE FROM segments_users su WHERE (su.userId = $1 and su.segmentId = $2)`
	queryInsertSegmentUser = `INSERT INTO segments_users (userId, segmentId) VALUES ($1,$2)`
)

const (
	defaultTimeout  = 5 * time.Second
	extendedTimeout = 30 * time.Second
)

var (
	errNothingChanged = errors.New("Nothing changed")
)
