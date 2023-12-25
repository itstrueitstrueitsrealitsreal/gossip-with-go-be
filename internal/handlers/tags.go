package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/itstrueitstrueitsrealitsreal/gossip-with-go/internal/dataaccess/tags"
	"net/http"

	"github.com/itstrueitstrueitsrealitsreal/gossip-with-go/internal/api"
	"github.com/itstrueitstrueitsrealitsreal/gossip-with-go/internal/database"
	"github.com/pkg/errors"
)

const (
	ListTags                  = "tags.HandleList"
	SuccessfulViewTagMessage  = "Successfully viewed tag"
	SuccessfulListTagsMessage = "Successfully listed tags"
	ErrRetrieveTags           = "Failed to retrieve tags in %s"
)

// HandleListTags returns all tags in a json format
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

// HandleGetTag retrieves a single tag by ID
func HandleGetTag(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	tagID := chi.URLParam(r, "id")
	if tagID == "" {
		return nil, errors.New("Tag ID is missing")
	}

	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to retrieve database in HandleGetTag")
	}

	tag, err := tags.GetTagByID(db, tagID)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to retrieve tag in HandleGetTag")
	}

	data, err := json.Marshal(tag)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to encode tag in HandleGetTag")
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulViewTagMessage},
	}, nil
}
