package controllers

import (
	"net/http"
	"github.com/adorechic/goblo/models"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		o := ViewObject{}
		render(w, "signup", o)
		return
	}

	r.ParseForm()

	err := models.CreateUser(r.Form["username"][0], r.Form["password"][0])
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	setFlash(w, r, "Account has created.")
	http.Redirect(w, r, "/signin", 301)
	return
}

func Signin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		messages, err := flashMessages(r)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		o := ViewObject{}
		if len(messages) > 0 {
			o.Error = messages[0].(string)
		}
		render(w, "signin", o)
		return
	}

	r.ParseForm()

	err := signin(w, r, r.Form["username"][0], r.Form["password"][0])

	if err != nil {
		o := ViewObject{Error: "Invalid credentials."}
		render(w, "signin", o)
		return
	}

	http.Redirect(w, r, "/", 301)
}

func Signout(w http.ResponseWriter, r *http.Request) {
	err := clearSession(w, r)
	if err == nil {
		http.Redirect(w, r, "/", 301)
	} else {
		http.Error(w, err.Error(), 500)
	}
}
