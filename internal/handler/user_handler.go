package handler

import (
	"chatapp/internal/dto"
	"chatapp/internal/repository"
	"chatapp/internal/service"
	"chatapp/utils"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(db *sql.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var userDTO dto.CreateUserDTO

		if err := c.ShouldBindJSON(&userDTO); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		userRepository := repository.NewUserRepository(db)
		userService := service.NewUserService(userRepository)

		if err := userService.CreateUser(c.Request.Context(), userDTO); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, utils.SimpleSuccessResponse(nil))
	}
}
