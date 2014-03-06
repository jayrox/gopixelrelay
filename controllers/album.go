package controllers

import (
	"fmt"

	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"

	"pixelrelay/db"
	"pixelrelay/encoder"
	"pixelrelay/models"
)

type AlbumVars struct {
	ImageLinks []ImageLink
	AlbumUser  models.User
}

func Album(args martini.Params, su models.User, session sessions.Session, r render.Render, dbh *db.Dbh, p *models.Page) {
	album := args["name"]
	auser := args["user"]

	if auser != "" {
		fmt.Println("auser: ", auser)
	}

	images := dbh.GetAllAlbumImages(album)

	var imageLinks []ImageLink
	for _, f := range images {
		imageLinks = append(imageLinks, ImageLink{Title: f.Name, FileName: f.Name})
	}

	p.SetTitle("Album", album)
	p.SetUser(su)
	p.Data = AlbumVars{ImageLinks: imageLinks}

	encoder.Render(p.Encoding, 200, "image_link", p, r)
}
