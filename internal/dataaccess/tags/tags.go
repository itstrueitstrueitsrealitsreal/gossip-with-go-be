package tags

import (
	"github.com/CVWO/sample-go-app/internal/database"
	"github.com/CVWO/sample-go-app/internal/models"
)

func List(db *database.Database) ([]models.Tag, error) {
	tags := []models.Tag{
		{
			ID:   1,
			Name: "SampleTag1",
		},
		{
			ID:   2,
			Name: "SampleTag2",
		},
	}
	return tags, nil
}
