package main

import (
	"github.com/mahdi-cpp/api-go-gallery/cache"
	"github.com/mahdi-cpp/api-go-gallery/config"
	"github.com/mahdi-cpp/api-go-gallery/repository_chat"
	"github.com/mahdi-cpp/api-go-gallery/repository_drawer"
	"github.com/mahdi-cpp/api-go-gallery/repository_photos"
)

func main() {

	config.LayoutInit()

	repository_photos.InitPhotos()
	repository_chat.InitModels()
	repository_drawer.InitDrawers()

	cache.ReadIcons()

	Run()
}
