package db

import (
	"context"
	"fmt"
	"segmenty/app/db/models"
	"time"
)

func (r *Database) insertHistory(ctx context.Context, userId int, historyType, segmentName string) error {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	if _, err := r.db.ExecContext(ctx, queryInsertHistory, userId, time.Now(), historyType, segmentName); err != nil {
		return fmt.Errorf("while inserting into history: %w", err)
	}

	return nil

}

func (r *Database) fetchUserhistory(ctx context.Context, userId int, startTime, endTime time.Time) (*[]models.History, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	var hist []models.History

	rows, err := r.db.QueryContext(ctx, queryFetchUserHistory, userId, startTime, endTime)
	if err != nil {
		return nil, fmt.Errorf("while fetching all users: %w", err)
	}

	defer rows.Close()

	var tempHist models.History

	for rows.Next() {
		fmt.Println(1)
		err := rows.Scan(&tempHist.ID, &tempHist.UserId, &tempHist.Time, &tempHist.Type, &tempHist.SegmentName)
		if err != nil {
			return &hist, fmt.Errorf("while scanning user: %w", err)
		}

		hist = append(hist, tempHist)
	}

	return &hist, rows.Err()
}

func (r *Database) ListAllUserHistory(ctx context.Context, userId int) (*[]models.History, error) {
	return r.fetchUserhistory(ctx, userId, time.Time{}, time.Now())
}
