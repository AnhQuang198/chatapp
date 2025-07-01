package dto

import "time"

type RoomDTO struct {
	Id        int64      `json:"id"`
	RoomName  string     `json:"room_name"`
	UserIds   string     `json:"user_ids"`
	IsGroup   bool       `json:"is_group"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
