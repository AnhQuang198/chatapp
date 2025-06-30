-- name: ListUsers :many
SELECT * FROM users ORDER BY id DESC;

-- name: CreateUser :one
INSERT INTO users (username, full_name)
VALUES ($1, $2)
    RETURNING id;