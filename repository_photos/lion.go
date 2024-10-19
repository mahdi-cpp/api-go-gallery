package repository_photos

import (
	"github.com/mahdi-cpp/api-go-gallery/cache"
	"github.com/mahdi-cpp/api-go-gallery/model"
)

var lionDTO LionDTO

type LionDTO struct {
	Photos    []model.PhotoBase `json:"photos"`
	Icons     []model.IconBase  `json:"icons"`
	Titles    []string          `json:"titles"`
	SubTitles []string          `json:"subTitles"`
}

func GetLion(folder string) {

	var file = "data.txt"
	photos := cache.ReadOfFile(folder, file)
	var count = len(photos)

	var index = 0

	for i := 0; i < count; i++ {
		var photo = model.PhotoBase{}
		photo = photos[index]
		photo.Key = -1
		photo.Crop = true
		photo.ThumbSize = 540
		lionDTO.Photos = append(lionDTO.Photos, photo)
		index++
	}

	lionDTO.Titles = []string{"", "", "", "Videos", "Favourites", "Suggestions", "Crater Lake", "", "", "", "", ""}
	lionDTO.SubTitles = []string{"", "MEDIA TYPES", "LIBRARY", "Reza", "TRIPS", "", "", "", "", ""}

	var iconsName = []string{"camera_51.png", "favourite_50.png", "camera_photo_51.png", "trip_50.png", "trip_50.png"}
	for i := 0; i < 5; i++ {
		var icon model.IconBase
		icon.Name = iconsName[i]
		icon.Key = -1
		lionDTO.Icons = append(lionDTO.Icons, icon)
	}

}
