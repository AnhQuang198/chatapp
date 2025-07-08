package dto

import (
	"chatapp/models"
	"chatapp/utils"
)

type CreateUserDTO struct {
	Username string `json:"username" binding:"required"`
	FullName string `json:"full_name" binding:"required"`
}

type UserDTO struct {
	Id        int64  `json:"id"`
	Username  string `json:"username"`
	FullName  string `json:"full_name"`
	AvatarUrl string `json:"avatar_url"`
}

func ConvertUserToDTO(user models.User) UserDTO {
	return UserDTO{
		Id:       user.ID,
		Username: utils.NullStringToString(user.Username),
		FullName: utils.NullStringToString(user.FullName),
		//AvatarUrl:    utils.NullStringToString(user.FullName),
	}
}
