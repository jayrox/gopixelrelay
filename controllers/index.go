package controllers

import (
	"net/http"

	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"

	"pixelrelay/models"
)

type IndexVars struct {
	Page *models.Page
}

func Index(args martini.Params, su models.User, res http.ResponseWriter, req *http.Request, ren render.Render, p *models.Page) {
	var indexVars IndexVars

	indexVars.Page = p
	indexVars.Page.SetUser(su)
	indexVars.Page.SetTitle("JustRiot!")
	ren.HTML(200, "index", indexVars)
}
