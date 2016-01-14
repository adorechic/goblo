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

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}
