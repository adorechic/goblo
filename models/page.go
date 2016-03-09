package models

import (
	"time"
)

type Page struct {
	Id int64 `db:"pk"`
	Title string
	Body string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (p *Page) TableName() string {
	return "pages"
}
