package models

import "time"

type Post struct {
	ID        int       `json:"id"`
	ThreadID  int       `json:"thread"`
	AuthorID  int       `json:"author"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}
