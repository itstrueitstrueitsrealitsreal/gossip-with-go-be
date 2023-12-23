package tags

import (
	"github.com/CVWO/sample-go-app/internal/database"
	"github.com/CVWO/sample-go-app/internal/models"
)

func List(db *database.Database) ([]models.Tag, error) {
	tags := []models.Tag{
		{
			ID:   1,
			Name: "Opinion",
		},
		{
			ID:   2,
			Name: "Satirical",
		},
	}
	return tags, nil
}

//
//// GetTagByID retrieves a tag by ID from the database
//func GetTagByID(db *database.Database, tagID string) (*models.Tag, error) {
//	query := "SELECT id, name FROM tags WHERE id = ?"
//	var tag models.Tag
//
//	err := db.DB.QueryRow(query, tagID).Scan(&tag.ID, &tag.Name)
//	if err != nil {
//		return nil, err
//	}
//
//	return &tag, nil
//}

// GetTagByID retrieves a tag by ID, first checking a local set and then the database
func GetTagByID(db *database.Database, tagID string) (*models.Tag, error) {
	localTags := map[string]models.Tag{
		"1": {ID: 1, Name: "Opinion"},
		"2": {ID: 2, Name: "Satirical"},
	}

	// Check if the tag ID exists in the local set
	if localTag, ok := localTags[tagID]; ok {
		return &localTag, nil
	}

	// If not found in the local set, proceed with the database query
	query := "SELECT id, name FROM tags WHERE id = ?"
	var tag models.Tag

	err := db.DB.QueryRow(query, tagID).Scan(&tag.ID, &tag.Name)
	if err != nil {
		return nil, err
	}

	return &tag, nil
}
