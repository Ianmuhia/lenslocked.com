package models

import (
	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm/logger"

	"gorm.io/gorm"
)

var (
	/**
	 * ErrNotFound is returned when a reasouce is not found
	 */
	ErrNotFound = errors.New("models: reasource not found")
)

func NewUserService(connectionInfo string) (*UserService, error) {

	// dsn := "host=localhost user=postgres password=postgres dbname=lenslocked_dev port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(connectionInfo), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	return &UserService{
		db: db,
	}, nil

}

type UserService struct {
	db *gorm.DB
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
	// err := us.db.Where("id = ?", id).First(&user).Error
	// switch err {
	// case nil:
	// 	return &user, nil
	// case gorm.ErrRecordNotFound:
	// 	return nil, ErrNotFound
	// default:
	// 	return nil, err

	// }

}
func (us *UserService) ByEmail(email string) (*User, error) {
	var user User
	db := us.db.Where("email = ?", email)
	// err := db.First(&user).Error
	err := first(db, &user)
	return &user, err
	// switch err {
	// case nil:
	// 	return &user, nil
	// case gorm.ErrRecordNotFound:
	// 	return nil, ErrNotFound
	// default:
	// 	return nil, err

	// }
	// return nil
	// return us.db.Save(user).Error
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
	// return nil
	return us.db.Create(user).Error
}
func (us *UserService) Update(user *User) error {
	// return nil
	return us.db.Save(user).Error
}

func (us *UserService) DestructiveReset() {
	us.db.Migrator().DropTable(&User{})
	us.db.AutoMigrate(&User{})
}

func (us *UserService) Close(error) {
	db, _ := us.db.DB()
	db.Close()

}

type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;uniqueIndex"`
}
