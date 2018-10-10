package main

import (
	"testing"
	"io/ioutil"
	"os"
	"strings"
	"bytes"
	"github.com/stretchr/testify/require"
)

func TestStartBankUI(t *testing.T) {
	t.Run("should be able open an account", func(t *testing.T) {
		commands := []string{"2", "1", "20", "2", "0"}
		expectedLines := []string{
			"Welcome to the Golang bank",
			"You have the folllowing choices:",
			"0. Exit",
			"1. Open account",
			"2. Do I have an opened account?",
			"No",
			"How much money?",
			"Account opened",
			"Yes",
		}
		testCase := strings.Join(commands, "\n")
		ioutil.WriteFile("/tmp/testin", []byte(testCase), os.ModePerm)
		in, _ := os.Open("/tmp/testin")

		out := new(bytes.Buffer)
		startBankUI(in, out)
		output, _ := ioutil.ReadAll(out)

		require.Equal(t, strings.Join(expectedLines, "\n"), string(output))
	})
}

