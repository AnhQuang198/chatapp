-- name: CreateRoom :one
INSERT INTO rooms (room_name, user_ids, is_group)
VALUES ($1, $2, $3)
    RETURNING id;

-- name: ListRooms :many
SELECT * FROM rooms;

-- name: GetRoomById :one
SELECT * FROM rooms r WHERE r.id = $1;

-- name: GetRoomByIds :many
SELECT * FROM rooms r WHERE r.id = ANY($1::bigint[]);

-- name: RemoveRoom :exec
DELETE FROM rooms WHERE id = $1;

-- name: ListRoomsByUserId :many
SELECT * FROM rooms WHERE $1::BIGINT = ANY(user_ids);

-- name: UpdateRoomUserId :exec
UPDATE rooms SET user_ids = $2, updated_at = CURRENT_TIMESTAMP WHERE id = $1;