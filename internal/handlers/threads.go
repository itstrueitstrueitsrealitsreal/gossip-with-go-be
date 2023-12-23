package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/CVWO/sample-go-app/internal/dataaccess/threads"
	"github.com/go-chi/chi/v5"

	"github.com/CVWO/sample-go-app/internal/api"
	"github.com/CVWO/sample-go-app/internal/database"
	"github.com/pkg/errors"
	"net/http"
)

const (
	ListThreads = "threads.HandleList"

	SuccessfulListThreadsMessage = "Successfully listed threads"
	ErrRetrieveThreads           = "Failed to retrieve threads in %s"
)

func HandleListThreads(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	db, err := database.GetDB()

	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, ListThreads))
	}

	threadList, err := threads.List(db)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveThreads, ListThreads))
	}

	data, err := json.Marshal(threadList)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrEncodeView, ListThreads))
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulListThreadsMessage},
	}, nil
}

// HandleGetThread retrieves a single thread by ID
func HandleGetThread(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	threadID := chi.URLParam(r, "id")
	if threadID == "" {
		return nil, errors.New("Thread ID is missing")
	}

	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to retrieve database in HandleGetThread")
	}

	thread, err := threads.GetThreadByID(db, threadID)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to retrieve thread in HandleGetThread")
	}

	data, err := json.Marshal(thread)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to encode thread in HandleGetThread")
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
	}, nil
}
