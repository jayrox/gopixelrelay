package controllers

import (
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	
	"pixelrelay/db"
)

func Tagged(args martini.Params, r render.Render) {
	tag := args["name"]
	d := db.InitDB()
	images := db.GetImagesWithTag(&d, tag)

	var imageLinks []ImageLink
	for _, f := range images {
		imageLinks = append(imageLinks, ImageLink{Title: f.Name, FileName: f.Name})
	}

	r.HTML(200, "image_link", imageLinks)
}
