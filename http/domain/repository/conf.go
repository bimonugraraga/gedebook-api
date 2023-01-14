package repository

import (
	"context"
	"time"
)

type repositoryPool struct {
	Admin AdminRepository
}

func InitRepositoryInstance() *repositoryPool {
	return &repositoryPool{
		Admin: NewAdminRepository(),
	}
}
func NewContext(parent context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(parent, 5*time.Second)
}
