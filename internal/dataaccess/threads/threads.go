package threads

import (
	"github.com/CVWO/sample-go-app/internal/database"
	"github.com/CVWO/sample-go-app/internal/models"
)

// List retrieves a list of threads from the database.
func List(db *database.Database) ([]models.Thread, error) {
	// Query to select threads from the database
	query := "SELECT id, author_id, tag_id, title, content FROM threads"

	// Execute the query
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Initialize a slice to store the retrieved threads
	var threads []models.Thread

	// Iterate through the rows and populate the threads slice
	for rows.Next() {
		var thread models.Thread
		err := rows.Scan(&thread.ID, &thread.AuthorID, &thread.TagID, &thread.Title, &thread.Content)
		if err != nil {
			return nil, err
		}
		threads = append(threads, thread)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return threads, nil
}

// GetThreadByID retrieves a thread by ID from the database
func GetThreadByID(db *database.Database, threadID string) (*models.Thread, error) {
	query := "SELECT id, author_id, tag_id, title, content FROM threads WHERE id = ?"
	var thread models.Thread

	err := db.DB.QueryRow(query, threadID).Scan(&thread.ID, &thread.AuthorID, &thread.TagID, &thread.Title, &thread.Content)
	if err != nil {
		return nil, err
	}
	return &thread, nil
}
