package dto

type CreateMessageDTO struct {
	RoomID   int32  `json:"room_id"`
	SenderID int32  `json:"sender_id"`
	ImageUrl string `json:"image_url"`
	Content  string `json:"content"`
}
