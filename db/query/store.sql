-- name: CreateStore :one
INSERT INTO store (user_email, name, description, address, phone, pic_url)
VALUES ($1, $2, $3, $4, $5, $6) 
RETURNING *;

-- name: GetStore :one
SELECT * FROM store
WHERE id = $1;

-- name: GetStoreByUserEmail :one
SELECT * FROM store
WHERE user_email = $1;

-- name: ListStores :many
SELECT * FROM store
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: UpdateStore :one
UPDATE store
SET name = $2, description = $3, address = $4, phone = $5, pic_url = $6
WHERE id = $1
RETURNING *;

-- name: DeleteStore :exec
DELETE FROM store 
WHERE id = $1;