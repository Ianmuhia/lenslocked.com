package models

import (
	"testing"
	"time"

	"gorm.io/gorm/logger"
)

func testingUserService() (*UserService, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=lenslocked_dev port=5432 sslmode=disable"
	us, err := NewUserService(dsn)
	if err != nil {
		return nil, err
	}
	us.db.Logger.LogMode(logger.Silent)
	//clear the users table between tests
	us.DestructiveReset()
	return us, nil
}

func TestCreateUser(t *testing.T) {
	us, err := testingUserService()
	if err != nil {
		t.Fatal(err)
	}
	user := User{
		Name:  "ian muhia",
		Email: "ianmuhia3@gmail.com",
	}
	err = us.Create(&user)
	if err != nil {
		t.Fatal(err)
	}

	if user.ID == 0 {
		t.Errorf("Expected ID > 0 .Recieved %d", user.ID)
	}

	if time.Since(user.CreatedAt) > time.Duration(5*time.Second) {
		t.Errorf("Expected CreatedAt to be recent. Recieved %s", user.CreatedAt)
	}
	if time.Since(user.UpdatedAt) > time.Duration(5*time.Second) {
		t.Errorf("Expected UpdatedAt to be recent. Recieved %s", user.UpdatedAt)
	}

}
