package repository

import (
	"github.com/mahdi-cpp/api-go-gallery/cache"
	"github.com/mahdi-cpp/api-go-gallery/model"
	"github.com/mahdi-cpp/api-go-gallery/utils"
)

var albumDTO AlbumDTO

var shareAlbumDTO AlbumDTO

type AlbumDTO struct {
	Albums []Album `json:"albums"`
}

type Album struct {
	AlbumName1  string            `json:"albumName1"`
	AlbumName2  string            `json:"albumName2"`
	PhotoLarge1 model.PhotoBase   `json:"photoLarge1"`
	PhotoLarge2 model.PhotoBase   `json:"photoLarge2"`
	PhotosTiny  []model.PhotoBase `json:"photosTiny"`
}

func GetAlbums(folder string) AlbumDTO {

	var file = "data.txt"
	var photos = cache.ReadOfFile(folder, file)
	var albumDTO AlbumDTO

	var count = len(photos) / 10
	var index = 0
	var nameIndex = 0

	if count > 50 {
		count = 50
	}

	for i := 0; i < count; i++ {
		var album = Album{}
		if nameIndex+1 >= len(utils.FackNames) {
			nameIndex = 0
		}
		album.AlbumName1 = utils.FackTrips[nameIndex]
		album.AlbumName2 = utils.ShareAlbumTitles[nameIndex+1]

		album.PhotoLarge1 = photos[index+1]
		album.PhotoLarge1.ThumbSize = 270
		album.PhotoLarge1.Crop = 1
		album.PhotoLarge1.Key = -1
		album.PhotoLarge1.Dx = 0
		album.PhotoLarge1.Dy = 0

		album.PhotoLarge2 = photos[index+2]
		album.PhotoLarge2.ThumbSize = 270
		album.PhotoLarge2.Crop = 1
		album.PhotoLarge2.Key = -1
		album.PhotoLarge2.Dx = 0
		album.PhotoLarge2.Dy = 0

		for j := 0; j < 8; j++ {
			var photoBase model.PhotoBase
			photoBase = photos[index+2+j]
			photoBase.ThumbSize = 135
			photoBase.Crop = 1
			photoBase.Key = -1
			album.PhotosTiny = append(album.PhotosTiny, photoBase)
		}

		albumDTO.Albums = append(albumDTO.Albums, album)

		index += 10
		nameIndex += 2
	}

	return albumDTO
}
