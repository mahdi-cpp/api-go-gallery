package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-gallery/model"
	"github.com/mahdi-cpp/api-go-gallery/utils"
	"math"
	"sync"
)

var root = "/"

func InitPhotos() {

	LayoutInit()

	//GetGalleries("/var/cloud/id/mahan/")
	GetGalleries("/var/cloud/00-instagram/razzle/")

	GetRecently("/var/cloud/00-instagram/razzle-photo/")
	GetPeoples("/var/cloud/00-instagram/ashtonhall/")
	GetTrips("/var/cloud/id/mahan/")
	GetPinned("/var/cloud/id/trip/")

	GetYears("/var/cloud/fa/")

	albumDTO = GetAlbums("/var/cloud/id/me/")
	shareAlbumDTO = GetAlbums("/var/cloud/00-instagram/lucaspopan/")
	cameraDTO = GetCameras("/var/cloud/00-instagram/nickloveswildlife/")

	//GetGrid("/var/cloud/00-instagram/video/")

	utils.GetCities()
	utils.GetNames()
}

func RestLayout() map[string]any {
	return gin.H{
		"views": layout.Views,
	}
}

func RestGallery() map[string]any {
	return gin.H{
		"avatar": galleryDTO.Avatar,
		"photos": galleryDTO.Photos,
	}
}
func RestYears() map[string]any {
	return gin.H{
		"years": yearsDTO.Years,
	}
}

func RestPeople() map[string]any {
	return gin.H{
		"personGroups": peopleDTO.PersonGroups,
	}
}

func RestRecently() map[string]any {
	return gin.H{
		"days": recentlyDTO.Days,
	}
}
func RestTrip() map[string]any {
	return gin.H{
		"trips": tripDTO.Trips,
	}
}
func RestPinnedCollection() map[string]any {
	return gin.H{
		"pinnedCollections": pinnedCollectionDTO.PinnedCollections,
	}
}
func RestAlbums() map[string]any {
	return gin.H{"albums": albumDTO.Albums}
}
func RestShareAlbums() map[string]any {
	return gin.H{"albums": shareAlbumDTO.Albums}
}
func RestCameraDTO() map[string]any {
	return gin.H{
		"cameras": cameraDTO.Cameras,
	}
}

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func RestUser() map[string]any {

	return gin.H{
		"Id":    "12",
		"name":  "Mahdi",
		"email": "mahdi.cpp@gmail.com",
	}
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
