package domain

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel
	ID             int64     `bun:"id,pk,autoincrement" json:"id"`
	Email          string    `bun:"email" json:"email"`
	Password       string    `bun:"password" json:"password"`
	Name           string    `bun:"name" json:"name"`
	Profile        *string   `bun:"profile" json:"profile"`
	ProfilePicture *string   `bun:"profile_picture" json:"profile_picture"`
	LifePoint      int32     `bun:"life_point" json:"life_point"`
	Book           *[]Book   `bun:"rel:has-many,join:id=user_id" json:"books"`
	UpdatedAt      time.Time `bun:"updated_at,nullzero" json:"updated_at"`
	CreatedAt      time.Time `bun:"created_at,nullzero" json:"created_at"`
}
