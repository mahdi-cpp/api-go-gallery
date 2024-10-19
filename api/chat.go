package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-gallery/repository_chat"
)

func AddChatRoute(rg *gin.RouterGroup) {

	route := rg.Group("/chat")

	route.GET("/chatV2", func(context *gin.Context) {
		context.JSON(210, repository_chat.RestChatV2())
	})

}