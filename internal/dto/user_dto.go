package dto

type CreateUserDTO struct {
	Username string `json:"username" binding:"required"`
	FullName string `json:"full_name" binding:"required"`
}
