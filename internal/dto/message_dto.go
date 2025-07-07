package dto

import (
	"chatapp/models"
	"chatapp/utils"
	"time"
)

type CreateMessageDTO struct {
	RoomId     int64  `json:"room_id"`
	SenderId   int64  `json:"sender_id"`
	ReceiverId int64  `json:"receiver_id"`
	ImageUrl   string `json:"image_url"`
	ParentId   int64  `json:"parent_id"`
	TreePath   string `json:"tree_path"`
	Content    string `json:"content"`
}

type MessageDTO struct {
	Id          int64      `json:"id"`
	SenderId    int64      `json:"sender_id"`
	ReceiverId  int64      `json:"receiver_id"`
	RoomId      int64      `json:"room_id"`
	IsSendGroup bool       `json:"is_send_group"`
	Content     string     `json:"content"`
	TreePath    string     `json:"tree_path"`
	ImageUrl    string     `json:"image_url"`
	ParentId    int64      `json:"parent_id"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

func ConvertMessageToDTO(msg models.Message) MessageDTO {
	return MessageDTO{
		Id:        msg.ID,
		SenderId:  msg.SenderID,
		RoomId:    msg.RoomID,
		ParentId:  utils.NullIntToInt64(msg.ParentID, 0),
		ImageUrl:  utils.NullStringToString(msg.ImageUrl),
		TreePath:  utils.NullStringToString(msg.TreePath),
		Content:   utils.NullStringToString(msg.Content),
		CreatedAt: utils.NullTimeToTime(msg.CreatedAt),
		UpdatedAt: utils.NullTimeToTime(msg.UpdatedAt),
	}
}

func MapMessageToCreateDTO(msg MessageDTO) CreateMessageDTO {
	return CreateMessageDTO{
		RoomId:     msg.RoomId,
		SenderId:   msg.SenderId,
		ReceiverId: msg.ReceiverId,
		ParentId:   msg.ParentId,
		ImageUrl:   msg.ImageUrl,
		TreePath:   msg.TreePath,
		Content:    msg.Content,
	}
}
