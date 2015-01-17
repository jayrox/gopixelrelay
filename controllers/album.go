/*
 Routes:
  - /album/:name
  - /album/:name.json
  - /:user/album/:name
  - /:user/album/:name.json
  - /album/:name/:key
  - /album/:name/:key.json

 Method: GET

 Return:
  - AngularJS album template
  - JSON encoded image list

 Note:
  - In the case that a key is provided, a private album may be accessed where
  - it would otherwise prompt for login.
*/

package controllers

import (
	"log"
	"net/http"
	"strings"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"

	"pixelrelay/db"
	"pixelrelay/encoder"
	"pixelrelay/models"
	"pixelrelay/utils"
)

type AlbumVars struct {
	ImageLinks  []imageLink `json:"images"`
	AlbumUser   models.User
	Title       string
	Description string
}

type imageLink struct {
	Id       int64
	Title    string
	FileName string
	Owner    int64
}

func Album(args martini.Params, su models.User, session sessions.Session, r render.Render, dbh *db.Dbh, p *models.Page) {
	name := args["name"]
	auser := args["user"]
	key := args["key"]

	if auser != "" {
		log.Println("album user: ", auser)
	}

	album := dbh.GetAlbumByName(name)
	if su.Id != album.User && album.Private && album.Privatekey != key {
		session.Set("flash", "Login Required")
		r.Redirect(strings.Join([]string{utils.AppCfg.Url(), "login"}, "/"), http.StatusFound)
		return
	}

	images := dbh.GetAllImagesByAlbumId(album.Id)

	var imageLinks []imageLink
	for _, f := range images {
		if f.Trashed {
			continue
		}
		var file_name = f.Name
		if f.HashId != "" {
			file_name = f.HashId
		}
		log.Println(file_name, f)
		imageLinks = append(imageLinks, imageLink{Id: f.Id, Title: f.Name, FileName: file_name, Owner: f.User})
	}

	p.SetTitle("Album", name)
	p.SetUser(su)
	p.Data = AlbumVars{Title: name, Description: album.Description, ImageLinks: imageLinks}

	encoder.Render(p.Encoding, 200, "image_link", p, r)
}
