package models

import "time"

type Post struct {
	ID        int       `json:"id"`
	Thread    Thread    `json:"thread"`
	Author    User      `json:"author"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}
