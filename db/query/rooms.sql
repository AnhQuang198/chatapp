-- name: CreateRoom :one
INSERT INTO rooms (room_name, user_ids, is_group)
VALUES ($1, $2, $3)
    RETURNING id;

-- name: ListRooms :many
SELECT * FROM rooms;

-- name: RemoveRoom :exec
DELETE FROM rooms WHERE id = $1;