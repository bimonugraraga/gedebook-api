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

type BookController interface {
	CreateBook(ctx *gin.Context)
	UpdateBook(ctx *gin.Context)
	GetOneBook(ctx *gin.Context)
	GetAllBook(ctx *gin.Context)
}
type bookController struct {
	bookSrv services.BookService
}

func NewBookController(bookSrv services.BookService) BookController {
	return &bookController{
		bookSrv: bookSrv,
	}
}

func (ctl *bookController) CreateBook(ctx *gin.Context) {
	value, exists := ctx.Get("user")
	if !exists {
		errs.ErrorHandler(ctx, 401, "Login Is Needed")
		return
	}
	user, ok := value.(constants.AuthnPayload)
	if !ok {
		errs.ErrorHandler(ctx, 401, "Login Is Needed")
	}

	var src requests.BookRequest
	if err := ctl.ParseRequestBookEntity(ctx, &src); err != nil {
		errs.ErrorHandler(ctx, 400, "Invalid Input")
		return
	}
	err := ctl.bookSrv.CreateBook(ctx, user, src)
	if err == nil {
		ctx.JSON(http.StatusCreated, responses.R{
			Code:    http.StatusCreated,
			Message: "Success Create Book",
		})
	}
}

func (ctl *bookController) ParseRequestBookEntity(ctx *gin.Context, src *requests.BookRequest) error {
	if err := ctx.ShouldBindBodyWith(src, binding.JSON); err != nil {
		return err
	}
	return nil
}

func (ctl *bookController) UpdateBook(ctx *gin.Context) {
	value, exists := ctx.Get("user")
	if !exists {
		errs.ErrorHandler(ctx, 401, "Login Is Needed")
		return
	}
	user, ok := value.(constants.AuthnPayload)
	if !ok {
		errs.ErrorHandler(ctx, 401, "Login Is Needed")
	}
	id, err := strconv.Atoi(ctx.Param("book_id"))
	if err != nil {
		errs.ErrorHandler(ctx, 400, "Failed To Get Params")
	}
	var src requests.BookRequest
	if err := ctl.ParseRequestBookEntity(ctx, &src); err != nil {
		errs.ErrorHandler(ctx, 400, "Invalid Input")
		return
	}
	err = ctl.bookSrv.UpdateBook(ctx, user, src, id)
	if err == nil {
		ctx.JSON(http.StatusCreated, responses.R{
			Code:    http.StatusCreated,
			Message: "Success Update Book",
		})
	}
}

func (ctl *bookController) GetOneBook(ctx *gin.Context) {
	var targetBook interface{}
	id, err := strconv.Atoi(ctx.Param("book_id"))
	if err != nil {
		errs.ErrorHandler(ctx, 400, "Failed To Get Params")
	}
	value, exists := ctx.Get("user")
	if !exists {
		targetBook, err = ctl.bookSrv.GetOneBook(ctx, id, constants.AuthnPayload{})
	} else {
		user, ok := value.(constants.AuthnPayload)
		if !ok {
			errs.ErrorHandler(ctx, 401, "Login Is Needed")
			return
		}
		targetBook, err = ctl.bookSrv.GetOneBook(ctx, id, user)
	}

	if err == nil {
		ctx.JSON(http.StatusOK, responses.R{
			Code:    http.StatusOK,
			Message: "Success Fetch Book",
			Data:    targetBook,
		})
	}
}

func (ctl *bookController) GetAllBook(ctx *gin.Context) {
	var query requests.BookList
	var allBook interface{}
	var err error
	if err := ctx.ShouldBindQuery(&query); err != nil {
		errs.ErrorHandler(ctx, 400, "Invalid Query")
		return
	}
	value, exists := ctx.Get("user")
	if !exists {
		allBook, err = ctl.bookSrv.GetAllBook(ctx, &query, constants.AuthnPayload{})
	} else {
		//! To Get My Book user.ID and query.UserID must be the same
		user, ok := value.(constants.AuthnPayload)
		if !ok {
			errs.ErrorHandler(ctx, 401, "Login Is Needed")
			return
		}
		allBook, err = ctl.bookSrv.GetAllBook(ctx, &query, user)
	}
	if err == nil {
		ctx.JSON(http.StatusOK, responses.R{
			Code:    http.StatusOK,
			Message: "Success Fetch Books",
			Data:    allBook,
		})
	}
}
