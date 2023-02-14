package controllers

import (
	"net/http"
	"strconv"

	"gedebook.com/api/constants"
	"gedebook.com/api/dto/responses"
	"gedebook.com/api/errs"
	"gedebook.com/api/services"
	"github.com/gin-gonic/gin"
)

type GuestController interface {
	GetOneChapter(ctx *gin.Context)
}

type guestController struct {
	userSrv    services.UserService
	bookSrv    services.BookService
	chapterSrv services.ChapterService
}

func NewGuestController(userSrv services.UserService, bookSrv services.BookService, chapterSrv services.ChapterService) GuestController {
	return &guestController{
		userSrv:    userSrv,
		bookSrv:    bookSrv,
		chapterSrv: chapterSrv,
	}
}

func (ctl *guestController) GetOneChapter(ctx *gin.Context) {
	chapter_id, err := strconv.Atoi(ctx.Param("chapter_id"))
	if err != nil {
		errs.ErrorHandler(ctx, 400, "Failed To Get Params")
	}
	book_id, err := strconv.Atoi(ctx.Param("book_id"))
	if err != nil {
		errs.ErrorHandler(ctx, 400, "Failed To Get Params")
	}
	responseChapter, err := ctl.chapterSrv.GetOneChapter(ctx, constants.AuthnPayload{}, chapter_id, book_id)
	if err == nil {
		ctx.JSON(http.StatusOK, responses.R{
			Code:    http.StatusOK,
			Message: "Success Fetch Chapter",
			Data:    responseChapter,
		})
	}
}
