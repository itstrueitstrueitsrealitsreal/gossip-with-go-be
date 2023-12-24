package threads

// ThreadInput represents the input for creating or updating a thread
type ThreadInput struct {
	AuthorID int    `json:"author_id"`
	TagID    int    `json:"tag_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}
