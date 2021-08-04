package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Color string
	Name  string
	Email string `gorm:"not null;uniqueIndex"`
}

func main() {
	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	// 	logger.Config{
	// 		SlowThreshold:             time.Second, // Slow SQL threshold
	// 		LogLevel:                  logger.Info, // Log level
	// 		IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
	// 		Colorful:                  true,        // Disable color
	// 	},
	// )

	dsn := "host=localhost user=postgres password=postgres dbname=lenslocked_dev port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Logger: newLogger,
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!!")
	err = db.AutoMigrate(&User{})
	if err != nil {
		return
	}
	// db.Logger{}
	name, email, color := getInfo()
	u := User{
		Name:  name,
		Email: email,
		Color: color,
	}
	if err = db.Create(&u).Error; err != nil {
		panic(err)
	}
	fmt.Println(u)

}

func getInfo() (name, email, color string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your name")
	name, _ = reader.ReadString('\n')

	fmt.Println("What is your email")
	email, _ = reader.ReadString('\n')
	fmt.Println("What is your fav color")
	color, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)
	color = strings.TrimSpace(color)

	return name, email, color
}
