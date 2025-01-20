package controller

import "github.com/Yandelf00/leela/queries"

func BookController() {
	inpt := input_instance
	if len(inpt.args) == 1 {
		switch inpt.args[0] {
		case "add":
			queries.AddBook()
		}
	}

}
