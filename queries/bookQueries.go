package queries

import (
	"fmt"

	"github.com/Yandelf00/leela/database"
	"github.com/Yandelf00/leela/models"
)

func AddBook() {
	var name string
	var category string
	var author string
	fmt.Println("Enter the book's name : ")

	fmt.Println("Enter the book's category (can be left blank) :")
	fmt.Scanln(&category)
	fmt.Println("Enter the book's author (can be left blank) :")
	fmt.Scan(&author)

	book := models.Book{
		Name:     name,
		Category: category,
		Author:   author,
		IsRead:   false,
		Rating:   0,
	}
	database.DB.Create(&book)
}
