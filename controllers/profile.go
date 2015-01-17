/*
 Route:  /profile/:name

 Method: GET

 Params:
  - user string

 Return:
  - User profile
*/

package controllers

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"

	"pixelrelay/db"
	"pixelrelay/encoder"
	"pixelrelay/models"
)

func Profile(args martini.Params, su models.User, session sessions.Session, r render.Render, dbh *db.Dbh, p *models.Page) {
	name := args["user"]

	p.SetTitle("Profile", name)
	p.SetUser(su)
	p.Data = ""

	encoder.Render(p.Encoding, 200, "profile", p, r)
}
