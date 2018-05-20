package server

import (
	"github.com/gin-gonic/gin"
	"github.com/syncifyme/news_api/parser"
)

// NewsControllerV1 Handle news get
func NewsControllerV1(context *gin.Context) {
	context.JSON(200, parser.Parse())
}
