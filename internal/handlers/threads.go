package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/itstrueitstrueitsrealitsreal/gossip-with-go/internal/api"
	"github.com/itstrueitstrueitsrealitsreal/gossip-with-go/internal/dataaccess/threads"
	"github.com/itstrueitstrueitsrealitsreal/gossip-with-go/internal/database"
	"github.com/pkg/errors"
	"net/http"
)

const (
	ListThreads                   = "threads.HandleList"
	ErrCreateThread               = "Failed to create thread"
	SuccessfulCreateThreadMessage = "Successfully created thread"
	ErrUpdateThread               = "Failed to update thread"
	SuccessfulUpdateThreadMessage = "Successfully updated thread"
	ErrDeleteThread               = "Failed to delete thread"
	SuccessfulDeleteThreadMessage = "Successfully deleted thread"
	SuccessfulListThreadsMessage  = "Successfully listed threads"
	ErrRetrieveThreads            = "Failed to retrieve threads in %s"
	SuccessfulViewThreadMessage   = "Successfully viewed thread"
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
		Messages: []string{SuccessfulViewThreadMessage},
	}, nil
}

// HandleCreateThread creates a new thread and inserts it into the database
func HandleCreateThread(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	var threadInput threads.ThreadInput

	if err := json.NewDecoder(r.Body).Decode(&threadInput); err != nil {
		return nil, errors.Wrap(err, "Failed to decode thread input")
	}

	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to retrieve database in HandleCreateThread")
	}

	thread, err := threads.Create(db, threadInput)
	if err != nil {
		return nil, errors.Wrap(err, ErrCreateThread)
	}

	// Convert the thread to JSON
	threadJSON, err := json.Marshal(thread)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to encode thread to JSON")
	}

	// Use json.RawMessage to assign threadJSON to api.Payload.Data
	threadData := json.RawMessage(threadJSON)

	return &api.Response{
		Payload: api.Payload{
			Data: threadData,
		},
		Messages: []string{SuccessfulCreateThreadMessage},
	}, nil
}

// HandleUpdateThread updates a thread's information in the database
func HandleUpdateThread(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	threadID := chi.URLParam(r, "id")
	if threadID == "" {
		return nil, errors.New("Thread ID is missing")
	}

	var threadInput threads.ThreadInput

	if err := json.NewDecoder(r.Body).Decode(&threadInput); err != nil {
		return nil, errors.Wrap(err, "Failed to decode updated thread input")
	}

	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to retrieve database in HandleUpdateThread")
	}

	thread, err := threads.Update(db, threadID, threadInput)
	if err != nil {
		return nil, errors.Wrap(err, ErrUpdateThread)
	}

	// Convert the updated thread to JSON
	threadJSON, err := json.Marshal(thread)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to encode updated thread to JSON")
	}

	// Use json.RawMessage to assign threadJSON to api.Payload.Data
	threadData := json.RawMessage(threadJSON)

	return &api.Response{
		Payload: api.Payload{
			Data: threadData,
		},
		Messages: []string{SuccessfulUpdateThreadMessage},
	}, nil
}

// HandleDeleteThread deletes a thread from the database
func HandleDeleteThread(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	threadID := chi.URLParam(r, "id")
	if threadID == "" {
		return nil, errors.New("Thread ID is missing")
	}

	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to retrieve database in HandleDeleteThread")
	}

	err = threads.Delete(db, threadID)
	if err != nil {
		return nil, errors.Wrap(err, ErrDeleteThread)
	}

	return &api.Response{
		Messages: []string{SuccessfulDeleteThreadMessage},
	}, nil
}
