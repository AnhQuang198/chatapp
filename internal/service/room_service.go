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

type RoomService interface {
	CreateRoom(ctx context.Context, roomDTO dto.RoomDTO, createUserId int64) error
}

type roomService struct {
	repo repository.RoomRepository
}

func NewRoomService(repo repository.RoomRepository) RoomService {
	return &roomService{repo: repo}
}

func (m *roomService) CreateRoom(ctx context.Context, roomDTO dto.RoomDTO, createUserId int64) error {
	userIds := roomDTO.UserIds + "," + strconv.FormatInt(createUserId, 10)
	//TODO: validate info DTO after save
	arg := models.CreateRoomParams{
		RoomName: utils.ToNullString(roomDTO.RoomName),
		UserIds:  utils.ToNullString(userIds),
		IsGroup:  utils.ToNullBool(roomDTO.IsGroup),
	}
	if err := m.repo.Create(ctx, arg); err != nil {
		return fmt.Errorf("create new message: %w", err)
	}
	return nil
}
