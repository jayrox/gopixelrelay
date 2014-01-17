package controllers

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"net/http"
)

func Index(args martini.Params, res http.ResponseWriter, req *http.Request, ren render.Render) {
	ren.HTML(200, "index", "")
}
