package service

import (
	"chatapp/internal/dto"
	"chatapp/internal/repository"
	"chatapp/models"
	"chatapp/utils"
	"context"
	"fmt"
)

type UserService interface {
	CreateUser(ctx context.Context, userDTO dto.CreateUserDTO) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (m *userService) CreateUser(ctx context.Context, userDTO dto.CreateUserDTO) error {
	//TODO: validate info DTO after save
	arg := models.CreateUserParams{
		Username: utils.ToNullString(userDTO.Username),
		FullName: utils.ToNullString(userDTO.FullName),
	}
	if err := m.repo.CreateUser(ctx, arg); err != nil {
		return fmt.Errorf("create new message: %w", err)
	}
	return nil
}
