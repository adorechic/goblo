package controllers

import (
	"github.com/adorechic/goblo/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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

	validation_errors, err := page.ValidationErrors()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if len(validation_errors) != 0 {
		pages := []models.Page{*page}
		//TODO Show all errors
		o := ViewObject{CurrentUser: user, Pages: &pages, Error: validation_errors[0]}
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

	page, err := models.FindPageByTitle(vars["title"])
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

	page, err := models.FindPageByTitle(vars["title"])
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

func UpdatePage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", 405)
		return
	}

	user, err := currentUser(r)
	if err != nil {
		http.Redirect(w, r, "/signin", 301)
		return
	}

	vars := mux.Vars(r)

	page_id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	page, err := models.FindPage(page_id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if page == nil {
		http.Redirect(w, r, "/newpage?title="+vars["title"], 301)
		return
	}

	r.ParseForm()

	page.Title = r.Form["title"][0]
	page.Body = r.Form["body"][0]

	validation_errors, err := page.ValidationErrors()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if len(validation_errors) != 0 {
		pages := []models.Page{*page}
		//TODO Show all errors
		o := ViewObject{CurrentUser: user, Pages: &pages, Error: validation_errors[0]}
		render(w, "edit_pages", o)
		return
	}

	err = page.Update()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	setFlash(w, r, "Page has updated.")
	http.Redirect(w, r, "/pages/"+page.Title, 301)
	return
}

func DeletePage(w http.ResponseWriter, r *http.Request) {
	//TODO CSRF
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", 405)
		return
	}

	_, err := currentUser(r)
	if err != nil {
		http.Redirect(w, r, "/signin", 301)
		return
	}

	vars := mux.Vars(r)

	page_id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	page, err := models.FindPage(page_id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if page == nil {
		http.Redirect(w, r, "/newpage?title="+vars["title"], 301)
		return
	}

	err = page.Delete()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	setFlash(w, r, "Page has removed.")
	http.Redirect(w, r, "/pages", 301)
	return
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
