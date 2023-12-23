package database

import (
	"github.com/CVWO/sample-go-app/internal/models"
	"gorm.io/gorm"
	"time"
)

func Seed(db *gorm.DB) error {
	// Create users
	users := []models.Tag{
		{
			ID:   1,
			Name: "CVWO",
		},
		{
			ID:   2,
			Name: "Kenneth",
		},
	}
	db.Create(&users)

	// Create tags
	tags := []models.Tag{
		{
			ID:   1,
			Name: "Opinion",
		},
		{
			ID:   2,
			Name: "Satirical",
		},
	}
	db.Create(&tags)

	// Create threads
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
	db.Create(&threads)

	// Create posts
	posts := []models.Post{
		{
			ID:        1,
			ThreadID:  1,
			AuthorID:  1,
			Content:   "Sample post content 1",
			Timestamp: time.Now(),
		},
		{
			ID:        2,
			ThreadID:  1,
			AuthorID:  2,
			Content:   "Sample post content 2",
			Timestamp: time.Now(),
		},
	}
	db.Create(&posts)

	return nil
}
