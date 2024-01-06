package users

import (
	"database/sql"
	"fmt"

	"github.com/itstrueitstrueitsrealitsreal/gossip-with-go-be/internal/database"
	"github.com/itstrueitstrueitsrealitsreal/gossip-with-go-be/internal/models"
)

// List retrieves a list of all users from the database.
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
	query := "SELECT id, name FROM users WHERE id = $1"
	var user models.Tag

	err := db.DB.QueryRow(query, userID).Scan(&user.ID, &user.Name)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUserByName retrieves a user by their name from the database.
func GetUserByName(db *database.Database, name string) (*models.User, error) {
	query := "SELECT id, name FROM users WHERE name = $1"
	var user models.User

	err := db.DB.QueryRow(query, name).Scan(&user.ID, &user.Name)
	if err == sql.ErrNoRows {
		// User not found
		return nil, nil
	} else if err != nil {
		// Other error
		return nil, fmt.Errorf("error executing query: %v", err)
	}

	return &user, nil
}

// Create inserts a new user into the database.
func Create(db *database.Database, userInput models.UserInput) (*models.User, error) {
	// Check if the user with the same name already exists
	existingUser, err := GetUserByName(db, userInput.Name)
	if err != nil {
		return nil, fmt.Errorf("error checking existing user: %v", err)
	}

	if existingUser != nil {
		return nil, fmt.Errorf("user with name %s already exists", userInput.Name)
	}

	// Insert the new user into the database
	query := "INSERT INTO users (name) VALUES ($1) RETURNING id"
	var userID string

	err = db.DB.QueryRow(query, userInput.Name).Scan(&userID)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}

	// Return the newly created user
	return &models.User{
		ID:   userID,
		Name: userInput.Name,
	}, nil
}

// Update updates the user with the specified ID in the database.
func Update(db *database.Database, userID string, userInput models.UserInput) (*models.User, error) {
	// Check if the user with the given ID exists
	existingUser, err := GetUserByID(db, userID)
	if err != nil {
		return nil, fmt.Errorf("error checking existing user: %v", err)
	}

	if existingUser == nil {
		return nil, fmt.Errorf("user with ID %s not found", userID)
	}

	// Update the user's name
	query := "UPDATE users SET name = $1 WHERE id = $2"
	_, err = db.DB.Exec(query, userInput.Name, userID)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}

	// Return the updated user
	return &models.User{
		ID:   existingUser.ID,
		Name: userInput.Name,
	}, nil
}

// Delete removes a user from the database by ID.
func Delete(db *database.Database, userID string) error {
	deleteUserQuery := "DELETE FROM users WHERE id = $1"
	stmt, err := db.DB.Prepare(deleteUserQuery)
	if err != nil {
		return fmt.Errorf("error preparing delete user statement: %v", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(userID)
	if err != nil {
		return fmt.Errorf("error executing delete user statement: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected after delete user statement: %v", err)
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
