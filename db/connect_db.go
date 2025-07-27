package db

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

type Conn_DB struct {
	DB *sql.DB
}

func NewDBConnection(dsn string) (*Conn_DB, error) {

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database connection %w", err)
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("database closed due to error")
	}

	log.Print("Database connection success")

	return &Conn_DB{DB: db}, nil
}