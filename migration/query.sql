-- name: FindUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUsers :many
SELECT * FROM users;

-- name: Login :one
SELECT * FROM users
WHERE username = $1 AND password = $2 LIMIT 1;

-- name: Register :one
INSERT INTO users (username, email, phone_number, password, created_at, updated_at, name)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: FindUserByUsername :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: FindUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: FindUserByPhoneNumber :one
SELECT * FROM users
WHERE phone_number = $1 LIMIT 1;

-- name: FindUserByUsernameOrEmailOrPhoneNumber :one
SELECT * FROM users
WHERE username = $1 OR( email = $2 and email IS NOT NULL) OR (phone_number = $3 and phone_number IS NOT NULL) LIMIT 1;

