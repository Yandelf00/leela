package controller

import (
	"strings"
)

// this is our input structure
// we have the command and the arguments
// following the command
type Input struct {
	command string
	args    []string
}

// this is a global variable
// so we can access our input
// from everywhere
var input_instance Input

// this function instantiates our input
func SetInputInstance(inpt string) {
	if len(inpt) > 0 {
		// eg : inpt = go run main.go
		// 		tmp = [go run main.go]
		tmp := strings.Fields(inpt)

		// input_instance.command = go
		// input_instance.args=[run main.go]
		input_instance.command = tmp[0]
		input_instance.args = tmp[1:]
		Controller()
	}
}

// this function is to control our input
// to eventually get us to the actual function
// that we should call
func Controller() {
	switch input_instance.command {
	case "book":
		BookController()
	}
}
