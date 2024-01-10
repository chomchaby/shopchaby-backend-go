package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all functions to execute db queries and transactions
type StoreTx interface {
	Querier
	CreateProductListingTx(ctx context.Context, arg CreateProductListingTxParams) (ProductListingTxResult, error)
	UpdateProductListingTx(ctx context.Context, arg UpdateProductListingTxParams) (ProductListingTxResult, error)
}

// Store provides all functions to execute db queries and transactions
type SQLStoreTx struct {
	db *sql.DB
	*Queries
}

// NewStore creates a new Store
func NewStoreTx(db *sql.DB) StoreTx {
	return &SQLStoreTx{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a function within a database transaction
func (storeTx *SQLStoreTx) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := storeTx.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}
