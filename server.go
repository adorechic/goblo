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

type Page struct {
	Title string
}

type Users struct {
	Id int64
	Name string
	Password string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

var store = sessions.NewCookieStore([]byte("goblo-session"))

func topAction(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "POST" {
		fmt.Println("body", r.Form)
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	} else {
		p := Page{"title"}
		t := template.Must(template.ParseFiles("top.html"))
		t.Execute(w, p)
	}
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

		db, err := genmai.New(&genmai.SQLite3Dialect{}, "./development.db")
		if err != nil {
			panic(err)
		}
		defer db.Close()

		var users []Users

		err = db.Select(&users,
			db.Where(
				"name", "=", r.Form["username"][0]).And(
					db.Where("password", "=", r.Form["password"][0])))
		if err != nil {
			panic(err)
		}

		if len(users) == 1 {
			session, err := store.Get(r, "goblo-session")
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}

			session.Values["uid"] = users[0].Id
			session.Save(r, w)
			http.Redirect(w, r, "/", 301)
		} else {
			fmt.Println("not found")
		}

	} else {
		t := template.Must(template.ParseFiles("signin.html"))
		t.Execute(w, nil)
	}
}

func main() {
	http.HandleFunc("/", topAction)
	http.HandleFunc("/signup", signupAction)
	http.HandleFunc("/signin", signinAction)
	http.ListenAndServe(":3000", nil)
}
