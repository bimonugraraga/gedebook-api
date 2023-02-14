package services

import (
	"context"
	"database/sql"

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

type ChapterService interface {
	CreateChapter(ctx *gin.Context, user constants.AuthnPayload, src requests.ChapterRequest) (err error)
	UpdateChapter(ctx *gin.Context, user constants.AuthnPayload, src requests.ChapterRequest, id int) error
	GetOneChapter(ctx *gin.Context, user constants.AuthnPayload, chapter_id int, book_id int) (*responses.ChapterResponse, error)
}

type chapterService struct {
	chapterRepo repository.ChapterRepository
	bookRepo    repository.BookRepository
	userRepo    repository.UserRepository
}

func NewChapterService(chapterRepo repository.ChapterRepository, bookRepo repository.BookRepository, userRepo repository.UserRepository) ChapterService {
	return &chapterService{
		chapterRepo: chapterRepo,
		bookRepo:    bookRepo,
		userRepo:    userRepo,
	}
}

func (srv *chapterService) CreateChapter(ctx *gin.Context, user constants.AuthnPayload, src requests.ChapterRequest) (err error) {
	c, cancel := repository.NewContext(ctx)
	defer cancel()
	if err := db.GetConn().RunInTx(c, &sql.TxOptions{}, func(context context.Context, tx bun.Tx) (err error) {
		_, err = srv.bookRepo.GetUserBook(ctx, int(src.BookID), int(user.ID))
		if err != nil {
			errs.ErrorHandler(ctx, 404, "Book Not Found")
			return err
		}
		newChapter, err := src.AssignedChapterRequest()
		if err != nil {
			errs.ErrorHandler(ctx, 400, "Failed To Create Chapter")
			return err
		}
		err = srv.chapterRepo.CreateOneChapter(ctx, &newChapter)
		if err != nil {
			errs.ErrorHandler(ctx, 400, "Failed To Create Chapter")
			return err
		}
		return
	}); err != nil {
		return err
	}
	return nil
}

func (srv *chapterService) UpdateChapter(ctx *gin.Context, user constants.AuthnPayload, src requests.ChapterRequest, id int) error {
	c, cancel := repository.NewContext(ctx)
	defer cancel()
	if err := db.GetConn().RunInTx(c, &sql.TxOptions{}, func(context context.Context, tx bun.Tx) (err error) {
		_, err = srv.bookRepo.GetUserBook(ctx, int(src.BookID), int(user.ID))
		if err != nil {
			errs.ErrorHandler(ctx, 404, "Book Not Found")
			return err
		}
		var status_published []domain.ChapterPublishedStatus
		status_published = append(status_published, domain.ChapterPublishedStatusDraft, domain.ChapterPublishedStatusPublished, domain.ChapterPublishedStatusCancelled)
		targetChapter, err := srv.chapterRepo.GetOneChapter(ctx, int(src.BookID), id, status_published, int(user.ID))
		if err != nil {
			errs.ErrorHandler(ctx, 404, "Chapter Not Found")
			return err
		}
		updateChapter := targetChapter
		if len(src.ChapterTitle) != 0 {
			updateChapter.ChapterTitle = src.ChapterTitle
		}
		if len(src.ChapterText) != 0 {
			updateChapter.ChapterText = src.ChapterText
		}
		if src.ChapterCover != nil {
			updateChapter.ChapterCover = src.ChapterCover
		}

		err = srv.chapterRepo.UpdateOneChapter(ctx, int(src.BookID), id, &updateChapter)
		if err != nil {
			errs.ErrorHandler(ctx, 400, "Failed To Update Chapter")
			return err
		}
		return
	}); err != nil {
		return err
	}
	return nil
}

func (srv *chapterService) GetOneChapter(ctx *gin.Context, user constants.AuthnPayload, chapter_id int, book_id int) (*responses.ChapterResponse, error) {
	var status_published []domain.ChapterPublishedStatus
	if user.ID != 0 {
		_, err := srv.bookRepo.GetUserBook(ctx, book_id, int(user.ID))
		if err != nil {
			user.ID = 0
			status_published = append(status_published, domain.ChapterPublishedStatusPublished)
		} else {
			status_published = append(status_published, domain.ChapterPublishedStatusDraft, domain.ChapterPublishedStatusPublished, domain.ChapterPublishedStatusCancelled)
		}
	} else {
		status_published = append(status_published, domain.ChapterPublishedStatusPublished)
	}
	targetChapter, err := srv.chapterRepo.GetOneChapter(ctx, book_id, chapter_id, status_published, int(user.ID))
	if err != nil {
		errs.ErrorHandler(ctx, 404, "Chapter Not Found")
		return nil, err
	}
	targetUser, err := srv.userRepo.GetOneUserByID(ctx, int(targetChapter.Book.UserID))
	if err != nil {
		errs.ErrorHandler(ctx, 404, "User Not Found")
		return nil, err
	}
	responseChapter := responses.AssignedGetOneChapter(targetChapter, targetUser)
	return &responseChapter, nil
}
