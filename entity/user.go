package entity

import "time"

type User struct {
	Id          uint
	Username    string
	Password    string
	PhoneNumber string
	TotalScore  int
	CreatedAt   time.Time
}
