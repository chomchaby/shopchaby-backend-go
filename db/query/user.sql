-- name: CreateUser :one
INSERT INTO "user" (email, username, pwd_hash, phone, address)
VALUES ($1, $2, $3, $4, $5) 
RETURNING *;

-- name: GetUser :one
SELECT * FROM "user"
WHERE email = $1;

-- name: ListUsers :many
SELECT * FROM "user"
ORDER BY email;

-- name: UpdateUser :one
UPDATE "user" 
SET username = $2, pwd_hash = $3, phone = $4, address = $5, balance = $6
WHERE email = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM "user" 
WHERE email = $1;