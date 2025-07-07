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
	CreateRoom(ctx context.Context, roomDTO dto.RoomDTO, createUserId int64) (int64, error)
	GetRoomsByUserId(ctx context.Context, userId int64) ([]dto.RoomDTO, error)
	GetRoomById(ctx context.Context, roomId int64) (dto.RoomDTO, error)
	UpdateRoomUserId(ctx context.Context, roomDTO dto.RoomDTO) error
}

type roomService struct {
	repo repository.RoomRepository
}

func NewRoomService(repo repository.RoomRepository) RoomService {
	return &roomService{repo: repo}
}

func (m *roomService) CreateRoom(ctx context.Context, roomDTO dto.RoomDTO, createUserId int64) (int64, error) {
	userIds := append(roomDTO.UserIds, createUserId)
	//TODO: validate info DTO after save
	arg := models.CreateRoomParams{
		RoomName: utils.ToNullString(roomDTO.RoomName),
		UserIds:  userIds,
		IsGroup:  utils.ToNullBool(roomDTO.IsGroup),
	}
	id, err := m.repo.Create(ctx, arg)
	if err != nil {
		return 0, fmt.Errorf("create new room: %w", err)
	}
	return id, nil
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

func (m *roomService) GetRoomById(ctx context.Context, roomId int64) (dto.RoomDTO, error) {
	roomEntity, err := m.repo.GetById(ctx, roomId)
	if err != nil {
		return dto.RoomDTO{}, fmt.Errorf("get room by id: %w", err)
	}

	return dto.ConvertToDTO(*roomEntity), nil
}

func (m *roomService) UpdateRoomUserId(ctx context.Context, roomDTO dto.RoomDTO) error {
	if roomDTO.Id <= 0 {
		return fmt.Errorf("room id is nil")
	}
	roomEntity, err := m.repo.GetById(ctx, roomDTO.Id)
	if err != nil {
		return fmt.Errorf("room is not existed (roomId): %w", err)
	}
	roomEntity.UserIds = roomDTO.UserIds

	updateParams := models.UpdateRoomUserIdParams{
		ID:      roomEntity.ID,
		UserIds: roomEntity.UserIds,
	}
	if err := m.repo.UpdateRoomUserId(ctx, updateParams); err != nil {
		return fmt.Errorf("update room user: %w", err)
	}
	return nil
}
