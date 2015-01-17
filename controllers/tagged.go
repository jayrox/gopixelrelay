/*
 Routes:
  - /tag/:name
  - /tag/:name.json

 Method: GET

 Params:
  - tag string

 Return:
  - List of images with tag
*/

package controllers

import (
	"fmt"
	"log"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"

	"pixelrelay/db"
	"pixelrelay/encoder"
	"pixelrelay/models"
)

type TaggedVars struct {
	ImageLinks  []imageLink `json:"images"`
	Description string
	Title       string
}

func Tagged(args martini.Params, r render.Render, su models.User, dbh *db.Dbh, p *models.Page) {

	tag := args["tag"]

	images := dbh.GetImagesWithTag(tag)

	var imageLinks []imageLink
	for _, f := range images {
		if f.Trashed {
			log.Println("Trashed: ", f)
			continue
		}
		imageLinks = append(imageLinks, imageLink{Title: f.Name, FileName: f.Name})
	}

	p.SetUser(su)
	p.SetTitle("Tagged", tag)
	description := fmt.Sprintf("Images tagged as %s", tag)
	p.Data = TaggedVars{Title: tag, Description: description, ImageLinks: imageLinks}
	p.Encoding = "json"

	encoder.Render(p.Encoding, 200, "image_link", p, r)
}
