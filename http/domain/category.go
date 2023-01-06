package domain

import (
	"time"

	"github.com/uptrace/bun"
)

type Category struct {
	bun.BaseModel
	ID           int64     `bun:"id,pk,autoincrement" json:"id"`
	CategoryName string    `bun:"category_name" json:"category_name"`
	UpdatedAt    time.Time `bun:"updated_at,nullzero" json:"updated_at"`
	CreatedAt    time.Time `bun:"created_at,nullzero" json:"created_at"`
}
