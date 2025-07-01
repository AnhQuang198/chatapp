package service

import (
	"chatapp/internal/dto"
	"chatapp/internal/repository"
	"chatapp/models"
	"chatapp/utils"
	"context"
	"fmt"
)

type MessageService interface {
	CreateUser(ctx context.Context, msgDto dto.CreateMessageDTO) error
}

type messageService struct {
	repo repository.MessageRepository
}

func NewMessageService(repo repository.MessageRepository) MessageService {
	return &messageService{repo: repo}
}

func (m *messageService) CreateUser(ctx context.Context, msgDto dto.CreateMessageDTO) error {
	//TODO: validate info DTO after save
	arg := models.CreateMessageParams{
		RoomID:   msgDto.RoomID,
		SenderID: msgDto.SenderID,
		Content:  utils.ToNullString(msgDto.Content),
	}
	if err := m.repo.CreateMessage(ctx, arg); err != nil {
		return fmt.Errorf("create new message: %w", err)
	}
	return nil
}
