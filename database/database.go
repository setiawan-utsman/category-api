package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func InitDB(connectionString string) (*sql.DB, error) {
	// Buka database
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// Coba koneksi
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Set Koneksi
	maxIdleConns := 10
	maxOpenConns := 100
	connMaxLifetime := 5 * time.Minute

	db.SetMaxIdleConns(maxIdleConns)
	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxLifetime(connMaxLifetime)

	log.Println("Database Connected")
	return db, nil
}
