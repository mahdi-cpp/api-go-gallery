package repository_chat

import (
	"github.com/mahdi-cpp/api-go-gallery/cache"
	"github.com/mahdi-cpp/api-go-gallery/model"
	"github.com/mahdi-cpp/api-go-gallery/utils"
)

var electronicDTO ElectronicDTO

type ElectronicDTO struct {
	Caption     string       `json:"name"`
	Electronics []Electronic `json:"electronics"`
}

type Electronic struct {
	Name  string          `json:"name"`
	Photo model.PhotoBase `json:"photo"`
}

func GetElectronic(folder string) ElectronicDTO {

	var file = "data.txt"
	var photos = cache.ReadOfFile(folder, file)
	var count = len(photos)
	var dto ElectronicDTO

	//if count > 50 {
	//	count = 50
	//}

	var index = 0
	var nameIndex = 0

	for i := 0; i < count; i++ {
		var electronic Electronic
		if nameIndex >= len(utils.MovieNames) {
			nameIndex = 0
		}

		electronic.Name = utils.MovieNames[nameIndex]

		electronic.Photo = photos[index]
		electronic.Photo.Key = -1
		electronic.Photo.ThumbSize = 540

		dto.Electronics = append(dto.Electronics, electronic)
		nameIndex++
		index++
	}

	index = 0

	return dto
}
