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
	User       models.User
	ImageLinks []ImageLink
	AlbumUser  models.User
}

func Album(args martini.Params, su models.User, session sessions.Session, r render.Render) {
	var albumVars AlbumVars
	
	if su.Id > 0 {
		albumVars.User = su
	}

	album := args["name"]
	auser := args["user"]

	if auser != "" {
		fmt.Println("auser: ", auser)
	}

	d := db.InitDB()
	images := db.GetAllAlbumImages(&d, album)

	var imageLinks []ImageLink
	for _, f := range images {
		imageLinks = append(imageLinks, ImageLink{Title: f.Name, FileName: f.Name})
	}

	albumVars.ImageLinks = imageLinks

	r.HTML(200, "image_link", albumVars)
}
