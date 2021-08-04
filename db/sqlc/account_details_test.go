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