package store

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/pressly/goose"
)

func Open() (*sql.DB, error) {
	db, err := sql.Open("pgx", "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable")

	if err != nil {
		return nil, fmt.Errorf("db : open %w", err)
	}

	fmt.Println("Connected to database....")
	return db, nil
}

func Migrate(db *sql.DB, dir string) error {
	err := goose.SetDialect("postgres")
	if err != nil {
		return fmt.Errorf("migrate : %w", err)
	}
	err = goose.Up(db, dir)
	if err != nil {
		return fmt.Errorf("goose up :%w", err)
	}
	return err
}
