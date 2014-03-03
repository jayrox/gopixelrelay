package controllers

import (
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"

	"pixelrelay/db"
	"pixelrelay/models"
)

type TagVars struct {
	Page *models.Page
	Tags []models.Tag
}

func Tags(args martini.Params, r render.Render, su models.User, dbh *db.Dbh, p *models.Page) {
	var tags []models.Tag
	var tagVars TagVars

	tags = dbh.GetAllTags()

	tagVars.Page = p
	tagVars.Page.SetUser(su)
	tagVars.Page.SetTitle("Tags")
	tagVars.Tags = tags

	r.HTML(200, "tags", tagVars)
}
