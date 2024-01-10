-- name: CreateFavorite :one
INSERT INTO favorite (user_email, product_id)
VALUES ($1, $2) 
RETURNING *;

-- name: GetFavorite :one
SELECT * FROM favorite
WHERE user_email = $1 AND product_id = $2;

-- name: ListFavorites :many
SELECT * FROM favorite
ORDER BY product_id, timestamp;

-- name: ListFavoritesByUser :many
SELECT * FROM favorite
WHERE user_email = $1
ORDER BY timestamp;

-- name: DeleteFavorite :exec
DELETE FROM favorite
WHERE user_email = $1 AND product_id = $2;