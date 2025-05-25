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
	Name   string          `json:"name"`
	Images []model.UIImage `json:"images"`
}

func GetCameras(folder string) CameraDTO {

	var file = "data.txt"
	var uiImages = cache.ReadOfFile(folder, file)
	var cameraDTO CameraDTO

	var count = len(uiImages) / 6
	var index = 0

	var nameIndex = 0
	if count > 30 {
		count = 30
	}

	for i := 0; i < count; i++ {
		var camera = Camera{}
		camera.Name = utils.CameraNames[nameIndex]

		for j := 0; j < 6; j++ {
			var image model.UIImage
			image = uiImages[index+2+j]
			camera.Images = append(camera.Images, image)
		}

		cameraDTO.Cameras = append(cameraDTO.Cameras, camera)
		index += 6
		nameIndex += 1
	}

	return cameraDTO
}
