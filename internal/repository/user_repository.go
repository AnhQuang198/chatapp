package repository

import (
	"chatapp/models"
	"context"
	"database/sql"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user models.CreateUserParams) error
}

type userRepository struct {
	queries *models.Queries
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{queries: models.New(db)}
}

func (m *userRepository) CreateUser(ctx context.Context, user models.CreateUserParams) error {
	if _, err := m.queries.CreateUser(ctx, user); err != nil {
		return err
	}
	return nil
}
