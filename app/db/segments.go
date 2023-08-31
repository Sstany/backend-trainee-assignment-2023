package db

import (
	"context"
	"fmt"
	"segmenty/app/db/models"
	"segmenty/app/sdk"
)

func (r *Database) InsertSegment(ctx context.Context, segment *models.Segment) (int, bool, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	lastInsertId := 0
	if err := r.db.QueryRowContext(ctx, queryInsertSegment, segment.Name).Scan(&lastInsertId); err != nil {
		return lastInsertId, sdk.IsUniqueViolationErr(err), fmt.Errorf("while inserting segment: %w", err)
	}

	return lastInsertId, false, nil

}
func (r *Database) ListSegments(ctx context.Context) (*[]models.Segment, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	var allSegments []models.Segment

	rows, err := r.db.QueryContext(ctx, queryFetchAllSegments)
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
