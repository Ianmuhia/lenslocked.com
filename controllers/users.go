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
type LoginForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}
type Users struct {
	NewView   *views.View
	LoginView *views.View
	us        *models.UserService
}

func NewUsers(us *models.UserService) *Users {
	return &Users{
		NewView:   views.NewView("bootstrap", "auth/signup"),
		LoginView: views.NewView("bootstrap", "auth/login"),
		us:        us,
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
		Name:     form.Name,
		Email:    form.Email,
		Password: form.Password,
	}
	if err := u.us.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// _, _ = fmt.Fprintln(w, user)

}

/**
 *
 */
func (u *Users) Login(w http.ResponseWriter, r *http.Request) {
	form := LoginForm{}
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	user, err := u.us.Authenticate(form.Email, form.Password)
	if err != nil {

		switch err {
		case models.ErrInvalidPassword:
			fmt.Fprintln(w, "Invalid password provided")
		case models.ErrNotFound:
			fmt.Fprintln(w, "Invalid email address")

		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)

		}

		return
	}

	cookie := http.Cookie{
		Name:  "email",
		Value: user.Email,
	}
	http.SetCookie(w, &cookie)
	fmt.Fprintln(w, user)

	// fmt.Fprintln(w, form)
}

//Display the cookie of current user
func (u *Users) CookieTest(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("email")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
	fmt.Fprintln(w, "Email is : ", cookie.Value)
	// fmt.Fprintln(w, cookie)
}
