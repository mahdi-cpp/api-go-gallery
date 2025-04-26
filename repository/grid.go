package repository

import (
	"github.com/mahdi-cpp/api-go-gallery/cache"
	"github.com/mahdi-cpp/api-go-gallery/model"
	"github.com/mahdi-cpp/api-go-gallery/utils"
)

var gridDTO GridDTO

type GridDTO struct {
	TotalItems int               `json:"count"`
	Name       string            `json:"name"`
	Photos     []model.PhotoBase `json:"photos"`
}

func GetGrid(folder string) {

	var file = "data.txt"
	var photos = cache.ReadOfFile(folder, file)
	gridDTO.TotalItems = len(photos)

	var index = 0
	var nameIndex = 0

	gridDTO.Name = utils.FackNames[nameIndex]

	for i := 0; i < gridDTO.TotalItems; i++ {
		var photoBase model.PhotoBase
		photoBase = photos[index]
		photoBase.ThumbSize = 270
		photoBase.Crop = 1
		photoBase.Round = int(dp(10))
		photoBase.Key = -1
		photoBase.VideoFormat = "mp4"
		photoBase.IsVideo = true
		gridDTO.Photos = append(gridDTO.Photos, photoBase)
		index++
	}
}
