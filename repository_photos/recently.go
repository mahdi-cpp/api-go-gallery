package repository_photos

import (
	"github.com/mahdi-cpp/api-go-gallery/cache"
	"github.com/mahdi-cpp/api-go-gallery/model"
	"strconv"
)

type RecentlyDay struct {
	Name       string          `json:"name"`
	PhotoLarge model.PhotoBase `json:"photoLarge"`
	PhotoTiny1 model.PhotoBase `json:"photoTiny1"`
	PhotoTiny2 model.PhotoBase `json:"photoTiny2"`
	PhotoTiny3 model.PhotoBase `json:"photoTiny3"`
}

var recentlyDTO []RecentlyDay

func GetRecently(folder string) {

	var file = "data.txt"
	var photos = cache.ReadOfFile(folder, file)

	recentlyDTO = []RecentlyDay{}
	var count = (len(photos) / 4) - 4
	var index = 0

	var marginX = dp(38)
	var screenWidthPhotosCount float32 = 1.9

	var photoSize = (1080 - (marginX * (screenWidthPhotosCount + 1))) / screenWidthPhotosCount
	var tinyPhotoSize = photoSize / 3

	if count > 50 {
		count = 50
	}

	for i := 0; i < count; i++ {
		var album = RecentlyDay{}
		album.Name = "Camera " + strconv.Itoa(i)

		album.PhotoLarge = photos[index+1]
		album.PhotoTiny1 = photos[index+2]
		album.PhotoTiny2 = photos[index+3]
		album.PhotoTiny3 = photos[index+4]

		album.PhotoLarge.ThumbSize = 540
		album.PhotoTiny1.ThumbSize = 135
		album.PhotoTiny2.ThumbSize = 135
		album.PhotoTiny3.ThumbSize = 135

		album.PhotoLarge.Crop = true
		album.PhotoTiny1.Crop = true
		album.PhotoTiny2.Crop = true
		album.PhotoTiny3.Crop = true

		album.PhotoLarge.Key = -1
		album.PhotoTiny1.Key = -1
		album.PhotoTiny2.Key = -1
		album.PhotoTiny3.Key = -1

		album.PhotoLarge.PaintWidth = photoSize + dp(4)
		album.PhotoLarge.PaintHeight = photoSize + dp(6.5)

		album.PhotoTiny1.PaintWidth = tinyPhotoSize
		album.PhotoTiny1.PaintHeight = tinyPhotoSize

		album.PhotoTiny2.PaintWidth = tinyPhotoSize
		album.PhotoTiny2.PaintHeight = tinyPhotoSize

		album.PhotoTiny3.PaintWidth = tinyPhotoSize
		album.PhotoTiny3.PaintHeight = tinyPhotoSize

		album.PhotoTiny1.Dx = photoSize + dp(6.5)
		album.PhotoTiny1.Dy = 0

		album.PhotoTiny2.Dx = photoSize + dp(6.5)
		album.PhotoTiny2.Dy = tinyPhotoSize + dp(2)

		album.PhotoTiny3.Dx = photoSize + dp(6.5)
		album.PhotoTiny3.Dy = tinyPhotoSize + tinyPhotoSize + dp(4.5)

		recentlyDTO = append(recentlyDTO, album)

		index += 4
	}
}
