package controllers

import (
	"net/http"
	"strconv"

	"gedebook.com/api/domain"
	"gedebook.com/api/dto/responses"
	"gedebook.com/api/errs"
	"gedebook.com/api/services"
	"github.com/gin-gonic/gin"
)

type GuestController interface {
	GetOneBook(ctx *gin.Context)
}

type guestController struct {
	userSrv services.UserService
	bookSrv services.BookService
}

func NewGuestController(userSrv services.UserService, bookSrv services.BookService) GuestController {
	return &guestController{
		userSrv: userSrv,
		bookSrv: bookSrv,
	}
}

func (ctl *guestController) GetOneBook(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		errs.ErrorHandler(ctx, 400, "Failed To Get Params")
	}
	var published_status []domain.BookPublishedStatus
	published_status = append(published_status, domain.BookPublishedStatusPublished)

	targetBook, err := ctl.bookSrv.GetOneBook(ctx, id, published_status)
	if err == nil {
		ctx.JSON(http.StatusOK, responses.R{
			Code:    http.StatusOK,
			Message: "Success Fetch Book",
			Data:    targetBook,
		})
	}
}
