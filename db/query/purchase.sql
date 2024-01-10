-- name: CreatePurchase :one
INSERT INTO purchase (buyer_email, buyer_username, seller_email, seller_username, product_id, store_id, product_name, variation, price, quantity, store_name, pic_url, shipping_address, receiver)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) 
RETURNING *;

-- name: GetPurchase :one
SELECT * FROM purchase
WHERE id = $1;

-- name: ListPurchases :many
SELECT * FROM purchase
ORDER BY timestamp;

-- name: ListPurchasesByBuyer :many
SELECT * FROM purchase
WHERE buyer_email = $1
ORDER BY timestamp;

-- name: ListPurchasesByBuyerAndStore :many
SELECT * FROM purchase
WHERE buyer_email = $1 AND store_id = $2
ORDER BY timestamp;

-- name: DeletePurchase :exec
DELETE FROM purchase
WHERE id = $1;