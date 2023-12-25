package threads

import (
	"database/sql"
	"github.com/itstrueitstrueitsrealitsreal/gossip-with-go-be/internal/database"
	"github.com/itstrueitstrueitsrealitsreal/gossip-with-go-be/internal/models"
	"strconv"
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

// GetThreadByID retrieves a thread by ID from the database.
func GetThreadByID(db *database.Database, threadID string) (*models.Thread, error) {
	query := "SELECT id, author_id, tag_id, title, content FROM threads WHERE id = ?"
	var thread models.Thread

	err := db.DB.QueryRow(query, threadID).Scan(&thread.ID, &thread.AuthorID, &thread.TagID, &thread.Title, &thread.Content)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &thread, nil
}

// Create inserts a new thread into the database.
func Create(db *database.Database, input models.ThreadInput) (*models.Thread, error) {
	query := "INSERT INTO threads (author_id, tag_id, title, content) VALUES (?, ?, ?, ?)"
	result, err := db.DB.Exec(query, input.AuthorID, input.TagID, input.Title, input.Content)
	if err != nil {
		return nil, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	thread := &models.Thread{
		ID:       int(lastInsertID),
		AuthorID: input.AuthorID,
		TagID:    input.TagID,
		Title:    input.Title,
		Content:  input.Content,
	}

	return thread, nil
}

// Update updates an existing thread in the database.
func Update(db *database.Database, threadID string, input models.ThreadInput) (*models.Thread, error) {
	query := "UPDATE threads SET author_id = ?, tag_id = ?, title = ?, content = ? WHERE id = ?"
	_, err := db.DB.Exec(query, input.AuthorID, input.TagID, input.Title, input.Content, threadID)
	if err != nil {
		return nil, err
	}

	thread := &models.Thread{
		ID:       atoi(threadID),
		AuthorID: input.AuthorID,
		TagID:    input.TagID,
		Title:    input.Title,
		Content:  input.Content,
	}

	return thread, nil
}

// Delete removes a thread from the database.
func Delete(db *database.Database, threadID string) error {
	query := "DELETE FROM threads WHERE id = ?"
	_, err := db.DB.Exec(query, threadID)
	if err != nil {
		return err
	}

	return nil
}

// Atoi converts ASCII values to integers.
func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
