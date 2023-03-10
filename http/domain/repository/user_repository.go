package repository

import (
	"context"

	"gedebook.com/api/db"
	"gedebook.com/api/domain"
	"github.com/uptrace/bun"
)

type UserRepository interface {
	Register(ctx context.Context, src *domain.User) (err error)
	GetOneUser(ctx context.Context, email string) (res domain.User, err error)
	GetOneUserByID(ctx context.Context, id int) (res domain.User, err error)
	UpdateUser(ctx context.Context, src *domain.User, id int) (err error)
}

type userRepository struct {
	db *bun.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{
		db: db.GetConn(),
	}
}

func (r *userRepository) Register(ctx context.Context, src *domain.User) (err error) {
	_, err = r.db.NewInsert().
		Model(src).
		Returning("*").
		Exec(ctx)
	if err != nil {
		return err
	}
	return
}

func (r *userRepository) GetOneUser(ctx context.Context, email string) (res domain.User, err error) {
	if err := r.db.NewSelect().
		Model(&res).
		Where("email = ?", email).
		Scan(ctx); err != nil {
		return res, err
	}
	return
}

func (r *userRepository) GetOneUserByID(ctx context.Context, id int) (res domain.User, err error) {
	if err := r.db.NewSelect().
		Model(&res).
		Where("id = ?", id).
		Scan(ctx); err != nil {
		return res, err
	}
	return
}

func (r *userRepository) UpdateUser(ctx context.Context, src *domain.User, id int) (err error) {
	_, err = r.db.NewUpdate().
		Model(src).
		Where("id = ?", id).
		Returning("*").
		ExcludeColumn("id").
		Exec(ctx)

	if err != nil {
		return err
	}
	return nil

}
