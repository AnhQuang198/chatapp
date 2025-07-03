package handler

import (
	"chatapp/internal/dto"
	"chatapp/internal/repository"
	"chatapp/internal/service"
	"chatapp/utils"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateMessage(db *sql.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		userId, err := strconv.ParseInt(c.Param("userId"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var msgDto dto.CreateMessageDTO
		if err := c.ShouldBindJSON(&msgDto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		repo := repository.NewMessageRepository(db)
		messageService := service.NewMessageService(repo)

		if err := messageService.SendMessage(c.Request.Context(), userId, msgDto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, utils.SimpleSuccessResponse(nil))
	}
}

func GetMessageByRoomId(db *sql.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		roomId, err := strconv.ParseInt(c.Param("roomId"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		repo := repository.NewMessageRepository(db)
		messageService := service.NewMessageService(repo)

		result, err := messageService.GetMessageByRoomId(c.Request.Context(), roomId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, utils.SimpleSuccessResponse(result))
	}
}
