package controllers

import (
	"net/http"

	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"

	"pixelrelay/db"
	"pixelrelay/models"
)

type ImagePageVars struct {
	User  models.User
	Image models.Image
	Tags  []models.TagList
}

func ImagePage(args martini.Params, su models.User, res http.ResponseWriter, req *http.Request, ren render.Render, dbh *db.Dbh) {
	var ipv ImagePageVars
	name := args["name"]

	ipv.User = su
	ipv.Image.Name = name

	image := dbh.FirstImageByName(name)
	tags := dbh.GetAllTagsByImageId(image.Id)
	ipv.Tags = tags

	ren.HTML(200, "image", ipv)
}
