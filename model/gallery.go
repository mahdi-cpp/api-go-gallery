package model

type GalleryAlbum struct {
	AlbumName  string
	photoCount uint16
	Photos     []GalleryPhoto
}

type GalleryPhoto struct {
	Photo     string
	Thumbnail string
	Width     uint16
	Height    uint16
}

type Album struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Title     string `json:"title"`
	Title2    string `json:"title2"`
	AlbumName string `json:"album_name"`
	AlbumPath string `json:"album_path"`
}
