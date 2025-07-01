package repository

import (
	"chatapp/models"
	"context"
	"database/sql"
)

type MessageRepository interface {
	CreateMessage(ctx context.Context, msg models.CreateMessageParams) error
}

type messageRepository struct {
	queries *models.Queries
}

func NewMessageRepository(db *sql.DB) *messageRepository {
	return &messageRepository{queries: models.New(db)}
}

func (m *messageRepository) CreateMessage(ctx context.Context, msg models.CreateMessageParams) error {
	if _, err := m.queries.CreateMessage(ctx, msg); err != nil {
		return err
	}
	return nil
}
