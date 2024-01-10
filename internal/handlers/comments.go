package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/itstrueitstrueitsrealitsreal/gossip-with-go-be/internal/api"
	"github.com/itstrueitstrueitsrealitsreal/gossip-with-go-be/internal/dataaccess/comments"
	"github.com/itstrueitstrueitsrealitsreal/gossip-with-go-be/internal/database"
	"github.com/itstrueitstrueitsrealitsreal/gossip-with-go-be/internal/models"
	"github.com/pkg/errors"
)

const (
	ListComments                   = "comments.HandleList"
	SuccessfulListCommentsMessage  = "Successfully listed comments"
	ErrRetrieveComments            = "Failed to retrieve comments in %s"
	ErrEncodeComments              = "Failed to encode comments in %s"
	SuccessfulCreateCommentMessage = "Successfully created comment"
	ErrCreateComment               = "Failed to create comment"
	SuccessfulUpdateCommentMessage = "Successfully updated comment"
	ErrUpdateComment               = "Failed to update comment"
	SuccessfulDeleteCommentMessage = "Successfully deleted comment"
	ErrDeleteComment               = "Failed to delete comment"
	SuccessfulViewCommentMessage   = "Successfully viewed comment"
)

// HandleListComments returns all comments in JSON format.
func HandleListComments(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, ListComments))
	}

	commentList, err := comments.List(db)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveComments, ListComments))
	}
	data, err := json.Marshal(commentList)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrEncodeComments, ListComments))
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulListCommentsMessage},
	}, nil
}

// HandleGetComment retrieves a single comment by ID.
func HandleGetComment(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	commentID := chi.URLParam(r, "id")
	if commentID == "" {
		return nil, errors.New("Comment ID is missing")
	}

	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to retrieve database in HandleGetComment")
	}

	comment, err := comments.GetCommentByID(db, commentID)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to retrieve comment in HandleGetComment")
	}

	data, err := json.Marshal(comment)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to encode comment in HandleGetComment")
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulViewCommentMessage},
	}, nil
}

// HandleCreateComment creates a new comment and inserts it into the database.
func HandleCreateComment(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	var commentInput models.CommentInput

	if err := json.NewDecoder(r.Body).Decode(&commentInput); err != nil {
		return nil, errors.Wrap(err, "Failed to decode comment input")
	}

	// Parse the timestamp string into a time.Time object
	timestamp := time.Now() // Get the current timestamp

	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to retrieve database in HandleCreateComment")
	}

	comment, err := comments.Create(db, commentInput, timestamp)
	if err != nil {
		return nil, errors.Wrap(err, ErrCreateComment)
	}

	// Convert the comment to JSON
	commentJSON, err := json.Marshal(comment)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to encode comment to JSON")
	}

	// Use json.RawMessage to assign commentJSON to api.Payload.Data
	commentData := json.RawMessage(commentJSON)

	return &api.Response{
		Payload: api.Payload{
			Data: commentData,
		},
		Messages: []string{SuccessfulCreateCommentMessage},
	}, nil
}

// HandleUpdateComment updates a comment's information in the database.
func HandleUpdateComment(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	commentID := chi.URLParam(r, "id")
	if commentID == "" {
		return nil, errors.New("Comment ID is missing")
	}

	var commentInput models.CommentInput

	if err := json.NewDecoder(r.Body).Decode(&commentInput); err != nil {
		return nil, errors.Wrap(err, "Failed to decode updated comment input")
	}

	// Parse the timestamp string into a time.Time object
	timestamp := time.Now() // Get the current timestamp

	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to retrieve database in HandleUpdateComment")
	}

	comment, err := comments.Update(db, commentID, commentInput, timestamp)
	if err != nil {
		return nil, errors.Wrap(err, ErrUpdateComment)
	}

	// Convert the updated comment to JSON
	commentJSON, err := json.Marshal(comment)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to encode updated comment to JSON")
	}

	// Use json.RawMessage to assign commentJSON to api.Payload.Data
	commentData := json.RawMessage(commentJSON)

	return &api.Response{
		Payload: api.Payload{
			Data: commentData,
		},
		Messages: []string{SuccessfulUpdateCommentMessage},
	}, nil
}

// HandleDeleteComment deletes a comment from the database.
func HandleDeleteComment(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	commentID := chi.URLParam(r, "id")
	if commentID == "" {
		return nil, errors.New("Comment ID is missing")
	}

	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to retrieve database in HandleDeleteComment")
	}

	err = comments.Delete(db, commentID)
	if err != nil {
		return nil, errors.Wrap(err, ErrDeleteComment)
	}

	return &api.Response{
		Messages: []string{SuccessfulDeleteCommentMessage},
	}, nil
}

// HandleGetCommentsByThread retrieves all comments for a specific thread by thread ID.
func HandleGetCommentsByThread(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	threadID := chi.URLParam(r, "id")
	if threadID == "" {
		return nil, errors.New("Thread ID is missing")
	}

	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to retrieve database in HandleGetCommentsByThread")
	}

	comments, err := comments.GetCommentsByThreadID(db, threadID)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to retrieve comments in HandleGetCommentsByThread")
	}

	data, err := json.Marshal(comments)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to encode comments in HandleGetCommentsByThread")
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{"Successfully retrieved comments"},
	}, nil
}
