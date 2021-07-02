package main

import (
	"github.com/gorilla/mux"

	"github.com/ianmuhia/lenslocked.com/views"

	"net/http"
)

var (
	homeView    *views.View
	contactView *views.View
)

func home(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := homeView.Template.ExecuteTemplate(w, homeView.Layout, nil)
	if err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := contactView.Template.ExecuteTemplate(w, contactView.Layout, nil)
	if err != nil {
		panic(err)
	}
}

func main() {
	homeView = views.NewView("bootstrap", "views/home.html")
	contactView = views.NewView("bootstrap", "views/contact.html")

	r := mux.NewRouter()
	r.HandleFunc("/", home)

	r.HandleFunc("/contact", contact)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}

}
