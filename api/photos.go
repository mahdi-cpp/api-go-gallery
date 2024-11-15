package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-gallery/repository_drawer"
	"github.com/mahdi-cpp/api-go-gallery/repository_photos"
	"net/http"
	"strconv"
)

func AddPhotosRoutes(rg *gin.RouterGroup) {

	route := rg.Group("/photos")

	route.GET("/grid/:start/:end", func(context *gin.Context) {
		startIndex, _ := strconv.Atoi(context.Param("start"))
		endIndex, _ := strconv.Atoi(context.Param("end"))

		context.JSON(http.StatusOK, repository_photos.RestGrid(startIndex, endIndex))
	})

	route.GET("/people_drawer", func(context *gin.Context) {
		context.JSON(http.StatusOK, repository_drawer.RestPeople())
	})

	route.GET("/music", func(context *gin.Context) {
		context.JSON(http.StatusOK, repository_drawer.RestMusic())
	})

	route.GET("/test", func(context *gin.Context) {
		context.JSON(http.StatusOK, repository_photos.RestUser())
	})

}
