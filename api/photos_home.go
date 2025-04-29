package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-gallery/repository"
	"net/http"
)

func AddPhotosHomeRoutes(rg *gin.RouterGroup) {

	route := rg.Group("/photos")

	route.GET("/layout", func(context *gin.Context) {
		context.JSON(http.StatusOK, repository.RestLayout())
	})

	route.GET("/gallery", func(context *gin.Context) {
		context.JSON(http.StatusOK, repository.RestGallery())
	})
	route.GET("/years", func(context *gin.Context) {
		context.JSON(http.StatusOK, repository.RestYears())
	})
	//route.GET("/lion", func(context *gin.Context) {
	//	context.JSON(http.StatusOK, repository.RestLion())
	//})
	route.GET("/people", func(context *gin.Context) {
		context.JSON(http.StatusOK, repository.RestPeople())
	})
	route.GET("/recently", func(context *gin.Context) {
		context.JSON(http.StatusOK, repository.RestRecently())
	})
	route.GET("/trip", func(context *gin.Context) {
		context.JSON(http.StatusOK, repository.RestTrip())
	})
	route.GET("/pinned_collections", func(context *gin.Context) {
		context.JSON(http.StatusOK, repository.RestPinnedCollection())
	})
	route.GET("/albums", func(context *gin.Context) {
		context.JSON(http.StatusOK, repository.RestAlbums())
	})
	route.GET("/share_albums", func(context *gin.Context) {
		context.JSON(http.StatusOK, repository.RestShareAlbums())
	})
	route.GET("/cameras", func(context *gin.Context) {
		context.JSON(http.StatusOK, repository.RestCameraDTO())
	})
}
