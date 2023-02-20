package requests

type PaginationQuery struct {
	Page  int    `json:"page" form:"page,default=1" binding:"gte=0"`
	Limit int    `json:"limit" form:"limit,default=10" binding:"oneof=10 20 50 100"`
	Order string `json:"order" form:"order,default=desc" binding:"oneof=asc desc"`
	Skip  int    `json:"skip" form:"skip,default=1" binding:"numeric,gte=0"`
	Rel   string `json:"rel" form:"rel,default=next" binding:"required_with=cursor,oneof=next prev"`
}

type BookList struct {
	PaginationQuery
	StatusPublished string `json:"status_published" form:"status_published,default=Published"`
	ChapterStatus   string `json:"chapter_status" form:"chapter_status,default=Published"`
	Category        int64  `json:"category_id" form:"category_id"`
	BookTitle       string `json:"book_title" form:"book_title"`
	Type            string `json:"type" form:"type"`
	Status          string `json:"status" form:"status"`
	Name            string `json:"name" form:"name"`
	UserID          int64  `json:"user_id" form:"user_id"`
	OrderField      string `json:"order_field" form:"order_field,default=book.updated_at"`
	OrderBy         string `json:"order_by" form:"order_by,default=DESC" binding:"oneof=ASC DESC"`
}
