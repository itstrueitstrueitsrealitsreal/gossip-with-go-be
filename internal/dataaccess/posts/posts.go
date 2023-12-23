package posts

import (
	"github.com/CVWO/sample-go-app/internal/database"
	"github.com/CVWO/sample-go-app/internal/models"
)

func List(db *database.Database) ([]models.Post, error) {
	posts := []models.Post{
		{
			ID:      1,
			Content: "Sample post content 1",
		},
		{
			ID:      2,
			Content: "Sample post content 2",
		},
	}
	return posts, nil
}
