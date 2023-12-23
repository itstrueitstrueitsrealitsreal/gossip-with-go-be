package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/CVWO/sample-go-app/internal/dataaccess/tags"
	"net/http"

	"github.com/CVWO/sample-go-app/internal/api"
	"github.com/CVWO/sample-go-app/internal/database"
	"github.com/pkg/errors"
)

const (
	ListTags = "tags.HandleList"

	SuccessfulListTagsMessage = "Successfully listed tags"
	ErrRetrieveTags           = "Failed to retrieve tags in %s"
)

func HandleListTags(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	db, err := database.GetDB()

	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, ListTags))
	}

	tagList, err := tags.List(db)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveTags, ListTags))
	}

	data, err := json.Marshal(tagList)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrEncodeView, ListTags))
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulListTagsMessage},
	}, nil
}
