package main

import (
	"testing"

	"github.com/stretchr/testify/require"
	"io"
	"bufio"
)

func TestStartBankUI(t *testing.T) {
	t.Run("should be able open an account", func(t *testing.T) {
		cmdLine := NewFakeCmdLine()

		go startBankUI(cmdLine, cmdLine)

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

	t.Run("should be able open two accounts and check balance on each one", func(t *testing.T) {
		cmdLine := NewFakeCmdLine()

		go startBankUI(cmdLine, cmdLine)

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
		respond(cmdLine, "1")
		expectLine(t, cmdLine, "Enter account name")
		respond(cmdLine, "some-other-name")
		expectLine(t, cmdLine, "How much money?")
		respond(cmdLine, "50")
		expectLine(t, cmdLine, "Account opened")
		respond(cmdLine, "2")
		expectLine(t, cmdLine, "Yes")

		respond(cmdLine, "3")
		expectLine(t, cmdLine, "Enter account name")
		respond(cmdLine, "some-name")
		expectLine(t, cmdLine, "20")

		respond(cmdLine, "3")
		expectLine(t, cmdLine, "Enter account name")
		respond(cmdLine, "some-other-name")
		expectLine(t, cmdLine, "50")
	})

	t.Run("negative check", func(t *testing.T) {
		cmdLine := NewFakeCmdLine()

		go startBankUI(cmdLine, cmdLine)

		expectMenu(t, cmdLine)
		respond(cmdLine, "2")
		expectLine(t, cmdLine, "No")
		respond(cmdLine, "1")
		expectLine(t, cmdLine, "Enter account name")
		respond(cmdLine, "some-name")
		expectLine(t, cmdLine, "How much money?")
		respond(cmdLine, "-20")
		expectLine(t, cmdLine, "Cannot be negative")
		respond(cmdLine, "2")
		expectLine(t, cmdLine, "No")
	})

	t.Run("open account and check balance", func(t *testing.T) {
		cmdLine := NewFakeCmdLine()

		go startBankUI(cmdLine, cmdLine)

		expectMenu(t, cmdLine)
		respond(cmdLine, "1")
		expectLine(t, cmdLine, "Enter account name")
		respond(cmdLine, "some-name")
		expectLine(t, cmdLine, "How much money?")
		respond(cmdLine, "100")
		expectLine(t, cmdLine, "Account opened")
		respond(cmdLine, "3")
		expectLine(t, cmdLine, "Enter account name")
		respond(cmdLine, "some-name")
		expectLine(t, cmdLine, "100")
	})

	t.Run("open account and check balance for non-existant account", func(t *testing.T) {
		cmdLine := NewFakeCmdLine()

		go startBankUI(cmdLine, cmdLine)

		expectMenu(t, cmdLine)
		respond(cmdLine, "1")
		expectLine(t, cmdLine, "Enter account name")
		respond(cmdLine, "some-name")
		expectLine(t, cmdLine, "How much money?")
		respond(cmdLine, "100")
		expectLine(t, cmdLine, "Account opened")
		respond(cmdLine, "3")
		expectLine(t, cmdLine, "Enter account name")
		respond(cmdLine, "someone else")
		expectLine(t, cmdLine, "account doesnt exist")
	})

	t.Run("check balance on non-existant account", func(t *testing.T) {
		cmdLine := NewFakeCmdLine()

		go startBankUI(cmdLine, cmdLine)

		expectMenu(t, cmdLine)
		respond(cmdLine, "3")
		expectLine(t, cmdLine, "No account available")
	})

	t.Run("open account and withdraw money", func(t *testing.T){
		cmdLine := NewFakeCmdLine()

		go startBankUI(cmdLine, cmdLine)

		expectMenu(t, cmdLine)
		respond(cmdLine, "1")
		expectLine(t, cmdLine, "Enter account name")
		respond(cmdLine, "some-name")
		expectLine(t, cmdLine, "How much money?")
		respond(cmdLine, "100")
		expectLine(t, cmdLine, "Account opened")
		respond(cmdLine, "4")
		expectLine(t, cmdLine, "Enter account name")
		respond(cmdLine, "some-name")
		expectLine(t, cmdLine, "How much money to withdraw?")
		respond(cmdLine, "20")
		expectLine(t, cmdLine, "Successful")
		respond(cmdLine, "3")
		expectLine(t, cmdLine, "Enter account name")
		respond(cmdLine, "some-name")
		expectLine(t, cmdLine, "80")
	})


	t.Run("open account and withdraw money from non-existant account", func(t *testing.T){
		cmdLine := NewFakeCmdLine()

		go startBankUI(cmdLine, cmdLine)

		expectMenu(t, cmdLine)
		respond(cmdLine, "1")
		expectLine(t, cmdLine, "Enter account name")
		respond(cmdLine, "some-name")
		expectLine(t, cmdLine, "How much money?")
		respond(cmdLine, "100")
		expectLine(t, cmdLine, "Account opened")
		respond(cmdLine, "4")
		expectLine(t, cmdLine, "Enter account name")
		respond(cmdLine, "some other account")
		expectLine(t, cmdLine, "account doesnt exist")
	})
}

func expectMenu(t *testing.T, cmdLine *fakeCmdLine) {
	expectLine(t, cmdLine, "Welcome to the Golang bank")
	expectLine(t, cmdLine, "You have the folllowing choices:")
	expectLine(t, cmdLine, "0. Exit")
	expectLine(t, cmdLine, "1. Open account")
	expectLine(t, cmdLine, "2. Do I have an opened account?")
	expectLine(t, cmdLine, "3. Check Balance")
	expectLine(t, cmdLine, "4. Withdraw Money")
}

func expectLine(t *testing.T, buffer io.Reader, line string) {
	scanner := bufio.NewScanner(buffer)
	scanner.Scan()
	text := scanner.Text()
	require.Equal(t, line, text)
}

func respond(buffer io.Writer, line string) {
	buffer.Write([]byte(line + "\n"))
}

func NewFakeCmdLine() *fakeCmdLine {
	return &fakeCmdLine{console: make(chan string)}
}

type fakeCmdLine struct {
	console chan string
}

func (cmdLine *fakeCmdLine) Write(p []byte) (n int, err error) {
	cmdLine.console <- string(p)
	return len(p), nil
}

func (cmdLine *fakeCmdLine) Read(p []byte) (n int, err error) {
	content := <-cmdLine.console

	contentLen := len(p)
	if len(content) < len(p) {
		contentLen = len(content)
	}

	for i := 0; i < contentLen; i++ {
		p[i] = []byte(content)[i]
	}

	return contentLen, nil
}
