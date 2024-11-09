package repository_photos

import (
	"github.com/mahdi-cpp/api-go-gallery/cache"
	"github.com/mahdi-cpp/api-go-gallery/model"
	"github.com/mahdi-cpp/api-go-gallery/utils"
	"strconv"
)

var yearsDTO YearsDTO

type YearsDTO struct {
	Years []Year `json:"years"`
}

type Year struct {
	Title string          `json:"title"`
	Photo model.PhotoBase `json:"photo"`
}

func GetYears(folder string) {

	var file = "data.txt"
	photos := cache.ReadOfFile(folder, file)
	var count = len(photos)

	var index = 0

	if count > 20 {
		count = 20
	}

	var yearsCount = 1403

	for i := 0; i < count; i++ {
		var year = Year{}
		year.Photo = photos[index]
		year.Photo.Key = -1
		year.Photo.Crop = 1
		year.Photo.Round = 20
		year.Photo.ThumbSize = 540
		year.Photo.PaintWidth = float32(utils.ScreenWidth) - dp(30)
		year.Photo.PaintHeight = float32(utils.ScreenWidth * 0.70)
		year.Photo.Dx = dp(15)

		year.Title = "" + strconv.Itoa(yearsCount)

		yearsDTO.Years = append(yearsDTO.Years, year)
		index++
		yearsCount--
	}
}
