// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: purchase.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createPurchase = `-- name: CreatePurchase :one
INSERT INTO purchase (buyer_email, buyer_username, seller_email, seller_username, product_id, store_id, product_name, variation, price, quantity, store_name, pic_url, shipping_address, receiver)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) 
RETURNING id, buyer_email, buyer_username, seller_email, seller_username, product_id, store_id, product_name, variation, price, quantity, store_name, pic_url, shipping_address, receiver, timestamp
`

type CreatePurchaseParams struct {
	BuyerEmail      string        `json:"buyer_email"`
	BuyerUsername   string        `json:"buyer_username"`
	SellerEmail     string        `json:"seller_email"`
	SellerUsername  string        `json:"seller_username"`
	ProductID       uuid.NullUUID `json:"product_id"`
	StoreID         uuid.NullUUID `json:"store_id"`
	ProductName     string        `json:"product_name"`
	Variation       string        `json:"variation"`
	Price           int32         `json:"price"`
	Quantity        int32         `json:"quantity"`
	StoreName       string        `json:"store_name"`
	PicUrl          string        `json:"pic_url"`
	ShippingAddress string        `json:"shipping_address"`
	Receiver        string        `json:"receiver"`
}

func (q *Queries) CreatePurchase(ctx context.Context, arg CreatePurchaseParams) (Purchase, error) {
	row := q.db.QueryRowContext(ctx, createPurchase,
		arg.BuyerEmail,
		arg.BuyerUsername,
		arg.SellerEmail,
		arg.SellerUsername,
		arg.ProductID,
		arg.StoreID,
		arg.ProductName,
		arg.Variation,
		arg.Price,
		arg.Quantity,
		arg.StoreName,
		arg.PicUrl,
		arg.ShippingAddress,
		arg.Receiver,
	)
	var i Purchase
	err := row.Scan(
		&i.ID,
		&i.BuyerEmail,
		&i.BuyerUsername,
		&i.SellerEmail,
		&i.SellerUsername,
		&i.ProductID,
		&i.StoreID,
		&i.ProductName,
		&i.Variation,
		&i.Price,
		&i.Quantity,
		&i.StoreName,
		&i.PicUrl,
		&i.ShippingAddress,
		&i.Receiver,
		&i.Timestamp,
	)
	return i, err
}

const deletePurchase = `-- name: DeletePurchase :exec
DELETE FROM purchase
WHERE id = $1
`

func (q *Queries) DeletePurchase(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deletePurchase, id)
	return err
}

const getPurchase = `-- name: GetPurchase :one
SELECT id, buyer_email, buyer_username, seller_email, seller_username, product_id, store_id, product_name, variation, price, quantity, store_name, pic_url, shipping_address, receiver, timestamp FROM purchase
WHERE id = $1
`

func (q *Queries) GetPurchase(ctx context.Context, id uuid.UUID) (Purchase, error) {
	row := q.db.QueryRowContext(ctx, getPurchase, id)
	var i Purchase
	err := row.Scan(
		&i.ID,
		&i.BuyerEmail,
		&i.BuyerUsername,
		&i.SellerEmail,
		&i.SellerUsername,
		&i.ProductID,
		&i.StoreID,
		&i.ProductName,
		&i.Variation,
		&i.Price,
		&i.Quantity,
		&i.StoreName,
		&i.PicUrl,
		&i.ShippingAddress,
		&i.Receiver,
		&i.Timestamp,
	)
	return i, err
}

const listPurchases = `-- name: ListPurchases :many
SELECT id, buyer_email, buyer_username, seller_email, seller_username, product_id, store_id, product_name, variation, price, quantity, store_name, pic_url, shipping_address, receiver, timestamp FROM purchase
ORDER BY timestamp
`

func (q *Queries) ListPurchases(ctx context.Context) ([]Purchase, error) {
	rows, err := q.db.QueryContext(ctx, listPurchases)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Purchase{}
	for rows.Next() {
		var i Purchase
		if err := rows.Scan(
			&i.ID,
			&i.BuyerEmail,
			&i.BuyerUsername,
			&i.SellerEmail,
			&i.SellerUsername,
			&i.ProductID,
			&i.StoreID,
			&i.ProductName,
			&i.Variation,
			&i.Price,
			&i.Quantity,
			&i.StoreName,
			&i.PicUrl,
			&i.ShippingAddress,
			&i.Receiver,
			&i.Timestamp,
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

const listPurchasesByBuyer = `-- name: ListPurchasesByBuyer :many
SELECT id, buyer_email, buyer_username, seller_email, seller_username, product_id, store_id, product_name, variation, price, quantity, store_name, pic_url, shipping_address, receiver, timestamp FROM purchase
WHERE buyer_email = $1
ORDER BY timestamp
`

func (q *Queries) ListPurchasesByBuyer(ctx context.Context, buyerEmail string) ([]Purchase, error) {
	rows, err := q.db.QueryContext(ctx, listPurchasesByBuyer, buyerEmail)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Purchase{}
	for rows.Next() {
		var i Purchase
		if err := rows.Scan(
			&i.ID,
			&i.BuyerEmail,
			&i.BuyerUsername,
			&i.SellerEmail,
			&i.SellerUsername,
			&i.ProductID,
			&i.StoreID,
			&i.ProductName,
			&i.Variation,
			&i.Price,
			&i.Quantity,
			&i.StoreName,
			&i.PicUrl,
			&i.ShippingAddress,
			&i.Receiver,
			&i.Timestamp,
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

const listPurchasesByBuyerAndStore = `-- name: ListPurchasesByBuyerAndStore :many
SELECT id, buyer_email, buyer_username, seller_email, seller_username, product_id, store_id, product_name, variation, price, quantity, store_name, pic_url, shipping_address, receiver, timestamp FROM purchase
WHERE buyer_email = $1 AND store_id = $2
ORDER BY timestamp
`

type ListPurchasesByBuyerAndStoreParams struct {
	BuyerEmail string        `json:"buyer_email"`
	StoreID    uuid.NullUUID `json:"store_id"`
}

func (q *Queries) ListPurchasesByBuyerAndStore(ctx context.Context, arg ListPurchasesByBuyerAndStoreParams) ([]Purchase, error) {
	rows, err := q.db.QueryContext(ctx, listPurchasesByBuyerAndStore, arg.BuyerEmail, arg.StoreID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Purchase{}
	for rows.Next() {
		var i Purchase
		if err := rows.Scan(
			&i.ID,
			&i.BuyerEmail,
			&i.BuyerUsername,
			&i.SellerEmail,
			&i.SellerUsername,
			&i.ProductID,
			&i.StoreID,
			&i.ProductName,
			&i.Variation,
			&i.Price,
			&i.Quantity,
			&i.StoreName,
			&i.PicUrl,
			&i.ShippingAddress,
			&i.Receiver,
			&i.Timestamp,
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
