package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"segmenty/app/sdk"

	"go.uber.org/zap"

	_ "github.com/lib/pq"
)

const envPostgres = "POSTGRES_URI"

type Database struct {
	db     *sql.DB
	logger *zap.Logger
}

func NewDB() *Database {
	DB := &Database{}

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	DB.logger = logger

	db, err := sql.Open("postgres", os.Getenv(envPostgres))
	if err != nil {
		DB.logger.Error("", zap.Error(err))

		return nil
	}

	DB.db = db

	return DB
}

func (r *Database) Init(ctx context.Context) bool {
	ctx, cancel := context.WithTimeout(ctx, extendedTimeout)
	defer cancel()

	var err error

	_, err = r.db.ExecContext(ctx, queryCreateUsersTable)
	if err != nil && !sdk.IsDublicateTableErr(err) {
		r.logger.Error("", zap.Error(fmt.Errorf("while creating users table: %w", err)))

		return false
	}

	_, err = r.db.ExecContext(ctx, queryCreateSegmentsTable)
	if err != nil && !sdk.IsDublicateTableErr(err) {
		r.logger.Error("", zap.Error(fmt.Errorf("while creating segments table: %w", err)))

		return false
	}
	return true

}
