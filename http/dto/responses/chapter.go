package responses

import (
	"time"

	"gedebook.com/api/domain"
)

type BookChapterInfo struct {
	ID              int64   `json:"id"`
	Title           string  `json:"title"`
	PublishedStatus string  `json:"published_status"`
	BookCover       *string `json:"book_cover"`
}
type ChapterResponse struct {
	ID              int64     `json:"id"`
	ChapterTitle    string    `json:"chapter_title"`
	PublishedStatus string    `json:"published_status"`
	ChapterCover    *string   `json:"chapter_cover"`
	ChapterText     string    `json:"chapter_text"`
	UpdatedAt       time.Time `json:"updated_at"`
	AuthorName      string    `json:"author"`
	AuthorProfile   AuthorResponse
	Book            BookChapterInfo
}

func AssignedGetOneChapter(src domain.Chapter, user domain.User) (res ChapterResponse) {
	res.ID = src.ID
	res.ChapterTitle = src.ChapterTitle
	res.PublishedStatus = *src.PublishedStatus
	res.ChapterCover = src.ChapterCover
	res.ChapterText = src.ChapterText
	res.UpdatedAt = src.UpdatedAt
	res.AuthorName = user.Name
	res.AuthorProfile.ID = user.ID
	res.AuthorProfile.Name = user.Name
	res.AuthorProfile.ProfilePicture = user.ProfilePicture
	res.Book.ID = src.BookID
	res.Book.Title = src.Book.Title
	res.Book.BookCover = src.Book.BookCover
	res.Book.PublishedStatus = *src.Book.PublishedStatus

	return
}
