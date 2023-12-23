package models

type Thread struct {
	ID       int    `json:"id"`
	AuthorID int    `json:"author"`
	TagID    int    `json:"tag"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}
