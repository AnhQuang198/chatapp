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

func CreateRoom(db *sql.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var roomDTO dto.RoomDTO

		if err := c.ShouldBindJSON(&roomDTO); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		userId, err := strconv.ParseInt(c.Param("userId"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		repo := repository.NewRoomRepository(db)
		roomService := service.NewRoomService(repo)

		if err := roomService.CreateRoom(c.Request.Context(), roomDTO, userId); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, utils.SimpleSuccessResponse(nil))
	}
}

func GetRoomsByUserId(db *sql.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		userId, err := strconv.ParseInt(c.Param("userId"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		repo := repository.NewRoomRepository(db)
		roomService := service.NewRoomService(repo)

		result, err := roomService.GetRoomsByUserId(c.Request.Context(), userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, utils.SimpleSuccessResponse(result))
	}
}
