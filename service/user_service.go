package service

import (
	"game/api/request"
	"game/api/response"
)

type UserService interface {
	register(request request.UserRegisterRequest) (response.UserRegisterResponse, error)
	login(request request.UserLoginRequest) (response.UserLoginResponse, error)
	profile(request request.UserProfileRequest) (response.UserProfileResponse, error)
}
