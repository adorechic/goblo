package main

import (
	"net/http"
	"html/template"
	"fmt"
)

type Page struct {
	Title string
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
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])

	} else {
		t := template.Must(template.ParseFiles("signup.html"))
		t.Execute(w, nil)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/signup", signupAction)
	http.ListenAndServe(":3000", nil)
}
