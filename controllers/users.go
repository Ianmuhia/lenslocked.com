package controllers

import (
	"fmt"
	"net/http"

	"github.com/ianmuhia/lenslocked.com/models"
	"github.com/ianmuhia/lenslocked.com/views"
)

type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
	Name     string `schema:"name"`
}
type Users struct {
	NewView *views.View
	us      *models.UserService
}

func NewUsers(us *models.UserService) *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "auth/signup"),
		us:      us,
	}

}

/**
 * GET /signup
 */
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	err := u.NewView.Render(w, nil)
	if err != nil {
		return
	}
}

// Create /**
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {

	var form SignupForm
	if err := parseForm(r, &form); err != nil {

		panic(err)
	}
	user := models.User{
		Name:  form.Name,
		Email: form.Email,
	}
	if err := u.us.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = fmt.Fprintln(w, form)

}
