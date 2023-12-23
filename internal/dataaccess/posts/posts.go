package posts

import (
	"github.com/CVWO/sample-go-app/internal/database"
	"github.com/CVWO/sample-go-app/internal/models"
	"time"
)

func List(db *database.Database) ([]models.Post, error) {
	posts := []models.Post{
		{
			ID:        1,
			ThreadID:  1,
			AuthorID:  1,
			Content:   "Sample post content 1",
			Timestamp: time.Now(),
		},
		{
			ID:        2,
			ThreadID:  1,
			AuthorID:  2,
			Content:   "Sample post content 2",
			Timestamp: time.Now(),
		},
	}
	return posts, nil
}

//// GetPostByID retrieves a post by ID from the database
//func GetPostByID(db *database.Database, postID string) (*models.Post, error) {
//	query := "SELECT id, thread_id, author_id, content, timestamp FROM posts WHERE id = ?"
//	var post models.Post
//
//	err := db.DB.QueryRow(query, postID).Scan(&post.ID, &post.ThreadID, &post.AuthorID, &post.Content, &post.Timestamp)
//	if err != nil {
//		return nil, err
//	}
//
//	return &post, nil
//}

// GetPostByID retrieves a post by ID, first checking a local set and then the database
func GetPostByID(db *database.Database, postID string) (*models.Post, error) {
	localPosts := map[string]models.Post{
		"1": {
			ID:        1,
			ThreadID:  1,
			AuthorID:  1,
			Content:   "Sample post content 1",
			Timestamp: time.Now(),
		},
		"2": {
			ID:        2,
			ThreadID:  1,
			AuthorID:  2,
			Content:   "Sample post content 2",
			Timestamp: time.Now(),
		},
	}

	// Check if the post ID exists in the local set
	if localPost, ok := localPosts[postID]; ok {
		return &localPost, nil
	}

	// If not found in the local set, proceed with the database query
	query := "SELECT id, author_id, thread_id, content FROM posts WHERE id = ?"
	var post models.Post

	err := db.DB.QueryRow(query, postID).Scan(&post.ID, &post.AuthorID, &post.ThreadID, &post.Content)
	if err != nil {
		return nil, err
	}

	return &post, nil
}
