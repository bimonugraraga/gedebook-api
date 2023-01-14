package repository

import (
	"gedebook.com/api/db"
	"github.com/uptrace/bun"
)

type AdminRepository interface {
	//!Attach Your Function Here
}

type adminRepository struct {
	db *bun.DB
}

func NewAdminRepository() AdminRepository {
	return &adminRepository{
		db: db.GetConn(),
	}
}

//!Code Your Function Here
