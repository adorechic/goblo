package main

import (
	"time"
	"github.com/naoina/genmai"
)

type Users struct {
	Id int64 `db:"pk"`
	Name string
	Password string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func findUser(id int) (*Users, error) {
	db, err := genmai.New(&genmai.SQLite3Dialect{}, "./development.db")
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
