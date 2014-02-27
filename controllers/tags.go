package controllers

import (
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"

	"pixelrelay/db"
	"pixelrelay/models"
)

func Tags(args martini.Params, r render.Render, dbh *db.Dbh) {
	var tags []models.Tag

	tags = dbh.GetAllTags()

	r.HTML(200, "tags", tags)
}
