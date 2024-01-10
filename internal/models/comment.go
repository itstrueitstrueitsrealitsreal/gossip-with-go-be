package models

import "time"

type Comment struct {
	ID        string    `json:"id"`
	ThreadID  string    `json:"thread_id"`
	AuthorID  string    `json:"author_id"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}

// CommentJSON is a struct used for JSON marshaling with a custom timestamp format
type CommentJSON struct {
	ID        string `json:"id"`
	ThreadID  string `json:"thread_id"`
	AuthorID  string `json:"author_id"`
	Content   string `json:"content"`
	Timestamp string `json:"timestamp"`
}

type CommentInput struct {
	ID        string `json:"id"`
	ThreadID  string `json:"thread_id"`
	Author    string `json:"author"`
	Content   string `json:"content"`
	Timestamp string `json:"timestamp"`
}
