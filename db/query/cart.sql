-- name: CreateCart :one
INSERT INTO cart (user_email, subproduct_id, quantity)
VALUES ($1, $2, $3) 
RETURNING *;

-- name: GetCart :one
SELECT * FROM cart
WHERE user_email = $1 AND subproduct_id = $2;

-- name: ListCarts :many
SELECT * FROM cart
ORDER BY user_email, created_at;

-- name: ListCartsByUser :many
SELECT * FROM cart
WHERE user_email = $1
ORDER BY created_at;

-- name: UpdateCart :one
UPDATE cart
SET quantity = $3
WHERE user_email = $1 AND subproduct_id = $2
RETURNING *;

-- name: DeleteCart :exec
DELETE FROM cart
WHERE user_email = $1 AND subproduct_id = $2;