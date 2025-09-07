package mysql

import (
	"database/sql"
	"errors"
	"game/entity"
	"time"
)

func (db *DB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	user := entity.User{}
	var createdAt time.Time
	row := db.connection.QueryRow(`select * from users where phone_number = ?`, phoneNumber)
	err := row.Scan(&user.Id, &user.PhoneNumber, &user.Username, &user.Password, &user.TotalScore, &createdAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func (db *DB) Save(user entity.User) (entity.User, error) {
	exec, err := db.connection.Exec(`insert into users (username, phone_number, password) values (?, ?, ?)`,
		user.Username, user.PhoneNumber, user.Password)
	if err != nil {
		return entity.User{}, err
	}
	id, _ := exec.LastInsertId()
	user.Id = uint(id)
	return user, nil
}

func (db *DB) FindByUsernameOrPhoneNumber(username, phoneNumber string) (*entity.User, error) {
	user := entity.User{}
	var createdAt time.Time
	row := db.connection.QueryRow(`select * from users where username = ? or phone_number = ?`, username, phoneNumber)
	err := row.Scan(&user.Id, &user.PhoneNumber, &user.Username, &user.Password, &user.TotalScore, &createdAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &entity.User{}, nil
		}
	}
	return &user, nil
}
