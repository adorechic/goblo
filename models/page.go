package models

import (
	"github.com/naoina/genmai"
	"strings"
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

func (p *Page) Delete() error {
	db, err := connect()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Delete(p)
	if err != nil {
		return err
	}

	return nil
}

func (p *Page) ValidationErrors() ([]string, error) {
	errors := []string{}

	if len(strings.TrimSpace(p.Title)) == 0 {
		errors = append(errors, "Title is required")
	} else {
		pages, err := FindPagesByTitle(p.Title)
		if err != nil {
			return errors, err
		}

		for i := 0; i < len(pages); i++ {
			if pages[i].Id != p.Id {
				errors = append(errors, "Same title article exists")
			}
		}
	}

	if len(strings.TrimSpace(p.Body)) == 0 {
		errors = append(errors, "Body is required")
	}
	return errors, nil
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
	pages, err := FindPagesByTitle(title)

	if err != nil {
		return nil, err
	}

	if len(pages) == 0 {
		return nil, nil
	}

	return &pages[0], nil
}

func FindPagesByTitle(title string) ([]Page, error) {
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

	return pages, nil
}

func AllPage() (*[]Page, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var pages []Page

	err = db.Select(&pages, db.OrderBy("updated_at", genmai.DESC))
	if err != nil {
		return nil, err
	}

	return &pages, nil
}
