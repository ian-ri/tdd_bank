package main

import (
	"testing"
	"github.com/mmircea16/tdd_bank/cmd/account_service"
)

func TestBank(t *testing.T) {
	t.Run("should be able open an account", func(t *testing.T) {
		cmdLine := NewFakeCmdLine()
		b := NewBankUI(account_service.NewAccountService())

		go b.start(cmdLine, cmdLine)

		expectMenu(t, cmdLine)
		respond(cmdLine, "2")
		expectLine(t, cmdLine, "No")
		respond(cmdLine, "1")
		expectLine(t, cmdLine, "Enter account name")
		respond(cmdLine, "some-name")
		expectLine(t, cmdLine, "How much money?")
		respond(cmdLine, "20")
		expectLine(t, cmdLine, "Account opened")
		respond(cmdLine, "2")
		expectLine(t, cmdLine, "Yes")
	})

	t.Run("should be able to open an account, withdraw money, and check remaining balance", func(t *testing.T) {
		cmdLine := NewFakeCmdLine()
		b := NewBankUI(account_service.NewAccountService())

		accountName := "some-name"

		go b.start(cmdLine, cmdLine)

		// open account
		expectMenu(t, cmdLine)
		respond(cmdLine, "2")
		expectLine(t, cmdLine, "No")
		respond(cmdLine, "1")
		expectLine(t, cmdLine, "Enter account name")
		respond(cmdLine, accountName)
		expectLine(t, cmdLine, "How much money?")
		respond(cmdLine, "20")
		expectLine(t, cmdLine, "Account opened")
		respond(cmdLine, "2")
		expectLine(t, cmdLine, "Yes")

		// withdraw money
		respond(cmdLine, "4")
		expectLine(t, cmdLine, "Enter account name")
		respond(cmdLine, accountName)
		expectLine(t, cmdLine, "How much money to withdraw?")
		respond(cmdLine, "5")
		expectLine(t, cmdLine, "Successful")

		// check remaining balance
		respond(cmdLine, "3")
		expectLine(t, cmdLine, "Enter account name")
		respond(cmdLine, accountName)
		expectLine(t, cmdLine, "15")
	})


}
