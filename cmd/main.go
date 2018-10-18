package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"github.com/mmircea16/tdd_bank/cmd/account"
	"fmt"
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
	writer.Write([]byte("2. Do I have an opened account?\n"))
	writer.Write([]byte("3. Check Balance\n"))

	var myAccount *account.Account

	for {
		input := readFromCmdLine(scanner)

		if input == "0" {
			break
		}

		if input == "1" {
			writer.Write([]byte("How much money?\n"))
			amount := readIntFromCmdLine(writer, scanner)
			myAccount = account.NewAccount(amount)

			if myAccount == nil {
				writer.Write([]byte("Cannot be negative\n"))
			} else {
				writer.Write([]byte("Account opened\n"))

				continue
			}
		}

		if input == "2" {
			if myAccount != nil {
				writer.Write([]byte("Yes\n"))
			} else {
				writer.Write([]byte("No\n"))
			}
		}

		if input == "3" {
			if myAccount != nil {
				writer.Write([]byte(fmt.Sprintf("%d\n", myAccount.CheckBalance())))
			} else {
				writer.Write([]byte("No account available\n"))
			}

		}

	}
}

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
