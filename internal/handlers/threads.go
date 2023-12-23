package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/CVWO/sample-go-app/internal/dataaccess/threads"

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
