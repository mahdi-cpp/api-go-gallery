package repository

import (
	"github.com/mahdi-cpp/api-go-gallery/cache"
	"github.com/mahdi-cpp/api-go-gallery/model"
	"github.com/mahdi-cpp/api-go-gallery/utils"
)

var albumDTO AlbumDTO

type AlbumDTO struct {
	Albums []Album `json:"albums"`
}

type Album struct {
	Name   string          `json:"name"`
	Images []model.UIImage `json:"images"`
}

func GetAlbums(folder string) AlbumDTO {

	var file = "data.txt"
	var uiImages = cache.ReadOfFile(folder, file)
	var albumDTO AlbumDTO

	var count = len(uiImages) / 10
	var index = 0
	var nameIndex = 0

	if count > 125 {
		count = 125
	}

	for i := 0; i < count; i++ {
		var album = Album{}
		if nameIndex+1 >= len(utils.FackNames) {
			nameIndex = 0
		}
		album.Name = utils.FackTrips[nameIndex]

		for j := 0; j < 5; j++ {
			var image model.UIImage
			image = uiImages[index+1+j]
			album.Images = append(album.Images, image)
		}

		albumDTO.Albums = append(albumDTO.Albums, album)
		index += 5
		nameIndex += 1
	}

	return albumDTO
}
