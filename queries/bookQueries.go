package queries

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Yandelf00/leela/database"
	"github.com/Yandelf00/leela/models"
)

func AddBook() {
	var name string
	var category string
	var author string

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the book's name : ")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Println("Enter the book's category (can be left blank) :")
	category, _ = reader.ReadString('\n')
	category = strings.TrimSpace(category)

	fmt.Println("Enter the book's author (can be left blank) :")
	author, _ = reader.ReadString('\n')
	author = strings.TrimSpace(author)

	book := models.Book{
		Name:     name,
		Category: category,
		Author:   author,
		IsRead:   false,
		Rating:   0,
	}
	database.DB.Create(&book)
}

func GetAllBooks() {
	var books []models.Book
	database.DB.Find(&books)
	for _, book := range books {
		fmt.Printf("book name : %s \n", book.Name)
	}
}

func GetBookByName(name string) {
	var books []models.Book
	database.DB.Where("name = ?", name).Find(&books)
	for _, book := range books {
		fmt.Printf("book name : %s \n", book.Name)
	}
}

func GetBookByCategory(category string) {
	var books []models.Book
	database.DB.Where("category = ?", category).Find(&books)
	for _, book := range books {
		fmt.Printf("book name : %s \n", book.Name)
	}
}

func GetBookByAuthor(author string) {
	var books []models.Book
	database.DB.Where("author = ?", author).Find(&books)
	for _, book := range books {
		fmt.Printf("book name : %s \n", book.Name)
	}
}

func GetReadBooks() {
	var books []models.Book
	database.DB.Where("is_read = ?", true).Find(&books)
	for _, book := range books {
		fmt.Printf("book name : %s \n", book.Name)
	}
}

func GetUnreadBooks() {
	var books []models.Book
	database.DB.Where("is_read = ?", false).Find(&books)
	for _, book := range books {
		fmt.Printf("book name : %s \n", book.Name)
	}

}
