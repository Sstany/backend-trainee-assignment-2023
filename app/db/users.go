package db

import (
	"context"
	"database/sql"
	"fmt"
	"segmenty/app/db/models"

	"go.uber.org/zap"
)

func (r *Database) FetchUser(ctx context.Context, userID int) (*models.User, bool, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	var user models.User

	err := r.db.QueryRowContext(ctx, queryFetchUser, userID).Scan(&user.ID, &user.Username)
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

func (r *Database) UpdateUserSegments(
	ctx context.Context,
	user *models.User,
	update *models.Update,
) *models.Response {
	ctx, cancel := context.WithTimeout(ctx, extendedTimeout)
	defer cancel()

	var response models.Response

	if len(update.Add) > 0 {
		response.Added = r.addUserSegments(ctx, user.ID, update.Add)
	}

	if len(update.Delete) > 0 {
		response.Deleted = r.deleteUserSegments(ctx, user.ID, update.Delete)
	}

	return &response
}

func (r *Database) ListAllUserSegments(ctx context.Context, user *models.User) (*[]models.Segment, error) {
	ctx, cancel := context.WithTimeout(ctx, extendedTimeout)
	defer cancel()

	var allSegments []models.Segment

	rows, err := r.db.QueryContext(ctx, queryFetchAllUserSegments, user.ID)
	if err != nil {
		return nil, fmt.Errorf("while fetching all segments: %w", err)
	}

	defer rows.Close()

	var tempSegment models.Segment

	for rows.Next() {
		err := rows.Scan(&tempSegment.ID, &tempSegment.Name)
		if err != nil {
			return &allSegments, fmt.Errorf("while scanning segment: %w", err)
		}

		allSegments = append(allSegments, tempSegment)
	}

	return &allSegments, rows.Err()
}

func (r *Database) addUserSegments(
	ctx context.Context,
	userId int,
	addUpdate []string,
) *models.UpdateStats {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	var updateStats models.UpdateStats

	for _, segmentName := range addUpdate {
		segment, isNotExists, err := r.FetchSegment(ctx, segmentName)
		if err != nil || isNotExists {
			r.logger.Error("", zap.Error(err))
			updateStats.Skipped = append(updateStats.Skipped, segmentName)

			continue
		}

		if err := r.insertSegmentUser(ctx, segment, userId); err != nil {
			r.logger.Error("", zap.Error(err))
			updateStats.Failed = append(updateStats.Failed, segmentName)

			continue
		}

		updateStats.Processed = append(updateStats.Processed, segmentName)
	}

	return &updateStats
}

func (r *Database) deleteUserSegments(
	ctx context.Context,
	userId int,
	deleteUpdate []string,
) *models.UpdateStats {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	var updateStats models.UpdateStats

	for _, segmentName := range deleteUpdate {
		segment, isNotExists, err := r.FetchSegment(ctx, segmentName)
		if err != nil || isNotExists {
			r.logger.Error("", zap.Error(err))
			updateStats.Skipped = append(updateStats.Skipped, segmentName)

			continue
		}

		if err := r.deleteSegmentUser(ctx, segment, userId); err != nil {
			r.logger.Error("", zap.Error(err))
			if err == errNothingChanged {
				updateStats.Skipped = append(updateStats.Skipped, segmentName)

				continue
			}

			updateStats.Failed = append(updateStats.Failed, segmentName)

			continue
		}

		updateStats.Processed = append(updateStats.Processed, segmentName)
	}

	return &updateStats
}

func (r *Database) insertSegmentUser(ctx context.Context, segment *models.Segment, userId int) error {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	if _, err := r.db.ExecContext(ctx, queryInsertSegmentUser, userId, segment.ID); err != nil {
		return fmt.Errorf("while inserting segment into user: %w", err)
	}

	if err := r.insertHistory(ctx, userId, models.TypeAdd, segment.Name); err != nil {
		return err
	}

	return nil
}

func (r *Database) deleteSegmentUser(ctx context.Context, segment *models.Segment, userId int) error {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	res, err := r.db.ExecContext(ctx, queryDeleteSegmentUser, userId, segment.ID)
	if err != nil {
		return fmt.Errorf("while deleting segment from user: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("while deleting segment from user: %w", err)
	}

	if rowsAffected == 0 {
		return errNothingChanged
	}

	if err := r.insertHistory(ctx, userId, models.TypeDelete, segment.Name); err != nil {
		return err
	}

	return nil
}

func (r *Database) ListUsers(ctx context.Context) (*[]models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	var allUsers []models.User

	rows, err := r.db.QueryContext(ctx, queryFetchAllUsers)
	if err != nil {
		return nil, fmt.Errorf("while fetching all users: %w", err)
	}

	defer rows.Close()

	var tempUser models.User

	for rows.Next() {
		err := rows.Scan(&tempUser.ID, &tempUser.Username)
		if err != nil {
			return &allUsers, fmt.Errorf("while scanning user: %w", err)
		}

		allUsers = append(allUsers, tempUser)
	}

	return &allUsers, rows.Err()
}

func (r *Database) DeleteUser(ctx context.Context, userId int) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	var user models.User

	if err := r.db.QueryRowContext(ctx, queryDeleteUser, userId).Scan(&user.ID, &user.Username); err != nil {
		return &user, fmt.Errorf("while deleting user: %w", err)
	}

	return &user, nil

}
