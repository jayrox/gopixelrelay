package controllers

import (
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	"pixelrelay/db"
	"pixelrelay/models"
)

func Albums(args martini.Params, r render.Render) {
	d := db.InitDB()
	albums := db.GetAllAlbums(&d)

	var albumList []models.AlbumList
	for _, f := range albums {
		i := db.FirstImage(&d, f.Name)
		albumList = append(albumList, models.AlbumList{Name: f.Name, Poster: i[0].Name})
	}

	r.HTML(200, "albums", albumList)
}
