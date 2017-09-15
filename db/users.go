package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	ChatID int64 
	FirstName string
	LastName string
}

func Init() *gorm.DB {
	db, err := gorm.Open("sqlite3", "users.db")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})

	return db
}

func Registered(ChatID int64) bool {
	db := Init()
	defer db.Close()

	var user User
	res := db.First(&user, "chat_id = ?", ChatID)

	if res.Error != nil {
		errorMsg := res.Error.Error()

		if errorMsg == "record not found" {
			return false
		}
		fmt.Println(errorMsg)
	}
	return true
}


func Save(user *User) {
	db := Init()
	defer db.Close()

	result := db.Create(&user)

	if result.Error != nil {
		errorMsg := result.Error.Error()
		fmt.Println(errorMsg)
	}
}