package controllers

import (
	"github.com/adorechic/goblo/models"
	"github.com/gorilla/mux"
	"net/http"
)

func NewPage(w http.ResponseWriter, r *http.Request) {
	user, err := currentUser(r)
	if err != nil {
		http.Redirect(w, r, "/signin", 301)
		return
	}

	if r.Method != "POST" {
		o := ViewObject{CurrentUser: user}
		render(w, "new_pages", o)
		return
	}

	r.ParseForm()

	title, body := r.Form["title"][0], r.Form["body"][0]
	err = models.CreatePage(title, body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	setFlash(w, r, "Page has created.")
	http.Redirect(w, r, "/pages/"+title, 301)
	return
}

func ShowPage(w http.ResponseWriter, r *http.Request) {
	user, err := currentUser(r)
	if err != nil {
		http.Redirect(w, r, "/signin", 301)
		return
	}

	vars := mux.Vars(r)

	page, err := models.FindPage(vars["title"])
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	pages := []models.Page{*page}
	o := ViewObject{CurrentUser: user, Pages: &pages}
	render(w, "pages", o)
}

func EditPage(w http.ResponseWriter, r *http.Request) {
	user, err := currentUser(r)
	if err != nil {
		http.Redirect(w, r, "/signin", 301)
		return
	}

	vars := mux.Vars(r)

	page, err := models.FindPage(vars["title"])
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	pages := []models.Page{*page}
	o := ViewObject{CurrentUser: user, Pages: &pages}
	render(w, "edit_pages", o)
}

func IndexPage(w http.ResponseWriter, r *http.Request) {
	user, err := currentUser(r)
	if err != nil {
		http.Redirect(w, r, "/signin", 301)
		return
	}

	pages, err := models.AllPage()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	o := ViewObject{CurrentUser: user, Pages: pages}
	render(w, "page_index", o)
}
