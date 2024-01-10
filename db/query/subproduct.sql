-- name: CreateSubproduct :one
INSERT INTO subproduct (product_id, variation, stock_amount, price, sale_price)
VALUES ($1, $2, $3, $4, $5) 
RETURNING *;

-- name: GetSubproduct :one
SELECT * FROM subproduct
WHERE id = $1;

-- name: ListSubproductsByProductID :many
SELECT * FROM subproduct
WHERE product_id = $1
ORDER BY price;

-- name: UpdateSubproduct :one
UPDATE subproduct
SET variation = $2, stock_amount = $3, price = $4, sale_price = $5
WHERE id = $1
RETURNING *;

-- name: DeleteSubproduct :exec
DELETE FROM subproduct
WHERE id = $1;

