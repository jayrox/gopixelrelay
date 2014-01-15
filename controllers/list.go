package controllers

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"io/ioutil"
	"strings"
)

type ImageLink struct {
	Title    string
	FileName string
}

func List(args martini.Params, r render.Render) {
	files, _ := ioutil.ReadDir("./tmp/")

	var imageLinks []ImageLink

	for _, f := range files {
		if strings.Contains(f.Name(), ".") && !strings.Contains(f.Name(), ".git") {
			imageLinks = append(imageLinks, ImageLink{Title: f.Name(), FileName: f.Name()})
		}
	}

	r.HTML(200, "image_link", imageLinks)
}

func (il *ImageLink) SetFile(file string) {
	il.FileName = file
}
