package posts

// PostInput represents the input data structure for creating or updating a post.
type PostInput struct {
	ThreadID  int    `json:"thread_id"`
	AuthorID  int    `json:"author_id"`
	Content   string `json:"content"`
	Timestamp string `json:"timestamp"`
}
