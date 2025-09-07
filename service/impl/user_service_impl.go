package service

import (
	"errors"
	"game/api/request"
	"game/api/response"
	"game/entity"
	"game/repository"
	"game/utility"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	repository repository.UserRepository
}

func NewUser(repository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{repository: repository}
}

func (u *UserServiceImpl) Register(request request.RegisterUserRequest) (response.RegisterUserResponse, error) {
	if len(request.Username) < 3 {
		return response.RegisterUserResponse{}, errors.New("username should be at least 3 characters long")
	}

	//TODO check the password with regex
	if request.Password == "" {
		return response.RegisterUserResponse{}, errors.New("password is empty")
	}

	if !utility.IsPhoneNumberValid(request.PhoneNumber) {
		return response.RegisterUserResponse{}, errors.New("phone number is invalid")
	}

	if isUnique, err := u.repository.IsPhoneNumberUnique(request.PhoneNumber); err != nil || !isUnique {
		if err != nil {
			return response.RegisterUserResponse{}, err
		}
		return response.RegisterUserResponse{}, errors.New("phone number is not unique")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return response.RegisterUserResponse{}, err
	}

	user := entity.User{
		Username:    request.Username,
		PhoneNumber: request.PhoneNumber,
		Password:    string(hashPassword),
	}
	user, err = u.repository.Save(user)
	if err != nil {
		return response.RegisterUserResponse{}, err
	}

	return response.RegisterUserResponse{Id: user.Id}, nil
}

func (u *UserServiceImpl) Login(request request.LoginUserDto) error {
	existUser, err := u.repository.FindByUsernameOrPhoneNumber(request.Username, request.PhoneNumber)
	if err != nil {
		return err
	}
	if existUser == nil {
		return errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existUser.Password), []byte(request.Password)); err != nil {
		return errors.New("username,phoneNumber or password is wrong")
	}
	return nil
}
