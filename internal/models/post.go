package models

import "time"

type Post struct {
	ID        int       `json:"id"`
	ThreadID  int       `json:"thread_id"`
	AuthorID  int       `json:"author_id"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}
