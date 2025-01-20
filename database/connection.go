package database

import (
	"fmt"

	"github.com/Yandelf00/leela/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// this creates a global variable that will
// make the database accessible from anywhere
var DB *gorm.DB

// this function connects to the database an migrates the models
func Connection() {
	// this line connects to the database
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	// this basically reads, if there is an
	// error print it and return
	if err != nil {
		fmt.Println("There was a problem conneting to the database !")
		return
	}

	// here we instantiate our global variable DB
	// as the db that connected successfully
	DB = db

	// here we migrate the models
	db.AutoMigrate(&models.Book{})
	db.AutoMigrate(&models.Practice{})
	db.AutoMigrate(&models.Skill{})
	db.AutoMigrate(&models.Task{})
}
