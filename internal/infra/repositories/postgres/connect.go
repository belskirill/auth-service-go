package postgres

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"

	"go.uber.org/zap"
)

func Connect(ctx context.Context, dsn string, logger *zap.Logger) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Error("failed to connect to postgres", zap.Error(err))
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		if err := db.Close(); err != nil {
			logger.Warn("failed to close database", zap.Error(err))
			return nil, err
		}
		logger.Error("failed to ping postgres", zap.Error(err))
		return nil, err
	}

	logger.Info("successfully connected to postgres")
	return db, nil
}

func MustConnect(ctx context.Context, dsn string, logger *zap.Logger) *sql.DB {
	db, err := Connect(ctx, dsn, logger)
	if err != nil {
		panic(err)
	}

	return db
}
