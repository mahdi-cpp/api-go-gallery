package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-gallery/repository"
)

func AddGalleryRoutes(rg *gin.RouterGroup) {

	gallery := rg.Group("/gallery")

	gallery.GET("/feed", func(context *gin.Context) {

		days := repository.GetGalleryDays()
		peoples := repository.GetGalleryPeoples()
		trips := repository.GetGalleryTrips()
		playlists := repository.GetGalleryPlaylists()

		context.JSON(210, gin.H{
			"days":      days,
			"peoples":   peoples,
			"trips":     trips,
			"playlists": playlists,
		})
	})
}
