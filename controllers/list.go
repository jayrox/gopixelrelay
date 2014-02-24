package controllers

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"

	"pixelrelay/utils"
	"pixelrelay/models"
)

type ImageLink struct {
	Title    string
	FileName string
}

func List(args martini.Params, r render.Render) {

	files, _ := ioutil.ReadDir(utils.ImageCfg.Root())

	var imageLinks []ImageLink

	for _, f := range files {
		if strings.Contains(f.Name(), ".") && !strings.HasPrefix(f.Name(), ".") {
			imageLinks = append(imageLinks, ImageLink{Title: f.Name(), FileName: f.Name()})
		}
	}

	r.HTML(200, "image_link", imageLinks)
}
