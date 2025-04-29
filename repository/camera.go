package repository

import (
	"github.com/mahdi-cpp/api-go-gallery/cache"
	"github.com/mahdi-cpp/api-go-gallery/model"
	"github.com/mahdi-cpp/api-go-gallery/utils"
)

var cameraDTO CameraDTO

type CameraDTO struct {
	Cameras []Camera `json:"cameras"`
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

	var nameIndex = 0

	for i := 0; i < count; i++ {
		var camera = Camera{}
		camera.Name = utils.CameraNames[nameIndex]

		camera.PhotoLarge = photos[index+2]
		camera.PhotoLarge.ThumbSize = 540
		camera.PhotoLarge.Crop = 1
		camera.PhotoLarge.Key = -1
		camera.PhotoLarge.PaintWidth = 0
		camera.PhotoLarge.PaintHeight = 0
		camera.PhotoLarge.Dx = 0
		camera.PhotoLarge.Dy = 0

		for j := 0; j < 4; j++ {
			var photoBase model.PhotoBase
			photoBase = photos[index+1+j]
			photoBase.ThumbSize = 270
			photoBase.Crop = 1
			photoBase.Key = -1
			photoBase.PaintWidth = 0
			photoBase.PaintHeight = 0
			camera.PhotosTiny = append(camera.PhotosTiny, photoBase)
		}

		cameraDTO.Cameras = append(cameraDTO.Cameras, camera)

		index += 5
		nameIndex += 1
	}

	return cameraDTO
}
