package psqlbd

import (
	"fmt"
	config "test-module/internal/config"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	*sqlx.DB
}

func Connect(config config.Config) (*DB, error) {
	db, err := sqlx.Open(
		"pgx",
		fmt.Sprintf(
			"user=%s dbname=%s password=%s host=%s port=%d",
			config.PSQLUser,
			config.PSQLName,
			config.PSQLPass,
			config.PSQLHost,
			config.PSQLPort,
		),
	)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func Close(db *DB) error {
	err := db.Close()
	return err
}
