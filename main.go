package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Yandelf00/leela/controller"
	"github.com/Yandelf00/leela/database"
)

func main() {
	database.Connection()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(">>> ")
		scanner.Scan()
		input := scanner.Text()
		if input == "exit" {
			break
		}
		controller.SetInputInstance(input)
	}
}
