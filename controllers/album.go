package controllers

import (
	"fmt"
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"pixelrelay/db"
)

func Album(args martini.Params, session sessions.Session, r render.Render) {
	album := args["name"]
	d := db.InitDB()
	images := db.GetAllAlbumImages(&d, album)
	session.Set("album", album)

	var imageLinks []ImageLink
	for _, f := range images {
		imageLinks = append(imageLinks, ImageLink{Title: f.Name, FileName: f.Name})
	}
	
	v := session.Get("album")
	fmt.Printf("sessions: %s\n", v)

	r.HTML(200, "image_link", imageLinks)
}
