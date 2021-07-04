package main

import (
	"github.com/gorilla/mux"

	"github.com/ianmuhia/lenslocked.com/views"

	"net/http"
)

var (
	homeView    *views.View
	contactView *views.View
	signupView  *views.View
)

func home(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
	// err := homeView.Template.ExecuteTemplate(w, homeView.Layout, nil)
	// if err != nil {
	// 	panic(err)
	// }
}

func contact(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(contactView.Render(w, nil))
	// err := contactView.Template.ExecuteTemplate(w, contactView.Layout, nil)
	// if err != nil {
	// 	panic(err)
	// }
}
func signup(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(signupView.Render(w, nil))
	// err := contactView.Template.ExecuteTemplate(w, contactView.Layout, nil)
	// if err != nil {
	// 	panic(err)
	// }
}

func main() {
	homeView = views.NewView("bootstrap", "views/home.html")
	contactView = views.NewView("bootstrap", "views/contact.html")
	signupView = views.NewView("bootstrap", "views/signup.html")

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/signup", signup)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}

}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
