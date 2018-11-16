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
}
