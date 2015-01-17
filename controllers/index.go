/*
 Route:  /

 Method: GET

 Return:
  - index page
*/

package controllers

import (
	"github.com/martini-contrib/render"

	"pixelrelay/encoder"
	"pixelrelay/models"
)

func Index(su models.User, r render.Render, p *models.Page) {

	p.SetUser(su)
	p.SetTitle("")
	encoder.Render(p.Encoding, 200, "index", p, r)
}
