package database

import "database/sql"

type Database struct {
	DB *sql.DB
}

// GetDB initialises a new Database instance.
func GetDB() (*Database, error) {
	//db, err := sql.Open("mysql", "gossip")
	//if err != nil {
	//	return nil, err
	//}

	//return &Database{DB: db}, nil
	return &Database{}, nil
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
			username VARCHAR(128) UNIQUE NOT NULL
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
