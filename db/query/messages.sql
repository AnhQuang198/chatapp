-- name: ListMessages :many
SELECT * FROM messages ORDER BY id DESC;

-- name: CreateMessage :one
INSERT INTO messages (room_id, sender_id, receiver_id, image_url, tree_path, level, parent_id, content)
VALUES ($1, $2, $3, $4, $5, $6,$7, $8)
    RETURNING id;

-- name: GetMessageById :one
SELECT * FROM messages m WHERE m.id = $1;

-- name: GetMessageByIds :many
SELECT * FROM messages m WHERE m.id = ANY($1::bigint[]);

-- name: DeleteMessage :exec
DELETE FROM messages WHERE id = $1;

-- name: GetMessageByRoomId :many
SELECT * FROM messages m WHERE m.room_id = $1;