// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: subproduct.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createSubproduct = `-- name: CreateSubproduct :one
INSERT INTO subproduct (product_id, variation, stock_amount, price, sale_price)
VALUES ($1, $2, $3, $4, $5) 
RETURNING id, product_id, variation, stock_amount, price, sale_price, created_at, updated_at
`

type CreateSubproductParams struct {
	ProductID   uuid.UUID     `json:"product_id"`
	Variation   string        `json:"variation"`
	StockAmount int32         `json:"stock_amount"`
	Price       int32         `json:"price"`
	SalePrice   sql.NullInt32 `json:"sale_price"`
}

func (q *Queries) CreateSubproduct(ctx context.Context, arg CreateSubproductParams) (Subproduct, error) {
	row := q.db.QueryRowContext(ctx, createSubproduct,
		arg.ProductID,
		arg.Variation,
		arg.StockAmount,
		arg.Price,
		arg.SalePrice,
	)
	var i Subproduct
	err := row.Scan(
		&i.ID,
		&i.ProductID,
		&i.Variation,
		&i.StockAmount,
		&i.Price,
		&i.SalePrice,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteSubproduct = `-- name: DeleteSubproduct :exec
DELETE FROM subproduct
WHERE id = $1
`

func (q *Queries) DeleteSubproduct(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteSubproduct, id)
	return err
}

const getSubproduct = `-- name: GetSubproduct :one
SELECT id, product_id, variation, stock_amount, price, sale_price, created_at, updated_at FROM subproduct
WHERE id = $1
`

func (q *Queries) GetSubproduct(ctx context.Context, id uuid.UUID) (Subproduct, error) {
	row := q.db.QueryRowContext(ctx, getSubproduct, id)
	var i Subproduct
	err := row.Scan(
		&i.ID,
		&i.ProductID,
		&i.Variation,
		&i.StockAmount,
		&i.Price,
		&i.SalePrice,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listSubproductsByProductID = `-- name: ListSubproductsByProductID :many
SELECT id, product_id, variation, stock_amount, price, sale_price, created_at, updated_at FROM subproduct
WHERE product_id = $1
ORDER BY price
`

func (q *Queries) ListSubproductsByProductID(ctx context.Context, productID uuid.UUID) ([]Subproduct, error) {
	rows, err := q.db.QueryContext(ctx, listSubproductsByProductID, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Subproduct{}
	for rows.Next() {
		var i Subproduct
		if err := rows.Scan(
			&i.ID,
			&i.ProductID,
			&i.Variation,
			&i.StockAmount,
			&i.Price,
			&i.SalePrice,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateSubproduct = `-- name: UpdateSubproduct :one
UPDATE subproduct
SET variation = $2, stock_amount = $3, price = $4, sale_price = $5
WHERE id = $1
RETURNING id, product_id, variation, stock_amount, price, sale_price, created_at, updated_at
`

type UpdateSubproductParams struct {
	ID          uuid.UUID     `json:"id"`
	Variation   string        `json:"variation"`
	StockAmount int32         `json:"stock_amount"`
	Price       int32         `json:"price"`
	SalePrice   sql.NullInt32 `json:"sale_price"`
}

func (q *Queries) UpdateSubproduct(ctx context.Context, arg UpdateSubproductParams) (Subproduct, error) {
	row := q.db.QueryRowContext(ctx, updateSubproduct,
		arg.ID,
		arg.Variation,
		arg.StockAmount,
		arg.Price,
		arg.SalePrice,
	)
	var i Subproduct
	err := row.Scan(
		&i.ID,
		&i.ProductID,
		&i.Variation,
		&i.StockAmount,
		&i.Price,
		&i.SalePrice,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}