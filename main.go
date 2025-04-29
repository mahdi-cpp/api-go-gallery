package main

import (
	"github.com/mahdi-cpp/api-go-gallery/cache"
	"github.com/mahdi-cpp/api-go-gallery/config"
	"github.com/mahdi-cpp/api-go-gallery/repository"
)

func main() {

	config.LayoutInit()
	repository.InitPhotos()
	cache.ReadIcons()

	Run()
}
