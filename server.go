package main

import (
	"net/http"
	"github.com/adorechic/goblo/controllers"
)

func main() {
	http.HandleFunc("/", controllers.Top)
	http.HandleFunc("/signup", controllers.Signup)
	http.HandleFunc("/signin", controllers.Signin)
	http.HandleFunc("/signout", controllers.Signout)
	http.ListenAndServe(":3000", nil)
}
