package repository

import (
	"context"

	"gedebook.com/api/db"
	"gedebook.com/api/domain"
	"github.com/uptrace/bun"
)

type UserRepository interface {
	Register(ctx context.Context, tx *bun.Tx, src *domain.User) (err error)
}

type userRepository struct {
	db *bun.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{
		db: db.GetConn(),
	}
}

func (r *userRepository) Register(ctx context.Context, tx *bun.Tx, src *domain.User) (err error) {
	_, err = tx.NewInsert().
		Model(src).
		Returning("*").
		Exec(ctx)
	if err != nil {
		return err
	}
	return
}
