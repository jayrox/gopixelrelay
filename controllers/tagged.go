package controllers

import (
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"

	"pixelrelay/db"
	"pixelrelay/encoder"
	"pixelrelay/models"
)

type TaggedVars struct {
	ImageLinks []ImageLink `json:"images"`
}

func Tagged(args martini.Params, r render.Render, su models.User, dbh *db.Dbh, p *models.Page) {

	tag := args["name"]

	images := dbh.GetImagesWithTag(tag)

	var imageLinks []ImageLink
	for _, f := range images {
		imageLinks = append(imageLinks, ImageLink{Title: f.Name, FileName: f.Name})
	}

	p.SetUser(su)
	p.SetTitle("Tagged", tag)
	p.Data = TaggedVars{ImageLinks: imageLinks}

	encoder.Render(p.Encoding, 200, "image_link", p, r)
}
