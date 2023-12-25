package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/itstrueitstrueitsrealitsreal/gossip-with-go/internal/api"
	"github.com/itstrueitstrueitsrealitsreal/gossip-with-go/internal/dataaccess/posts"
	"github.com/itstrueitstrueitsrealitsreal/gossip-with-go/internal/database"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

const (
	ListPosts                   = "posts.HandleList"
	SuccessfulListPostsMessage  = "Successfully listed posts"
	ErrRetrievePosts            = "Failed to retrieve posts in %s"
	ErrEncodePosts              = "Failed to encode posts in %s"
	SuccessfulCreatePostMessage = "Successfully created post"
	ErrCreatePost               = "Failed to create post"
	SuccessfulUpdatePostMessage = "Successfully updated post"
	ErrUpdatePost               = "Failed to update post"
	SuccessfulDeletePostMessage = "Successfully deleted post"
	ErrDeletePost               = "Failed to delete post"
	SuccessfulViewPostMessage   = "Successfully viewed post"
)

// HandleListPosts returns all posts in JSON format
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
		return nil, errors.Wrap(err, fmt.Sprintf(ErrEncodePosts, ListPosts))
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
		Messages: []string{SuccessfulViewPostMessage},
	}, nil
}

// HandleCreatePost creates a new post and inserts it into the database
func HandleCreatePost(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	var postInput posts.PostInput

	if err := json.NewDecoder(r.Body).Decode(&postInput); err != nil {
		return nil, errors.Wrap(err, "Failed to decode post input")
	}

	// Parse the timestamp string into a time.Time object
	timestamp := time.Now() // Get the current timestamp

	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to retrieve database in HandleCreatePost")
	}

	post, err := posts.Create(db, postInput, timestamp)
	if err != nil {
		return nil, errors.Wrap(err, ErrCreatePost)
	}

	// Convert the post to JSON
	postJSON, err := json.Marshal(post)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to encode post to JSON")
	}

	// Use json.RawMessage to assign postJSON to api.Payload.Data
	postData := json.RawMessage(postJSON)

	return &api.Response{
		Payload: api.Payload{
			Data: postData,
		},
		Messages: []string{SuccessfulCreatePostMessage},
	}, nil
}

// HandleUpdatePost updates a post's information in the database
func HandleUpdatePost(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	postID := chi.URLParam(r, "id")
	if postID == "" {
		return nil, errors.New("Post ID is missing")
	}

	var postInput posts.PostInput

	if err := json.NewDecoder(r.Body).Decode(&postInput); err != nil {
		return nil, errors.Wrap(err, "Failed to decode updated post input")
	}

	// Parse the timestamp string into a time.Time object
	timestamp := time.Now() // Get the current timestamp

	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to retrieve database in HandleUpdatePost")
	}

	post, err := posts.Update(db, postID, postInput, timestamp)
	if err != nil {
		return nil, errors.Wrap(err, ErrUpdatePost)
	}

	// Convert the updated post to JSON
	postJSON, err := json.Marshal(post)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to encode updated post to JSON")
	}

	// Use json.RawMessage to assign postJSON to api.Payload.Data
	postData := json.RawMessage(postJSON)

	return &api.Response{
		Payload: api.Payload{
			Data: postData,
		},
		Messages: []string{SuccessfulUpdatePostMessage},
	}, nil
}

// HandleDeletePost deletes a post from the database
func HandleDeletePost(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	postID := chi.URLParam(r, "id")
	if postID == "" {
		return nil, errors.New("Post ID is missing")
	}

	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to retrieve database in HandleDeletePost")
	}

	err = posts.Delete(db, postID)
	if err != nil {
		return nil, errors.Wrap(err, ErrDeletePost)
	}

	return &api.Response{
		Messages: []string{SuccessfulDeletePostMessage},
	}, nil
}
