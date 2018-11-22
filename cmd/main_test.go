package main

import (
	"testing"

	"github.com/stretchr/testify/require"
	"io"
	"bufio"
	"github.com/mmircea16/tdd_bank/cmd/account_service"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
)

func TestStartBankUI(t *testing.T) {
	t.Run("should be able open an account", func(t *testing.T) {
		cmdLine := NewFakeCmdLine()
		ctrl := gomock.NewController(t)
		mockAccountService := account_service.NewMockAccountService(ctrl)
		mockAccountService.EXPECT().Open("some-name", int64(20)).Return(nil)

		b := NewBankUI(mockAccountService)
		go b.start(cmdLine, cmdLine)

		expectMenu(t, cmdLine)
		respond(cmdLine, "1")
		expectLine(t, cmdLine, "Enter account name")
		respond(cmdLine, "some-name")
		expectLine(t, cmdLine, "How much money?")
		respond(cmdLine, "20")
		expectLine(t, cmdLine, "Account opened")
	})

	t.Run("should be able to exit bank", func(t *testing.T) {
		cmdLine := NewFakeCmdLine()
		ctrl := gomock.NewController(t)
		mockAccountService := account_service.NewMockAccountService(ctrl)

		b := NewBankUI(mockAccountService)
		go b.start(cmdLine, cmdLine)

		expectMenu(t, cmdLine)
		respond(cmdLine, "0")
		expectLine(t, cmdLine, "Goodbye!")
	})

	t.Run("should not be able to open account with a negative balance", func(t *testing.T) {
		cmdLine := NewFakeCmdLine()
		ctrl := gomock.NewController(t)
		mockAccountService := account_service.NewMockAccountService(ctrl)
		mockAccountService.EXPECT().Open(gomock.Any(), gomock.Any()).Times(1).Return(errors.New("An error"))

		b := NewBankUI(mockAccountService)
		go b.start(cmdLine, cmdLine)

		expectMenu(t, cmdLine)
		respond(cmdLine, "1")
		expectLine(t, cmdLine, "Enter account name")
		respond(cmdLine, "some-name")
		expectLine(t, cmdLine, "How much money?")
		respond(cmdLine, "-20")
		expectLine(t, cmdLine, "Cannot be negative")
	})

	t.Run("should be able to check if any account exists and it does", func(t *testing.T) {
		cmdLine := NewFakeCmdLine()
		ctrl := gomock.NewController(t)
		mockAccountService := account_service.NewMockAccountService(ctrl)
		mockAccountService.EXPECT().AnyAccountExists().Times(1).Return(true)

		b := NewBankUI(mockAccountService)
		go b.start(cmdLine, cmdLine)

		expectMenu(t, cmdLine)
		respond(cmdLine, "2")
		expectLine(t, cmdLine, "Yes")
	})

	t.Run("should be able to check if any account exists and it doesn't", func(t *testing.T) {
		cmdLine := NewFakeCmdLine()
		ctrl := gomock.NewController(t)
		mockAccountService := account_service.NewMockAccountService(ctrl)
		mockAccountService.EXPECT().AnyAccountExists().Times(1).Return(false)

		b := NewBankUI(mockAccountService)
		go b.start(cmdLine, cmdLine)

		expectMenu(t, cmdLine)
		respond(cmdLine, "2")
		expectLine(t, cmdLine, "No")
	})

	t.Run("should be able to try to check balance when no account exists", func(t *testing.T) {
		cmdLine := NewFakeCmdLine()
		ctrl := gomock.NewController(t)
		mockAccountService := account_service.NewMockAccountService(ctrl)
		mockAccountService.EXPECT().AnyAccountExists().Times(1).Return(false)

		b := NewBankUI(mockAccountService)
		go b.start(cmdLine, cmdLine)

		expectMenu(t, cmdLine)
		respond(cmdLine, "3")
		expectLine(t, cmdLine, "No account available")
	})

	t.Run("should be able to try to check balance when only a different account exists", func(t *testing.T) {
		cmdLine := NewFakeCmdLine()
		ctrl := gomock.NewController(t)
		mockAccountService := account_service.NewMockAccountService(ctrl)
		mockAccountService.EXPECT().AnyAccountExists().Times(1).Return(true)
		mockAccountService.EXPECT().CheckBalance("some-account").Times(1).Return(int64(0), errors.New("an error"))

		b := NewBankUI(mockAccountService)
		go b.start(cmdLine, cmdLine)

		expectMenu(t, cmdLine)
		respond(cmdLine, "3")
		expectLine(t, cmdLine, "Enter account name")
		respond(cmdLine, "some-account")
		expectLine(t, cmdLine, "account doesnt exist")
	})

	t.Run("should be able to check balance when account exists", func(t *testing.T) {
		cmdLine := NewFakeCmdLine()
		ctrl := gomock.NewController(t)
		mockAccountService := account_service.NewMockAccountService(ctrl)
		mockAccountService.EXPECT().AnyAccountExists().Times(1).Return(true)
		mockAccountService.EXPECT().CheckBalance("some-account").Times(1).Return(int64(20), nil)

		b := NewBankUI(mockAccountService)
		go b.start(cmdLine, cmdLine)

		expectMenu(t, cmdLine)
		respond(cmdLine, "3")
		expectLine(t, cmdLine, "Enter account name")
		respond(cmdLine, "some-account")
		expectLine(t, cmdLine, "20")
	})


	t.Run("should be able to try to withdraw money where no accounts exist", func(t *testing.T) {
		cmdLine := NewFakeCmdLine()
		ctrl := gomock.NewController(t)
		mockAccountService := account_service.NewMockAccountService(ctrl)
		mockAccountService.EXPECT().AnyAccountExists().Times(1).Return(false)

		b := NewBankUI(mockAccountService)
		go b.start(cmdLine, cmdLine)

		expectMenu(t, cmdLine)
		respond(cmdLine, "4")
		expectLine(t, cmdLine, "No account available")
	})

	t.Run("should be able to try to withdraw money when only a different account exists", func(t *testing.T) {
		cmdLine := NewFakeCmdLine()
		ctrl := gomock.NewController(t)
		mockAccountService := account_service.NewMockAccountService(ctrl)
		mockAccountService.EXPECT().AnyAccountExists().Times(1).Return(true)
		mockAccountService.EXPECT().AccountExists("some-account").Times(1).Return(false)

		b := NewBankUI(mockAccountService)
		go b.start(cmdLine, cmdLine)

		expectMenu(t, cmdLine)
		respond(cmdLine, "4")
		expectLine(t, cmdLine, "Enter account name")
		respond(cmdLine, "some-account")
		expectLine(t, cmdLine, "account doesnt exist")
	})

	t.Run("should be able to withdraw money when account exists", func(t *testing.T) {
		cmdLine := NewFakeCmdLine()
		ctrl := gomock.NewController(t)
		mockAccountService := account_service.NewMockAccountService(ctrl)
		mockAccountService.EXPECT().AnyAccountExists().Times(1).Return(true)
		mockAccountService.EXPECT().AccountExists("some-account").Times(1).Return(true)
		mockAccountService.EXPECT().Withdraw("some-account", int64(20)).Times(1).Return(nil)

		b := NewBankUI(mockAccountService)
		go b.start(cmdLine, cmdLine)

		expectMenu(t, cmdLine)
		respond(cmdLine, "4")
		expectLine(t, cmdLine, "Enter account name")
		respond(cmdLine, "some-account")
		expectLine(t, cmdLine, "How much money to withdraw?")
		respond(cmdLine, "20")
		expectLine(t, cmdLine, "Successful")
	})

}

func expectMenu(t *testing.T, cmdLine *fakeCmdLine) {
	expectLine(t, cmdLine, "Welcome to the Golang bank")
	expectLine(t, cmdLine, "You have the following choices:")
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
