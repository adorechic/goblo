package main

import (
	"net/http"
	"html/template"
	"fmt"
	"time"
	_ "github.com/mattn/go-sqlite3"
	"github.com/naoina/genmai"
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

func handler(w http.ResponseWriter, r *http.Request) {
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
	} else {
		t := template.Must(template.ParseFiles("signin.html"))
		t.Execute(w, nil)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/signup", signupAction)
	http.HandleFunc("/signin", signinAction)
	http.ListenAndServe(":3000", nil)
}
