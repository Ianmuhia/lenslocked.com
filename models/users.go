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
	err := us.db.Where("id = ?", id).First(&user).Error
	switch err {
	case nil:
		return &user, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	default:
		return nil, err

	}

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
