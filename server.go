package main

import (
	"net/http"
	"html/template"
	"fmt"
)

func topAction(w http.ResponseWriter, r *http.Request) {
	user, err := currentUser(r)
	if err != nil {
		http.Redirect(w, r, "/signin", 301)
	} else {
		t := template.Must(template.ParseFiles("top.html"))
		t.Execute(w, user)
	}
}

func signoutAction(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "goblo-session")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	delete(session.Values, "uid")
	session.Save(r, w)

	http.Redirect(w, r, "/", 301)
}

func signupAction(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()

		err := createUser(r.Form["username"][0], r.Form["password"][0])
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		//TODO redirect or compile view
		fmt.Println("insert:")

	} else {
		t := template.Must(template.ParseFiles("signup.html"))
		t.Execute(w, nil)
	}
}
func signinAction(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()

		user, err := findUserByCredential(r.Form["username"][0], r.Form["password"][0])

		if err != nil {
			//TODO handle
			panic(err)
		}

		session, err := store.Get(r, "goblo-session")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		session.Values["uid"] = user.Id
		session.Save(r, w)
		http.Redirect(w, r, "/", 301)

	} else {
		t := template.Must(template.ParseFiles("signin.html"))
		t.Execute(w, nil)
	}
}

func main() {
	http.HandleFunc("/", topAction)
	http.HandleFunc("/signup", signupAction)
	http.HandleFunc("/signin", signinAction)
	http.HandleFunc("/signout", signoutAction)
	http.ListenAndServe(":3000", nil)
}
