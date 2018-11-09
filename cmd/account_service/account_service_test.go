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


	t.Run("check balance on non-existant account", func(t *testing.T) {
		accountService := NewAccountService()
		_, err := accountService.CheckBalance("wrong account")
		require.Error(t, err)
	})


	t.Run("open account and withdraw money", func(t *testing.T) {
		accountService := NewAccountService()
		accountService.Open("first account", 20)
		err := accountService.Withdraw("first account", 10)
		require.NoError(t, err)

		balance, _ := accountService.CheckBalance("first account")

		require.Equal(t, int64(10), balance)
	})

	t.Run("open account and withdraw money from non existant account", func(t *testing.T) {
		accountService := NewAccountService()
		accountService.Open("first account", 20)
		err := accountService.Withdraw("non existant account", 10)
		require.Error(t, err)
	})

	t.Run("don't open account and withdraw money from non existant account", func(t *testing.T) {
		accountService := NewAccountService()
		err := accountService.Withdraw("non existant account", 10)
		require.Error(t, err)
	})

	t.Run("should be able open two accounts and check balance on each one", func(t *testing.T) {
		accountService := NewAccountService()
		accountService.Open("first account", 20)
		accountService.Open("second account", 30)
		firstAccountBalance, _ := accountService.CheckBalance("first account")
		require.Equal(t, int64(20), firstAccountBalance)

		secondAccountBalance, _ := accountService.CheckBalance("second account")
		require.Equal(t, int64(30), secondAccountBalance)

	})

	t.Run("open account and check if exists", func(t *testing.T) {
		accountService := NewAccountService()
		accountService.Open("first account", 20)
		found := accountService.AnyAccountExists()
		require.True(t,found)
	})

	t.Run("dont open account and check if any exist", func(t *testing.T) {
		accountService := NewAccountService()
		found := accountService.AnyAccountExists()
		require.False(t,found)
	})


}