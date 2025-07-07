package repository

import (
	"chatapp/models"
	"context"
	"database/sql"
)

type RoomRepository interface {
	BaseRepository[models.Room, models.CreateRoomParams]
	ListRoomsByUserId(ctx context.Context, userId int64) ([]models.Room, error)
	UpdateRoomUserId(ctx context.Context, room models.UpdateRoomUserIdParams) error
}

type roomRepository struct {
	queries *models.Queries
}

func NewRoomRepository(db *sql.DB) *roomRepository {
	return &roomRepository{queries: models.New(db)}
}

func (m *roomRepository) Create(ctx context.Context, room models.CreateRoomParams) (int64, error) {
	return m.queries.CreateRoom(ctx, room)
}

func (m *roomRepository) Delete(ctx context.Context, roomId int64) error {
	return m.queries.RemoveRoom(ctx, roomId)
}

func (m *roomRepository) GetById(ctx context.Context, roomId int64) (*models.Room, error) {
	roomData, err := m.queries.GetRoomById(ctx, roomId)
	if err != nil {
		return nil, err
	}
	return &roomData, nil
}

func (m *roomRepository) GetByIds(ctx context.Context, ids []int64) ([]models.Room, error) {
	return m.queries.GetRoomByIds(ctx, ids)
}

func (m *roomRepository) ListRoomsByUserId(ctx context.Context, userId int64) ([]models.Room, error) {
	return m.queries.ListRoomsByUserId(ctx, userId)
}

func (m *roomRepository) UpdateRoomUserId(ctx context.Context, room models.UpdateRoomUserIdParams) error {
	return m.queries.UpdateRoomUserId(ctx, room)
}
