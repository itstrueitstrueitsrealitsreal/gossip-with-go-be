// internal/database/database.go

package database

import (
	"database/sql"
	"fmt"
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
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"127.0.0.1", "5432", os.Getenv("DBUSER"), os.Getenv("DBPASS"), "gossip")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
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
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(128) UNIQUE NOT NULL
		);

		CREATE TABLE IF NOT EXISTS tags (
			id SERIAL PRIMARY KEY,
			name VARCHAR(16) NOT NULL
		);

		CREATE TABLE IF NOT EXISTS threads (
			id SERIAL PRIMARY KEY,
			author_id INT,
			tag_id INT,
			title VARCHAR(255) NOT NULL,
			content VARCHAR(1024) NOT NULL,
			FOREIGN KEY (author_id) REFERENCES users(id),
			FOREIGN KEY (tag_id) REFERENCES tags(id)
		);

		CREATE TABLE IF NOT EXISTS posts (
			id SERIAL PRIMARY KEY,
			thread_id INT,
			author_id INT,
			content VARCHAR(1024) NOT NULL,
			timestamp TIMESTAMP NOT NULL,
			FOREIGN KEY (thread_id) REFERENCES threads(id),
			FOREIGN KEY (author_id) REFERENCES users(id)
		);
	`)
	return err
}
