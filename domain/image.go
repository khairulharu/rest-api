package domain

import (
	"context"

	"github.com/khairulharu/restapi/dto"
)

type Image struct {
	ID     int64  `db:"id"`
	UserId int64  `db:"users_id"`
	Images string `db:"images"`
}

type ImageRepository interface {
	All(ctx context.Context) ([]Image, error)
	FindByUserId(ctx context.Context, UserId int64) (Image, error)
	Insert(ctx context.Context, image *Image) error
	Delete(ctx context.Context, id int64) error
}

type ImageService interface {
	GetAll(ctx context.Context) dto.ApiResponse
	Save(ctx context.Context, id int64, userId int64) dto.ApiResponse
	SaveData(ctx context.Context, image dto.ImageData) dto.ApiResponse
}
