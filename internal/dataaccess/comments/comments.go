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
func List(db *database.Database) ([]models.CommentJSON, error) {
	// Query to select comments from the database
	query := "SELECT id, thread_id, author_id, content, timestamp FROM comments"

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
	var comments []models.CommentJSON

	// Iterate through the rows and populate the comments slice
	for rows.Next() {
		var commentInput models.CommentJSON
		var timestampStr string

		err := rows.Scan(&commentInput.ID, &commentInput.ThreadID, &commentInput.AuthorID, &commentInput.Content, &timestampStr)
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
func GetCommentByID(db *database.Database, commentID string) (*models.CommentJSON, error) {
	query := "SELECT id, thread_id, author_id, content, timestamp FROM comments WHERE id = $1"
	var commentInput models.CommentJSON
	var timestampStr string

	err := db.DB.QueryRow(query, commentID).Scan(&commentInput.ID, &commentInput.ThreadID, &commentInput.AuthorID, &commentInput.Content, &timestampStr)
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
func Create(db *database.Database, commentInput models.CommentInput, timestamp time.Time) (*models.Comment, error) {
	// Insert the new comment into the database
	query := "INSERT INTO comments (thread_id, author_id, content, timestamp) VALUES ($1, $2, $3, $4) RETURNING id"
	row := db.DB.QueryRow(query, commentInput.ThreadID, commentInput.AuthorID, commentInput.Content, timestamp)

	// Get the ID of the newly inserted comment
	var commentID string
	err := row.Scan(&commentID)
	if err != nil {
		// Check if the error is due to duplicate key value violation
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return nil, fmt.Errorf("failed to create comment: comment with the same ID already exists")
		}
		return nil, fmt.Errorf("error getting last insert ID: %v", err)
	}

	// Return the newly created comment
	return &models.Comment{
		ID:        commentID,
		ThreadID:  commentInput.ThreadID,
		AuthorID:  commentInput.AuthorID,
		Content:   commentInput.Content,
		Timestamp: timestamp,
	}, nil
}

// Update updates the comment with the specified ID in the database.
func Update(db *database.Database, commentID string, commentInput models.CommentInput, timestamp time.Time) (*models.Comment, error) {
	// Check if the comment with the given ID exists
	existingComment, err := GetCommentByID(db, commentID)
	if err != nil {
		return nil, fmt.Errorf("error checking existing comment: %v", err)
	}

	if existingComment == nil {
		return nil, fmt.Errorf("comment with ID %s not found", commentID)
	}

	// Update the comment's information
	query := "UPDATE comments SET thread_id = $1, author_id = $2, content = $3, timestamp = $4 WHERE id = $5"
	_, err = db.DB.Exec(query, commentInput.ThreadID, commentInput.AuthorID, commentInput.Content, timestamp, commentID)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}

	// Return the updated comment
	return &models.Comment{
		ID:        existingComment.ID,
		ThreadID:  commentInput.ThreadID,
		AuthorID:  commentInput.AuthorID,
		Content:   commentInput.Content,
		Timestamp: timestamp,
	}, nil
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
func GetCommentsByThreadID(db *database.Database, threadID string) ([]models.CommentJSON, error) {
	query := "SELECT id, thread_id, author_id, content, timestamp FROM comments WHERE thread_id = $1"
	rows, err := db.DB.Query(query, threadID)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	var comments []models.CommentJSON
	for rows.Next() {
		var comment models.CommentJSON
		var timestampStr string
		err := rows.Scan(&comment.ID, &comment.ThreadID, &comment.AuthorID, &comment.Content, &timestampStr)
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
