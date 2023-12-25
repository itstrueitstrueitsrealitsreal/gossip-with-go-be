package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/CVWO/sample-go-app/internal/dataaccess/users"
	"github.com/go-chi/chi/v5"
	"net/http"

	"github.com/CVWO/sample-go-app/internal/api"
	"github.com/CVWO/sample-go-app/internal/database"
	"github.com/pkg/errors"
)

const (
	ListUsers                   = "users.HandleList"
	SuccessfulListUsersMessage  = "Successfully listed users"
	ErrRetrieveDatabase         = "Failed to retrieve database in %s"
	ErrRetrieveUsers            = "Failed to retrieve users in %s"
	ErrEncodeView               = "Failed to retrieve users in %s"
	SuccessfulCreateUserMessage = "Successfully created user"
	ErrCreateUser               = "Failed to create user"
	SuccessfulUpdateUserMessage = "Successfully updated user"
	ErrUpdateUser               = "Failed to update user"
	SuccessfulDeleteUserMessage = "Successfully deleted user"
	ErrDeleteUser               = "Failed to delete user"
	SuccessfulViewUserMessage   = "Successfully viewed user"
)

// HandleListUsers returns all users
func HandleListUsers(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	db, err := database.GetDB()

	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, ListUsers))
	}

	userList, err := users.List(db)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveUsers, ListUsers))
	}
	// Log userList using fmt.Printf
	fmt.Printf("User List: %+v\n", userList)

	data, err := json.Marshal(userList)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrEncodeView, ListUsers))
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulListUsersMessage},
	}, nil
}

// HandleGetUser retrieves a single user by ID
func HandleGetUser(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	userID := chi.URLParam(r, "id")
	if userID == "" {
		return nil, errors.New("User ID is missing")
	}

	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to retrieve database in HandleGetUser")
	}

	user, err := users.GetUserByID(db, userID)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to retrieve user in HandleGetUser")
	}

	data, err := json.Marshal(user)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to encode user in HandleGetUser")
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulViewUserMessage},
	}, nil
}

// HandleUpdateUser updates a user's information in the database
func HandleUpdateUser(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	userID := chi.URLParam(r, "id")
	if userID == "" {
		return nil, errors.New("User ID is missing")
	}

	var userInput users.UserInput

	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		return nil, errors.Wrap(err, "Failed to decode updated user input")
	}

	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to retrieve database in HandleUpdateUser")
	}

	user, err := users.Update(db, userID, userInput)
	if err != nil {
		return nil, errors.Wrap(err, ErrUpdateUser)
	}

	// Marshal the user object to obtain a json.RawMessage
	userJSON, err := json.Marshal(user)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to encode updated user to JSON")
	}

	return &api.Response{
		Payload: api.Payload{
			// Use json.RawMessage to assign userJSON to api.Payload.Data
			Data: json.RawMessage(userJSON),
		},
		Messages: []string{SuccessfulUpdateUserMessage},
	}, nil
}

// HandleCreateUser creates a new user and inserts it into the database
func HandleCreateUser(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	var userInput users.UserInput

	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		return nil, errors.Wrap(err, "Failed to decode user input")
	}

	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to retrieve database in HandleCreateUser")
	}

	user, err := users.Create(db, userInput)
	if err != nil {
		return nil, errors.Wrap(err, ErrCreateUser)
	}

	// Marshal the user object to obtain a json.RawMessage
	userJSON, err := json.Marshal(user)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to encode user to JSON")
	}

	return &api.Response{
		Payload: api.Payload{
			// Use json.RawMessage to assign userJSON to api.Payload.Data
			Data: json.RawMessage(userJSON),
		},
		Messages: []string{SuccessfulCreateUserMessage},
	}, nil
}

// HandleDeleteUser deletes a user from the database
func HandleDeleteUser(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	userID := chi.URLParam(r, "id")
	if userID == "" {
		return nil, errors.New("User ID is missing")
	}

	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to retrieve database in HandleDeleteUser")
	}

	err = users.Delete(db, userID)
	if err != nil {
		return nil, errors.Wrap(err, ErrDeleteUser)
	}

	return &api.Response{
		Messages: []string{SuccessfulDeleteUserMessage},
	}, nil
}
