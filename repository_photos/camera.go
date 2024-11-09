package repository_photos

import (
	"github.com/mahdi-cpp/api-go-gallery/cache"
	"github.com/mahdi-cpp/api-go-gallery/model"
	"github.com/mahdi-cpp/api-go-gallery/utils"
)

var cameraDTO CameraDTO

type CameraDTO struct {
	MarginX                float32  `json:"marginX"`
	ScreenWidthPhotosCount float32  `json:"screenWidthPhotosCount"`
	PhotoSize              float32  `json:"photoSize"`
	TinyPhoto              float32  `json:"tinyPhoto"`
	Cameras                []Camera `json:"cameras"`
}

type Camera struct {
	Name       string            `json:"name"`
	PhotoLarge model.PhotoBase   `json:"photoLarge"`
	PhotosTiny []model.PhotoBase `json:"photosTiny"`
}

func GetCameras(folder string) CameraDTO {

	var file = "data.txt"
	var photos = cache.ReadOfFile(folder, file)
	var cameraDTO CameraDTO

	var count = len(photos) / 5
	var index = 0

	var marginX = dp(20)
	var screenWidthPhotosCount float32 = 0.7
	var photoSize = (1080 - (marginX * (screenWidthPhotosCount + 1))) / screenWidthPhotosCount
	photoSize = photoSize / 3.8
	var tinyPhoto = (photoSize - dp(2)) / 2.0

	cameraDTO.MarginX = marginX
	cameraDTO.ScreenWidthPhotosCount = screenWidthPhotosCount
	cameraDTO.PhotoSize = photoSize
	cameraDTO.TinyPhoto = tinyPhoto

	//if count > 5 {
	//	count = 2
	//}

	var nameIndex = 0

	for i := 0; i < count; i++ {
		var camera = Camera{}
		camera.Name = utils.CameraNames[nameIndex]

		camera.PhotoLarge = photos[index+2]
		camera.PhotoLarge.ThumbSize = 540
		camera.PhotoLarge.Crop = 1
		camera.PhotoLarge.Key = -1
		camera.PhotoLarge.PaintWidth = photoSize
		camera.PhotoLarge.PaintHeight = photoSize
		camera.PhotoLarge.Dx = 0
		camera.PhotoLarge.Dy = 0

		for j := 0; j < 4; j++ {
			var photoBase model.PhotoBase
			photoBase = photos[index+1+j]
			photoBase.ThumbSize = 270
			photoBase.Crop = 1
			photoBase.Key = -1
			photoBase.PaintWidth = tinyPhoto
			photoBase.PaintHeight = tinyPhoto
			camera.PhotosTiny = append(camera.PhotosTiny, photoBase)
		}

		camera.PhotosTiny[0].Dx = photoSize + dp(2)
		camera.PhotosTiny[0].Dy = 0

		camera.PhotosTiny[1].Dx = photoSize + dp(2)
		camera.PhotosTiny[1].Dy = tinyPhoto + dp(2)

		camera.PhotosTiny[2].Dx = photoSize + tinyPhoto + dp(4)
		camera.PhotosTiny[2].Dy = 0

		camera.PhotosTiny[3].Dx = photoSize + tinyPhoto + dp(4)
		camera.PhotosTiny[3].Dy = tinyPhoto + dp(2)

		cameraDTO.Cameras = append(cameraDTO.Cameras, camera)

		index += 5
		nameIndex += 1
	}

	return cameraDTO
}
