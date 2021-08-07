package db

import (
	"context"
	"testing"

	"example.com/movielikes/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomTranfser(t *testing.T, amount int64, sourceAccountId int64, targetAccountId int64) Transfer {
	arg := CreateTransferParams{
		SourceAccountID: sourceAccountId,
		TargetAccountID: targetAccountId,
		Amount: amount,
	}
	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, sourceAccountId, transfer.SourceAccountID)
	require.Equal(t, targetAccountId, transfer.TargetAccountID)
	require.Equal(t, amount, transfer.Amount)
	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)
	return transfer;
}

func TestCreateTransfer(t *testing.T) {
	amount := util.RandomBalannce()
	sourceAccount := CreateRandomAccount(t)
	targetAccount := CreateRandomAccount(t)
	CreateRandomTranfser(t, amount, sourceAccount.ID, targetAccount.ID)
}

func TestGetTransfer(t *testing.T) {
	amount := util.RandomBalannce()
	sourceAccount := CreateRandomAccount(t)
	targetAccount := CreateRandomAccount(t)
	transfer := CreateRandomTranfser(t, amount, sourceAccount.ID, targetAccount.ID)
	actualTransfer, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, actualTransfer)
	require.Equal(t, actualTransfer.SourceAccountID, transfer.SourceAccountID)
	require.Equal(t, actualTransfer.TargetAccountID, transfer.TargetAccountID)
	require.Equal(t, actualTransfer.Amount, transfer.Amount)
	require.Equal(t, actualTransfer.ID, transfer.ID)
	require.Equal(t, actualTransfer.CreatedAt, transfer.CreatedAt)
}

func TestListTransfers(t *testing.T) {
	transfers := []Transfer{}
	sourceAccount := CreateRandomAccount(t)
	targetAccount := CreateRandomAccount(t)
	for i := 0; i <5; i++ {
		amount := util.RandomBalannce()
		transfers = append(transfers, CreateRandomTranfser(t, amount, sourceAccount.ID, targetAccount.ID))
	}
	arg := ListTransfersParams{
		SourceAccountID: sourceAccount.ID,
		TargetAccountID: targetAccount.ID,
		Limit: 5,
		Offset: 0,
	}
	actualTransfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, actualTransfers)
	require.Len(t, actualTransfers, 5)
	for _, transfer := range actualTransfers {
		require.NotEmpty(t, transfer)
		require.Equal(t, transfer.SourceAccountID, sourceAccount.ID)
		require.Equal(t, transfer.TargetAccountID, targetAccount.ID)
	}
}