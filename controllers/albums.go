package controllers

import (
	"net/http"

	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"

	"pixelrelay/db"
	"pixelrelay/models"
)

type AlbumsVars struct {
	User       models.User
	AlbumsList []models.AlbumList
	AlbumUser  models.User
}

func Albums(args martini.Params, su models.User, session sessions.Session, r render.Render, res http.ResponseWriter, req *http.Request, dbh *db.Dbh) {
	var albumsVars AlbumsVars
	if su.Id > 0 {
		albumsVars.User = su
	}

	auser := args["user"]
	var albumUser models.User
	if auser != "" {
		albumUser = dbh.GetUserByUserName(auser)
		albumsVars.AlbumUser = albumUser
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
	albumsVars.AlbumsList = albumList

	r.HTML(200, "albums", albumsVars)
}
