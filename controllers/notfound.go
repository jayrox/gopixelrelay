package controllers

import (
	"net/http"

	"github.com/martini-contrib/render"

	"pixelrelay/encoder"
	"pixelrelay/models"
)

func NotFound(su models.User, res http.ResponseWriter, req *http.Request, r render.Render, p *models.Page) {
	p.SetUser(su)
	p.SetTitle("404")
	encoder.Render(p.Encoding, 404, "notfound", p, r)
}
