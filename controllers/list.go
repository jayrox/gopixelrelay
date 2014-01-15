package controllers

import (
	"io/ioutil"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"strings"
	"pixelrelay/utils"
	"fmt"
)

type ImageLink struct {
	Title string
	FileName string
}

func List(args martini.Params, r render.Render) {
	files, _ := ioutil.ReadDir(utils.ImageCfg.Root())

	var imageLinks []ImageLink
	
	for _, f := range files {
		if strings.Contains(f.Name(), ".") && ! strings.Contains(f.Name(), ".git") {
			imageLinks = append(imageLinks, ImageLink{Title: f.Name(), FileName: f.Name()})
		}
	}

	r.HTML(200, "image_link", imageLinks)
}

func (il *ImageLink) SetFile(file string) {
    il.FileName = file
}