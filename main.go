package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ianmuhia/lenslocked.com/controllers"
)

func main() {
	staticC := controllers.NewStatic()
	userC := controllers.NewUsers()
	r := mux.NewRouter()

	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.HandleFunc("/signup", userC.New).Methods("GET")
	r.HandleFunc("/signup", userC.Create).Methods("POST")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}

}
