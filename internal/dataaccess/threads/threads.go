package threads

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/itstrueitstrueitsrealitsreal/gossip-with-go-be/internal/database"
	"github.com/itstrueitstrueitsrealitsreal/gossip-with-go-be/internal/models"
)

// List retrieves a list of threads from the database.
func List(db *database.Database) ([]models.Thread, error) {
	// Query to select threads from the database
	query := `
        SELECT threads.id, users.username, tags.name, threads.title, threads.content
        FROM threads
        INNER JOIN users ON threads.author_id = users.id
        INNER JOIN tags ON threads.tag_id = tags.id
    `
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
		err := rows.Scan(&thread.ID, &thread.Author, &thread.Tag, &thread.Title, &thread.Content)
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
	query := "SELECT id, author_id, tag_id, title, content FROM threads WHERE id = $1"
	var thread models.Thread

	err := db.DB.QueryRow(query, threadID).Scan(&thread.ID, &thread.Author, &thread.Tag, &thread.Title, &thread.Content)
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
	// Query to retrieve the author ID by name
	authorQuery := "SELECT id FROM users WHERE username = $1"
	var authorID string
	err := db.DB.QueryRow(authorQuery, input.Author).Scan(&authorID)
	if err != nil {
		return nil, err
	}

	// Query to retrieve the tag ID by name
	tagQuery := "SELECT id FROM tags WHERE name = $1"
	var tagID string
	err = db.DB.QueryRow(tagQuery, input.Tag).Scan(&tagID)
	if err != nil {
		return nil, err
	}

	query := "INSERT INTO threads (id, author_id, tag_id, title, content) VALUES ($1, $2, $3, $4, $5)"
	_, err = db.DB.Exec(query, input.ID, authorID, tagID, input.Title, input.Content)
	if err != nil {
		// Check if the error is due to duplicate key violation
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return nil, fmt.Errorf("failed to create thread: thread with the same ID already exists")
		}
		return nil, err
	}

	thread := &models.Thread{
		ID:      input.ID,
		Author:  input.Author,
		Tag:     input.Tag,
		Title:   input.Title,
		Content: input.Content,
	}

	return thread, nil
}

// Update updates an existing thread in the database.
func Update(db *database.Database, threadID string, input models.ThreadInput) (*models.Thread, error) {
	// Query to retrieve the author ID by name
	authorQuery := "SELECT id FROM users WHERE username = $1"
	var authorID string
	err := db.DB.QueryRow(authorQuery, input.Author).Scan(&authorID)
	if err != nil {
		return nil, err
	}

	// Query to retrieve the tag ID by name
	tagQuery := "SELECT id FROM tags WHERE name = $1"
	var tagID string
	err = db.DB.QueryRow(tagQuery, input.Tag).Scan(&tagID)
	if err != nil {
		return nil, err
	}
	query := "UPDATE threads SET author_id = $1, tag_id = $2, title = $3, content = $4 WHERE id = $5"
	_, err = db.DB.Exec(query, authorID, tagID, input.Title, input.Content, threadID)
	if err != nil {
		return nil, err
	}

	thread := &models.Thread{
		ID:      threadID,
		Author:  authorID,
		Tag:     tagID,
		Title:   input.Title,
		Content: input.Content,
	}

	return thread, nil
}

// Delete removes a thread from the database.
func Delete(db *database.Database, threadID string) error {
	query := "DELETE FROM threads WHERE id = $1"
	_, err := db.DB.Exec(query, threadID)
	if err != nil {
		return err
	}

	return nil
}
