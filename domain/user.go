package domain

import (
	"context"

	"github.com/khairulharu/restapi/dto"
)

type User struct {
	ID       int    `db:"id"`
	FullName string `db:"full_name"`
	Phone    string `db:"phone"`
	Username string `db:"username"`
	Password string `db:"password"`
}

type UserRepository interface {
	FindByUsername(ctx context.Context, username string) (User, error)
	FindAll(ctx context.Context) ([]User, error)
	Insert(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) (dto.UserData, error)
	Delete(ctx context.Context, id int64) error
}

type UserService interface {
	Deleting(ctx context.Context, id int64) dto.ApiResponse
	GetAll(ctx context.Context) []dto.UserData
	Save(ctx context.Context, user dto.UserData) dto.ApiResponse
	FindUserByUsername(ctx context.Context, user dto.UserData) dto.ApiResponse
	SaveUpdate(ctx context.Context, user dto.UserData) dto.ApiResponse
}
