package controllers

import (
	"log"
	"net/http"
	"strings"

	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"

	"pixelrelay/db"
	"pixelrelay/encoder"
	"pixelrelay/models"
	"pixelrelay/utils"
)

type AlbumVars struct {
	ImageLinks []ImageLink
	AlbumUser  models.User
}

func Album(args martini.Params, su models.User, session sessions.Session, r render.Render, dbh *db.Dbh, p *models.Page) {
	name := args["name"]
	auser := args["user"]
	key := args["key"]

	if auser != "" {
		log.Println("album user: ", auser)
	}

	album := dbh.GetAlbumByName(name)
	log.Println("private: ", album.Private)
	if su.Id != album.User && album.Private && album.Privatekey != key {
		session.Set("flash", "Login Required")
		r.Redirect(strings.Join([]string{utils.AppCfg.Url(), "login"}, "/"), http.StatusFound)
		return
	}

	log.Println("album: ", album)
	images := dbh.GetAllAlbumImages(name)

	var imageLinks []ImageLink
	for _, f := range images {
		imageLinks = append(imageLinks, ImageLink{Title: f.Name, FileName: f.Name})
	}

	p.SetTitle("Album", name)
	p.SetUser(su)
	p.Data = AlbumVars{ImageLinks: imageLinks}

	encoder.Render(p.Encoding, 200, "image_link", p, r)
}
