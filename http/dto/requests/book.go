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
	Synopsis       *string `json:"synopsis" form:"synopsis"`
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
