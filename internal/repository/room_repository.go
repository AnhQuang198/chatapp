package repository

import (
	"chatapp/models"
	"context"
	"database/sql"
)

type RoomRepository interface {
	CreateRoom(ctx context.Context, room models.CreateRoomParams) error
	DeleteRoom(ctx context.Context, roomId int64) error
	GetRoomById(ctx context.Context, roomId int64) (*models.Room, error)
}

type roomRepository struct {
	queries *models.Queries
}

func NewRoomRepository(db *sql.DB) *roomRepository {
	return &roomRepository{queries: models.New(db)}
}

func (m *roomRepository) CreateRoom(ctx context.Context, room models.CreateRoomParams) error {
	if _, err := m.queries.CreateRoom(ctx, room); err != nil {
		return err
	}
	return nil
}

func (m *roomRepository) DeleteRoom(ctx context.Context, roomId int64) error {
	if err := m.queries.RemoveRoom(ctx, roomId); err != nil {
		return err
	}
	return nil
}

func (m *roomRepository) GetRoomById(ctx context.Context, roomId int64) (*models.Room, error) {
	roomData, err := m.queries.GetRoomById(ctx, roomId)
	if err != nil {
		return nil, err
	}
	return &roomData, nil
}
