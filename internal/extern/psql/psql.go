package psql

import (
	"context"
	"database/sql"
	"os"
	"time"

	"brewnique.fdunlap.com/internal/jsonlog"
)

type PsqlConfig struct {
	Dsn             string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

type PostgresProvider struct {
	db *sql.DB
}

func NewPsqlProvider(config PsqlConfig) *PostgresProvider {
	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo, "psql")
	logger.PrintInfo("Opening DB connection", nil)
	db, err := sql.Open("postgres", config.Dsn)
	if err != nil {
		panic(err)
	}
	logger.PrintInfo("DB connection opened", nil)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetConnMaxLifetime(config.ConnMaxLifetime)

	err = db.PingContext(ctx)
	if err != nil {
		panic(err)
	}

	logger.PrintInfo("Successfully connected to database", nil)

	return &PostgresProvider{
		db: db,
	}
}

func (p PostgresProvider) Close() {
	p.db.Close()
}
