package main

import (
	"github.com/mahdi-cpp/api-go-gallery/cache"
	"github.com/mahdi-cpp/api-go-gallery/config"
	"github.com/mahdi-cpp/api-go-gallery/repository"
	"github.com/mahdi-cpp/api-go-gallery/repository_drawer"
)

func main() {

	config.LayoutInit()
	repository.InitPhotos()
	repository_drawer.InitDrawers()
	cache.ReadIcons()

	Run()
}
