package controllers

import (
	"net/http"
)

func NewPage(w http.ResponseWriter, r *http.Request) {
	user, err := currentUser(r)
	if err != nil {
		http.Redirect(w, r, "/signin", 301)
	} else {
		o := ViewObject{CurrentUser: user}
		render(w, "pages", o)
	}
}
