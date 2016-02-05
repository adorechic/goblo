package main

import (
	"net/http"
	"html/template"
	"fmt"
	"time"
	_ "github.com/mattn/go-sqlite3"
	"github.com/naoina/genmai"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("goblo-session"))

func topAction(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "goblo-session")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	uid := session.Values["uid"]

	if uid == nil {
		http.Redirect(w, r, "/signin", 301)
	} else {
		user, err := findUser(int(uid.(int64)))
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

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

		db, err := genmai.New(&genmai.SQLite3Dialect{}, "./development.db")
		if err != nil {
			panic(err)
		}
		defer db.Close()

		t := time.Now()

		obj := &Users{
			Name: r.Form["username"][0],
		  Password: r.Form["password"][0],
			CreatedAt: &t,
			UpdatedAt: &t,
		}
		n, err := db.Insert(obj)
		if err != nil {
			panic(err)
		}
		fmt.Println("insert:", n)
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
