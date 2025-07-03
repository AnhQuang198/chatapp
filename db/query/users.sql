-- name: ListUsers :many
SELECT * FROM users ORDER BY id DESC;

-- name: CreateUser :one
INSERT INTO users (username, full_name)
VALUES ($1, $2)
    RETURNING id;

-- name: GetUserById :one
SELECT * FROM users u WHERE u.id = $1;

-- name: GetUserByIds :many
SELECT * FROM users u WHERE u.id = ANY($1::bigint[]);

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;