package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTransaction struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *DBTransaction {
	return &DBTransaction {
		db: db,
		Queries: New(db),
	}
}

func (transaction *DBTransaction) executeTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := transaction.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	query := New(tx)
	err = fn(query)
	if err != nil {
		if rbkerr := tx.Rollback(); rbkerr != nil {
			return fmt.Errorf("Transaction Error : %v, Rollback Error : %v", err, rbkerr)
		}
		return err
	}
	return tx.Commit()
}

type AccountTransferTxParams struct {
	SourceAccountID int64 `json:"source_account_id"`
	TargetAccountID int64 `json:"target_account_id"`
	Amount         	int64 `json:"amount"`
}

// TransferTxResult is the result of the transfer transaction
type TransferTxResult struct {
	Transfer    		Transfer 		`json:"transfer"`
	SourceAccount 		AccountDetail  	`json:"source_account"`
	TargetAccount   	AccountDetail  	`json:"target_account"`
	SourceTransaction   Transaction    	`json:"source_transaction"`
	TargetTransaction   Transaction 	`json:"target_transaction"`
}

var txKey = struct{}{}

func (transaction *DBTransaction) transferTx(ctx context.Context, arg AccountTransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult
	err := transaction.executeTx(ctx, func(queries *Queries) error {
		var err error

		result.Transfer, err = queries.CreateTransfer(ctx, CreateTransferParams{
			SourceAccountID: arg.SourceAccountID,
			TargetAccountID: arg.TargetAccountID,
			Amount: arg.Amount,
		})
		if err != nil {
			return err
		}
		
		result.SourceTransaction, err = queries.CreateTransaction(ctx, CreateTransactionParams {
			AccountID: arg.SourceAccountID,
			Amount: -arg.Amount,
		})
		if err != nil {
			return err
		}

		result.TargetTransaction, err = queries.CreateTransaction(ctx, CreateTransactionParams {
			AccountID: arg.TargetAccountID,
			Amount: arg.Amount,
		})
		if err != nil {
			return err
		}

		if arg.SourceAccountID < arg.TargetAccountID {
			result.SourceAccount, result.TargetAccount, err = updateBalance(ctx, queries, arg.SourceAccountID, -arg.Amount, 
				arg.TargetAccountID, arg.Amount)
		} else {
			result.TargetAccount, result.SourceAccount, err = updateBalance(ctx, queries, arg.TargetAccountID, arg.Amount, 
				arg.SourceAccountID, -arg.Amount)
		}
		
		
		return nil
	})
	return result, err
}

func updateBalance(ctx context.Context, q *Queries, accountID1 int64, amount1 int64, accountID2 int64, amount2 int64,
) (account1 AccountDetail, account2 AccountDetail, err error) {
	account1, err = q.UpdateAccountBalance(ctx, UpdateAccountBalanceParams{
		ID:     accountID1,
		Amount: amount1,
	})
	if err != nil {
		return
	}

	account2, err = q.UpdateAccountBalance(ctx, UpdateAccountBalanceParams{
		ID:     accountID2,
		Amount: amount2,
	})
	return
}