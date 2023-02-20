package services

import (
	"context"
	"database/sql"
	"fmt"

	"gedebook.com/api/constants"
	"gedebook.com/api/db"
	"gedebook.com/api/domain"
	"gedebook.com/api/domain/repository"
	"gedebook.com/api/dto/requests"
	"gedebook.com/api/dto/responses"
	"gedebook.com/api/errs"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type BookService interface {
	CreateBook(ctx *gin.Context, user constants.AuthnPayload, src requests.BookRequest) (err error)
	UpdateBook(ctx *gin.Context, user constants.AuthnPayload, src requests.BookRequest, id int) (err error)
	GetOneBook(ctx *gin.Context, id int, user constants.AuthnPayload) (*responses.BookResponse, error)
	GetAllBook(ctx *gin.Context, query *requests.BookList, user constants.AuthnPayload) (res responses.PagingResponse, err error)
}

type bookService struct {
	bookRepo     repository.BookRepository
	userRepo     repository.UserRepository
	categoryRepo repository.CategoryRepository
}

func NewBookService(bookRepo repository.BookRepository, userRepo repository.UserRepository, categoryRepo repository.CategoryRepository) BookService {
	return &bookService{
		bookRepo:     bookRepo,
		userRepo:     userRepo,
		categoryRepo: categoryRepo,
	}
}

func (srv *bookService) CreateBook(ctx *gin.Context, user constants.AuthnPayload, src requests.BookRequest) error {
	if !domain.BookType.ValidBookType(domain.BookType(src.Type)) {
		errs.ErrorHandler(ctx, 400, "Invalid Book Type")
	}
	c, cancel := repository.NewContext(ctx)
	defer cancel()
	if err := db.GetConn().RunInTx(c, &sql.TxOptions{}, func(context context.Context, tx bun.Tx) (err error) {
		_, err = srv.categoryRepo.GetOneByID(ctx, int(src.MainCategoryID))
		if err != nil {
			errs.ErrorHandler(ctx, 404, "Category Not Found")
			return err
		}
		_, err = srv.userRepo.GetOneUserByID(ctx, int(user.ID))
		if err != nil {
			errs.ErrorHandler(ctx, 404, "User Not Found")
			return err
		}
		newBook, err := src.AssignedBookRequest()
		if err != nil {
			errs.ErrorHandler(ctx, 400, "Failed To Create Book")
			return err
		}
		newBook.UserID = user.ID
		newBook.Status = string(domain.BookStatusOngoing)
		err = srv.bookRepo.CreateOneBook(ctx, &newBook)
		if err != nil {
			errs.ErrorHandler(ctx, 400, "Failed To Create Book")
			return err
		}
		return
	}); err != nil {
		return err
	}
	return nil
}

func (srv *bookService) UpdateBook(ctx *gin.Context, user constants.AuthnPayload, src requests.BookRequest, id int) (err error) {
	if !domain.BookType.ValidBookType(domain.BookType(src.Type)) {
		errs.ErrorHandler(ctx, 400, "Invalid Book Type")
	}

	c, cancel := repository.NewContext(ctx)
	defer cancel()
	if err := db.GetConn().RunInTx(c, &sql.TxOptions{}, func(context context.Context, tx bun.Tx) (err error) {
		targetBook, err := srv.bookRepo.GetUserBook(ctx, id, int(user.ID))
		if err != nil {
			errs.ErrorHandler(ctx, 404, "Book Not Found")
			return err
		}
		_, err = srv.categoryRepo.GetOneByID(ctx, int(src.MainCategoryID))
		if err != nil {
			errs.ErrorHandler(ctx, 404, "Category Not Found")
			return err
		}

		updateBook := targetBook
		updateBook.MainCategoryID = src.MainCategoryID
		updateBook.Type = src.Type
		if src.BookCover != nil {
			updateBook.BookCover = src.BookCover
		}
		if len(src.Title) != 0 {
			updateBook.Title = src.Title
		}

		err = srv.bookRepo.UpdateOneBook(ctx, &updateBook, id)
		if err != nil {
			errs.ErrorHandler(ctx, 400, "Failed To Update Book")
			return err
		}
		return
	}); err != nil {
		return err
	}
	return nil
}

func (srv *bookService) GetOneBook(ctx *gin.Context, id int, user constants.AuthnPayload) (*responses.BookResponse, error) {
	var targetBook domain.Book
	var err error
	if user.ID == 0 {
		targetBook, err = srv.bookRepo.GetOneBook(ctx, id)

	} else {
		targetBook, err = srv.bookRepo.GetUserBook(ctx, id, int(user.ID))
		if err != nil {
			targetBook, err = srv.bookRepo.GetOneBook(ctx, id)
		}
	}
	if err != nil {
		errs.ErrorHandler(ctx, 404, "Book Not Found")
		return nil, err
	}
	responseBook := responses.AssignedGetOneBook(targetBook)
	return &responseBook, nil
}

func (srv *bookService) GetAllBook(ctx *gin.Context, query *requests.BookList, user constants.AuthnPayload) (res responses.PagingResponse, err error) {
	fmt.Println(query)

	if query.UserID == 0 {
		//!Guest
		query.UserID = 0
		query.ChapterStatus = "Published"
		query.StatusPublished = "Published"
	} else if query.UserID != 0 && query.UserID != user.ID {
		//!Login But Did Not Checking My Book
		query.UserID = 0
		query.ChapterStatus = "Published"
		query.StatusPublished = "Published"
	} else if query.UserID != 0 && query.UserID == user.ID {
		//!Login But Checking My Book
		query.UserID = user.ID
	}

	data, total, err := srv.bookRepo.GetAllBook(ctx, query)
	res = responses.NewPagingResponse(query.Limit, query.Page, total, len(data))
	res.Records, err = responses.ListPaginatedBook(data)
	if err != nil {
		errs.ErrorHandler(ctx, 400, "Something Went Wrong")
		return res, err
	}

	return
}
