package psql

import (
	"fmt"
	"strconv"
	config "test-module/internal/configs"

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
			"user=%s dbname=%s password=%s host=%s port=%s",
			config.PSQLUser,
			config.PSQLName,
			config.PSQLPass,
			config.PSQLHost,
			strconv.Itoa(int(config.PSQLPort)),
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
