package service

import (
	"chatapp/internal/dto"
	"chatapp/internal/repository"
	"chatapp/models"
	"chatapp/utils"
	"context"
	"fmt"
	"strconv"
)

type MessageService interface {
	SendMessage(ctx context.Context, senderId int64, msgDto dto.CreateMessageDTO) error
	GetMessageByRoomId(ctx context.Context, roomId int64) ([]dto.MessageDTO, error)
}

type messageService struct {
	repo repository.MessageRepository
}

func NewMessageService(repo repository.MessageRepository) MessageService {
	return &messageService{repo: repo}
}

func (m *messageService) SendMessage(ctx context.Context, senderId int64, msgDto dto.CreateMessageDTO) error {
	var level int32
	if msgDto.ParentId > 0 {
		if utils.IsEmpty(msgDto.TreePath) {
			level = 1
			msgDto.TreePath += strconv.FormatInt(msgDto.ParentId, 10)
		} else {
			msgDto.TreePath = utils.AppendWithSeparator(msgDto.TreePath, strconv.FormatInt(msgDto.ParentId, 10), utils.PathSeparator)
			level = utils.CountParts[int32](msgDto.TreePath, utils.PathSeparator)
		}
	}
	//TODO: validate info DTO after save
	arg := models.CreateMessageParams{
		RoomID:   msgDto.RoomId,
		SenderID: senderId,
		TreePath: utils.ToNullString(msgDto.TreePath),
		ImageUrl: utils.ToNullString(msgDto.ImageUrl),
		Level:    level,
		ParentID: utils.ToNullInt64(msgDto.ParentId),
		Content:  utils.ToNullString(msgDto.Content),
	}
	if _, err := m.repo.Create(ctx, arg); err != nil {
		return fmt.Errorf("create new message: %w", err)
	}
	return nil
}

func (m *messageService) GetMessageByRoomId(ctx context.Context, roomId int64) ([]dto.MessageDTO, error) {
	msgEntity, err := m.repo.GetMessageByRoomId(ctx, roomId)
	if err != nil {
		return nil, fmt.Errorf("get new message: %w", err)
	}
	var msgResult []dto.MessageDTO
	for _, msg := range msgEntity {
		msgDTO := dto.ConvertMessageToDTO(msg)
		msgResult = append(msgResult, msgDTO)
	}
	return msgResult, nil
}
