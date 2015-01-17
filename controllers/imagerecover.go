/*
 Route:  /image/recover/:album/:name

 Method: GET

 Params:
  - name string
  - album string

 Return:
  - JSON
  - Params:
   - Name string
   - Album string
   - Status string
*/

package controllers

import (
	"log"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"

	"pixelrelay/db"
	"pixelrelay/encoder"
	"pixelrelay/models"
)

type TrashRecover struct {
	Name   string
	Album  string
	Action string
	Status string
}

func ImageRecover(args martini.Params, su models.User, dbh *db.Dbh, r render.Render, p *models.Page) {
	album := args["album"]
	name := args["name"]

	// Default status
	var status string = "Permission Denied"
	var code int = 401

	log.Printf("Recovering image: %s from %s", name, album)

	image := dbh.FirstImageByName(name)
	if su.Id == image.User {
		image.Trashed = false
		dbh.UpdateImage(image)
		status = "Success"
		code = 200
	}

	p.SetUser(su)
	p.Data = TrashRecover{Name: name, Album: album, Action: "Recover", Status: status}

	encoder.Render("json", code, "", p, r)
}
