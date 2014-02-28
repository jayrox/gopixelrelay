package controllers

import (
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"

	"pixelrelay/db"
	"pixelrelay/models"
)

type TaggedVars struct {
	User       models.User
	ImageLinks []ImageLink
}

func Tagged(args martini.Params, r render.Render, su models.User, dbh *db.Dbh) {
	var taggedVars TaggedVars

	tag := args["name"]

	images := dbh.GetImagesWithTag(tag)

	var imageLinks []ImageLink
	for _, f := range images {
		imageLinks = append(imageLinks, ImageLink{Title: f.Name, FileName: f.Name})
	}

	taggedVars.User = su
	taggedVars.ImageLinks = imageLinks

	r.HTML(200, "image_link", taggedVars)
}
