package controllers

import (
	"net/http"
	"strconv"

	"gedebook.com/api/constants"
	"gedebook.com/api/dto/requests"
	"gedebook.com/api/dto/responses"
	"gedebook.com/api/errs"
	"gedebook.com/api/services"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type ChapterController interface {
	CreateChapter(ctx *gin.Context)
	UpdateChapter(ctx *gin.Context)
	GetUserChapter(ctx *gin.Context)
	TestChapter(ctx *gin.Context)
}

type chapterController struct {
	chapterSrv services.ChapterService
}

func NewChapterController(chapterSrv services.ChapterService) ChapterController {
	return &chapterController{
		chapterSrv: chapterSrv,
	}
}

func (ctl *chapterController) CreateChapter(ctx *gin.Context) {
	value, exists := ctx.Get("user")
	if !exists {
		errs.ErrorHandler(ctx, 401, "Login Is Needed")
		return
	}
	user, ok := value.(constants.AuthnPayload)
	if !ok {
		errs.ErrorHandler(ctx, 401, "Login Is Needed")
	}

	var src requests.ChapterRequest
	if err := ctl.ParseRequestChapterEntity(ctx, &src); err != nil {
		errs.ErrorHandler(ctx, 400, "Invalid Input")
		return
	}
	err := ctl.chapterSrv.CreateChapter(ctx, user, src)
	if err == nil {
		ctx.JSON(http.StatusCreated, responses.R{
			Code:    http.StatusCreated,
			Message: "Success Create Chapter",
		})
	}

}

func (ctl *chapterController) ParseRequestChapterEntity(ctx *gin.Context, src *requests.ChapterRequest) error {
	if err := ctx.ShouldBindBodyWith(src, binding.JSON); err != nil {
		return err
	}
	return nil
}

func (ctl *chapterController) UpdateChapter(ctx *gin.Context) {
	value, exists := ctx.Get("user")
	if !exists {
		errs.ErrorHandler(ctx, 401, "Login Is Needed")
		return
	}
	user, ok := value.(constants.AuthnPayload)
	if !ok {
		errs.ErrorHandler(ctx, 401, "Login Is Needed")
	}
	id, err := strconv.Atoi(ctx.Param("chapter_id"))
	if err != nil {
		errs.ErrorHandler(ctx, 400, "Failed To Get Params")
	}

	var src requests.ChapterRequest
	if err := ctl.ParseRequestChapterEntity(ctx, &src); err != nil {
		errs.ErrorHandler(ctx, 400, "Invalid Input")
		return
	}
	err = ctl.chapterSrv.UpdateChapter(ctx, user, src, id)
	if err == nil {
		ctx.JSON(http.StatusOK, responses.R{
			Code:    http.StatusOK,
			Message: "Success Update Chapter",
		})
	}
}

func (ctl *chapterController) GetUserChapter(ctx *gin.Context) {
	value, exists := ctx.Get("user")
	if !exists {
		errs.ErrorHandler(ctx, 401, "Login Is Needed")
	}
	user, ok := value.(constants.AuthnPayload)
	if !ok {
		errs.ErrorHandler(ctx, 401, "Login Is Needed")
	}
	chapter_id, err := strconv.Atoi(ctx.Param("chapter_id"))
	if err != nil {
		errs.ErrorHandler(ctx, 400, "Failed To Get Params")
	}
	book_id, err := strconv.Atoi(ctx.Param("book_id"))
	if err != nil {
		errs.ErrorHandler(ctx, 400, "Failed To Get Params")
	}
	responseChapter, err := ctl.chapterSrv.GetOneChapter(ctx, user, chapter_id, book_id)
	if err == nil {
		ctx.JSON(http.StatusOK, responses.R{
			Code:    http.StatusOK,
			Message: "Success Fetch Chapter",
			Data:    responseChapter,
		})
	}

}

func (ctl *chapterController) TestChapter(ctx *gin.Context) {
	var targetChapter interface{}
	chapter_id, err := strconv.Atoi(ctx.Param("chapter_id"))
	if err != nil {
		errs.ErrorHandler(ctx, 400, "Failed To Get Params")
		return
	}
	book_id, err := strconv.Atoi(ctx.Param("book_id"))
	if err != nil {
		errs.ErrorHandler(ctx, 400, "Failed To Get Params")
		return
	}
	value, exists := ctx.Get("user")
	if !exists {
		targetChapter, err = ctl.chapterSrv.GetOneChapter(ctx, constants.AuthnPayload{}, chapter_id, book_id)
	} else {
		user, ok := value.(constants.AuthnPayload)
		if !ok {
			errs.ErrorHandler(ctx, 401, "Login Is Needed")
		}
		targetChapter, err = ctl.chapterSrv.GetOneChapter(ctx, user, chapter_id, book_id)
	}
	if err == nil {
		ctx.JSON(http.StatusOK, responses.R{
			Code:    http.StatusOK,
			Message: "Success Fetch Chapter",
			Data:    targetChapter,
		})
	}
}
