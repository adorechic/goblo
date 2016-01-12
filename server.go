package main

import (
	"net/http"
	"html/template"
)

type Page struct {
	Title string
}

func handler(w http.ResponseWriter, r *http.Request) {
	p := Page{"title"}
	t := template.Must(template.ParseFiles("top.html"))
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}
