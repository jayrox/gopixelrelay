package controllers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"

	"pixelrelay/db"
	"pixelrelay/forms"
	"pixelrelay/models"
	"pixelrelay/utils"
)

type ImagePageVars struct {
	Image   models.Image
	Tags    []models.TagList
	Page    *models.Page
	TagForm template.HTML
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

	errs := make(map[string]string)
	log.Println("generating form")
	ipv.TagForm = utils.GenerateForm(&forms.Tag{Image: name}, "/tag", "POST", errs)

	ren.HTML(200, "image", ipv)
}
