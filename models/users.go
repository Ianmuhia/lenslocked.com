package models

import (
	"errors"

	"github.com/ianmuhia/lenslocked.com/hash"
	"github.com/ianmuhia/lenslocked.com/rand"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm/logger"

	"gorm.io/gorm"
)

var (
	/**
	 * ErrNotFound is returned when a reasouce is not found
	 */
	ErrNotFound        = errors.New("models: reasource not found")
	ErrInvalidID       = errors.New("models: ID provided was invalid")
	ErrInvalidEmail    = errors.New("models:  invalid email was provided ")
	ErrInvalidPassword = errors.New("models: incorrect password provided")
)

const userPwPepper = "secret-random-string"
const hmacSecretKey = "secret-hmac-key"

type UserService struct {
	db   *gorm.DB
	hmac hash.HMAC
}
type User struct {
	gorm.Model
	Name         string
	Email        string `gorm:"not null;uniqueIndex"`
	Password     string `gorm:"-"`
	PasswordHash string `gorm:"not null"`
	Remember     string `gorm:"-"`
	RememberHash string `gorm:"not null;uniqueIndex"`
}

func NewUserService(connectionInfo string) (*UserService, error) {

	// dsn := "host=localhost user=postgres password=postgres dbname=lenslocked_dev port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(connectionInfo), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	hmac := hash.NewHMAC(hmacSecretKey)
	return &UserService{
		db:   db,
		hmac: hmac,
	}, nil

}

/**ByID will lookup a user by the id provided
* 1 - user, nil
* 2 - nil, ErrNotFound
* 3 - nil, otherError
 */
func (us *UserService) ByID(id uint) (*User, error) {
	var user User
	db := us.db.Where("id = ?", id)
	// err := db.First(&user).Error
	err := first(db, &user)
	return &user, err

}
func (us *UserService) ByEmail(email string) (*User, error) {
	var user User
	db := us.db.Where("email = ?", email)
	// err := db.First(&user).Error
	err := first(db, &user)
	return &user, err

}
func (us *UserService) Delete(id uint) error {
	if id == 0 {
		return ErrInvalidID
	}
	user := User{Model: gorm.Model{ID: id}}
	return us.db.Delete(&user).Error
	// var user User
	// db := us.db.Where("email = ?", email)
	// err := db.First(&user).Error
	// err := first(db, &user)
	// return &user, err

}

func (us *UserService) ByRemember(token string) (*User, error) {
	rememberHash := us.hmac.Hash(token)
	var user User
	err := first(us.db.Where("remember_hash = ?", rememberHash), &user)
	if err != nil {
		return nil, err
	}
	return &user, nil

}

func first(db *gorm.DB, user *User) error {
	err := db.First(user).Error
	if err == gorm.ErrRecordNotFound {
		return ErrNotFound
	}
	return err

}

//Create the provided user using the data provided returns an erorr is any
func (us *UserService) Create(user *User) error {
	pwBytes := []byte(user.Password + userPwPepper)
	hashedBytes, err := bcrypt.GenerateFromPassword(pwBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashedBytes)
	user.Password = ""
	if user.Remember == "" {
		token, err := rand.RememberToken()
		if err != nil {
			return err
		}
		user.Remember = token
	}
	user.RememberHash = us.hmac.Hash(user.Remember)

	return us.db.Create(user).Error
}
func (us *UserService) Update(user *User) error {
	if user.Remember != "" {
		user.RememberHash = us.hmac.Hash(user.Remember)

	}
	// return nil
	return us.db.Save(user).Error
}

func (us *UserService) Authenticate(email string, password string) (*User, error) {
	foundUser, err := us.ByEmail(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.PasswordHash), []byte(password+userPwPepper))
	if err != nil {
		switch err {
		case bcrypt.ErrMismatchedHashAndPassword:
			return nil, ErrInvalidPassword
		case nil:
		default:
			return nil, err

		}
	}

	return foundUser, nil

}

func (us *UserService) DestructiveReset() error {

	if err := us.db.Migrator().DropTable(&User{}); err != nil {
		return err
	}
	return us.AutoMigrate()

}

func (us *UserService) AutoMigrate() error {
	if err := us.db.Migrator().AutoMigrate(&User{}); err != nil {
		return err
	}
	return nil
}

func (us *UserService) Close() {
	db, _ := us.db.DB()
	db.Close()

}
