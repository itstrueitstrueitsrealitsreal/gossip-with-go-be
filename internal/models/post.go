package models

import "time"

type Post struct {
	ID        int       `json:"id"`
	ThreadID  int       `json:"thread_id"`
	AuthorID  int       `json:"author_id"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}

// PostInput represents the input data structure for creating or updating a post.
type PostInput struct {
	ThreadID  int    `json:"thread_id"`
	AuthorID  int    `json:"author_id"`
	Content   string `json:"content"`
	Timestamp string `json:"timestamp"`
}

// PostJSON is a struct used for JSON marshaling with a custom timestamp format
type PostJSON struct {
	ID        int    `json:"id"`
	ThreadID  int    `json:"thread_id"`
	AuthorID  int    `json:"author_id"`
	Content   string `json:"content"`
	Timestamp string `json:"timestamp"`
}
