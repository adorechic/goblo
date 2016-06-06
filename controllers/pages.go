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

	r.ParseForm()

	titles, bodies := r.Form["title"], r.Form["body"]

	page := &models.Page{}
	if titles != nil {
		page.Title = titles[0]
	}
	if bodies != nil {
		page.Body = bodies[0]
	}

	if r.Method != "POST" {
		pages := []models.Page{*page}

		o := ViewObject{CurrentUser: user, Pages: &pages}
		render(w, "new_pages", o)
		return
	}

	err = page.Create()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	setFlash(w, r, "Page has created.")
	http.Redirect(w, r, "/pages/"+page.Title, 301)
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

	if page == nil {
		http.Redirect(w, r, "/newpage?title="+vars["title"], 301)
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

	if page == nil {
		http.Redirect(w, r, "/newpage?title="+vars["title"], 301)
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
