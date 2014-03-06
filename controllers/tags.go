package controllers

import (
	"github.com/martini-contrib/render"

	"pixelrelay/db"
	"pixelrelay/encoder"
	"pixelrelay/models"
)

func Tags(su models.User, dbh *db.Dbh, p *models.Page, r render.Render) {
	tags := dbh.GetAllTags()

	p.SetUser(su)
	p.SetTitle("Tags")
	p.Data = tags

	encoder.Render(p.Encoding, 200, "tags", p, r)
}
