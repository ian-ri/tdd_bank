package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	fmt.Println("Welcome to the Golang bank")
	fmt.Println("You have the folllowing choices:")
	fmt.Println("0. Exit")
	fmt.Println("1. Open account")
	reader := bufio.NewScanner(os.Stdin)
	for {
		reader.Scan()
		input := reader.Text()
		if input == "0" {
			break
		}

		if input == "1" {
			fmt.Println("Feature incomplete!")
			continue
		}

		fmt.Println("Unknown command")
	}
}
