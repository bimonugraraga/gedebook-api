package requests

import (
	"gedebook.com/api/domain"
	"github.com/jinzhu/copier"
)

type BookRequest struct {
	Title          string  `json:"title" form:"title" binding:"required"`
	Type           string  `json:"type" form:"type" binding:"required"`
	MainCategoryID int64   `json:"main_category_id" form:"main_category_id" binding:"required"`
	BookCover      *string `json:"book_cover" form:"book_cover"`
}

type ChapterRequest struct {
	BookID       int64  `json:"book_id" form:"book_id" binding:"required"`
	ChapterTitle string `json:"chapter_title" form:"chapter_title" binding:"required"`
	ChapterText  string `json:"chapter_text" form:"chapter_text" binding:"required"`
	ChapterCover string `json:"chapter_cover" form:"chapter_cover" binding:"required"`
}

type UpdateBookStatusRequest struct {
	Status string `json:"status" form:"status" binding:"required"`
}

func (src BookRequest) AssignedBookRequest() (res domain.Book, err error) {
	if err := copier.Copy(&res, &src); err != nil {
		return domain.Book{}, err
	}
	return
}
