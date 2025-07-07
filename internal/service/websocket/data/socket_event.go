package data

import (
	"chatapp/internal/service/websocket/enum"
	"encoding/json"
)

type WebSocketEvent struct {
	Type enum.EventType  `json:"type"`
	Data json.RawMessage `json:"data"`
}

type RoomPayload struct {
	RoomId int64 `json:"room_id"`
}
