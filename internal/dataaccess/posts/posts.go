package posts

import (
	"fmt"
	"github.com/CVWO/sample-go-app/internal/database"
	"time"
)

// PostJSON is a struct used for JSON marshaling with a custom timestamp format
type PostJSON struct {
	ID        int    `json:"id"`
	ThreadID  int    `json:"thread"`
	AuthorID  int    `json:"author"`
	Content   string `json:"content"`
	Timestamp string `json:"timestamp"`
}

// List retrieves a list of posts from the database.
func List(db *database.Database) ([]PostJSON, error) {
	// Query to select posts from the database
	query := "SELECT id, thread_id, author_id, content, timestamp FROM posts"

	// Execute the query
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	// Initialize a slice to store the retrieved posts
	var posts []PostJSON

	// Iterate through the rows and populate the posts slice
	for rows.Next() {
		var postJSON PostJSON
		var timestampStr string

		err := rows.Scan(&postJSON.ID, &postJSON.ThreadID, &postJSON.AuthorID, &postJSON.Content, &timestampStr)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}

		// Convert the timestamp from string to time.Time
		timestamp, err := time.Parse("2006-01-02 15:04:05", timestampStr)
		if err != nil {
			return nil, fmt.Errorf("error parsing timestamp: %v", err)
		}
		postJSON.Timestamp = timestamp.Format("2006-01-02 15:04:05")

		posts = append(posts, postJSON)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	// Print the number of retrieved posts for debugging
	fmt.Printf("Number of retrieved posts: %d\n", len(posts))

	return posts, nil
}

// GetPostByID retrieves a post by ID from the database
func GetPostByID(db *database.Database, postID string) (*PostJSON, error) {
	query := "SELECT id, thread_id, author_id, content, timestamp FROM posts WHERE id = ?"
	var postJSON PostJSON
	var timestampStr string

	err := db.DB.QueryRow(query, postID).Scan(&postJSON.ID, &postJSON.ThreadID, &postJSON.AuthorID, &postJSON.Content, &timestampStr)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}

	// Convert the timestamp from string to time.Time
	timestamp, err := time.Parse("2006-01-02 15:04:05", timestampStr)
	if err != nil {
		return nil, fmt.Errorf("error parsing timestamp: %v", err)
	}
	postJSON.Timestamp = timestamp.Format("2006-01-02 15:04:05")

	return &postJSON, nil
}
