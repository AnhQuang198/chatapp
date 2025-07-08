package repository

import (
	"chatapp/models"
	"context"
	"database/sql"
)

type UserRepository interface {
	BaseRepository[models.User, models.CreateUserParams]
	GetAllUsers(ctx context.Context) ([]models.User, error)
}

type userRepository struct {
	queries *models.Queries
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{queries: models.New(db)}
}

func (m *userRepository) Create(ctx context.Context, user models.CreateUserParams) (int64, error) {
	return m.queries.CreateUser(ctx, user)
}

func (m *userRepository) Delete(ctx context.Context, userId int64) error {
	return m.queries.DeleteUser(ctx, userId)
}

func (m *userRepository) GetById(ctx context.Context, userId int64) (*models.User, error) {
	userData, err := m.queries.GetUserById(ctx, userId)
	if err != nil {
		return nil, err
	}
	return &userData, nil
}

func (m *userRepository) GetByIds(ctx context.Context, ids []int64) ([]models.User, error) {
	return m.queries.GetUserByIds(ctx, ids)
}

func (m *userRepository) GetAllUsers(ctx context.Context) ([]models.User, error) {
	return m.queries.ListUsers(ctx)
}
