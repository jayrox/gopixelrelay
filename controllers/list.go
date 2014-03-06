package controllers

import (
	"io/ioutil"
	"strings"

	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"

	"pixelrelay/encoder"
	"pixelrelay/models"
	"pixelrelay/utils"
)

type ListVars struct {
	ImageLinks []ImageLink
}

type ImageLink struct {
	Title    string
	FileName string
}

func List(args martini.Params, su models.User, r render.Render, p *models.Page) {
	//var listVars ListVars

	files, _ := ioutil.ReadDir(utils.ImageCfg.Root())

	var imageLinks []ImageLink

	for _, f := range files {
		if strings.Contains(f.Name(), ".") && !strings.HasPrefix(f.Name(), ".") {
			imageLinks = append(imageLinks, ImageLink{Title: f.Name(), FileName: f.Name()})
		}
	}

	p.SetUser(su)
	p.SetTitle("List")
	p.Data = ListVars{ImageLinks: imageLinks}

	encoder.Render(p.Encoding, 200, "image_link", p, r)
}
