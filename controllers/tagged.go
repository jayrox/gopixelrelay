package controllers

import (
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"

	"pixelrelay/db"
	"pixelrelay/models"
)

type TaggedVars struct {
	ImageLinks []ImageLink
	Page       *models.Page
}

func Tagged(args martini.Params, r render.Render, su models.User, dbh *db.Dbh, p *models.Page) {
	var taggedVars TaggedVars

	tag := args["name"]

	images := dbh.GetImagesWithTag(tag)

	var imageLinks []ImageLink
	for _, f := range images {
		imageLinks = append(imageLinks, ImageLink{Title: f.Name, FileName: f.Name})
	}

	taggedVars.Page = p
	taggedVars.Page.SetUser(su)
	taggedVars.Page.SetTitle("Tagged", tag)
	taggedVars.ImageLinks = imageLinks

	r.HTML(200, "image_link", taggedVars)
}
