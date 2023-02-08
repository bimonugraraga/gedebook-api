package repository

import (
	"context"

	"gedebook.com/api/db"
	"gedebook.com/api/domain"
	"github.com/uptrace/bun"
)

type CategoryRepository interface {
	GetOneByID(ctx context.Context, id int) (res domain.Category, err error)
}

type categoryRepository struct {
	db *bun.DB
}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{
		db: db.GetConn(),
	}
}

func (r *categoryRepository) GetOneByID(ctx context.Context, id int) (res domain.Category, err error) {
	if err := r.db.NewSelect().
		Model(&res).
		Where("id = ?", id).
		Scan(ctx); err != nil {
		return res, err
	}
	return
}
