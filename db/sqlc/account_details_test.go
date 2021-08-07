package db

import (
	"context"
	"testing"

	"example.com/movielikes/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomAccount(t *testing.T) AccountDetail {
	user := CreateRandomUser(t)
	arg := CreateAccountDetailsParams{
		UserName: user.UserName,
		Balance: util.RandomBalannce(),
	}
	account, err := testQueries.CreateAccountDetails(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.UserName, account.UserName)
	require.Equal(t, arg.Balance, account.Balance)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	return account
}

func TestCreateAccount(t *testing.T)  {
	CreateRandomAccount(t)
}

func TestDeleteAccountDetails(t *testing.T) {
	account := CreateRandomAccount(t);
	err1 := testQueries.DeleteAccountDetails(context.Background(), account.ID)
	require.NoError(t, err1);
	expectedAccount, err := testQueries.GetAccountDetails(context.Background(), account.ID);
	require.Empty(t, expectedAccount);
	require.Error(t, err);
	require.Equal(t, err.Error(), "sql: no rows in result set")
}

func TestGetAccountDetails_Success(t *testing.T) {
	account := CreateRandomAccount(t)
	expectedAccount, err := testQueries.GetAccountDetails(context.Background(), account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, expectedAccount)
	require.Equal(t, expectedAccount.ID, account.ID)
	require.Equal(t, expectedAccount.Balance, account.Balance)
	require.Equal(t, expectedAccount.CreatedAt, account.CreatedAt)
	require.Equal(t, expectedAccount.UserName, account.UserName)
}

func TestGetAccountDetails_NotFound(t *testing.T) {
	expectedAccount, err := testQueries.GetAccountDetails(context.Background(), int64(-1))
	require.Error(t, err)
	require.Equal(t, err.Error(), "sql: no rows in result set")
	require.Empty(t, expectedAccount)
}

func TestGetAccountDetailsForUpdate(t *testing.T) {
	account := CreateRandomAccount(t)
	expectedAccount, err := testQueries.GetAccountDetailsForUpdate(context.Background(), account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, expectedAccount)
	require.Equal(t, expectedAccount.ID, account.ID)
	require.Equal(t, expectedAccount.Balance, account.Balance)
	require.Equal(t, expectedAccount.CreatedAt, account.CreatedAt)
	require.Equal(t, expectedAccount.UserName, account.UserName)
}

func TestUpdateAccountDetails(t *testing.T) {
	account := CreateRandomAccount(t)
	updatedAccount, err := testQueries.UpdateAccountDetails(context.Background(), 
	UpdateAccountDetailsParams{
		Balance: 1000,
		ID: account.ID,
	})
	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)
	require.Equal(t, updatedAccount.Balance, int64(1000))
}

func TestUpdateAccountBalance(t *testing.T) {
	account := CreateRandomAccount(t)
	updatedAccount, err := testQueries.UpdateAccountBalance(context.Background(), 
	UpdateAccountBalanceParams{
		Amount: 100,
		ID: account.ID,
	})
	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)
	require.Equal(t, updatedAccount.Balance, account.Balance + 100)
}