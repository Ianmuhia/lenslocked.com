package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ianmuhia/lenslocked.com/controllers"
	"github.com/ianmuhia/lenslocked.com/models"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=lenslocked_dev port=5432 sslmode=disable"
	us, err := models.NewUserService(dsn)

	must(err)
	defer us.Close()
	us.AutoMigrate()
	// us.DestructiveReset()

	staticC := controllers.NewStatic()
	userC := controllers.NewUsers(us)
	r := mux.NewRouter()

	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.Handle("/signup", userC.NewView).Methods("GET")
	r.HandleFunc("/signup", userC.Create).Methods("POST")
	r.Handle("/login", userC.LoginView).Methods("GET")
	r.HandleFunc("/login", userC.Login).Methods("POST")
	r.HandleFunc("/cookietest", userC.CookieTest).Methods("GET")

	fmt.Println("server running in port : 8080")
	err = http.ListenAndServe(":8080", r)

	must(err)

}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
