package controllers

import (
	"fmt"

	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"

	"pixelrelay/db"
	"pixelrelay/models"
)

type AlbumVars struct {
	ImageLinks []ImageLink
	AlbumUser  models.User
	Page       *models.Page
}

func Album(args martini.Params, su models.User, session sessions.Session, r render.Render, dbh *db.Dbh, p *models.Page) {
	var albumVars AlbumVars

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

	albumVars.ImageLinks = imageLinks

	albumVars.Page = p
	albumVars.Page.SetTitle("Album", album)
	albumVars.Page.SetUser(su)

	r.HTML(200, "image_link", albumVars)
}
