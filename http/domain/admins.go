package domain

import (
	"time"

	"github.com/uptrace/bun"
)

type Admin struct {
	bun.BaseModel
	ID        int64     `bun:"id,pk,autoincrement" json:"id"`
	Email     string    `bun:"email" json:"email"`
	Password  string    `bun:"password" json:"password"`
	Name      string    `bun:"name" json:"name"`
	UpdatedAt time.Time `bun:"updated_at,nullzero" json:"updated_at"`
	CreatedAt time.Time `bun:"created_at,nullzero" json:"created_at"`
}
