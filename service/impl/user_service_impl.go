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

func (u *UserServiceImpl) Register(request request.UserRegisterRequest) (response.UserRegisterResponse, error) {
	if len(request.Username) < 3 {
		return response.UserRegisterResponse{}, errors.New("username should be at least 3 characters long")
	}

	//TODO check the password with regex
	if request.Password == "" {
		return response.UserRegisterResponse{}, errors.New("password is empty")
	}

	if !utility.IsPhoneNumberValid(request.PhoneNumber) {
		return response.UserRegisterResponse{}, errors.New("phone number is invalid")
	}

	if isUnique, err := u.repository.IsPhoneNumberUnique(request.PhoneNumber); err != nil || !isUnique {
		if err != nil {
			return response.UserRegisterResponse{}, err
		}
		return response.UserRegisterResponse{}, errors.New("phone number is not unique")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return response.UserRegisterResponse{}, err
	}

	user := entity.User{
		Username:    request.Username,
		PhoneNumber: request.PhoneNumber,
		Password:    string(hashPassword),
	}
	user, err = u.repository.Save(user)
	if err != nil {
		return response.UserRegisterResponse{}, err
	}

	return response.UserRegisterResponse{Id: user.Id}, nil
}

func (u *UserServiceImpl) Login(request request.UserLoginRequest) (response.UserLoginResponse, error) {
	existUser, err := u.repository.FindByUsernameOrPhoneNumber(request.Username, request.PhoneNumber)
	if err != nil {
		return response.UserLoginResponse{}, err
	}
	if existUser == nil {
		return response.UserLoginResponse{}, errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existUser.Password), []byte(request.Password)); err != nil {
		return response.UserLoginResponse{}, errors.New("username,phoneNumber or password is wrong")
	}

	token, err := utility.GenerateToken(existUser.Id)
	if err != nil {
		return response.UserLoginResponse{}, err
	}

	return response.UserLoginResponse{Authorization: token}, nil
}

func (u *UserServiceImpl) Profile(request request.UserProfileRequest) (response.UserProfileResponse, error) {
	existUser, err := u.repository.FindById(request.Id)
	if err != nil {
		return response.UserProfileResponse{}, err
	}
	if existUser == nil {
		return response.UserProfileResponse{}, errors.New("user not found")
	}
	return response.UserProfileResponse{Username: existUser.Username, PhoneNumber: existUser.PhoneNumber}, nil
}
