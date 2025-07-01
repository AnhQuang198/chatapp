package internal

import (
	"chatapp/internal/handler"
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
	}

	return r
}
