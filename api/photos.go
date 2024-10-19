package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-gallery/repository_drawer"
	"github.com/mahdi-cpp/api-go-gallery/repository_photos"
)

func AddPhotosRoutes(rg *gin.RouterGroup) {

	route := rg.Group("/photos")

	route.GET("/gallery", func(context *gin.Context) {
		context.JSON(210, repository_photos.RestPhotosV2())
	})

	route.GET("/people", func(context *gin.Context) {
		context.JSON(210, repository_drawer.RestPeople())
	})

	route.GET("/music", func(context *gin.Context) {
		context.JSON(210, repository_drawer.RestMusic())
	})
}
