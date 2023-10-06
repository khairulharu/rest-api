package service

import (
	"context"

	"github.com/khairulharu/restapi/domain"
	"github.com/khairulharu/restapi/dto"
)

type userService struct {
	userRepository domain.UserRepository
}

func NewUser(userRepository domain.UserRepository) domain.UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (u userService) Save(ctx context.Context, user dto.UserData) dto.ApiResponse {
	var userData domain.User
	userData.ID = user.ID
	userData.FullName = user.FullName
	userData.Phone = user.Phone
	userData.Username = user.Username
	userData.Password = user.Password

	exist, _ := u.userRepository.FindByUsername(ctx, user.Username)
	if exist != (domain.User{}) {
		return dto.ApiResponse{
			Code:    "400",
			Message: "cannot using username",
		}
	}
	err := u.userRepository.Insert(ctx, &userData)
	if err != nil {
		return dto.ApiResponse{
			Code:    "911",
			Message: "insert database error",
			Error:   err.Error(),
		}
	}

	return dto.ApiResponse{
		Code:    "00",
		Message: "approve",
	}
}

func (u userService) FindUserByUsername(ctx context.Context, user dto.UserData) dto.ApiResponse {
	data, err := u.userRepository.FindByUsername(ctx, user.Username)

	if data == (domain.User{}) {
		return dto.ApiResponse{
			Code:    "400",
			Message: "username not found",
		}
	}

	if err != nil {
		return dto.ApiResponse{
			Code:    "400",
			Message: "username not found",
			Error:   err.Error(),
		}
	}

	var userRes dto.UserData
	userRes.ID = data.ID
	userRes.FullName = data.FullName
	userRes.Phone = data.Phone
	userRes.Username = data.Username
	userRes.Password = data.Password

	return dto.ApiResponse{
		Code:    "00",
		Message: "APPROVE",
		Data:    &userRes,
	}
}

func (u userService) SaveUpdate(ctx context.Context, user dto.UserData) dto.ApiResponse {

	var userRes domain.User
	userRes.ID = user.ID
	userRes.FullName = user.FullName
	userRes.Phone = user.Phone
	userRes.Username = user.Username
	userRes.Password = user.Password

	res, err := u.userRepository.Update(ctx, &userRes)
	if err != nil {
		return dto.ApiResponse{
			Code:    "401",
			Message: "faileds update",
			Error:   err.Error(),
		}
	}
	return dto.ApiResponse{
		Code:    "00",
		Message: "APPROVE",
		Data:    res,
	}
}

func (u userService) GetAll(ctx context.Context) []dto.UserData {
	res, err := u.userRepository.FindAll(ctx)
	if err != nil {
		return []dto.UserData{}
	}

	var dataUsers []dto.UserData

	for _, v := range res {
		dataUsers = append(dataUsers, dto.UserData{
			ID:       v.ID,
			FullName: v.FullName,
			Username: v.Username,
			Phone:    v.Phone,
			Password: v.Password,
		})
	}

	return dataUsers
}

func (u userService) Deleting(ctx context.Context, id int64) dto.ApiResponse {
	err := u.userRepository.Delete(ctx, id)
	if err != nil {
		return dto.ApiResponse{
			Code:    "500",
			Message: "erroro whem delete",
			Error:   err.Error(),
		}
	}
	return dto.ApiResponse{
		Code:    "200",
		Message: "APPROVE",
	}
}
