package controllers

import (
	"fmt"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"io/ioutil"
	"pixelrelay/db"
	"pixelrelay/utils"
	"strings"
)

func List(args martini.Params, r render.Render) {

	d := db.InitDB()
	images := db.GetAllAlbumImages(&d, "private")
	fmt.Println("images: ", images)

	albums := db.GetAllAlbums(&d)
	fmt.Println("albums: ", albums)

	files, _ := ioutil.ReadDir(utils.ImageCfg.Root())

	var imageLinks []ImageLink

	for _, f := range files {
		if strings.Contains(f.Name(), ".") && !strings.Contains(f.Name(), ".git") {
			imageLinks = append(imageLinks, ImageLink{Title: f.Name(), FileName: f.Name()})
		}
	}

	r.HTML(200, "image_link", imageLinks)
}

type ImageLink struct {
	Title    string
	FileName string
}
