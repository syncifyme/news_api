package server

import (
	"github.com/gin-gonic/gin"
)

// Start server at 8080
func Start() {
	router := gin.Default()

	v1 := router.Group("v1")
	v1.GET("/news", NewsControllerV1)

	router.Run(":8080")
}
