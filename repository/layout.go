package repository

import (
	"github.com/mahdi-cpp/api-go-gallery/config"
)

var layout Layout

type Layout struct {
	Views         []View `json:"views"`
	CellItemCount int    `json:"cellItemCount"`
	AspectRatio   bool   `json:"aspectRatio"`
	ThumbSize     int    `json:"thumbSize"`
}

type View struct {
	Name  string `json:"name"`
	Title string `json:"title"`
	Url   string `json:"url"`
}

func LayoutInit() {

	var PinnedCollections = View{Name: "PinnedCollections", Title: "Pinned Collections", Url: "pinned_collections"}
	var CameraTypeView = View{Name: "CameraTypeView", Title: "Cameras", Url: "cameras"}

	var RecentlyView = View{Name: "RecentlyView", Title: "Recently Days", Url: "recently"}

	var PeopleView = View{Name: "PeopleView", Title: "Peoples", Url: "people"}
	var TripView = View{Name: "TripView", Title: "Trip", Url: "trip"}
	var AlbumView = View{Name: "AlbumView", Title: "Albums", Url: "albums"}
	var ShareAlbumView = View{Name: "ShareAlbumView", Title: "Share Albums", Url: "share_albums"}

	var GalleryView = View{Url: "gallery"}
	//var YearsView = View{Name: "YearsView", Title: "Years View", Url: "years"}

	//var LionView = View{Name: "LionView", Title: "Lion View", Url: "lion"}
	var CustomView = View{Name: "CustomView", Title: "Custom View", Url: "CustomView"}

	layout.Views = append(layout.Views, GalleryView)
	//layout.Views = append(layout.Views, LionView)
	layout.Views = append(layout.Views, RecentlyView)
	layout.Views = append(layout.Views, PeopleView)

	layout.Views = append(layout.Views, TripView)
	layout.Views = append(layout.Views, PinnedCollections)
	layout.Views = append(layout.Views, AlbumView)
	layout.Views = append(layout.Views, ShareAlbumView)
	//layout.Views = append(layout.Views, YearsView)
	layout.Views = append(layout.Views, CameraTypeView)
	layout.Views = append(layout.Views, CustomView)

	layout.AspectRatio = config.CurrentLayout.AspectRatio
	layout.CellItemCount = config.CurrentLayout.CellItemCount
	layout.ThumbSize = config.CurrentLayout.ThumbSize

	//for i, _ := range layout.Views {
	//for i := len(layout.Views) - 1; i >= 0; i-- {
	//	offset -= layout.Views[i].Height
	//	layout.Views[i].ScrollY = offset
	//}

	layout = layout

}
