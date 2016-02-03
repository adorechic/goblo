package main

import (
	"time"
)

type Users struct {
	Id int64 `db:"pk"`
	Name string
	Password string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
