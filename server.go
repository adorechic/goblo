package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/adorechic/goblo/controllers"
	"github.com/codegangsta/negroni"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Top)
	r.HandleFunc("/signup", controllers.Signup)
	r.HandleFunc("/signin", controllers.Signin)
	r.HandleFunc("/signout", controllers.Signout)
	r.HandleFunc("/pages", controllers.IndexPage)
	r.HandleFunc("/pages/{title}", controllers.ShowPage)
	r.HandleFunc("/pages/{title}/edit", controllers.EditPage)
	r.HandleFunc("/newpage", controllers.NewPage)

	n := negroni.Classic()
	n.UseHandler(r)

	http.ListenAndServe(":3000", n)
}
