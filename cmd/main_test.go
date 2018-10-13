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

		expectLine(t, cmdLine, "Welcome to the Golang bank")
		expectLine(t, cmdLine, "You have the folllowing choices:")
		expectLine(t, cmdLine, "0. Exit")
		expectLine(t, cmdLine, "1. Open account")
		expectLine(t, cmdLine, "2. Do I have an opened account?")
		respond(cmdLine, "2")
		expectLine(t, cmdLine, "No")
		respond(cmdLine, "1")
		expectLine(t, cmdLine, "How much money?")
		respond(cmdLine, "20")
		expectLine(t, cmdLine, "Account opened")
		respond(cmdLine, "2")
		expectLine(t, cmdLine, "Yes")
	})

	t.Run("negative check", func(t *testing.T) {
		cmdLine := NewFakeCmdLine()

		go startBankUI(cmdLine, cmdLine)

		expectLine(t, cmdLine, "Welcome to the Golang bank")
		expectLine(t, cmdLine, "You have the folllowing choices:")
		expectLine(t, cmdLine, "0. Exit")
		expectLine(t, cmdLine, "1. Open account")
		expectLine(t, cmdLine, "2. Do I have an opened account?")
		respond(cmdLine, "2")
		expectLine(t, cmdLine, "No")
		respond(cmdLine, "1")
		expectLine(t, cmdLine, "How much money?")
		respond(cmdLine, "-20")
		expectLine(t, cmdLine, "Cannot be negative")
		respond(cmdLine, "2")
		expectLine(t, cmdLine, "No")
	})
}

func expectLine(t *testing.T, buffer io.Reader, line string) {
	scanner := bufio.NewScanner(buffer)
	scanner.Scan()
	require.Equal(t, line, scanner.Text())
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
