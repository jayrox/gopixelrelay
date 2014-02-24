package controllers

import (
	"fmt"
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

func Albums(args martini.Params, su models.User, session sessions.Session, r render.Render, res http.ResponseWriter, req *http.Request) {
	var albumsVars AlbumsVars
	if su.Id > 0 {
		albumsVars.User = su
	}

	d := db.InitDB()
	albums := db.GetAllAlbums(&d)

	loggedin := session.Get("loggedin")
	if loggedin != nil {
		fmt.Println("logged in")
	}

	uid := session.Get("uid")
	if uid != nil {
		albumsVars.User.Id = uid.(int64)
	}

	// If email set, apply to form
	email := session.Get("email")
	if email != nil {
		albumsVars.User.Email = email.(string)
	}

	var albumList []models.AlbumList
	for _, f := range albums {
		i := db.FirstImage(&d, f.Name)
		albumList = append(albumList, models.AlbumList{Name: f.Name, Poster: i[0].Name})
	}
	albumsVars.AlbumsList = albumList

	r.HTML(200, "albums", albumsVars)
}
