package controllers

import (
	"net/http"

	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"

	"pixelrelay/models"
)

type IndexVars struct {
	User models.User
}

func Index(args martini.Params, su models.User, res http.ResponseWriter, req *http.Request, ren render.Render) {
	var indexVars IndexVars
	indexVars.User = su
	ren.HTML(200, "index", indexVars)
}
