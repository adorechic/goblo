package models

import (
	"time"
	_ "github.com/mattn/go-sqlite3"
	"errors"
)

type Users struct {
	Id int64 `db:"pk"`
	Name string
	Password string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func FindUser(id int) (*Users, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var users []Users

	err = db.Select(&users, db.Where("id", "=", id))
	if err != nil {
		return nil, err
	}

	return &users[0], nil
}

func FindUserByCredential(username, password string) (*Users, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}

	defer db.Close()
	var users []Users

	err = db.Select(&users,
		db.Where("name", "=", username).And(
			db.Where("password", "=", password)))

	if err != nil {
		return nil, err
	}

	if len(users) == 1 {
		return &users[0], nil
	} else {
		return nil, errors.New("Invalid Credential")
	}
}

func CreateUser(username, password string) error {
	db, err := connect()
	if err != nil {
		return err
	}
	defer db.Close()

	t := time.Now()

	user := &Users{
		Name: username,
		Password: password,
		CreatedAt: &t,
		UpdatedAt: &t,
	}
	_, err = db.Insert(user)
	if err != nil {
		return err
	}

	return nil
}
