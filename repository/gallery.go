package repository

import (
	"github.com/mahdi-cpp/api-go-gallery/cache"
	"github.com/mahdi-cpp/api-go-gallery/model"
)

type GalleryDTO struct {
	Avatar model.PhotoBase   `json:"avatar"`
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

		photo.IsVideo = false
		photo.VideoFormat = "mp4"

		galleryDTO.Photos = append(galleryDTO.Photos, photo)
		index++
	}

	var photo = model.PhotoBase{}
	photo.Key = -1
	photo.Name = "00044"
	photo.ThumbSize = 540
	photo.FileType = ".jpg"
	photo.Width = int(dp(33))
	photo.Height = int(dp(33))
	photo.PaintWidth = dp(33)
	photo.PaintHeight = dp(33)
	photo.Circle = true
	galleryDTO.Avatar = photo

}
