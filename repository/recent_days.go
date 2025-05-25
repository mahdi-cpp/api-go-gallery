package repository

import (
	"github.com/mahdi-cpp/api-go-gallery/cache"
	"github.com/mahdi-cpp/api-go-gallery/model"
)

var recentDaysDTO RecentDaysDTO

type RecentDaysDTO struct {
	RecentDays []RecentDays `json:"recentDays"`
}

type RecentDays struct {
	Name   string          `json:"name"`
	Images []model.UIImage `json:"images"`
}

func GetRecently(folder string) {

	var file = "data.txt"
	var uiImages = cache.ReadOfFile(folder, file)

	recentDaysDTO = RecentDaysDTO{}
	var count = ((len(uiImages) - 4) / 4) + 1
	var index = 0

	if count > 25 {
		count = 25
	}

	for i := 0; i < count; i++ {
		var album = RecentDays{}

		for j := 0; j < 4; j++ {
			var image model.UIImage
			image = uiImages[index+j]
			album.Images = append(album.Images, image)
		}

		recentDaysDTO.RecentDays = append(recentDaysDTO.RecentDays, album)
		index += 4
	}
}
