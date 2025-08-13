package repository

import "game/entity"

type UserRepository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	Save(user entity.User) (entity.User, error)
}
