package db

import (
	"context"
	"database/sql"
	"fmt"
	"segmenty/app/db/models"
)

func (r *Database) FetchUser(ctx context.Context, userID int) (*models.User, bool, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	var user models.User

	err := r.db.QueryRowContext(ctx, queryFetchUser, userID).Scan(&user.Username)
	if err != nil {
		return nil, err == sql.ErrNoRows, fmt.Errorf("while query user: %w", err)
	}

	return &user, err == sql.ErrNoRows, nil
}

func (r *Database) InsertUser(ctx context.Context, user *models.User) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	lastInsertId := 0
	if err := r.db.QueryRowContext(ctx, queryInsertUser, user.Username).Scan(&lastInsertId); err != nil {
		return lastInsertId, fmt.Errorf("while inserting user: %w", err)
	}

	return lastInsertId, nil

}
