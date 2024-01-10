-- name: CreateProduct :one
INSERT INTO product (category, name, description, store_id, suspend)
VALUES ($1, $2, $3, $4, $5) 
RETURNING *;

-- name: GetProduct :one
SELECT * FROM product
WHERE id = $1;

-- name: ListProducts :many
SELECT * FROM product
ORDER BY category, id;

-- name: ListProductsByCategory :many
SELECT * FROM product
WHERE category = $1
ORDER BY id;

-- name: ListProductsByStoreAndCategory :many
SELECT * FROM product
WHERE store_id = $1 AND category = $2
ORDER BY name;

-- name: ListProductsByFilter :many
SELECT * FROM product
WHERE category = $1 AND ((min_price <= 32 AND min_price >= $3) OR (max_price <= $2 AND max_price >= $3))
ORDER BY onsale DESC;

-- name: ListProductsByName :many
SELECT * FROM product
WHERE name = $1 
ORDER BY onsale DESC;

-- name: UpdateProduct :one
UPDATE product
SET category = $2, name = $3, description = $4, suspend = $5
WHERE id = $1
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM product
WHERE id = $1;