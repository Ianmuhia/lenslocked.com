package main

import (
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	gorm.Model
	Color  string
	Name   string
	Email  string `gorm:"not null;uniqueIndex"`
	Orders []Order
}

type Order struct {
	gorm.Model
	UserID      uint
	Amount      int
	Description string
}

func main() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)

	dsn := "host=localhost user=postgres password=postgres dbname=lenslocked_dev port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!!")
	err = db.AutoMigrate(&User{}, &Order{})
	if err != nil {
		return
	}
	var u User
	if err := db.Preload("Orders").First(&u).Error; err != nil {
		panic(err)
	}
	fmt.Println(u)
	fmt.Println(u.Orders)

	// createOrder(db, u, 1001, "Fake Description #1")
	// createOrder(db, u, 4534, "Fake Description #2")
	// createOrder(db, u, 2215, "Fake Description #3")
	// createOrder(db, u, 5656, "Fake Description #4")

}

func createOrder(db *gorm.DB, user User, amount int, desc string) {
	err := db.Create(&Order{
		UserID:      user.ID,
		Amount:      amount,
		Description: desc,
	}).Error

	if err != nil {
		panic(err)
	}

}
