package responses

import (
	"time"

	"gedebook.com/api/domain"
)

type CategoryResponse struct {
	Name string `json:"name"`
	ID   int64  `json:"id"`
}
type AuthorResponse struct {
	Name           string  `json:"name"`
	ID             int64   `json:"id"`
	ProfilePicture *string `json:"profile_picture"`
}
type PreviewChapter struct {
	ID              int64     `json:"id"`
	ChapterTitle    string    `json:"chapter_title"`
	PublishedStatus string    `json:"published_status"`
	UpdatedAt       time.Time `json:"updated_at"`
}
type BookResponse struct {
	ID             int64   `json:"id"`
	Title          string  `json:"title"`
	BookCover      *string `json:"book_cover"`
	Type           string  `json:"type"`
	MainCategoryID int64   `json:"main_category_id"`
	Category       CategoryResponse
	AuthorName     string `json:"author"`
	AuthorProfile  AuthorResponse
	ChapterCount   int `json:"chapter_count"`
	Chapters       []PreviewChapter
}

func AssignedGetOneBook(src domain.Book) (res BookResponse) {
	res.ID = src.ID
	res.Title = src.Title
	res.BookCover = src.BookCover
	res.Type = src.Type
	res.MainCategoryID = src.MainCategoryID
	res.Category.ID = src.Category.ID
	res.Category.Name = src.Category.CategoryName
	res.AuthorName = src.User.Name
	res.AuthorProfile.ID = src.UserID
	res.AuthorProfile.Name = src.User.Name
	res.AuthorProfile.ProfilePicture = src.User.ProfilePicture

	for i := 0; i < len(src.Chapter); i++ {
		chapter := PreviewChapter{
			ID:              src.Chapter[i].ID,
			ChapterTitle:    src.Chapter[i].ChapterTitle,
			UpdatedAt:       src.Chapter[i].UpdatedAt,
			PublishedStatus: *src.Chapter[i].PublishedStatus,
		}
		res.Chapters = append(res.Chapters, chapter)
	}
	if len(res.Chapters) == 0 {
		res.Chapters = make([]PreviewChapter, 0)
	}
	res.ChapterCount = len(res.Chapters)
	return
}
