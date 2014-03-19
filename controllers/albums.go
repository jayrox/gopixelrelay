package controllers

import (
	"net/http"

	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"

	"pixelrelay/db"
	"pixelrelay/encoder"
	"pixelrelay/models"
)

type AlbumsVars struct {
	AlbumUser  models.User        `json:"owner"`
	AlbumsList []models.AlbumList `json:"albums"`
}

func Albums(args martini.Params, su models.User, session sessions.Session, r render.Render, res http.ResponseWriter, req *http.Request, dbh *db.Dbh, p *models.Page) {
	auser := args["user"]
	var albumUser models.User
	if su.Id == 0 {
		albumUser.Name = "Public"
	}
	if auser != "" {
		albumUser = dbh.GetUserByUserName(auser)
	}

	if auser != "" && albumUser.Id == 0 {
		// handle user not found a little better?
		http.NotFound(res, req)
		return
	}

	var albums []models.Album
	if albumUser.Id > 0 {
		albums = dbh.GetAllAlbumsByUserId(albumUser.Id)
	} else {
		albums = dbh.GetAllAlbums("DESC")
	}

	var albumList []models.AlbumList
	for _, f := range albums {
		image := dbh.FirstImageByAlbum(f.Name)
		if f.Private && su.Id != f.User {
			continue
		}
		var pk string
		if su.Id == f.User {
			pk = f.Privatekey
		}

		albumList = append(albumList, models.AlbumList{Id: f.Id, Name: f.Name, Poster: image[0].Name, Private: f.Private, Privatekey: pk, Description: f.Description, Owner: f.User})
	}

	p.SetTitle("Albums")
	p.SetUser(su)
	p.Data = AlbumsVars{AlbumUser: albumUser, AlbumsList: albumList}

	encoder.Render(p.Encoding, 200, "albums", p, r)
}
