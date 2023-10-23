package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	stringConnect := os.Getenv("DB_CONNECTION_STRING")

	if stringConnect == "" {
		return nil, fmt.Errorf("environment variable not defined or found")
	}

	db, err := sql.Open("mysql", stringConnect)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
