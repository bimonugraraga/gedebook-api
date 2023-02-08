package domain

import (
	"time"

	"github.com/uptrace/bun"
)

type ChapterPublishedStatus string

const (
	ChapterPublishedStatusDraft     ChapterPublishedStatus = "Draft"
	ChapterPublishedStatusPublished ChapterPublishedStatus = "Published"
	ChapterPublishedStatusCancelled ChapterPublishedStatus = "Cancelled"
)

func (Type ChapterPublishedStatus) ValidChapterPublishedStatus() bool {
	switch Type {
	case ChapterPublishedStatusDraft, ChapterPublishedStatusPublished, ChapterPublishedStatusCancelled:
		return true
	default:
		return false
	}
}

type Chapter struct {
	bun.BaseModel
	ID              int64     `bun:"id,pk,autoincrement" json:"id"`
	BookID          int64     `bun:"book_id,nullzero" json:"book_id"`
	Book            *Book     `bun:"rel:belongs-to,join:book_id=id" json:"book"`
	ChapterTitle    string    `bun:"chapter_title" json:"chapter_title"`
	ChapterText     string    `bun:"chapter_text" json:"chapter_text"`
	ChapterCover    *string   `bun:"chapter_cover" json:"chapter_cover"`
	PublishedStatus *string   `bun:"published_status" json:"published_status"`
	UpdatedAt       time.Time `bun:"updated_at,nullzero" json:"updated_at"`
	CreatedAt       time.Time `bun:"created_at,nullzero" json:"created_at"`
}
