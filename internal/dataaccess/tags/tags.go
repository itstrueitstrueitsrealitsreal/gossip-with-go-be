package tags

import (
	"github.com/CVWO/sample-go-app/internal/database"
	"github.com/CVWO/sample-go-app/internal/models"
)

func List(db *database.Database) ([]models.Tag, error) {
	// Query to select tags from the database
	query := "SELECT id, name FROM tags"

	// Execute the query
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Initialize a slice to store the retrieved tags
	var tags []models.Tag

	// Iterate through the rows and populate the tags slice
	for rows.Next() {
		var tag models.Tag
		err := rows.Scan(&tag.ID, &tag.Name)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tags, nil
}

// GetTagByID retrieves a tag by ID from the database
func GetTagByID(db *database.Database, tagID string) (*models.Tag, error) {
	query := "SELECT id, name FROM tags WHERE id = ?"
	var tag models.Tag

	err := db.DB.QueryRow(query, tagID).Scan(&tag.ID, &tag.Name)
	if err != nil {
		return nil, err
	}

	return &tag, nil
}
