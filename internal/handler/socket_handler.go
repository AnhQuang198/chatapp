package handler

import (
	websocket2 "chatapp/internal/service/websocket"
	"context"
	"database/sql"
	"github.com/coder/websocket"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func JoinRoom(db *sql.DB, hub *websocket2.Hub) func(*gin.Context) {
	return func(c *gin.Context) {
		userId, err := strconv.ParseInt(c.Query("userId"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		roomId, err := strconv.ParseInt(c.Query("roomId"), 10, 64)
		if err != nil {
			roomId = 0
		}

		conn, err := websocket.Accept(c.Writer, c.Request, &websocket.AcceptOptions{
			InsecureSkipVerify: true, // bỏ kiểm tra origin, nên cấu hình trong production
		})

		if err != nil {
			log.Println("WebSocket accept error:", err)
			return
		}
		defer conn.Close(websocket.StatusInternalError, "internal error")

		ctx, cancel := context.WithCancel(c.Request.Context())
		defer cancel()

		client := websocket2.NewClient(conn, roomId, userId)
		hub.Join(ctx, db, roomId, client)

		done := make(chan struct{})

		go func() {
			client.ReadLoop(ctx, db, hub)
			cancel()    // hủy context để WriteLoop dừng
			close(done) // thông báo main thread dừng
		}()

		go client.WriteLoop(ctx)

		<-done
		conn.Close(websocket.StatusNormalClosure, "")
	}
}
