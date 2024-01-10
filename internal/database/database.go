// internal/database/database.go

package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// Database represents the database connection.
type Database struct {
	DB *sql.DB
}

// GetDB initializes a new Database instance and opens a PostgreSQL database connection.
func GetDB() (*Database, error) {
	// Capture connection properties.
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		os.Getenv("DBHOST"),
		os.Getenv("DBPORT"),
		os.Getenv("DBUSER"),
		os.Getenv("DBPASS"),
		os.Getenv("DBNAME"),
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	fmt.Println("Connected to the database")

	return &Database{DB: db}, nil
}

// Close closes the database connection.
func (d *Database) Close() error {
	return d.DB.Close()
}

// CreateTables creates necessary tables in the database.
func (d *Database) CreateTables() error {
	_, err := d.DB.Exec(`
		CREATE TABLE users (
			id VARCHAR(50) PRIMARY KEY,
			username VARCHAR(128) UNIQUE NOT NULL,
			password VARCHAR(128) NOT NULL
		);

		CREATE TABLE tags (
			id VARCHAR(50) PRIMARY KEY,
			name VARCHAR(32) NOT NULL
		);

		CREATE TABLE threads (
			id VARCHAR(50) PRIMARY KEY,
			author_id VARCHAR(50),
			tag_id VARCHAR(50),
			title VARCHAR(255) NOT NULL,
			content VARCHAR(1024) NOT NULL,
			FOREIGN KEY (author_id) REFERENCES users(id),
			FOREIGN KEY (tag_id) REFERENCES tags(id)
		);

		CREATE TABLE comments (
			id VARCHAR(50) PRIMARY KEY,
			thread_id VARCHAR(50),
			author_id VARCHAR(50),
			content VARCHAR(1024) NOT NULL,
			timestamp TIMESTAMPTZ NOT NULL,
			FOREIGN KEY (thread_id) REFERENCES threads(id),
			FOREIGN KEY (author_id) REFERENCES users(id)
		);
	`)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
