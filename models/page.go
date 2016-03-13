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

func CreatePage(title, body string) error {
	db, err := connect()
	if err != nil {
		return err
	}
	defer db.Close()

	t := time.Now()

	page := &Page{
		Title: title,
		Body: body,
		CreatedAt: &t,
		UpdatedAt: &t,
	}
	_, err = db.Insert(page)
	if err != nil {
		return err
	}

	return nil
}
