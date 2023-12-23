package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/CVWO/sample-go-app/internal/dataaccess/posts"
	"github.com/go-chi/chi/v5"
	"net/http"

	"github.com/CVWO/sample-go-app/internal/api"
	"github.com/CVWO/sample-go-app/internal/database"
	"github.com/pkg/errors"
)

const (
	ListPosts = "posts.HandleList"

	SuccessfulListPostsMessage = "Successfully listed posts"
	ErrRetrievePosts           = "Failed to retrieve posts in %s"
)

func HandleListPosts(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	db, err := database.GetDB()

	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, ListPosts))
	}

	postList, err := posts.List(db)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrievePosts, ListPosts))
	}

	data, err := json.Marshal(postList)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrEncodeView, ListPosts))
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulListPostsMessage},
	}, nil
}

// HandleGetPost retrieves a single post by ID
func HandleGetPost(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	postID := chi.URLParam(r, "id")
	if postID == "" {
		return nil, errors.New("Post ID is missing")
	}

	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to retrieve database in HandleGetPost")
	}

	post, err := posts.GetPostByID(db, postID)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to retrieve post in HandleGetPost")
	}

	data, err := json.Marshal(post)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to encode post in HandleGetPost")
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
	}, nil
}
