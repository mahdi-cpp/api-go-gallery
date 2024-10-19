package repository_photos

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-gallery/model"
	"github.com/mahdi-cpp/api-go-gallery/utils"
	"math"
	"sync"
)

var root = "/"

func RestPhotosV2() map[string]any {
	return gin.H{
		"galleryDTO":          galleryDTO,
		"lionDTO":             lionDTO,
		"recentlyDTO":         recentlyDTO,
		"peopleDTO":           peopleDTO,
		"tripDTO":             tripDTO,
		"pinnedCollectionDTO": pinnedCollectionDTO,
		"albums":              albums,
		"shareAlbums":         shareAlbums,
		"cameraDTO":           cameraDTO,
	}
}

func InitPhotos() {

	GetGalleries("/var/cloud/fa/")
	GetLion("/var/cloud/lion/")

	GetRecently("/var/cloud/bb/")
	GetPeoples("/var/cloud/id/face/")
	GetTrips("/var/cloud/id/trip/")
	GetPinned("/var/cloud/bb/")
	albums = GetAlbums("/var/cloud/bb/")
	shareAlbums = GetAlbums("/var/cloud/id/go/")
	cameraDTO = GetCameras("/var/cloud/id/go/")

	utils.GetCities()
	utils.GetNames()
}

type PhotoBaseCache struct {
	sync.RWMutex
	Cache map[int]model.PhotoBase
}

func dp(value float32) float32 {
	if value == 0 {
		return 0
	}
	return float32(math.Ceil(float64(2.625 * value)))
}
