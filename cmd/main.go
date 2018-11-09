package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"github.com/mmircea16/tdd_bank/cmd/account"
	"fmt"
	"github.com/mmircea16/tdd_bank/cmd/account_service"
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
	writer.Write([]byte("4. Withdraw Money\n"))

	//myAccounts := make(map[string]*account.Account)

	myAccountServices := account_service.NewAccountService()

	for {
		input := readFromCmdLine(scanner)

		if input == "0" {
			break
		}

		if input == "1" {
			writer.Write([]byte("Enter account name\n"))
			name := readFromCmdLine(scanner)
			writer.Write([]byte("How much money?\n"))
			amount := readIntFromCmdLine(writer, scanner)
			//myAccount := account.NewAccount(name, amount)
			err := myAccountServices.Open(name,amount)
			if err != nil {
				writer.Write([]byte("Cannot be negative\n")) //refactor error message TODO
			} else {

				writer.Write([]byte("Account opened\n"))

				continue
			}
		}

		if input == "2" {


			if myAccountServices.AnyAccountExists() {
				writer.Write([]byte("Yes\n"))
			} else {
				writer.Write([]byte("No\n"))
			}
		}

		if input == "3" {
			if account, found := isAccountExist(writer, scanner, myAccounts); found {
				writer.Write([]byte(fmt.Sprintf("%d\n", account.CheckBalance())))
			}
		}

		if input == "4" {
			if  account, found := isAccountExist(writer, scanner, myAccounts); found {
				writer.Write([]byte("How much money to withdraw?\n"))
				amount := readIntFromCmdLine(writer, scanner)
				account.Withdraw(amount)
				writer.Write([]byte("Successful\n"))
			}
		}

	}
}

func isAccountExist(writer io.Writer, scanner *bufio.Scanner, accounts map[string]*account.Account) (*account.Account, bool) {
	if len(accounts) > 0 {
		writer.Write([]byte("Enter account name\n"))
		name := readFromCmdLine(scanner)

		if account, found := accounts[name]; found {
			return account, true
		}else {
			writer.Write([]byte("account doesnt exist\n"))
		}
	} else {
		writer.Write([]byte("No account available\n"))
	}
	return nil, false
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
