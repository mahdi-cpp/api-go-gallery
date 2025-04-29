package repository

var layout Layout

type Layout struct {
	Views []View `json:"views"`
	//CellItemCount int    `json:"cellItemCount"`
	//AspectRatio   bool   `json:"aspectRatio"`
	//ThumbSize     int    `json:"thumbSize"`
}

type View struct {
	Name  string `json:"name"`
	Title string `json:"title"`
	Url   string `json:"url"`
}

func LayoutInit() {

	var GalleryView = View{Url: "gallery"}

	var RecentDaysView = View{Name: "RecentDaysView", Title: "Recent Days", Url: "recent_days"}
	var PeopleView = View{Name: "PeopleView", Title: "Peoples", Url: "people"}
	var TripView = View{Name: "TripView", Title: "Trip", Url: "trip"}
	var PinnedCollections = View{Name: "PinnedCollections", Title: "Pinned Collections", Url: "pinned_collections"}
	var AlbumView = View{Name: "AlbumView", Title: "Albums", Url: "albums"}
	var ShareAlbumView = View{Name: "ShareAlbumView", Title: "Share Albums", Url: "share_albums"}
	var CameraTypeView = View{Name: "CameraTypeView", Title: "Cameras", Url: "cameras"}
	var CustomView = View{Name: "CustomView", Title: "Custom View", Url: "CustomView"}

	layout.Views = append(layout.Views, GalleryView)

	layout.Views = append(layout.Views, RecentDaysView)
	layout.Views = append(layout.Views, PeopleView)
	layout.Views = append(layout.Views, TripView)
	layout.Views = append(layout.Views, PinnedCollections)
	layout.Views = append(layout.Views, AlbumView)
	layout.Views = append(layout.Views, ShareAlbumView)
	layout.Views = append(layout.Views, CameraTypeView)
	layout.Views = append(layout.Views, CustomView)

	layout = layout

}
