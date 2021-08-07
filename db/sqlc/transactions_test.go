package db

import (
	"context"
	"testing"

	"example.com/movielikes/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomTransaction(t *testing.T) (Transaction) {
	accountDetail := CreateRandomAccount(t)
	arg := CreateTransactionParams{
		AccountID: accountDetail.ID,
		Amount: util.RandomBalannce(),
	}
	transaction, err := testQueries.CreateTransaction(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transaction)
	require.Equal(t, arg.AccountID, transaction.AccountID)
	require.Equal(t, arg.Amount, transaction.Amount)
	require.NotZero(t, transaction.ID)
	require.NotZero(t, transaction.CreatedAt)
	return transaction;
}

func CreateTransactionForAccount(t *testing.T, accountId int64) (Transaction) {
	arg := CreateTransactionParams{
		AccountID: accountId,
		Amount: util.RandomBalannce(),
	}
	transaction, err := testQueries.CreateTransaction(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transaction)
	require.Equal(t, arg.AccountID, transaction.AccountID)
	require.Equal(t, arg.Amount, transaction.Amount)
	require.NotZero(t, transaction.ID)
	require.NotZero(t, transaction.CreatedAt)
	return transaction;
}

func TestCreateTransaction(t *testing.T) {
	CreateRandomTransaction(t);
}

func TestCreateTransaction_Failure(t *testing.T) {
	arg := CreateTransactionParams{
		AccountID: int64(-1),
		Amount: util.RandomBalannce(),
	}
	transaction, err := testQueries.CreateTransaction(context.Background(), arg)
	require.Error(t, err)
	require.Equal(t, err.Error(), "pq: insert or update on table \"transactions\" violates foreign key constraint \"transactions_account_id_fkey\"")
	require.Empty(t, transaction)
}

func TestGetTransaction(t *testing.T) {
	transaction := CreateRandomTransaction(t)
	expectedTransaction, err := testQueries.GetTransaction(context.Background(), transaction.ID)
	require.NoError(t, err)
	require.NotEmpty(t, expectedTransaction)
	require.Equal(t, expectedTransaction.ID, transaction.ID)
	require.Equal(t, expectedTransaction.AccountID, transaction.AccountID)
	require.Equal(t, expectedTransaction.Amount, transaction.Amount)
	require.Equal(t, expectedTransaction.CreatedAt, transaction.CreatedAt)
}

func TestGetTransactions(t *testing.T) {
	account := CreateRandomAccount(t)
	transactions := []Transaction{}
	for i := 0; i < 5; i++ {
		transactions = append(transactions, CreateTransactionForAccount(t, account.ID))
	}
	arg := GetTransactionsParams{
		AccountID: account.ID,
		Limit: 5,
		Offset: 0,
	}
	actualTransactions, err := testQueries.GetTransactions(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, actualTransactions)
	require.Len(t, actualTransactions, 5)
	for _, transaction := range actualTransactions {
		require.NotEmpty(t, transaction)
		require.Equal(t, transaction.AccountID, account.ID)
	}
}