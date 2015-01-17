/*
 Routes:
  - /albums
  - /albums.json
  - /:user/albums
  - /:user/albums.json

 Method: GET

 Return:
  - AngularJS albums template
  - JSON encoded album list

 Note:
  - By default this route only provides access to the entry point for AngularJS
  - Once AngularJS kicks in, it is called again but this time it returns a list
  - of albums in JSON format.
*/

package controllers

import (
	"net/http"

	"github.com/go-martini/martini"
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

	p.SetTitle("Albums")
	p.SetUser(su)

	if p.Encoding == "json" {
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
		var image []models.Image
		for _, f := range albums {
			if f.Private && su.Id != f.User {
				continue
			}
			var pk string
			var editable bool = false
			if su.Id == f.User {
				pk = f.Privatekey
				editable = true
			}

			// Select the best album poster image
			var poster string = "/gallery.png"
			image = dbh.FirstImageByAlbumId(f.Id)
			if len(image) > 0 {
				var file_name = image[0].Name
				if image[0].HashId != "" {
					file_name = image[0].HashId
				}
				poster = "/t/" + file_name
			}
			albumList = append(albumList, models.AlbumList{Id: f.Id, Name: f.Name, Poster: poster, Private: f.Private, Privatekey: pk, Description: f.Description, Owner: f.User, Editable: editable})
		}

		p.Data = AlbumsVars{AlbumUser: albumUser, AlbumsList: albumList}
	}
	encoder.Render(p.Encoding, 200, "albums", p, r)
}
