package main

import (
	"github.com/gin-gonic/gin"
	// "github.com/gorilla/mux"
	"github.com/ianmuhia/lenslocked.com/views"

	"net/http"
)

var (
	homeView    *views.View
	contactView *views.View
)

func home(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := homeView.Template.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := contactView.Template.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

// func main() {
// 	router := gin.Default()
// 	router.LoadHTMLGlob("templates/*")
// 	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
// 	router.GET("/index", func(c *gin.Context) {
// 		c.HTML(http.StatusOK, "test.gohtml", gin.H{
// 			"title": "Main website",
// 		})
// 	})
// 	router.Run(":8080")
// }

func main() {
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")
	//handlers := BaseHandlers()
	handlers := BaseHandlers()

	mux := mux.NewRouter()

	mux.HandleFunc("/", handlers.Landing)
	// mux.HandleFunc("/signup", handlers.signup)

	// router.HandleFunc("/contact", contact)
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		panic(err)
	}

}
