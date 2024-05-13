package db

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jmoiron/sqlx"
)

func NewConnection() (*sqlx.DB, error) {
	return sqlx.Connect(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		),
	)
}

func GenerateSchema(conn *sqlx.DB) error {
	path := filepath.Join("scripts", "init.sql")

	content, ioErr := os.ReadFile(path)
	if ioErr != nil {
		return ioErr
	}
	sql := string(content)
	if _, execErr := conn.Exec(sql); execErr != nil {
		return execErr
	}
	return nil
}
