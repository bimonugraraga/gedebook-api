package repository

import (
	"context"
	"fmt"

	"gedebook.com/api/db"
	"gedebook.com/api/domain"
	"github.com/uptrace/bun"
)

type BookRepository interface {
	CreateOneBook(ctx context.Context, src *domain.Book) (err error)
	UpdateOneBook(ctx context.Context, src *domain.Book, id int) (err error)
	GetUserBook(ctx context.Context, book_id int, user_id int) (domain.Book, error)
	GetOneBook(ctx context.Context, id int, published_status []domain.BookPublishedStatus) (domain.Book, error)
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

func (r *bookRepository) GetUserBook(ctx context.Context, book_id int, user_id int) (domain.Book, error) {
	res := domain.Book{}
	if err := r.db.NewSelect().
		Model(&res).
		Where(fmt.Sprintf("id = %d AND user_id = %d", book_id, user_id)).
		Scan(ctx); err != nil {
		return res, err
	}
	return res, nil
}

func (r *bookRepository) GetOneBook(ctx context.Context, id int, published_status []domain.BookPublishedStatus) (domain.Book, error) {
	res := domain.Book{}
	if err := r.db.NewSelect().
		Model(&res).
		Where("Book.id = ?", id).
		Where("Book.published_status IN (?)", bun.In(published_status)).
		Relation("User").
		Relation("Category").
		Relation("Chapter").
		Scan(ctx); err != nil {
		return res, err
	}
	return res, nil
}
