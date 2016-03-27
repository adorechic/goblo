package controllers

import (
	"net/http"
	"regexp"
)

var url_pattern = regexp.MustCompile(`\A/pages/*`)

func Top(w http.ResponseWriter, r *http.Request) {
	user, err := currentUser(r)
	if err != nil {
		http.Redirect(w, r, "/signin", 301)
		return
	}

	if url_pattern.MatchString(r.URL.Path) {
		ShowPage(w, r)
	} else {
		o := ViewObject{CurrentUser: user}
		render(w, "top", o)
	}
}
