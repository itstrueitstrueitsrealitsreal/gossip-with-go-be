package users

import (
	"github.com/CVWO/sample-go-app/internal/database"
	"github.com/CVWO/sample-go-app/internal/models"
)

func List(db *database.Database) ([]models.User, error) {
	users := []models.User{
		{
			ID:   1,
			Name: "CVWO",
		},
		{
			ID:   2,
			Name: "Kenneth",
		},
	}
	return users, nil
}

//
//// GetUserByID retrieves a user by ID from the database
//func GetUserByID(db *database.Database, userID string) (*models.User, error) {
//	query := "SELECT id, username FROM users WHERE id = ?"
//	var user models.User
//
//	err := db.DB.QueryRow(query, userID).Scan(&user.ID, &user.Name)
//	if err != nil {
//		return nil, err
//	}
//
//	return &user, nil
//}

// GetUserByID retrieves a user by ID, first checking a local set and then the database
func GetUserByID(db *database.Database, userID string) (*models.User, error) {
	localUsers := map[string]models.User{
		"1": {ID: 1, Name: "CVWO"},
		"2": {ID: 2, Name: "Kenneth"},
	}

	// Check if the user ID exists in the local set
	if localUser, ok := localUsers[userID]; ok {
		return &localUser, nil
	}

	// If not found in the local set, proceed with the database query
	query := "SELECT id, username FROM users WHERE id = ?"
	var user models.User

	err := db.DB.QueryRow(query, userID).Scan(&user.ID, &user.Name)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
