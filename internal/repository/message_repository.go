package repository

import (
	"chatapp/models"
	"context"
	"database/sql"
)

type MessageRepository interface {
	BaseRepository[models.Message, models.CreateMessageParams]
	GetMessageByRoomId(ctx context.Context, roomId int64) ([]models.Message, error)
}

type messageRepository struct {
	queries *models.Queries
}

func NewMessageRepository(db *sql.DB) *messageRepository {
	return &messageRepository{queries: models.New(db)}
}

func (m *messageRepository) Create(ctx context.Context, msg models.CreateMessageParams) error {
	if _, err := m.queries.CreateMessage(ctx, msg); err != nil {
		return err
	}
	return nil
}

func (m *messageRepository) Delete(ctx context.Context, msgId int64) error {
	return m.queries.DeleteMessage(ctx, msgId)
}

func (m *messageRepository) GetById(ctx context.Context, msgId int64) (*models.Message, error) {
	msgData, err := m.queries.GetMessageById(ctx, msgId)
	if err != nil {
		return nil, err
	}
	return &msgData, nil
}

func (m *messageRepository) GetByIds(ctx context.Context, ids []int64) ([]models.Message, error) {
	return m.queries.GetMessageByIds(ctx, ids)
}

func (m *messageRepository) GetMessageByRoomId(ctx context.Context, roomId int64) ([]models.Message, error) {
	return m.queries.GetMessageByRoomId(ctx, roomId)
}
