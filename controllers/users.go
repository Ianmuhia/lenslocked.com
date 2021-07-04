package controllers

import (
	"fmt"
	"net/http"

	"github.com/ianmuhia/lenslocked.com/views"
)

type Users struct {
	NewView *views.View
}

func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/auth/signup.html"),
	}

}

/**
 * GET /signup
 */
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	u.NewView.Render(w, nil)
}

/**
 * creates new user account
 * POST /signup
 */
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "this is a fake message pretend that we have created")
}
