package db

import (
	"database/sql"
	"fmt"

	// NOTE: used for mysql client plugin
	_ "github.com/go-sql-driver/mysql"
)

func NewClient() (*sql.DB, error) {
	// TODO: make this configurable
	db, err := sql.Open("mysql", "root:@(localhost:3306)/practice")
	if err != nil {
		return nil, fmt.Errorf("failed to create MySQL client: %w", err)
	}
	return db, nil
}
