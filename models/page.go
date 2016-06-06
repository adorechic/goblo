package models

import (
	"time"
)

type Page struct {
	Id        int64 `db:"pk"`
	Title     string
	Body      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (p *Page) TableName() string {
	return "pages"
}

func (p *Page) Create() error {
	db, err := connect()
	if err != nil {
		return err
	}
	defer db.Close()

	t := time.Now()

	p.CreatedAt = &t
	p.UpdatedAt = &t

	_, err = db.Insert(p)
	if err != nil {
		return err
	}

	return nil
}

func (p *Page) Update() error {
	db, err := connect()
	if err != nil {
		return err
	}
	defer db.Close()

	t := time.Now()

	p.UpdatedAt = &t

	_, err = db.Update(p)
	if err != nil {
		return err
	}

	return nil
}

func FindPage(id int) (*Page, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var pages []Page

	err = db.Select(&pages, db.Where("id", "=", id))
	if err != nil {
		return nil, err
	}

	if len(pages) == 0 {
		return nil, nil
	}

	return &pages[0], nil
}

func FindPageByTitle(title string) (*Page, error) {
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

	if len(pages) == 0 {
		return nil, nil
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
