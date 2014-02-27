package controllers

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"

	"pixelrelay/models"
	"pixelrelay/utils"
)

type ListVars struct {
	ImageLinks []ImageLink
	User       models.User
}

type ImageLink struct {
	Title    string
	FileName string
}

func List(args martini.Params, su models.User, r render.Render) {
	var listVars ListVars

	files, _ := ioutil.ReadDir(utils.ImageCfg.Root())

	var imageLinks []ImageLink

	for _, f := range files {
		if strings.Contains(f.Name(), ".") && !strings.HasPrefix(f.Name(), ".") {
			imageLinks = append(imageLinks, ImageLink{Title: f.Name(), FileName: f.Name()})
		}
	}

	listVars.User = su
	listVars.ImageLinks = imageLinks

	fmt.Println(su)

	r.HTML(200, "image_link", listVars)
}
