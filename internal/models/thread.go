package models

type Thread struct {
	ID      string `json:"id"`
	Author  string `json:"author"`
	Tag     string `json:"tag"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// ThreadInput represents the input for creating or updating a thread
type ThreadInput struct {
	AuthorID string `json:"author_id"`
	TagID    string `json:"tag_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}
