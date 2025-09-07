package repository

import "game/entity"

type UserRepository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	Save(user entity.User) (entity.User, error)
	FindByUsernameOrPhoneNumber(username, phoneNumber string) (*entity.User, error)
	FindById(id uint) (*entity.User, error)
}
