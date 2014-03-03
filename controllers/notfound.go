package controllers

import (
	"log"
	"net/http"

	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"

	"pixelrelay/models"
)

type fourohfour struct {
	Page *models.Page
}

func NotFound(args martini.Params, su models.User, res http.ResponseWriter, req *http.Request, ren render.Render, p *models.Page) {
	log.Println("404")

	var fof fourohfour
	fof.Page = p
	fof.Page.SetUser(su)
	fof.Page.SetTitle("404")
	ren.HTML(404, "notfound", fof)
}
