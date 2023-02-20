package domain

import (
	"time"

	"github.com/uptrace/bun"
)

type BookType string

const (
	BookTypeNovel      BookType = "Novel"
	BookTypeShortStory BookType = "Short Story"
)

func (Type BookType) ValidBookType() bool {
	switch Type {
	case BookTypeNovel, BookTypeShortStory:
		return true
	default:
		return false
	}
}

type BookStatus string

const (
	BookStatusOngoing   BookStatus = "Ongoing"
	BookStatusCompleted BookStatus = "Completed"
)

func (Type BookStatus) ValidBookStatus() bool {
	switch Type {
	case BookStatusOngoing, BookStatusCompleted:
		return true
	default:
		return false
	}
}

type BookPublishedStatus string

const (
	BookPublishedStatusDraft     BookPublishedStatus = "Draft"
	BookPublishedStatusPublished BookPublishedStatus = "Published"
	BookPublishedStatusCancelled BookPublishedStatus = "Cancelled"
)

func (Type BookPublishedStatus) ValidBookPublishedStatus() bool {
	switch Type {
	case BookPublishedStatusDraft, BookPublishedStatusPublished, BookPublishedStatusCancelled:
		return true
	default:
		return false
	}
}

type Book struct {
	bun.BaseModel
	ID              int64     `bun:"id,pk,autoincrement" json:"id"`
	UserID          int64     `bun:"user_id,nullzero" json:"user_id"`
	User            User      `bun:"rel:belongs-to,join:user_id=id" json:"user"`
	Title           string    `bun:"title" json:"title"`
	BookCover       *string   `bun:"book_cover" json:"book_cover"`
	Type            string    `bun:"type" json:"type"`
	MainCategoryID  int64     `bun:"main_category_id,nullzero" json:"main_category_id"`
	Category        Category  `bun:"rel:belongs-to,join:main_category_id=id" json:"category"`
	Status          string    `bun:"status" json:"status"`
	Chapter         []Chapter `bun:"rel:has-many,join:id=book_id" json:"chapters"`
	PublishedStatus *string   `bun:"published_status" json:"published_status"`
	Synopsis        *string   `bun:"synopsis" json:"synopsis"`
	UpdatedAt       time.Time `bun:"updated_at,nullzero" json:"updated_at"`
	CreatedAt       time.Time `bun:"created_at,nullzero" json:"created_at"`
}
