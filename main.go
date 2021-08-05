package main

import (
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
	r.HandleFunc("/signup", userC.New).Methods("GET")
	r.HandleFunc("/signup", userC.Create).Methods("POST")

	err = http.ListenAndServe(":8080", r)
	must(err)

}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
