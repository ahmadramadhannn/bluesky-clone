package user

import (
	"context"
	"time"
)

type User struct {
	ID            int       `json:"id" db:"id"`
	Username      string    `json:"username" db:"username"`
	Email         string    `json:"email" db:"email"`
	Password      string    `json:"password" db:"password"`
	Bio           string    `json:"bio" db:"bio"`
	ProfilePicUrl string    `json:"profilePicUrl" db:"profile_pic_url"`
	CoverPicUrl   string    `json:"coverPicUrl" db:"cover_pic_url"`
	CreatedAt     time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt     time.Time `json:"updatedAt" db:"updated_at"`
}

type UserRepository interface {
	GetUserByID(ctx context.Context, id int) (string, error)
}

type UserService interface {
	GetUserByID(ctx context.Context, id int) (string, error)
}
