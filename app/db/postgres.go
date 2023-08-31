package db

import (
	"database/sql"
	"os"

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
