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
	err := clearSession(w, r)
	if err == nil {
		http.Redirect(w, r, "/", 301)
	} else {
		http.Error(w, err.Error(), 500)
	}
}

func signupAction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		t := template.Must(template.ParseFiles("signup.html"))
		t.Execute(w, nil)
		return
	}

	r.ParseForm()

	err := createUser(r.Form["username"][0], r.Form["password"][0])
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//TODO redirect or compile view
	fmt.Println("insert:")
}
func signinAction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		t := template.Must(template.ParseFiles("signin.html"))
		t.Execute(w, nil)
		return
	}

	r.ParseForm()

	user, err := findUserByCredential(r.Form["username"][0], r.Form["password"][0])

	if err != nil {
		t := template.Must(template.ParseFiles("signin.html"))
		t.Execute(w, "Invalid credentials.")
		return
	}

	err = createSession(user, w, r)
	if err != nil {
		t := template.Must(template.ParseFiles("signin.html"))
		t.Execute(w, "Invalid credentials.")
		return
	}

	http.Redirect(w, r, "/", 301)
}

func main() {
	http.HandleFunc("/", topAction)
	http.HandleFunc("/signup", signupAction)
	http.HandleFunc("/signin", signinAction)
	http.HandleFunc("/signout", signoutAction)
	http.ListenAndServe(":3000", nil)
}
