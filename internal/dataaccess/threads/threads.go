package threads

import (
	"github.com/CVWO/sample-go-app/internal/database"
	"github.com/CVWO/sample-go-app/internal/models"
)

func List(db *database.Database) ([]models.Thread, error) {
	threads := []models.Thread{
		{
			ID:      1,
			Title:   "Sample Thread 1",
			Content: "Sample thread content 1",
		},
		{
			ID:      2,
			Title:   "Sample Thread 2",
			Content: "Sample thread content 2",
		},
	}
	return threads, nil
}
