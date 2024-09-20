package router

import (
	"user-management-api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/user/:id", handlers.GetUser)
	r.POST("/user", handlers.CreateUser)
	r.DELETE("/user/:id", handlers.DeleteUser)

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "Gin is healthy"})
	})

	return r
}
