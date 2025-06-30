-- name: ListMessages :many
SELECT * FROM messages ORDER BY id DESC;

-- name: CreateMessage :one
INSERT INTO messages (room_id, sender_id, image_url, tree_path, level, parent_id, content)
VALUES ($1, $2, $3, $4, $5, $6,$7)
    RETURNING id;