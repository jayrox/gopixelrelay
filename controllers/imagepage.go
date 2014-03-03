package controllers

import (
	"net/http"

	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"

	"pixelrelay/db"
	"pixelrelay/models"
)

type ImagePageVars struct {
	Image models.Image
	Tags  []models.TagList
	Page  *models.Page
}

func ImagePage(args martini.Params, su models.User, res http.ResponseWriter, req *http.Request, ren render.Render, dbh *db.Dbh, p *models.Page) {
	var ipv ImagePageVars
	name := args["name"]

	ipv.Page = p
	ipv.Page.SetUser(su)
	ipv.Page.SetTitle("Image")
	ipv.Image.Name = name

	image := dbh.FirstImageByName(name)
	tags := dbh.GetAllTagsByImageId(image.Id)
	ipv.Tags = tags

	ren.HTML(200, "image", ipv)
}
