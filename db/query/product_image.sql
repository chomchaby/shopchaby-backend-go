-- name: CreateProductImage :one
INSERT INTO product_image (product_id, title, pic_url, is_default)
VALUES ($1, $2, $3, $4) 
RETURNING *;

-- name: GetProductImage :one
SELECT * FROM product_image
WHERE id = $1;

-- name: ListProductImages :many
SELECT * FROM product_image
ORDER BY product_id;

-- name: ListProductImagesByProductID :many
SELECT * FROM product_image
WHERE product_id = $1
ORDER BY is_default DESC;


-- name: UpdateProductImage :one
UPDATE product_image
SET product_id = $2, title = $3, pic_url= $4
WHERE id = $1
RETURNING *;

-- name: DeleteProductImage :exec
DELETE FROM product_image
WHERE id = $1;

-- name: DeleteProductImagesByProductID :exec
DELETE FROM product_image
WHERE product_id = $1;