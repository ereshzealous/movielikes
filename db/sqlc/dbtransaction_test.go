package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTransaction(t *testing.T) {
	store := NewStore(testDB)
	
	sourceAccount := CreateRandomAccount(t)
	targetAccount := CreateRandomAccount(t)

	fmt.Println(">> before:", sourceAccount.Balance, targetAccount.Balance)

	n := 5
	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferTxResult)

	// run n concurrent transfer transaction
	for i := 0; i < n; i++ {
		txName := fmt.Sprintf("TX %d", i + 1)
		go func() {
			ctx := context.WithValue(context.Background(), txKey, txName)
			result, err := store.transferTx(ctx, AccountTransferTxParams{
				SourceAccountID: 	sourceAccount.ID,
				TargetAccountID:   	targetAccount.ID,
				Amount:        		amount,
			})

			errs <- err
			results <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// check transfer
		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, sourceAccount.ID, transfer.SourceAccountID)
		require.Equal(t, targetAccount.ID, transfer.TargetAccountID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		// check entries
		fromEntry := result.SourceTransaction
		require.NotEmpty(t, fromEntry)
		require.Equal(t, sourceAccount.ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)

		_, err = store.GetTransaction(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		toEntry := result.TargetTransaction
		require.NotEmpty(t, toEntry)
		require.Equal(t, targetAccount.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)

		_, err = store.GetTransaction(context.Background(), toEntry.ID)
		require.NoError(t, err)

		// check accounts
		sourceAccountComputed := result.SourceAccount
		require.NotEmpty(t, sourceAccountComputed)
		require.Equal(t, sourceAccount.ID, sourceAccountComputed.ID)

		targetAccountComputed := result.TargetAccount
		require.NotEmpty(t, targetAccountComputed)
		require.Equal(t, targetAccount.ID, targetAccountComputed.ID)

		// check balances
		fmt.Println(">> tx:", sourceAccountComputed.Balance, targetAccountComputed.Balance)

		diff1 := sourceAccount.Balance - sourceAccountComputed.Balance
		diff2 := targetAccountComputed.Balance - targetAccount.Balance
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff1%amount == 0) // 1 * amount, 2 * amount, 3 * amount, ..., n * amount
	}

	sourceAccountResult, err := store.GetAccountDetails(context.Background(), sourceAccount.ID)
	require.NoError(t, err)

	targetAccountResult, err := store.GetAccountDetails(context.Background(), targetAccount.ID)
	require.NoError(t, err)

	fmt.Println(">> after:", sourceAccountResult.Balance, targetAccountResult.Balance)

	require.Equal(t, sourceAccount.Balance-int64(n)*amount, sourceAccountResult.Balance)
	require.Equal(t, targetAccount.Balance+int64(n)*amount, targetAccountResult.Balance)
}

func TestTransferTxDeadlock(t *testing.T) {
	store := NewStore(testDB)

	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)
	fmt.Println(">> before:", account1.Balance, account2.Balance)

	n := 10
	amount := int64(10)
	errs := make(chan error)

	for i := 0; i < n; i++ {
		fromAccountID := account1.ID
		toAccountID := account2.ID

		if i%2 == 1 {
			fromAccountID = account2.ID
			toAccountID = account1.ID
		}

		go func() {
			_, err := store.transferTx(context.Background(), AccountTransferTxParams{
				SourceAccountID: fromAccountID,
				TargetAccountID:   toAccountID,
				Amount:        amount,
			})

			errs <- err
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
	}

	// check the final updated balance
	updatedAccount1, err := store.GetAccountDetails(context.Background(), account1.ID)
	require.NoError(t, err)

	updatedAccount2, err := store.GetAccountDetails(context.Background(), account2.ID)
	require.NoError(t, err)

	fmt.Println(">> after:", updatedAccount1.Balance, updatedAccount2.Balance)
	require.Equal(t, account1.Balance, updatedAccount1.Balance)
	require.Equal(t, account2.Balance, updatedAccount2.Balance)
}
