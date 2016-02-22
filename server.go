package main

import (
	"net/http"
	"html/template"
)

func render(w http.ResponseWriter, name string, data interface{}) {
	t := template.Must(template.ParseFiles("views/layout.html", "views/" + name + ".html"))
	t.Execute(w, data)
}

func topAction(w http.ResponseWriter, r *http.Request) {
	user, err := currentUser(r)
	if err != nil {
		http.Redirect(w, r, "/signin", 301)
	} else {
		render(w, "top", user)
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
		render(w, "signup", nil)
		return
	}

	r.ParseForm()

	err := createUser(r.Form["username"][0], r.Form["password"][0])
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	setFlash(w, r, "Account has created.")
	http.Redirect(w, r, "/signin", 301)
	return

}
func signinAction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		messages, err := flashMessages(r)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		if len(messages) > 0 {
			render(w, "signin", messages[0])
		} else {
			render(w, "signin", nil)
		}
		return
	}

	r.ParseForm()

	err := signin(w, r, r.Form["username"][0], r.Form["password"][0])

	if err != nil {
		render(w, "signin", "Invalid credentials.")
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
