package account_service

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestAccountService(t *testing.T) {

	t.Run("should be able to open account", func(t *testing.T) {
		accountService := NewAccountService()

		require.NoError(t, accountService.Open("first account", 20))
	})

	t.Run("should not create account with negative initial amount", func(t *testing.T) {
		accountService := NewAccountService()

		require.Error(t, accountService.Open("first account", -20))
	})

	t.Run("open account and check balance", func(t *testing.T) {
		accountService := NewAccountService()
		accountService.Open("first account", 20)
		balance, _ := accountService.CheckBalance("first account")
		require.Equal(t, int64(20), balance)

	})

	t.Run("open account and check balance for non-existant account", func(t *testing.T) {
		accountService := NewAccountService()
		accountService.Open("first account", 20)
		_, err := accountService.CheckBalance("second account")
		require.Error(t, err)
	})
}