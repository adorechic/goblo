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

func FindPage(title string) (*Page, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var pages []Page

	err = db.Select(&pages, db.Where("title", "=", title))
	if err != nil {
		return nil, err
	}

	return &pages[0], nil
}

func AllPage() (*[]Page, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var pages []Page

	err = db.Select(&pages)
	if err != nil {
		return nil, err
	}

	return &pages, nil
}
