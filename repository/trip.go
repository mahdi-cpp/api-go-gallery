package repository

import (
	"github.com/mahdi-cpp/api-go-gallery/cache"
	"github.com/mahdi-cpp/api-go-gallery/model"
	"github.com/mahdi-cpp/api-go-gallery/utils"
)

var tripDTO TropeDTO
var tripBigDTO TropeDTO

type TropeDTO struct {
	Trips []Trip `json:"trips"`
}

type Trip struct {
	Name  string        `json:"name"`
	Date  string        `json:"date"`
	Image model.UIImage `json:"image"`
}

func GetTrips(folder string) {

	var file = "data.txt"
	var uiImages = cache.ReadOfFile(folder, file)
	var count = len(uiImages)

	if count > 50 {
		count = 50
	}

	var index = 0
	var nameIndex = 0

	for i := 0; i < count; i++ {
		var trip = Trip{}

		if nameIndex >= len(utils.FackTrips) {
			nameIndex = 0
		}

		trip.Name = utils.FackTrips[nameIndex]
		trip.Image = uiImages[index]
		tripDTO.Trips = append(tripDTO.Trips, trip)

		nameIndex++
		index++
	}

	index = 0
}
