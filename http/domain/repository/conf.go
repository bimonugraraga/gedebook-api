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
	Chapter  ChapterRepository
}

func InitRepositoryInstance() *repositoryPool {
	return &repositoryPool{
		Admin:    NewAdminRepository(),
		User:     NewUserRepository(),
		Book:     NewBookRepository(),
		Category: NewCategoryRepository(),
		Chapter:  NewChapterRepository(),
	}
}
func NewContext(parent context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(parent, 5*time.Second)
}
