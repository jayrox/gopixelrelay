package controllers

import (
	"log"
	"net/http"

	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"

	"pixelrelay/models"
)

type fourohfour struct {
	User models.User
}

func NotFound(args martini.Params, su models.User, res http.ResponseWriter, req *http.Request, ren render.Render) {
	log.Println("404")

	var fof fourohfour
	fof.User = su
	ren.HTML(404, "notfound", fof)
}
