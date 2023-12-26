package posts

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/itstrueitstrueitsrealitsreal/gossip-with-go-be/internal/database"
	"github.com/itstrueitstrueitsrealitsreal/gossip-with-go-be/internal/models"
	"github.com/pkg/errors"
)

// List retrieves a list of posts from the database.
func List(db *database.Database) ([]models.PostJSON, error) {
	// Query to select posts from the database
	query := "SELECT id, thread_id, author_id, content, timestamp FROM posts"

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

	// Initialize a slice to store the retrieved posts
	var posts []models.PostJSON

	// Iterate through the rows and populate the posts slice
	for rows.Next() {
		var postInput models.PostJSON
		var timestampStr string

		err := rows.Scan(&postInput.ID, &postInput.ThreadID, &postInput.AuthorID, &postInput.Content, &timestampStr)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}

		// Convert the timestamp from string to time.Time
		timestamp, err := time.Parse("2006-01-02T15:04:05.999999999Z07:00", timestampStr)
		if err != nil {
			return nil, fmt.Errorf("error parsing timestamp: %v", err)
		}
		postInput.Timestamp = timestamp.Format("2006-01-02 15:04:05")

		posts = append(posts, postInput)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	return posts, nil
}

// GetPostByID retrieves a post by ID from the database
func GetPostByID(db *database.Database, postID string) (*models.PostJSON, error) {
	query := "SELECT id, thread_id, author_id, content, timestamp FROM posts WHERE id = $1"
	var postInput models.PostJSON
	var timestampStr string

	err := db.DB.QueryRow(query, postID).Scan(&postInput.ID, &postInput.ThreadID, &postInput.AuthorID, &postInput.Content, &timestampStr)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}

	// Convert the timestamp from string to time.Time
	timestamp, err := time.Parse("2006-01-02T15:04:05.999999999Z07:00", timestampStr)
	if err != nil {
		return nil, fmt.Errorf("error parsing timestamp: %v", err)
	}
	postInput.Timestamp = timestamp.Format("2006-01-02 15:04:05")

	return &postInput, nil
}

// Create inserts a new post into the database.
func Create(db *database.Database, postInput models.PostInput, timestamp time.Time) (*models.Post, error) {
	// Insert the new post into the database
	query := "INSERT INTO posts (thread_id, author_id, content, timestamp) VALUES ($1, $2, $3, $4) RETURNING id"
	row := db.DB.QueryRow(query, postInput.ThreadID, postInput.AuthorID, postInput.Content, timestamp)

	// Get the ID of the newly inserted post
	var postID int
	err := row.Scan(&postID)
	if err != nil {
		// Check if the error is due to duplicate key value violation
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return nil, fmt.Errorf("failed to create post: post with the same ID already exists")
		}
		return nil, fmt.Errorf("error getting last insert ID: %v", err)
	}

	// Return the newly created post
	return &models.Post{
		ID:        postID,
		ThreadID:  postInput.ThreadID,
		AuthorID:  postInput.AuthorID,
		Content:   postInput.Content,
		Timestamp: timestamp,
	}, nil
}

// Update updates the post with the specified ID in the database.
func Update(db *database.Database, postID string, postInput models.PostInput, timestamp time.Time) (*models.Post, error) {
	// Check if the post with the given ID exists
	existingPost, err := GetPostByID(db, postID)
	if err != nil {
		return nil, fmt.Errorf("error checking existing post: %v", err)
	}

	if existingPost == nil {
		return nil, fmt.Errorf("post with ID %s not found", postID)
	}

	// Update the post's information
	query := "UPDATE posts SET thread_id = $1, author_id = $2, content = $3, timestamp = $4 WHERE id = $5"
	_, err = db.DB.Exec(query, postInput.ThreadID, postInput.AuthorID, postInput.Content, timestamp, postID)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}

	// Return the updated post
	return &models.Post{
		ID:        existingPost.ID,
		ThreadID:  postInput.ThreadID,
		AuthorID:  postInput.AuthorID,
		Content:   postInput.Content,
		Timestamp: timestamp,
	}, nil
}

// Delete removes a post from the database by ID.
func Delete(db *database.Database, postID string) error {
	deletePostQuery := "DELETE FROM posts WHERE id = $1"
	stmt, err := db.DB.Prepare(deletePostQuery)
	if err != nil {
		return errors.Wrap(err, "error preparing delete post statement")
	}
	defer stmt.Close()

	result, err := stmt.Exec(postID)
	if err != nil {
		return errors.Wrap(err, "error executing delete post statement")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "error getting rows affected after delete post statement")
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
