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
