package controller

import (
	"github.com/Yandelf00/leela/queries"
)

func BookController() {
	inpt := input_instance
	if len(inpt.args) > 0 {
		switch inpt.args[0] {
		case "get":
			BookGetController()
		case "add":
			BookAddController()
		}
	}

}

func BookGetController() {
	inpt := input_instance
	if len(inpt.args) == 3 {
		switch inpt.args[1] {
		case "-n":
			queries.GetBookByName(inpt.args[2])
		case "-c":
			queries.GetBookByCategory(inpt.args[2])
		case "-a":
			queries.GetBookByAuthor(inpt.args[2])
		}
	}
	if len(inpt.args) == 2 {
		switch inpt.args[1] {
		case "read":
			queries.GetReadBooks()
		case "unread":
			queries.GetUnreadBooks()
		case "all":
			queries.GetAllBooks()
		}
	}
}

func BookAddController() {

}
