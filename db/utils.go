package db

import (
	"context"
	"database/sql"
	_ "embed"

	_ "modernc.org/sqlite"
)

//go:embed schema.sql
var ddl string

func InitDB() error {
	ctx := context.Background()

	// will allow user to set db location via flag on init
	database, err := sql.Open("sqlite", "file:reqord.db")
	if err != nil {
		return err
	}

	// create tables
	if _, err := database.ExecContext(ctx, ddl); err != nil {
		return err
	}

	return nil
}

func GetQueries() (*Queries, error) {
	database, err := sql.Open("sqlite", "file:reqord.db")
	if err != nil {
		return nil, err
	}

	queries := New(database)
	return queries, nil
}
