package threads

import (
	"github.com/CVWO/sample-go-app/internal/database"
	"github.com/CVWO/sample-go-app/internal/models"
)

func List(db *database.Database) ([]models.Thread, error) {
	threads := []models.Thread{
		{
			ID:       1,
			AuthorID: 2,
			TagID:    2,
			Title:    "Sample Thread 1",
			Content:  "Sample thread content 1",
		},
		{
			ID:       2,
			AuthorID: 1,
			TagID:    2,
			Title:    "Sample Thread 2",
			Content:  "Sample thread content 2",
		},
	}
	return threads, nil
}

//// GetThreadByID retrieves a thread by ID from the database
//func GetThreadByID(db *database.Database, threadID string) (*models.Thread, error) {
//	query := "SELECT id, author_id, tag_id, title, content FROM threads WHERE id = ?"
//	var thread models.Thread
//
//	err := db.DB.QueryRow(query, threadID).Scan(&thread.ID, &thread.AuthorID, &thread.TagID, &thread.Title, &thread.Content)
//	if err != nil {
//		return nil, err
//	}
//	return &thread, nil
//}

// GetThreadByID retrieves a thread by ID, first checking a local set and then the database
func GetThreadByID(db *database.Database, threadID string) (*models.Thread, error) {
	localThreads := map[string]models.Thread{
		"1": {
			ID:       1,
			AuthorID: 2,
			TagID:    2,
			Title:    "Sample Thread 1",
			Content:  "Sample thread content 1",
		},
		"2": {
			ID:       2,
			AuthorID: 1,
			TagID:    2,
			Title:    "Sample Thread 2",
			Content:  "Sample thread content 2",
		},
	}

	// Check if the thread ID exists in the local set
	if localThread, ok := localThreads[threadID]; ok {
		return &localThread, nil
	}

	// If not found in the local set, proceed with the database query
	query := "SELECT id, author_id, tag_id, title, content FROM threads WHERE id = ?"
	var thread models.Thread

	err := db.DB.QueryRow(query, threadID).Scan(&thread.ID, &thread.AuthorID, &thread.TagID, &thread.Title, &thread.Content)
	if err != nil {
		return nil, err
	}

	return &thread, nil
}
