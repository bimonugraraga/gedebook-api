package repository

import (
	"context"
	"fmt"
	"strings"

	"gedebook.com/api/db"
	"gedebook.com/api/domain"
	"gedebook.com/api/dto/requests"
	"github.com/uptrace/bun"
)

type BookRepository interface {
	CreateOneBook(ctx context.Context, src *domain.Book) (err error)
	UpdateOneBook(ctx context.Context, src *domain.Book, id int) (err error)
	GetUserBook(ctx context.Context, book_id int, user_id int) (domain.Book, error)
	GetOneBook(ctx context.Context, id int) (domain.Book, error)
	GetAllBook(ctx context.Context, paging *requests.BookList) (res []domain.Book, total int, err error)
}

type bookRepository struct {
	db *bun.DB
}

func NewBookRepository() BookRepository {
	return &bookRepository{
		db: db.GetConn(),
	}
}

func (r *bookRepository) CreateOneBook(ctx context.Context, src *domain.Book) (err error) {
	_, err = r.db.NewInsert().
		Model(src).
		Returning("*").
		Exec(ctx)
	if err != nil {
		return err
	}
	return
}

func (r *bookRepository) UpdateOneBook(ctx context.Context, src *domain.Book, id int) (err error) {
	_, err = r.db.NewUpdate().
		Model(src).
		Where("id = ?", id).
		Returning("*").
		Exec(ctx)
	if err != nil {
		return err
	}
	return
}

//!For User No Restriction
func (r *bookRepository) GetUserBook(ctx context.Context, book_id int, user_id int) (domain.Book, error) {
	res := domain.Book{}
	if err := r.db.NewSelect().
		Model(&res).
		Where(fmt.Sprintf("id = %d AND user_id = %d", book_id, user_id)).
		Relation("Chapter").
		Scan(ctx); err != nil {
		return res, err
	}
	return res, nil
}

//!For Guest Restriction Must Be Published
func (r *bookRepository) GetOneBook(ctx context.Context, id int) (domain.Book, error) {
	res := domain.Book{}
	if err := r.db.NewSelect().
		Model(&res).
		Where("Book.id = ?", id).
		Where("Book.published_status = ?", domain.BookPublishedStatusPublished).
		Relation("User").
		Relation("Category").
		Relation("Chapter", func(query *bun.SelectQuery) *bun.SelectQuery {
			return query.Where("published_status = ?", domain.ChapterPublishedStatusPublished).Order("id DESC")
		}).
		Scan(ctx); err != nil {
		return res, err
	}

	return res, nil
}

func (r *bookRepository) GetAllBook(ctx context.Context, paging *requests.BookList) (res []domain.Book, total int, err error) {
	res = []domain.Book{}
	offset := (paging.Page - 1) * paging.Limit
	q := r.db.NewSelect().
		Model(&res).
		Offset(offset).
		Limit(paging.Limit).
		Relation("User").
		Relation("Chapter", func(query *bun.SelectQuery) *bun.SelectQuery {
			chapter_status := strings.Split(strings.ToLower(paging.ChapterStatus), ",")
			return query.Where("lower(published_status) IN (?)", bun.In(chapter_status))
		}).
		Relation("Category")

	if len(paging.StatusPublished) != 0 {
		fmt.Println(paging.StatusPublished)
		status_published := strings.Split(strings.ToLower(paging.StatusPublished), ",")
		q.Where("lower(Book.published_status) IN (?)", bun.In(status_published))
	}
	if paging.Category != 0 {
		q.Where("Book.main_category_id = ?", paging.Category)
	}
	if len(paging.BookTitle) != 0 {
		q.Where("Book.title ILIKE ?", fmt.Sprintf("%%%s%%", paging.BookTitle))
	}
	if len(paging.Status) != 0 {
		status := strings.Split(strings.ToLower(paging.Status), ",")
		q.Where("lower(Book.status) IN (?)", bun.In(status))
	}
	if len(paging.Name) != 0 {
		q.Where("name ILIKE ?", fmt.Sprintf("%%%s%%", paging.Name))
	}
	if paging.UserID != 0 {
		q.Where("user_id = ?", paging.UserID)
	}
	total, err = q.ScanAndCount(ctx)

	return res, total, err
}
