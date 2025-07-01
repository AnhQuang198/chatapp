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
	DeleteRoom(ctx context.Context, roomId int64) error
	GetRoomByIds(ctx context.Context, roomIds []int64) ([]*dto.RoomDTO, error)
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
	if err := m.repo.CreateRoom(ctx, arg); err != nil {
		return fmt.Errorf("create new message: %w", err)
	}
	return nil
}

func (m *roomService) DeleteRoom(ctx context.Context, roomId int64) error {
	return nil
}

func (m *roomService) GetRoomByIds(ctx context.Context, roomIds []int64) ([]*dto.RoomDTO, error) {
	return nil, nil
}
