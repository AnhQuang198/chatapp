package dto

import (
	"chatapp/models"
	"chatapp/utils"
	"time"
)

type RoomDTO struct {
	Id        int64      `json:"id"`
	RoomName  string     `json:"room_name"`
	UserIds   []int64    `json:"user_ids"`
	IsGroup   bool       `json:"is_group"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func ConvertToDTO(room models.Room) RoomDTO {
	return RoomDTO{
		Id:        room.ID,
		RoomName:  utils.NullStringToString(room.RoomName),
		UserIds:   room.UserIds,
		IsGroup:   utils.NullBoolToBool(room.IsGroup),
		CreatedAt: utils.NullTimeToTime(room.CreatedAt),
		UpdatedAt: utils.NullTimeToTime(room.UpdatedAt),
	}
}
