package controllers

import (
	"net/http"
)

func Top(w http.ResponseWriter, r *http.Request) {
	_, err := currentUser(r)
	if err != nil {
		http.Redirect(w, r, "/signin", 301)
		return
	} else {
		http.Redirect(w, r, "/pages", 301)
	}
}
