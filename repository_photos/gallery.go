package repository_photos

import (
	"github.com/mahdi-cpp/api-go-gallery/cache"
	"github.com/mahdi-cpp/api-go-gallery/model"
)

type GalleryDTO struct {
	Photos []model.PhotoBase `json:"photos"`
}

var galleryDTO GalleryDTO
var photoBaseArray []model.PhotoBase

func GetGalleries(folder string) {

	var file = "data.txt"
	photos := cache.ReadOfFile(folder, file)
	var count = len(photos)
	var index = 0

	for i := 0; i < count; i++ {
		var photo = model.PhotoBase{}
		photo = photos[index]
		photo.Key = -1
		galleryDTO.Photos = append(galleryDTO.Photos, photo)
		index++
	}
}
