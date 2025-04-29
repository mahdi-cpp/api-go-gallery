package main

import (
	"github.com/mahdi-cpp/api-go-gallery/cache"
	"github.com/mahdi-cpp/api-go-gallery/repository"
)

func main() {

	repository.InitPhotos()
	cache.ReadIcons()

	Run()
}
