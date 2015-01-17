/*
 Route:  /image/:name

 Method: GET

 Return:
  - Image page with image, tags, tag form
*/

package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"

	"pixelrelay/db"
	"pixelrelay/encoder"
	"pixelrelay/forms"
	"pixelrelay/models"
	"pixelrelay/utils"
)

type ImagePageVars struct {
	Name string
	Full string
	Tags []models.TagList
	Form template.HTML `json:"-"`
}

func ImagePage(args martini.Params, su models.User, res http.ResponseWriter, req *http.Request, r render.Render, dbh *db.Dbh, p *models.Page) {
	name := args["name"]

	image := dbh.FirstImageByName(name)
	tags := dbh.GetAllTagsByImageId(image.Id)

	errs := make(map[string]string)
	form := utils.GenerateForm(&forms.Tag{Image: name}, "/tag", "POST", errs)

	p.SetUser(su)
	p.SetTitle("Image")

	src := fmt.Sprintf("%s/i/%s", utils.AppCfg.Url(), name)
	fullsrc := fmt.Sprintf("%s/o/%s", utils.AppCfg.Url(), name)
	p.Data = ImagePageVars{Name: src, Full: fullsrc, Tags: tags, Form: form}

	encoder.Render(p.Encoding, 200, "image", p, r)
}
