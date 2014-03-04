package controllers

import (
	"net/http"

	"github.com/martini-contrib/render"

	"pixelrelay/models"
)

type fourohfour struct {
	Page *models.Page
}

func NotFound(su models.User, res http.ResponseWriter, req *http.Request, ren render.Render, p *models.Page) {
	var fof fourohfour
	fof.Page = p
	fof.Page.SetUser(su)
	fof.Page.SetTitle("404")
	ren.HTML(404, "notfound", fof)
}
