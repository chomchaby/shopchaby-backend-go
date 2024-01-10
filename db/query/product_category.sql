-- name: CreateProductCategory :one
INSERT INTO product_category (name, pic_url)
VALUES ($1, $2) 
RETURNING *;

-- name: GetProductCategory :one
SELECT * FROM product_category
WHERE name = $1;

-- name: ListProductCategories :many
SELECT * FROM product_category
ORDER BY name;

-- name: UpdateProductCategory :one
UPDATE product_category
SET name = $2, pic_url= $3
WHERE name = $1
RETURNING *;

-- name: DeleteProductCategory :exec
DELETE FROM product_category
WHERE name = $1;