package repository

import (
	"github.com/mahdi-cpp/api-go-gallery/cache"
	"github.com/mahdi-cpp/api-go-gallery/model"
	"github.com/mahdi-cpp/api-go-gallery/utils"
)

var tripDTO TropeDTO

type TropeDTO struct {
	Trips []Trip `json:"trips"`
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

	for i := 0; i < count; i++ {
		var trip = Trip{}

		if nameIndex >= len(utils.FackTrips) {
			nameIndex = 0
		}

		trip.Name = utils.FackTrips[nameIndex]
		trip.Photo = photos[index]
		trip.Photo.ThumbSize = 540
		trip.Photo.Crop = 1
		trip.Photo.Round = 20
		trip.Photo.Key = -1
		tripDTO.Trips = append(tripDTO.Trips, trip)

		nameIndex++
		index++
	}

	index = 0
}
