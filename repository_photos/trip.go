package repository_photos

import (
	"github.com/mahdi-cpp/api-go-gallery/cache"
	"github.com/mahdi-cpp/api-go-gallery/model"
	"github.com/mahdi-cpp/api-go-gallery/utils"
)

var tripDTO TropeDTO

type TropeDTO struct {
	Trips           []Trip            `json:"trips"`
	PhotoAnimations []model.PhotoBase `json:"photoAnimations"`
}

type Trip struct {
	Name  string          `json:"name"`
	Date  string          `json:"date"`
	Photo model.PhotoBase `json:"photo"`
}

func GetTrips(folder string) {

	var file = "data.txt"
	var photos = cache.ReadOfFile(folder, file)
	var count = len(photos)

	if count > 50 {
		count = 50
	}

	var index = 0
	var nameIndex = 0

	var marginX = dp(35)
	var screenWidthPhotosCount float32 = 1.55
	var photoSize = (1080 - (marginX * (screenWidthPhotosCount + 1))) / screenWidthPhotosCount

	for i := 0; i < count; i++ {
		var trip = Trip{}

		if nameIndex >= len(utils.FackTrips) {
			nameIndex = 0
		}

		trip.Name = utils.FackTrips[nameIndex]
		trip.Photo = photos[index]
		trip.Photo.ThumbSize = 540
		trip.Photo.Crop = 1
		trip.Photo.Round = 15
		trip.Photo.Key = -1
		trip.Photo.PaintWidth = photoSize
		trip.Photo.PaintHeight = photoSize
		tripDTO.Trips = append(tripDTO.Trips, trip)

		nameIndex++
		index++
	}

	index = 0

	for i := 0; i < 10; i++ {
		var photoBase model.PhotoBase
		photoBase = photos[i]
		photoBase.ThumbSize = 540
		photoBase.Crop = 1
		photoBase.Key = -1
		photoBase.PaintWidth = photoSize
		photoBase.PaintHeight = photoSize
		tripDTO.PhotoAnimations = append(tripDTO.PhotoAnimations, photoBase)
	}
}
