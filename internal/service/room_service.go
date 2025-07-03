package service

import (
	"chatapp/internal/dto"
	"chatapp/internal/repository"
	"chatapp/models"
	"chatapp/utils"
	"context"
	"fmt"
)

type RoomService interface {
	CreateRoom(ctx context.Context, roomDTO dto.RoomDTO, createUserId int64) error
	GetRoomsByUserId(ctx context.Context, userId int64) ([]dto.RoomDTO, error)
}

type roomService struct {
	repo repository.RoomRepository
}

func NewRoomService(repo repository.RoomRepository) RoomService {
	return &roomService{repo: repo}
}

func (m *roomService) CreateRoom(ctx context.Context, roomDTO dto.RoomDTO, createUserId int64) error {
	userIds := append(roomDTO.UserIds, createUserId)
	//TODO: validate info DTO after save
	arg := models.CreateRoomParams{
		RoomName: utils.ToNullString(roomDTO.RoomName),
		UserIds:  userIds,
		IsGroup:  utils.ToNullBool(roomDTO.IsGroup),
	}
	if err := m.repo.Create(ctx, arg); err != nil {
		return fmt.Errorf("create new room: %w", err)
	}
	return nil
}

func (m *roomService) GetRoomsByUserId(ctx context.Context, userId int64) ([]dto.RoomDTO, error) {
	roomEntity, err := m.repo.ListRoomsByUserId(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("list rooms: %w", err)
	}

	var results []dto.RoomDTO
	for _, room := range roomEntity {
		roomDTO := dto.ConvertToDTO(room)
		results = append(results, roomDTO)
	}
	return results, nil
}
