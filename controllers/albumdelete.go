/*
 Route:  /album/create

 Method: POST

 Return:
  - JSON
  - Params
   - Name string
   - Result string
   - Status string
*/

package controllers

import (
	"log"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"

	"pixelrelay/db"
	"pixelrelay/encoder"
	"pixelrelay/models"
)

func AlbumDelete(args martini.Params, req *http.Request, r render.Render, su models.User, dbh *db.Dbh, p *models.Page) {
	// Default status
	var result string = "Not Deleted"
	var status string = "Permission Denied"
	var code int = 401

	name := args["name"]

	album := dbh.GetAlbumByName(name)
	if album.User == su.Id {
		status = "Success"
		result = "Deleted"
		dAlbum := dbh.AlbumDelete(album)
		log.Printf("%+v", dAlbum)
	}
	if album.Id == 0 {
		status = "Album not found"
		code = 404
	}

	p.Data = models.AlbumResult{Name: name, Result: result, Status: status}
	log.Println("Album Delete: User: ", su.Id, " Album: ", name, " Result: ", result, " Status: ", status)
	encoder.Render("json", code, "", p, r)
}
