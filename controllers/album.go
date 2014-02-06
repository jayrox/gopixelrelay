package controllers

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"pixelrelay/db"
)

func Album(args martini.Params, r render.Render) {
	album := args["name"]
	d := db.InitDB()
	images := db.GetAllAlbumImages(&d, album)

	var imageLinks []ImageLink
	for _, f := range images {
		imageLinks = append(imageLinks, ImageLink{Title: f.Name, FileName: f.Name})
	}

	r.HTML(200, "image_link", imageLinks)
}
