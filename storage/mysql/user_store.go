package mysql

import (
	"database/sql"
	"errors"
	"game/entity"
)

func (db *DB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	user, err := db.fetchUser(`SELECT * FROM users WHERE phone_number = ?`, phoneNumber)
	if err != nil {
		return false, err
	}
	return user == nil, nil
}

func (db *DB) FindByUsernameOrPhoneNumber(username, phoneNumber string) (*entity.User, error) {
	return db.fetchUser(`SELECT * FROM users WHERE username = ? OR phone_number = ? LIMIT 1`, username, phoneNumber)
}

func (db *DB) FindById(id uint) (*entity.User, error) {
	return db.fetchUser(`SELECT * FROM users WHERE id = ?`, id)
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

func (db *DB) fetchUser(query string, args ...any) (*entity.User, error) {
	row := db.connection.QueryRow(query, args...)
	user, err := scanUser(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func scanUser(row *sql.Row) (*entity.User, error) {
	user := &entity.User{}
	err := row.Scan(&user.Id, &user.PhoneNumber, &user.Username, &user.Password, &user.TotalScore, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}
