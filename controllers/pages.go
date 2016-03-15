package controllers

import (
	"net/http"
	"github.com/adorechic/goblo/models"
)

func NewPage(w http.ResponseWriter, r *http.Request) {
	user, err := currentUser(r)
	if err != nil {
		http.Redirect(w, r, "/signin", 301)
		return
	}

	if r.Method != "POST" {
		o := ViewObject{CurrentUser: user}
		render(w, "pages", o)
		return
	}

	title, body := r.Form["title"][0], r.Form["body"][0]
	err = models.CreatePage(title, body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	setFlash(w, r, "Page has created.")
	http.Redirect(w, r, "/page/" + title, 301)
	return
}

func ShowPage(w http.ResponseWriter, r *http.Request) {
	user, err := currentUser(r)
	if err != nil {
		http.Redirect(w, r, "/signin", 301)
		return
	}

	page, err := models.FindPage(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	page, err := models.FindPage(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	o := ViewObject{CurrentUser: user, Page: page}
	render(w, "pages", o)
}
