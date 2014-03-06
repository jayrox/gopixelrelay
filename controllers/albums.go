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
		albums = dbh.GetAllAlbums()
	}

	var albumList []models.AlbumList
	for _, f := range albums {
		i := dbh.FirstImageByAlbum(f.Name)
		if f.Private && su.Id != f.User {
			continue
		}
		albumList = append(albumList, models.AlbumList{Name: f.Name, Poster: i[0].Name, Private: f.Private})
	}

	p.SetTitle("Albums")
	p.SetUser(su)
	p.Data = AlbumsVars{AlbumUser: albumUser, AlbumsList: albumList}

	encoder.Render(p.Encoding, 200, "albums", p, r)
}
