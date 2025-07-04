package internal

import (
	"chatapp/internal/handler"
	"chatapp/internal/service/websocket"
	"database/sql"
	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		users := v1.Group("/users")
		{
			users.POST("/register", handler.Register(db))
		}
		rooms := v1.Group("/rooms")
		{
			rooms.POST("/:userId", handler.CreateRoom(db))
			rooms.GET("/:userId", handler.GetRoomsByUserId(db))
		}
		messages := v1.Group("/messages")
		{
			messages.GET("/:roomId", handler.GetMessageByRoomId(db))
			messages.POST("/:userId", handler.CreateMessage(db))
		}
		hub := websocket.NewHub()
		chat := r.Group("/chat")
		{
			chat.GET("", handler.JoinRoom(db, hub))
		}
	}
	return r
}
