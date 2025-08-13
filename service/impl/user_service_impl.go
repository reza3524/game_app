package service

import (
	"errors"
	"game/api/request"
	"game/api/response"
	"game/entity"
	"game/repository"
	"game/utility"
)

type UserServiceImpl struct {
	repository repository.UserRepository
}

func NewUser(repository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{repository: repository}
}

func (u *UserServiceImpl) Register(request request.RegisterUserRequest) (response.RegisterUserResponse, error) {

	if !utility.IsPhoneNumberValid(request.PhoneNumber) {
		return response.RegisterUserResponse{}, errors.New("phone number is invalid")
	}
	if isUnique, err := u.repository.IsPhoneNumberUnique(request.PhoneNumber); err != nil || !isUnique {
		if err != nil {
			return response.RegisterUserResponse{}, err
		}
		return response.RegisterUserResponse{}, errors.New("phone number is not unique")
	}

	user := entity.User{Username: request.Username, PhoneNumber: request.PhoneNumber}
	user, err := u.repository.Save(user)
	if err != nil {
		return response.RegisterUserResponse{}, err
	}

	return response.RegisterUserResponse{Id: user.Id}, nil
}
