package models

type Thread struct {
	ID       int    `json:"id"`
	AuthorID int    `json:"author_id"`
	TagID    int    `json:"tag_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}
