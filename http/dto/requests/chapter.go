package requests

import (
	"gedebook.com/api/domain"
	"github.com/jinzhu/copier"
)

type ChapterRequest struct {
	BookID       int64   `json:"book_id" form:"book_id" binding:"required"`
	ChapterTitle string  `json:"chapter_title" form:"chapter_title" binding:"required"`
	ChapterText  string  `json:"chapter_text" form:"chapter_text" binding:"required"`
	ChapterCover *string `json:"chapter_cover" form:"chapter_cover"`
}

func (src ChapterRequest) AssignedChapterRequest() (res domain.Chapter, err error) {
	if err := copier.Copy(&res, &src); err != nil {
		return domain.Chapter{}, err
	}
	return
}
