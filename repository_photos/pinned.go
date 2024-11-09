package repository_photos

import (
	"github.com/mahdi-cpp/api-go-gallery/cache"
	"github.com/mahdi-cpp/api-go-gallery/model"
	"github.com/mahdi-cpp/api-go-gallery/utils"
)

var pinnedCollectionDTO PinnedCollectionDTO

type PinnedCollectionDTO struct {
	PinnedCollections []PinnedCollection `json:"pinnedCollections"`
}

type PinnedCollection struct {
	Name  string          `json:"name"`
	Photo model.PhotoBase `json:"photo"`
}

func GetPinned(folder string) {

	var file = "data.txt"
	var photos = cache.ReadOfFile(folder, file)
	var count = len(photos)

	if count > 50 {
		count = 50
	}

	var index = 0
	var nameIndex = 0

	var marginX = dp(15)
	var screenWidthPhotosCount float32 = 3.2
	var photoSize = (1080 - (marginX * (screenWidthPhotosCount + 1))) / screenWidthPhotosCount

	for i := 0; i < count; i++ {
		if nameIndex >= len(utils.FackNames) {
			nameIndex = 0
		}

		var pinned = PinnedCollection{}
		pinned.Name = utils.FackNames[nameIndex]
		pinned.Photo = photos[index]
		pinned.Photo.ThumbSize = 270
		pinned.Photo.Crop = 1
		pinned.Photo.Round = int(dp(10))
		pinned.Photo.Key = -1
		pinned.Photo.PaintWidth = photoSize
		pinned.Photo.PaintHeight = photoSize

		pinnedCollectionDTO.PinnedCollections = append(pinnedCollectionDTO.PinnedCollections, pinned)

		nameIndex++
		index++
	}

	index = 0
}
