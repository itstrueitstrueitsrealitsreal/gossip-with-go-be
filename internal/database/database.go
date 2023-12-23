package database

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"os"
)

type Database struct {
	DB *sql.DB
}

// GetDB initializes a new Database instance and opens a MySQL database connection.
func GetDB() (*Database, error) {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "gossip",
	}
	dataSourceName := cfg.FormatDSN()

	db, err := sql.Open("mysql", dataSourceName)
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
			id INT PRIMARY KEY AUTO_INCREMENT,
			name VARCHAR(128) UNIQUE NOT NULL
		);

		CREATE TABLE IF NOT EXISTS tags (
			id INT PRIMARY KEY AUTO_INCREMENT,
			name VARCHAR(16) NOT NULL
		);

		CREATE TABLE IF NOT EXISTS threads (
			id INT PRIMARY KEY AUTO_INCREMENT,
			author_id INT,
			tag_id INT,
			title VARCHAR(255) NOT NULL,
			content VARCHAR(1024) NOT NULL,
			FOREIGN KEY (author_id) REFERENCES users(id),
			FOREIGN KEY (tag_id) REFERENCES tags(id)
		);

		CREATE TABLE IF NOT EXISTS posts (
			id INT PRIMARY KEY AUTO_INCREMENT,
			thread_id INT,
			author_id INT,
			content VARCHAR(1024) NOT NULL,
			timestamp DATETIME NOT NULL,
			FOREIGN KEY (thread_id) REFERENCES threads(id),
			FOREIGN KEY (author_id) REFERENCES users(id)
		);
	`)
	return err
}
