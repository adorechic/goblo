package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/adorechic/goblo/controllers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Top)
	r.HandleFunc("/signup", controllers.Signup)
	r.HandleFunc("/signin", controllers.Signin)
	r.HandleFunc("/signout", controllers.Signout)
	r.HandleFunc("/pages", controllers.IndexPage)
	r.HandleFunc("/pages/{title}", controllers.ShowPage)
	r.HandleFunc("/newpage", controllers.NewPage)
	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)
}
