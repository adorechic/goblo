package controllers

import (
	"net/http"
)

func Top(w http.ResponseWriter, r *http.Request) {
	user, err := currentUser(r)
	if err != nil {
		http.Redirect(w, r, "/signin", 301)
		return
	}

	o := ViewObject{CurrentUser: user}
	render(w, "top", o)
}
