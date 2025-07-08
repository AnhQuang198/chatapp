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
	GetAllUsers(ctx context.Context) ([]dto.UserDTO, error)
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
	if _, err := m.repo.Create(ctx, arg); err != nil {
		return fmt.Errorf("create new message: %w", err)
	}
	return nil
}

func (m *userService) GetAllUsers(ctx context.Context) ([]dto.UserDTO, error) {
	lstUsers, err := m.repo.GetAllUsers(ctx)

	if err != nil {
		return nil, fmt.Errorf("get all users: %w", err)
	}

	var results []dto.UserDTO
	for _, user := range lstUsers {
		userDTO := dto.ConvertUserToDTO(user)
		results = append(results, userDTO)
	}
	return results, nil
}
