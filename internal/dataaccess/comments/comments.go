package comments

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/itstrueitstrueitsrealitsreal/gossip-with-go-be/internal/database"
	"github.com/itstrueitstrueitsrealitsreal/gossip-with-go-be/internal/models"
	"github.com/pkg/errors"
)

// List retrieves a list of comments from the database.
func List(db *database.Database) ([]models.CommentInput, error) {
	// Query to select comments from the database
	query := `
		SELECT c.id, c.thread_id, u.username AS author_name, c.content, c.timestamp
		FROM comments c
		INNER JOIN users u ON c.author_id = u.id
	`

	// Execute the query
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			fmt.Printf("error closing rows: %v\n", err)
		}
	}()

	// Initialize a slice to store the retrieved comments
	var comments []models.CommentInput

	// Iterate through the rows and populate the comments slice
	for rows.Next() {
		var commentInput models.CommentInput
		var timestampStr string

		err := rows.Scan(&commentInput.ID, &commentInput.ThreadID, &commentInput.Author, &commentInput.Content, &timestampStr)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}

		// Convert the timestamp from string to time.Time
		timestamp, err := time.Parse("2006-01-02T15:04:05.999999999Z07:00", timestampStr)
		if err != nil {
			return nil, fmt.Errorf("error parsing timestamp: %v", err)
		}
		commentInput.Timestamp = timestamp.Format("2006-01-02 15:04:05")

		comments = append(comments, commentInput)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	return comments, nil
}

// GetCommentByID retrieves a comment by ID from the database
func GetCommentByID(db *database.Database, commentID string) (*models.CommentInput, error) {
	query := `
		SELECT c.id, c.thread_id, u.username AS author_name, c.content, c.timestamp
		FROM comments c
		INNER JOIN users u ON c.author_id = u.id
		WHERE c.id = $1
	`
	var commentInput models.CommentInput
	var timestampStr string

	err := db.DB.QueryRow(query, commentID).Scan(&commentInput.ID, &commentInput.ThreadID, &commentInput.Author, &commentInput.Content, &timestampStr)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}

	// Convert the timestamp from string to time.Time
	timestamp, err := time.Parse("2006-01-02T15:04:05.999999999Z07:00", timestampStr)
	if err != nil {
		return nil, fmt.Errorf("error parsing timestamp: %v", err)
	}
	commentInput.Timestamp = timestamp.Format("2006-01-02 15:04:05")

	return &commentInput, nil
}

// Create inserts a new comment into the database.
func Create(db *database.Database, commentInput models.CommentInput, timestamp time.Time) (*models.CommentInput, error) {
	// Get the author ID for the given author name
	authorID, err := getAuthorIDByName(db, commentInput.Author)
	if err != nil {
		return nil, fmt.Errorf("error getting author ID: %v", err)
	}

	// Insert the new comment into the database
	query := "INSERT INTO comments (id, thread_id, author_id, content, timestamp) VALUES ($1, $2, $3, $4, $5)"
	_, err = db.DB.Exec(query, commentInput.ID, commentInput.ThreadID, authorID, commentInput.Content, timestamp)
	if err != nil {
		// Check if the error is due to duplicate key value violation
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return nil, fmt.Errorf("failed to create comment: comment with the same ID already exists")
		}
		return nil, fmt.Errorf("error executing query: %v", err)
	}

	// Return the newly created comment
	return &commentInput, nil
}

// getAuthorIDByName retrieves the author ID for the given author name from the database
func getAuthorIDByName(db *database.Database, authorName string) (string, error) {
	query := "SELECT id FROM users WHERE username = $1"
	var authorID string
	err := db.DB.QueryRow(query, authorName).Scan(&authorID)
	if err != nil {
		return "", fmt.Errorf("error getting author ID: %v", err)
	}
	return authorID, nil
}

// Update updates the comment with the specified ID in the database.
func Update(db *database.Database, commentID string, commentInput models.CommentInput, timestamp time.Time) (*models.CommentInput, error) {
	// Check if the comment with the given ID exists
	existingComment, err := GetCommentByID(db, commentID)
	if err != nil {
		return nil, fmt.Errorf("error checking existing comment: %v", err)
	}

	if existingComment == nil {
		return nil, fmt.Errorf("comment with ID %s not found", commentID)
	}

	// Get the author ID for the given author name
	authorID, err := getAuthorIDByName(db, commentInput.Author)
	if err != nil {
		return nil, fmt.Errorf("error getting author ID: %v", err)
	}

	// Update the comment's information
	query := "UPDATE comments SET thread_id = $1, author_id = $2, content = $3, timestamp = $4 WHERE id = $5"
	_, err = db.DB.Exec(query, commentInput.ThreadID, authorID, commentInput.Content, timestamp, commentID)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}

	// Return the updated comment
	return &commentInput, nil
}

// Delete removes a comment from the database by ID.
func Delete(db *database.Database, commentID string) error {
	deleteCommentQuery := "DELETE FROM comments WHERE id = $1"
	stmt, err := db.DB.Prepare(deleteCommentQuery)
	if err != nil {
		return errors.Wrap(err, "error preparing delete comment statement")
	}
	defer stmt.Close()

	result, err := stmt.Exec(commentID)
	if err != nil {
		return errors.Wrap(err, "error executing delete comment statement")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "error getting rows affected after delete comment statement")
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// GetCommentsByThreadID retrieves all comments for a specific thread by thread ID from the database
func GetCommentsByThreadID(db *database.Database, threadID string) ([]models.CommentInput, error) {
	query := `
        SELECT comments.id, comments.thread_id, users.username, comments.content, comments.timestamp 
        FROM comments 
        INNER JOIN users ON comments.author_id = users.id 
        WHERE comments.thread_id = $1
    `
	rows, err := db.DB.Query(query, threadID)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	var comments []models.CommentInput
	for rows.Next() {
		var comment models.CommentInput
		var timestampStr string
		err := rows.Scan(&comment.ID, &comment.ThreadID, &comment.Author, &comment.Content, &timestampStr)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}

		// Convert the timestamp from string to time.Time
		timestamp, err := time.Parse("2006-01-02T15:04:05.999999999Z07:00", timestampStr)
		if err != nil {
			return nil, fmt.Errorf("error parsing timestamp: %v", err)
		}
		comment.Timestamp = timestamp.Format("2006-01-02 15:04:05")

		comments = append(comments, comment)
	}

	return comments, nil
}
