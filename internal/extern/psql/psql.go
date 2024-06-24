package psql

import (
	"context"
	"database/sql"
	"log"
	"time"
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
	db, err := sql.Open("postgres", config.Dsn)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetConnMaxLifetime(config.ConnMaxLifetime)

	err = db.PingContext(ctx)
	if err != nil {
		panic(err)
	}

	log.Printf("Connected to database")

	return &PostgresProvider{
		db: db,
	}
}

func (p PostgresProvider) Close() {
	p.db.Close()
}
