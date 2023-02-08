package repository

import (
	"context"
	"time"
)

type repositoryPool struct {
	Admin    AdminRepository
	User     UserRepository
	Book     BookRepository
	Category CategoryRepository
}

func InitRepositoryInstance() *repositoryPool {
	return &repositoryPool{
		Admin:    NewAdminRepository(),
		User:     NewUserRepository(),
		Book:     NewBookRepository(),
		Category: NewCategoryRepository(),
	}
}
func NewContext(parent context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(parent, 5*time.Second)
}
