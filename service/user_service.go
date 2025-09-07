package service

import (
	"game/api/request"
	"game/api/response"
)

type UserService interface {
	register(request request.RegisterUserRequest) (response.RegisterUserResponse, error)
	login(request request.LoginUserDto) error
}
