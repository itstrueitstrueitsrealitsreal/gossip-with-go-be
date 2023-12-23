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
	ListUsers = "users.HandleList"

	SuccessfulListUsersMessage = "Successfully listed users"
	ErrRetrieveDatabase        = "Failed to retrieve database in %s"
	ErrRetrieveUsers           = "Failed to retrieve users in %s"
	ErrEncodeView              = "Failed to retrieve users in %s"
)

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
	}, nil
}
