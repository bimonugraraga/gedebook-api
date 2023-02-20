package responses

import (
	"math"

	"gedebook.com/api/utils"
)

type PagingResponse struct {
	TotalRecord  int         `json:"total_record"`
	NumberRecord int         `json:"number_record"`
	TotalPage    int         `json:"total_page"`
	CurrentPage  int         `json:"current_page"`
	PerPage      int         `json:"per_page"`
	HasNext      bool        `json:"has_next"`
	HasPrev      bool        `json:"has_prev"`
	Records      interface{} `json:"records"`
}

func NewPagingResponse(limit int, page int, totalRecord int, totalPageRecord int) (res PagingResponse) {
	offset := utils.CalculateOffset(limit, page)

	res.TotalRecord = totalRecord
	res.NumberRecord = totalPageRecord
	res.TotalPage = int(math.Ceil(float64(totalRecord) / float64(limit)))
	res.CurrentPage = (offset / limit) + 1
	res.PerPage = limit
	res.HasNext = res.TotalPage >= 1 && res.TotalPage > res.CurrentPage
	res.HasPrev = res.TotalPage >= 1 && res.TotalPage+1 >= res.CurrentPage && res.CurrentPage > 1

	return
}
