package main

import (
	"os"
	"bufio"
	"strconv"
	"io"
)

func main() {
	startBankUI(os.Stdin, os.Stdout)
}

func startBankUI(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	writer.Write([]byte("Welcome to the Golang bank\n"))
	writer.Write([]byte("You have the folllowing choices:\n"))
	writer.Write([]byte("0. Exit\n"))
	writer.Write([]byte("1. Open account\n"))
	for {
		input := readFromCmdLine(scanner)
		if input == "0" {
			break
		}

		if input == "1" {
			writer.Write([]byte("How much money?\n"))
			amount := readIntFromCmdLine(writer, scanner)
			doSomething(amount)
			writer.Write([]byte("Feature incomplete!\n"))
			continue
		}

		writer.Write([]byte("Unknown command\n"))
	}
}

func doSomething(something interface{}) {}

func readIntFromCmdLine(writer io.Writer, scanner *bufio.Scanner) int64 {
	input := readFromCmdLine(scanner)
	amount, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		writer.Write([]byte("You need to insert an integer\n"))
		os.Exit(1)
	}
	return amount
}

func readFromCmdLine(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}
