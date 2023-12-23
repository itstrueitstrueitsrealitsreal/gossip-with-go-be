package users

import (
	"github.com/CVWO/sample-go-app/internal/database"
	"github.com/CVWO/sample-go-app/internal/models"
)

func List(db *database.Database) ([]models.Tag, error) {
	// Query to select users from the database
	query := "SELECT id, name FROM users"

	// Execute the query
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Initialize a slice to store the retrieved users
	var users []models.Tag

	// Iterate through the rows and populate the users slice
	for rows.Next() {
		var user models.Tag
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// GetUserByID retrieves a user by ID from the database
func GetUserByID(db *database.Database, userID string) (*models.Tag, error) {
	query := "SELECT id, name FROM users WHERE id = ?"
	var user models.Tag

	err := db.DB.QueryRow(query, userID).Scan(&user.ID, &user.Name)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
