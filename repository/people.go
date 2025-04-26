package repository

import (
	"github.com/mahdi-cpp/api-go-gallery/cache"
	"github.com/mahdi-cpp/api-go-gallery/model"
	"github.com/mahdi-cpp/api-go-gallery/utils"
)

var peopleDTO PeopleDTO

type PeopleDTO struct {
	PersonGroups    []PersonGroup     `json:"personGroups"`
	PhotoAnimations []model.PhotoBase `json:"photoAnimations"`
}

type PersonGroup struct {
	Names  []string        `json:"names"`
	Photo1 model.PhotoBase `json:"photo1"`
	Photo2 model.PhotoBase `json:"photo2"`
	Photo3 model.PhotoBase `json:"photo3"`
	Photo4 model.PhotoBase `json:"photo4"`
}

func GetPeoples(folder string) {

	var file = "data.txt"
	var photos = cache.ReadOfFile(folder, file)

	var count = (len(photos) / 4) - 4

	if count > 5 {
		count = 5
	}

	var index = 0
	var nameIndex = 0

	for i := 0; i < count; i++ {
		var personGroup = PersonGroup{}

		if nameIndex+4 >= len(utils.FackNames) {
			nameIndex = 0
		}

		personGroup.Names = append(personGroup.Names, utils.FackNames[nameIndex])
		personGroup.Names = append(personGroup.Names, utils.FackNames[nameIndex+1])
		personGroup.Names = append(personGroup.Names, utils.FackNames[nameIndex+2])
		personGroup.Names = append(personGroup.Names, utils.FackNames[nameIndex+3])

		personGroup.Photo1 = photos[index+1]
		personGroup.Photo2 = photos[index+2]
		personGroup.Photo3 = photos[index+3]
		personGroup.Photo4 = photos[index+4]

		personGroup.Photo1.ThumbSize = 270
		personGroup.Photo2.ThumbSize = 270
		personGroup.Photo3.ThumbSize = 270
		personGroup.Photo4.ThumbSize = 270

		personGroup.Photo1.Round = int(dp(10))
		personGroup.Photo2.Round = int(dp(10))
		personGroup.Photo3.Round = int(dp(10))
		personGroup.Photo4.Round = int(dp(10))

		personGroup.Photo1.Crop = 1
		personGroup.Photo2.Crop = 1
		personGroup.Photo3.Crop = 1
		personGroup.Photo4.Crop = 1

		personGroup.Photo1.Key = -1
		personGroup.Photo2.Key = -1
		personGroup.Photo3.Key = -1
		personGroup.Photo4.Key = -1

		peopleDTO.PersonGroups = append(peopleDTO.PersonGroups, personGroup)

		nameIndex++
		index += 4
	}

	index = 0

	for i := 0; i < 10; i++ {
		var photoBase model.PhotoBase
		photoBase = photos[i]
		photoBase.ThumbSize = 540
		photoBase.Crop = 1
		photoBase.Key = -1
		peopleDTO.PhotoAnimations = append(peopleDTO.PhotoAnimations, photoBase)
	}
}
