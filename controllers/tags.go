package controllers

import (
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"

	"pixelrelay/db"
	"pixelrelay/models"
)

type TagVars struct {
	User models.User
	Tags []models.Tag
}

func Tags(args martini.Params, r render.Render, su models.User, dbh *db.Dbh) {
	var tags []models.Tag
	var tagVars TagVars

	tags = dbh.GetAllTags()

	tagVars.User = su
	tagVars.Tags = tags

	r.HTML(200, "tags", tagVars)
}
