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
	writer.Write([]byte("Welcome to the Golang bank\n"))
	writer.Write([]byte("You have the folllowing choices:\n"))
	writer.Write([]byte("0. Exit\n"))
	writer.Write([]byte("1. Open account\n"))
	for {
		input := readFromCmdLine(reader)
		if input == "0" {
			break
		}

		if input == "1" {
			writer.Write([]byte("How much money?\n"))
			amount := readIntFromCmdLine(writer, reader)
			doSomething(amount)
			writer.Write([]byte("Feature incomplete!\n"))
			continue
		}

		writer.Write([]byte("Unknown command\n"))
	}
}

func doSomething(something interface{}) {}

func readIntFromCmdLine(writer io.Writer, reader io.Reader) int64 {
	input := readFromCmdLine(reader)
	amount, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		writer.Write([]byte("You need to insert an integer\n"))
		os.Exit(1)
	}
	return amount
}

func readFromCmdLine(reader io.Reader) string {
	buffReader := bufio.NewScanner(reader)
	buffReader.Scan()
	return buffReader.Text()
}
