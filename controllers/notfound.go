/*
 Route: 404

 Method: *

 Return:
  - 404 page
*/

package controllers

import (
	"github.com/martini-contrib/render"

	"pixelrelay/encoder"
	"pixelrelay/models"
)

func NotFound(su models.User, r render.Render, p *models.Page) {
	p.SetUser(su)
	p.SetTitle("404")
	encoder.Render(p.Encoding, 404, "notfound", p, r)
}
