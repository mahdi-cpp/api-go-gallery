package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-gallery/repository_photos"
	"net/http"
)

func AddPhotosHomeRoutes(rg *gin.RouterGroup) {

	route := rg.Group("/photos")

	route.GET("/layout", func(context *gin.Context) {
		context.JSON(http.StatusOK, repository_photos.RestLayout())
	})

	route.GET("/gallery", func(context *gin.Context) {
		context.JSON(http.StatusOK, repository_photos.RestGallery())
	})
	route.GET("/years", func(context *gin.Context) {
		context.JSON(http.StatusOK, repository_photos.RestYears())
	})
	route.GET("/lion", func(context *gin.Context) {
		context.JSON(http.StatusOK, repository_photos.RestLion())
	})
	route.GET("/people", func(context *gin.Context) {
		context.JSON(http.StatusOK, repository_photos.RestPeople())
	})
	route.GET("/recently", func(context *gin.Context) {
		context.JSON(http.StatusOK, repository_photos.RestRecently())
	})
	route.GET("/trip", func(context *gin.Context) {
		context.JSON(http.StatusOK, repository_photos.RestTrip())
	})
	route.GET("/pinned_collections", func(context *gin.Context) {
		context.JSON(http.StatusOK, repository_photos.RestPinnedCollection())
	})
	route.GET("/albums", func(context *gin.Context) {
		context.JSON(http.StatusOK, repository_photos.RestAlbums())
	})
	route.GET("/share_albums", func(context *gin.Context) {
		context.JSON(http.StatusOK, repository_photos.RestShareAlbums())
	})
	route.GET("/cameras", func(context *gin.Context) {
		context.JSON(http.StatusOK, repository_photos.RestCameraDTO())
	})
}
