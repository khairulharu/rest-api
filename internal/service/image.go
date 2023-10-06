package service

import (
	"context"

	"github.com/khairulharu/restapi/domain"
	"github.com/khairulharu/restapi/dto"
)

type imageService struct {
	imageRepository domain.ImageRepository
}

func NewImage(imageRepository domain.ImageRepository) domain.ImageService {
	return &imageService{
		imageRepository: imageRepository,
	}
}

func (i imageService) GetAll(ctx context.Context) dto.ApiResponse {
	res, err := i.imageRepository.All(ctx)
	if err != nil {
		return dto.ApiResponse{
			Code:    "500",
			Message: "err when get images",
			Error:   err.Error(),
		}
	}
	var imageRespon []dto.ImageData

	for _, v := range res {
		imageRespon = append(imageRespon, dto.ImageData{
			ID:     v.ID,
			UserId: v.UserId,
			Image:  v.Images,
		})
	}

	return dto.ApiResponse{
		Code:    "200",
		Message: "APPROVE",
		Data:    imageRespon,
	}

}

func (i imageService) SaveData(ctx context.Context, image dto.ImageData) dto.ApiResponse {
	images := domain.Image{
		ID:     image.ID,
		UserId: image.UserId,
		Images: image.Image,
	}

	if images == (domain.Image{}) {
		return dto.ApiResponse{
			Code:    "400",
			Message: "err bad input",
		}
	}

	if err := i.imageRepository.Insert(ctx, &images); err != nil {
		return dto.ApiResponse{
			Code:    "5000",
			Message: "erorrr",
			Error:   err.Error(),
		}
	}

	return dto.ApiResponse{
		Code:    "200",
		Message: "APPROVE",
	}
}

func (i imageService) Save(ctx context.Context, id int64, userId int64) dto.ApiResponse {
	user, err := i.imageRepository.FindByUserId(ctx, userId)
	if err != nil {
		return dto.ApiResponse{
			Code:    "400",
			Message: "error when get user for delete",
			Error:   err.Error(),
		}
	}
	if user.ID != id {
		return dto.ApiResponse{
			Code:    "400",
			Message: "cannot delete, not your image",
		}
	}
	err = i.imageRepository.Delete(ctx, id)
	if err != nil {
		return dto.ApiResponse{
			Code:    "500",
			Message: "erorr when Delete",
			Error:   err.Error(),
		}
	}
	return dto.ApiResponse{
		Code:    "200",
		Message: "DELETE SUCCES",
	}
}
