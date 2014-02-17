package controllers

import (
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	"pixelrelay/db"
	"pixelrelay/models"
)

func Tags(args martini.Params, r render.Render) {
	d := db.InitDB()

	var tags []models.Tag
	tags = db.GetAllTags(&d)

	r.HTML(200, "tags", tags)
}
